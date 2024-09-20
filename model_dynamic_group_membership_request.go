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

// checks if the DynamicGroupMembershipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicGroupMembershipRequest{}

// DynamicGroupMembershipRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type DynamicGroupMembershipRequest struct {
	Operator OperatorEnum `json:"operator"`
	Weight int32 `json:"weight"`
	Group BulkWritableCableRequestStatus `json:"group"`
	ParentGroup BulkWritableCableRequestStatus `json:"parent_group"`
	AdditionalProperties map[string]interface{}
}

type _DynamicGroupMembershipRequest DynamicGroupMembershipRequest

// NewDynamicGroupMembershipRequest instantiates a new DynamicGroupMembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicGroupMembershipRequest(operator OperatorEnum, weight int32, group BulkWritableCableRequestStatus, parentGroup BulkWritableCableRequestStatus) *DynamicGroupMembershipRequest {
	this := DynamicGroupMembershipRequest{}
	this.Operator = operator
	this.Weight = weight
	this.Group = group
	this.ParentGroup = parentGroup
	return &this
}

// NewDynamicGroupMembershipRequestWithDefaults instantiates a new DynamicGroupMembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicGroupMembershipRequestWithDefaults() *DynamicGroupMembershipRequest {
	this := DynamicGroupMembershipRequest{}
	return &this
}

// GetOperator returns the Operator field value
func (o *DynamicGroupMembershipRequest) GetOperator() OperatorEnum {
	if o == nil {
		var ret OperatorEnum
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *DynamicGroupMembershipRequest) GetOperatorOk() (*OperatorEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *DynamicGroupMembershipRequest) SetOperator(v OperatorEnum) {
	o.Operator = v
}

// GetWeight returns the Weight field value
func (o *DynamicGroupMembershipRequest) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *DynamicGroupMembershipRequest) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *DynamicGroupMembershipRequest) SetWeight(v int32) {
	o.Weight = v
}

// GetGroup returns the Group field value
func (o *DynamicGroupMembershipRequest) GetGroup() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *DynamicGroupMembershipRequest) GetGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *DynamicGroupMembershipRequest) SetGroup(v BulkWritableCableRequestStatus) {
	o.Group = v
}

// GetParentGroup returns the ParentGroup field value
func (o *DynamicGroupMembershipRequest) GetParentGroup() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.ParentGroup
}

// GetParentGroupOk returns a tuple with the ParentGroup field value
// and a boolean to check if the value has been set.
func (o *DynamicGroupMembershipRequest) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentGroup, true
}

// SetParentGroup sets field value
func (o *DynamicGroupMembershipRequest) SetParentGroup(v BulkWritableCableRequestStatus) {
	o.ParentGroup = v
}

func (o DynamicGroupMembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicGroupMembershipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	toSerialize["weight"] = o.Weight
	toSerialize["group"] = o.Group
	toSerialize["parent_group"] = o.ParentGroup

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DynamicGroupMembershipRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operator",
		"weight",
		"group",
		"parent_group",
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

	varDynamicGroupMembershipRequest := _DynamicGroupMembershipRequest{}

	err = json.Unmarshal(data, &varDynamicGroupMembershipRequest)

	if err != nil {
		return err
	}

	*o = DynamicGroupMembershipRequest(varDynamicGroupMembershipRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group")
		delete(additionalProperties, "parent_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDynamicGroupMembershipRequest struct {
	value *DynamicGroupMembershipRequest
	isSet bool
}

func (v NullableDynamicGroupMembershipRequest) Get() *DynamicGroupMembershipRequest {
	return v.value
}

func (v *NullableDynamicGroupMembershipRequest) Set(val *DynamicGroupMembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicGroupMembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicGroupMembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicGroupMembershipRequest(val *DynamicGroupMembershipRequest) *NullableDynamicGroupMembershipRequest {
	return &NullableDynamicGroupMembershipRequest{value: val, isSet: true}
}

func (v NullableDynamicGroupMembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicGroupMembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


