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

// checks if the NestedLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedLocation{}

// NestedLocation Add a `tree_depth` field to non-nested model serializers based on django-tree-queries.
type NestedLocation struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _NestedLocation NestedLocation

// NewNestedLocation instantiates a new NestedLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedLocation(id string, objectType string, url string) *NestedLocation {
	this := NestedLocation{}
	this.Id = id
	this.ObjectType = objectType
	this.Url = url
	return &this
}

// NewNestedLocationWithDefaults instantiates a new NestedLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedLocationWithDefaults() *NestedLocation {
	this := NestedLocation{}
	return &this
}

// GetId returns the Id field value
func (o *NestedLocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedLocation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedLocation) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *NestedLocation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NestedLocation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NestedLocation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *NestedLocation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedLocation) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedLocation) SetUrl(v string) {
	o.Url = v
}

func (o NestedLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"url",
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

	varNestedLocation := _NestedLocation{}

	err = json.Unmarshal(data, &varNestedLocation)

	if err != nil {
		return err
	}

	*o = NestedLocation(varNestedLocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedLocation struct {
	value *NestedLocation
	isSet bool
}

func (v NullableNestedLocation) Get() *NestedLocation {
	return v.value
}

func (v *NullableNestedLocation) Set(val *NestedLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedLocation(val *NestedLocation) *NullableNestedLocation {
	return &NullableNestedLocation{value: val, isSet: true}
}

func (v NullableNestedLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


