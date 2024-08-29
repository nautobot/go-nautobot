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

// checks if the BulkWritableVLANLocationAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableVLANLocationAssignmentRequest{}

// BulkWritableVLANLocationAssignmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableVLANLocationAssignmentRequest struct {
	Id string `json:"id"`
	Vlan BulkWritableCableRequestStatus `json:"vlan"`
	Location BulkWritableCableRequestStatus `json:"location"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableVLANLocationAssignmentRequest BulkWritableVLANLocationAssignmentRequest

// NewBulkWritableVLANLocationAssignmentRequest instantiates a new BulkWritableVLANLocationAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableVLANLocationAssignmentRequest(id string, vlan BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus) *BulkWritableVLANLocationAssignmentRequest {
	this := BulkWritableVLANLocationAssignmentRequest{}
	this.Id = id
	this.Vlan = vlan
	this.Location = location
	return &this
}

// NewBulkWritableVLANLocationAssignmentRequestWithDefaults instantiates a new BulkWritableVLANLocationAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableVLANLocationAssignmentRequestWithDefaults() *BulkWritableVLANLocationAssignmentRequest {
	this := BulkWritableVLANLocationAssignmentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableVLANLocationAssignmentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableVLANLocationAssignmentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableVLANLocationAssignmentRequest) SetId(v string) {
	o.Id = v
}

// GetVlan returns the Vlan field value
func (o *BulkWritableVLANLocationAssignmentRequest) GetVlan() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *BulkWritableVLANLocationAssignmentRequest) GetVlanOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *BulkWritableVLANLocationAssignmentRequest) SetVlan(v BulkWritableCableRequestStatus) {
	o.Vlan = v
}

// GetLocation returns the Location field value
func (o *BulkWritableVLANLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *BulkWritableVLANLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *BulkWritableVLANLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = v
}

func (o BulkWritableVLANLocationAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableVLANLocationAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["vlan"] = o.Vlan
	toSerialize["location"] = o.Location

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableVLANLocationAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"vlan",
		"location",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBulkWritableVLANLocationAssignmentRequest := _BulkWritableVLANLocationAssignmentRequest{}

	err = json.Unmarshal(data, &varBulkWritableVLANLocationAssignmentRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableVLANLocationAssignmentRequest(varBulkWritableVLANLocationAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableVLANLocationAssignmentRequest struct {
	value *BulkWritableVLANLocationAssignmentRequest
	isSet bool
}

func (v NullableBulkWritableVLANLocationAssignmentRequest) Get() *BulkWritableVLANLocationAssignmentRequest {
	return v.value
}

func (v *NullableBulkWritableVLANLocationAssignmentRequest) Set(val *BulkWritableVLANLocationAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableVLANLocationAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableVLANLocationAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableVLANLocationAssignmentRequest(val *BulkWritableVLANLocationAssignmentRequest) *NullableBulkWritableVLANLocationAssignmentRequest {
	return &NullableBulkWritableVLANLocationAssignmentRequest{value: val, isSet: true}
}

func (v NullableBulkWritableVLANLocationAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableVLANLocationAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

