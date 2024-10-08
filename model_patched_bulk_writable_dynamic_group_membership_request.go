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

// checks if the PatchedBulkWritableDynamicGroupMembershipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableDynamicGroupMembershipRequest{}

// PatchedBulkWritableDynamicGroupMembershipRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableDynamicGroupMembershipRequest struct {
	Id string `json:"id"`
	Operator *OperatorEnum `json:"operator,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Group *BulkWritableCableRequestStatus `json:"group,omitempty"`
	ParentGroup *BulkWritableCableRequestStatus `json:"parent_group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableDynamicGroupMembershipRequest PatchedBulkWritableDynamicGroupMembershipRequest

// NewPatchedBulkWritableDynamicGroupMembershipRequest instantiates a new PatchedBulkWritableDynamicGroupMembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableDynamicGroupMembershipRequest(id string) *PatchedBulkWritableDynamicGroupMembershipRequest {
	this := PatchedBulkWritableDynamicGroupMembershipRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableDynamicGroupMembershipRequestWithDefaults instantiates a new PatchedBulkWritableDynamicGroupMembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableDynamicGroupMembershipRequestWithDefaults() *PatchedBulkWritableDynamicGroupMembershipRequest {
	this := PatchedBulkWritableDynamicGroupMembershipRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetId(v string) {
	o.Id = v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetOperator() OperatorEnum {
	if o == nil || IsNil(o.Operator) {
		var ret OperatorEnum
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetOperatorOk() (*OperatorEnum, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given OperatorEnum and assigns it to the Operator field.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetOperator(v OperatorEnum) {
	o.Operator = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetGroup() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Group) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Group field.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetGroup(v BulkWritableCableRequestStatus) {
	o.Group = &v
}

// GetParentGroup returns the ParentGroup field value if set, zero value otherwise.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetParentGroup() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.ParentGroup) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.ParentGroup
}

// GetParentGroupOk returns a tuple with the ParentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.ParentGroup) {
		return nil, false
	}
	return o.ParentGroup, true
}

// HasParentGroup returns a boolean if a field has been set.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasParentGroup() bool {
	if o != nil && !IsNil(o.ParentGroup) {
		return true
	}

	return false
}

// SetParentGroup gets a reference to the given BulkWritableCableRequestStatus and assigns it to the ParentGroup field.
func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetParentGroup(v BulkWritableCableRequestStatus) {
	o.ParentGroup = &v
}

func (o PatchedBulkWritableDynamicGroupMembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableDynamicGroupMembershipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.ParentGroup) {
		toSerialize["parent_group"] = o.ParentGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableDynamicGroupMembershipRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableDynamicGroupMembershipRequest := _PatchedBulkWritableDynamicGroupMembershipRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableDynamicGroupMembershipRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableDynamicGroupMembershipRequest(varPatchedBulkWritableDynamicGroupMembershipRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group")
		delete(additionalProperties, "parent_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableDynamicGroupMembershipRequest struct {
	value *PatchedBulkWritableDynamicGroupMembershipRequest
	isSet bool
}

func (v NullablePatchedBulkWritableDynamicGroupMembershipRequest) Get() *PatchedBulkWritableDynamicGroupMembershipRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableDynamicGroupMembershipRequest) Set(val *PatchedBulkWritableDynamicGroupMembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableDynamicGroupMembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableDynamicGroupMembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableDynamicGroupMembershipRequest(val *PatchedBulkWritableDynamicGroupMembershipRequest) *NullablePatchedBulkWritableDynamicGroupMembershipRequest {
	return &NullablePatchedBulkWritableDynamicGroupMembershipRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableDynamicGroupMembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableDynamicGroupMembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


