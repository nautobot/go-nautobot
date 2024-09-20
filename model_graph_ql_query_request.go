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

// checks if the GraphQLQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLQueryRequest{}

// GraphQLQueryRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type GraphQLQueryRequest struct {
	Name string `json:"name"`
	Query string `json:"query"`
	AdditionalProperties map[string]interface{}
}

type _GraphQLQueryRequest GraphQLQueryRequest

// NewGraphQLQueryRequest instantiates a new GraphQLQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLQueryRequest(name string, query string) *GraphQLQueryRequest {
	this := GraphQLQueryRequest{}
	this.Name = name
	this.Query = query
	return &this
}

// NewGraphQLQueryRequestWithDefaults instantiates a new GraphQLQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLQueryRequestWithDefaults() *GraphQLQueryRequest {
	this := GraphQLQueryRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GraphQLQueryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GraphQLQueryRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GraphQLQueryRequest) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *GraphQLQueryRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GraphQLQueryRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *GraphQLQueryRequest) SetQuery(v string) {
	o.Query = v
}

func (o GraphQLQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GraphQLQueryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"query",
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

	varGraphQLQueryRequest := _GraphQLQueryRequest{}

	err = json.Unmarshal(data, &varGraphQLQueryRequest)

	if err != nil {
		return err
	}

	*o = GraphQLQueryRequest(varGraphQLQueryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGraphQLQueryRequest struct {
	value *GraphQLQueryRequest
	isSet bool
}

func (v NullableGraphQLQueryRequest) Get() *GraphQLQueryRequest {
	return v.value
}

func (v *NullableGraphQLQueryRequest) Set(val *GraphQLQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLQueryRequest(val *GraphQLQueryRequest) *NullableGraphQLQueryRequest {
	return &NullableGraphQLQueryRequest{value: val, isSet: true}
}

func (v NullableGraphQLQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


