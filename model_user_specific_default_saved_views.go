/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the UserSpecificDefaultSavedViews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSpecificDefaultSavedViews{}

// UserSpecificDefaultSavedViews struct for UserSpecificDefaultSavedViews
type UserSpecificDefaultSavedViews struct {
	Id *BulkWritableCableRequestStatusId `json:"id,omitempty"`
	ObjectType *string `json:"object_type,omitempty" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSpecificDefaultSavedViews UserSpecificDefaultSavedViews

// NewUserSpecificDefaultSavedViews instantiates a new UserSpecificDefaultSavedViews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSpecificDefaultSavedViews() *UserSpecificDefaultSavedViews {
	this := UserSpecificDefaultSavedViews{}
	return &this
}

// NewUserSpecificDefaultSavedViewsWithDefaults instantiates a new UserSpecificDefaultSavedViews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSpecificDefaultSavedViewsWithDefaults() *UserSpecificDefaultSavedViews {
	this := UserSpecificDefaultSavedViews{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSpecificDefaultSavedViews) GetId() BulkWritableCableRequestStatusId {
	if o == nil || IsNil(o.Id) {
		var ret BulkWritableCableRequestStatusId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSpecificDefaultSavedViews) GetIdOk() (*BulkWritableCableRequestStatusId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSpecificDefaultSavedViews) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given BulkWritableCableRequestStatusId and assigns it to the Id field.
func (o *UserSpecificDefaultSavedViews) SetId(v BulkWritableCableRequestStatusId) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *UserSpecificDefaultSavedViews) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSpecificDefaultSavedViews) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *UserSpecificDefaultSavedViews) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *UserSpecificDefaultSavedViews) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserSpecificDefaultSavedViews) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSpecificDefaultSavedViews) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserSpecificDefaultSavedViews) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserSpecificDefaultSavedViews) SetUrl(v string) {
	o.Url = &v
}

func (o UserSpecificDefaultSavedViews) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSpecificDefaultSavedViews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSpecificDefaultSavedViews) UnmarshalJSON(data []byte) (err error) {
	varUserSpecificDefaultSavedViews := _UserSpecificDefaultSavedViews{}

	err = json.Unmarshal(data, &varUserSpecificDefaultSavedViews)

	if err != nil {
		return err
	}

	*o = UserSpecificDefaultSavedViews(varUserSpecificDefaultSavedViews)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSpecificDefaultSavedViews struct {
	value *UserSpecificDefaultSavedViews
	isSet bool
}

func (v NullableUserSpecificDefaultSavedViews) Get() *UserSpecificDefaultSavedViews {
	return v.value
}

func (v *NullableUserSpecificDefaultSavedViews) Set(val *UserSpecificDefaultSavedViews) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSpecificDefaultSavedViews) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSpecificDefaultSavedViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSpecificDefaultSavedViews(val *UserSpecificDefaultSavedViews) *NullableUserSpecificDefaultSavedViews {
	return &NullableUserSpecificDefaultSavedViews{value: val, isSet: true}
}

func (v NullableUserSpecificDefaultSavedViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSpecificDefaultSavedViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


