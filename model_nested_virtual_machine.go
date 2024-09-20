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

// checks if the NestedVirtualMachine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVirtualMachine{}

// NestedVirtualMachine Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type NestedVirtualMachine struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _NestedVirtualMachine NestedVirtualMachine

// NewNestedVirtualMachine instantiates a new NestedVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVirtualMachine(id string, objectType string, url string) *NestedVirtualMachine {
	this := NestedVirtualMachine{}
	this.Id = id
	this.ObjectType = objectType
	this.Url = url
	return &this
}

// NewNestedVirtualMachineWithDefaults instantiates a new NestedVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVirtualMachineWithDefaults() *NestedVirtualMachine {
	this := NestedVirtualMachine{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVirtualMachine) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVirtualMachine) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *NestedVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NestedVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *NestedVirtualMachine) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualMachine) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVirtualMachine) SetUrl(v string) {
	o.Url = v
}

func (o NestedVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVirtualMachine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVirtualMachine) UnmarshalJSON(data []byte) (err error) {
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

	varNestedVirtualMachine := _NestedVirtualMachine{}

	err = json.Unmarshal(data, &varNestedVirtualMachine)

	if err != nil {
		return err
	}

	*o = NestedVirtualMachine(varNestedVirtualMachine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVirtualMachine struct {
	value *NestedVirtualMachine
	isSet bool
}

func (v NullableNestedVirtualMachine) Get() *NestedVirtualMachine {
	return v.value
}

func (v *NullableNestedVirtualMachine) Set(val *NestedVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVirtualMachine(val *NestedVirtualMachine) *NullableNestedVirtualMachine {
	return &NullableNestedVirtualMachine{value: val, isSet: true}
}

func (v NullableNestedVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


