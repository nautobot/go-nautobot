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

// checks if the CloudNetworkPrefixAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudNetworkPrefixAssignmentRequest{}

// CloudNetworkPrefixAssignmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CloudNetworkPrefixAssignmentRequest struct {
	CloudNetwork BulkWritableCableRequestStatus `json:"cloud_network"`
	Prefix BulkWritableCableRequestStatus `json:"prefix"`
	AdditionalProperties map[string]interface{}
}

type _CloudNetworkPrefixAssignmentRequest CloudNetworkPrefixAssignmentRequest

// NewCloudNetworkPrefixAssignmentRequest instantiates a new CloudNetworkPrefixAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkPrefixAssignmentRequest(cloudNetwork BulkWritableCableRequestStatus, prefix BulkWritableCableRequestStatus) *CloudNetworkPrefixAssignmentRequest {
	this := CloudNetworkPrefixAssignmentRequest{}
	this.CloudNetwork = cloudNetwork
	this.Prefix = prefix
	return &this
}

// NewCloudNetworkPrefixAssignmentRequestWithDefaults instantiates a new CloudNetworkPrefixAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkPrefixAssignmentRequestWithDefaults() *CloudNetworkPrefixAssignmentRequest {
	this := CloudNetworkPrefixAssignmentRequest{}
	return &this
}

// GetCloudNetwork returns the CloudNetwork field value
func (o *CloudNetworkPrefixAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.CloudNetwork
}

// GetCloudNetworkOk returns a tuple with the CloudNetwork field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkPrefixAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudNetwork, true
}

// SetCloudNetwork sets field value
func (o *CloudNetworkPrefixAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus) {
	o.CloudNetwork = v
}

// GetPrefix returns the Prefix field value
func (o *CloudNetworkPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *CloudNetworkPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus) {
	o.Prefix = v
}

func (o CloudNetworkPrefixAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudNetworkPrefixAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_network"] = o.CloudNetwork
	toSerialize["prefix"] = o.Prefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudNetworkPrefixAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_network",
		"prefix",
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

	varCloudNetworkPrefixAssignmentRequest := _CloudNetworkPrefixAssignmentRequest{}

	err = json.Unmarshal(data, &varCloudNetworkPrefixAssignmentRequest)

	if err != nil {
		return err
	}

	*o = CloudNetworkPrefixAssignmentRequest(varCloudNetworkPrefixAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_network")
		delete(additionalProperties, "prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudNetworkPrefixAssignmentRequest struct {
	value *CloudNetworkPrefixAssignmentRequest
	isSet bool
}

func (v NullableCloudNetworkPrefixAssignmentRequest) Get() *CloudNetworkPrefixAssignmentRequest {
	return v.value
}

func (v *NullableCloudNetworkPrefixAssignmentRequest) Set(val *CloudNetworkPrefixAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkPrefixAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkPrefixAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkPrefixAssignmentRequest(val *CloudNetworkPrefixAssignmentRequest) *NullableCloudNetworkPrefixAssignmentRequest {
	return &NullableCloudNetworkPrefixAssignmentRequest{value: val, isSet: true}
}

func (v NullableCloudNetworkPrefixAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkPrefixAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


