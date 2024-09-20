/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the JobResultStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobResultStatus{}

// JobResultStatus struct for JobResultStatus
type JobResultStatus struct {
	Value *JobResultStatusValue `json:"value,omitempty"`
	Label *JobResultStatusValue `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JobResultStatus JobResultStatus

// NewJobResultStatus instantiates a new JobResultStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResultStatus() *JobResultStatus {
	this := JobResultStatus{}
	return &this
}

// NewJobResultStatusWithDefaults instantiates a new JobResultStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResultStatusWithDefaults() *JobResultStatus {
	this := JobResultStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *JobResultStatus) GetValue() JobResultStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret JobResultStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResultStatus) GetValueOk() (*JobResultStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *JobResultStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given JobResultStatusValue and assigns it to the Value field.
func (o *JobResultStatus) SetValue(v JobResultStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *JobResultStatus) GetLabel() JobResultStatusValue {
	if o == nil || IsNil(o.Label) {
		var ret JobResultStatusValue
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResultStatus) GetLabelOk() (*JobResultStatusValue, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *JobResultStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given JobResultStatusValue and assigns it to the Label field.
func (o *JobResultStatus) SetLabel(v JobResultStatusValue) {
	o.Label = &v
}

func (o JobResultStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobResultStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobResultStatus) UnmarshalJSON(data []byte) (err error) {
	varJobResultStatus := _JobResultStatus{}

	err = json.Unmarshal(data, &varJobResultStatus)

	if err != nil {
		return err
	}

	*o = JobResultStatus(varJobResultStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobResultStatus struct {
	value *JobResultStatus
	isSet bool
}

func (v NullableJobResultStatus) Get() *JobResultStatus {
	return v.value
}

func (v *NullableJobResultStatus) Set(val *JobResultStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResultStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResultStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResultStatus(val *JobResultStatus) *NullableJobResultStatus {
	return &NullableJobResultStatus{value: val, isSet: true}
}

func (v NullableJobResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResultStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


