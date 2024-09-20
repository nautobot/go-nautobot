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

// checks if the NestedModuleBayTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedModuleBayTemplate{}

// NestedModuleBayTemplate Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type NestedModuleBayTemplate struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _NestedModuleBayTemplate NestedModuleBayTemplate

// NewNestedModuleBayTemplate instantiates a new NestedModuleBayTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedModuleBayTemplate(id string, objectType string, url string) *NestedModuleBayTemplate {
	this := NestedModuleBayTemplate{}
	this.Id = id
	this.ObjectType = objectType
	this.Url = url
	return &this
}

// NewNestedModuleBayTemplateWithDefaults instantiates a new NestedModuleBayTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedModuleBayTemplateWithDefaults() *NestedModuleBayTemplate {
	this := NestedModuleBayTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *NestedModuleBayTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBayTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedModuleBayTemplate) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *NestedModuleBayTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBayTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NestedModuleBayTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *NestedModuleBayTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBayTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedModuleBayTemplate) SetUrl(v string) {
	o.Url = v
}

func (o NestedModuleBayTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedModuleBayTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedModuleBayTemplate) UnmarshalJSON(data []byte) (err error) {
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

	varNestedModuleBayTemplate := _NestedModuleBayTemplate{}

	err = json.Unmarshal(data, &varNestedModuleBayTemplate)

	if err != nil {
		return err
	}

	*o = NestedModuleBayTemplate(varNestedModuleBayTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedModuleBayTemplate struct {
	value *NestedModuleBayTemplate
	isSet bool
}

func (v NullableNestedModuleBayTemplate) Get() *NestedModuleBayTemplate {
	return v.value
}

func (v *NullableNestedModuleBayTemplate) Set(val *NestedModuleBayTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedModuleBayTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedModuleBayTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedModuleBayTemplate(val *NestedModuleBayTemplate) *NullableNestedModuleBayTemplate {
	return &NullableNestedModuleBayTemplate{value: val, isSet: true}
}

func (v NullableNestedModuleBayTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedModuleBayTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


