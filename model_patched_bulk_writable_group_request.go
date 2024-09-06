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

// checks if the PatchedBulkWritableGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableGroupRequest{}

// PatchedBulkWritableGroupRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableGroupRequest struct {
	Id int32 `json:"id"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableGroupRequest PatchedBulkWritableGroupRequest

// NewPatchedBulkWritableGroupRequest instantiates a new PatchedBulkWritableGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableGroupRequest(id int32) *PatchedBulkWritableGroupRequest {
	this := PatchedBulkWritableGroupRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableGroupRequestWithDefaults instantiates a new PatchedBulkWritableGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableGroupRequestWithDefaults() *PatchedBulkWritableGroupRequest {
	this := PatchedBulkWritableGroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableGroupRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableGroupRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableGroupRequest) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableGroupRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedBulkWritableGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableGroupRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableGroupRequest := _PatchedBulkWritableGroupRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableGroupRequest(varPatchedBulkWritableGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableGroupRequest struct {
	value *PatchedBulkWritableGroupRequest
	isSet bool
}

func (v NullablePatchedBulkWritableGroupRequest) Get() *PatchedBulkWritableGroupRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableGroupRequest) Set(val *PatchedBulkWritableGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableGroupRequest(val *PatchedBulkWritableGroupRequest) *NullablePatchedBulkWritableGroupRequest {
	return &NullablePatchedBulkWritableGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


