/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedMetadataChoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedMetadataChoiceRequest{}

// PatchedMetadataChoiceRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedMetadataChoiceRequest struct {
	Value *string `json:"value,omitempty"`
	// Higher weights appear later in the list
	Weight *int32 `json:"weight,omitempty"`
	MetadataType *BulkWritableCableRequestStatus `json:"metadata_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedMetadataChoiceRequest PatchedMetadataChoiceRequest

// NewPatchedMetadataChoiceRequest instantiates a new PatchedMetadataChoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedMetadataChoiceRequest() *PatchedMetadataChoiceRequest {
	this := PatchedMetadataChoiceRequest{}
	return &this
}

// NewPatchedMetadataChoiceRequestWithDefaults instantiates a new PatchedMetadataChoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedMetadataChoiceRequestWithDefaults() *PatchedMetadataChoiceRequest {
	this := PatchedMetadataChoiceRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchedMetadataChoiceRequest) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMetadataChoiceRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchedMetadataChoiceRequest) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PatchedMetadataChoiceRequest) SetValue(v string) {
	o.Value = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedMetadataChoiceRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMetadataChoiceRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedMetadataChoiceRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedMetadataChoiceRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *PatchedMetadataChoiceRequest) GetMetadataType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.MetadataType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMetadataChoiceRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.MetadataType) {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *PatchedMetadataChoiceRequest) HasMetadataType() bool {
	if o != nil && !IsNil(o.MetadataType) {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the MetadataType field.
func (o *PatchedMetadataChoiceRequest) SetMetadataType(v BulkWritableCableRequestStatus) {
	o.MetadataType = &v
}

func (o PatchedMetadataChoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedMetadataChoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.MetadataType) {
		toSerialize["metadata_type"] = o.MetadataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedMetadataChoiceRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedMetadataChoiceRequest := _PatchedMetadataChoiceRequest{}

	err = json.Unmarshal(data, &varPatchedMetadataChoiceRequest)

	if err != nil {
		return err
	}

	*o = PatchedMetadataChoiceRequest(varPatchedMetadataChoiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "metadata_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedMetadataChoiceRequest struct {
	value *PatchedMetadataChoiceRequest
	isSet bool
}

func (v NullablePatchedMetadataChoiceRequest) Get() *PatchedMetadataChoiceRequest {
	return v.value
}

func (v *NullablePatchedMetadataChoiceRequest) Set(val *PatchedMetadataChoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedMetadataChoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedMetadataChoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedMetadataChoiceRequest(val *PatchedMetadataChoiceRequest) *NullablePatchedMetadataChoiceRequest {
	return &NullablePatchedMetadataChoiceRequest{value: val, isSet: true}
}

func (v NullablePatchedMetadataChoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedMetadataChoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


