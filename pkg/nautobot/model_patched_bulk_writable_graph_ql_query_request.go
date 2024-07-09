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

// checks if the PatchedBulkWritableGraphQLQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableGraphQLQueryRequest{}

// PatchedBulkWritableGraphQLQueryRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableGraphQLQueryRequest struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Query *string `json:"query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableGraphQLQueryRequest PatchedBulkWritableGraphQLQueryRequest

// NewPatchedBulkWritableGraphQLQueryRequest instantiates a new PatchedBulkWritableGraphQLQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableGraphQLQueryRequest(id string) *PatchedBulkWritableGraphQLQueryRequest {
	this := PatchedBulkWritableGraphQLQueryRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableGraphQLQueryRequestWithDefaults instantiates a new PatchedBulkWritableGraphQLQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableGraphQLQueryRequestWithDefaults() *PatchedBulkWritableGraphQLQueryRequest {
	this := PatchedBulkWritableGraphQLQueryRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableGraphQLQueryRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableGraphQLQueryRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableGraphQLQueryRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableGraphQLQueryRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableGraphQLQueryRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableGraphQLQueryRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableGraphQLQueryRequest) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *PatchedBulkWritableGraphQLQueryRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableGraphQLQueryRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *PatchedBulkWritableGraphQLQueryRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *PatchedBulkWritableGraphQLQueryRequest) SetQuery(v string) {
	o.Query = &v
}

func (o PatchedBulkWritableGraphQLQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableGraphQLQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableGraphQLQueryRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableGraphQLQueryRequest := _PatchedBulkWritableGraphQLQueryRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableGraphQLQueryRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableGraphQLQueryRequest(varPatchedBulkWritableGraphQLQueryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableGraphQLQueryRequest struct {
	value *PatchedBulkWritableGraphQLQueryRequest
	isSet bool
}

func (v NullablePatchedBulkWritableGraphQLQueryRequest) Get() *PatchedBulkWritableGraphQLQueryRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableGraphQLQueryRequest) Set(val *PatchedBulkWritableGraphQLQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableGraphQLQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableGraphQLQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableGraphQLQueryRequest(val *PatchedBulkWritableGraphQLQueryRequest) *NullablePatchedBulkWritableGraphQLQueryRequest {
	return &NullablePatchedBulkWritableGraphQLQueryRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableGraphQLQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableGraphQLQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

