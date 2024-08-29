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

// checks if the PatchedBulkWritableMetadataChoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableMetadataChoiceRequest{}

// PatchedBulkWritableMetadataChoiceRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableMetadataChoiceRequest struct {
	Id string `json:"id"`
	Value *string `json:"value,omitempty"`
	// Higher weights appear later in the list
	Weight *int32 `json:"weight,omitempty"`
	MetadataType *BulkWritableCableRequestStatus `json:"metadata_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableMetadataChoiceRequest PatchedBulkWritableMetadataChoiceRequest

// NewPatchedBulkWritableMetadataChoiceRequest instantiates a new PatchedBulkWritableMetadataChoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableMetadataChoiceRequest(id string) *PatchedBulkWritableMetadataChoiceRequest {
	this := PatchedBulkWritableMetadataChoiceRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableMetadataChoiceRequestWithDefaults instantiates a new PatchedBulkWritableMetadataChoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableMetadataChoiceRequestWithDefaults() *PatchedBulkWritableMetadataChoiceRequest {
	this := PatchedBulkWritableMetadataChoiceRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableMetadataChoiceRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableMetadataChoiceRequest) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PatchedBulkWritableMetadataChoiceRequest) SetValue(v string) {
	o.Value = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedBulkWritableMetadataChoiceRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetMetadataType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.MetadataType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.MetadataType) {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *PatchedBulkWritableMetadataChoiceRequest) HasMetadataType() bool {
	if o != nil && !IsNil(o.MetadataType) {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the MetadataType field.
func (o *PatchedBulkWritableMetadataChoiceRequest) SetMetadataType(v BulkWritableCableRequestStatus) {
	o.MetadataType = &v
}

func (o PatchedBulkWritableMetadataChoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableMetadataChoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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

func (o *PatchedBulkWritableMetadataChoiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPatchedBulkWritableMetadataChoiceRequest := _PatchedBulkWritableMetadataChoiceRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableMetadataChoiceRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableMetadataChoiceRequest(varPatchedBulkWritableMetadataChoiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "value")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "metadata_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableMetadataChoiceRequest struct {
	value *PatchedBulkWritableMetadataChoiceRequest
	isSet bool
}

func (v NullablePatchedBulkWritableMetadataChoiceRequest) Get() *PatchedBulkWritableMetadataChoiceRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableMetadataChoiceRequest) Set(val *PatchedBulkWritableMetadataChoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableMetadataChoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableMetadataChoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableMetadataChoiceRequest(val *PatchedBulkWritableMetadataChoiceRequest) *NullablePatchedBulkWritableMetadataChoiceRequest {
	return &NullablePatchedBulkWritableMetadataChoiceRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableMetadataChoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableMetadataChoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

