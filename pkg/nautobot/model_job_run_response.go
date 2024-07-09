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

// checks if the JobRunResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobRunResponse{}

// JobRunResponse Serializer representing responses from the JobModelViewSet.run() POST endpoint.
type JobRunResponse struct {
	Schedule ScheduledJob `json:"schedule"`
	JobResult JobResult `json:"job_result"`
	AdditionalProperties map[string]interface{}
}

type _JobRunResponse JobRunResponse

// NewJobRunResponse instantiates a new JobRunResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRunResponse(schedule ScheduledJob, jobResult JobResult) *JobRunResponse {
	this := JobRunResponse{}
	this.Schedule = schedule
	this.JobResult = jobResult
	return &this
}

// NewJobRunResponseWithDefaults instantiates a new JobRunResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRunResponseWithDefaults() *JobRunResponse {
	this := JobRunResponse{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *JobRunResponse) GetSchedule() ScheduledJob {
	if o == nil {
		var ret ScheduledJob
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *JobRunResponse) GetScheduleOk() (*ScheduledJob, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *JobRunResponse) SetSchedule(v ScheduledJob) {
	o.Schedule = v
}

// GetJobResult returns the JobResult field value
func (o *JobRunResponse) GetJobResult() JobResult {
	if o == nil {
		var ret JobResult
		return ret
	}

	return o.JobResult
}

// GetJobResultOk returns a tuple with the JobResult field value
// and a boolean to check if the value has been set.
func (o *JobRunResponse) GetJobResultOk() (*JobResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobResult, true
}

// SetJobResult sets field value
func (o *JobRunResponse) SetJobResult(v JobResult) {
	o.JobResult = v
}

func (o JobRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobRunResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schedule"] = o.Schedule
	toSerialize["job_result"] = o.JobResult

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobRunResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schedule",
		"job_result",
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

	varJobRunResponse := _JobRunResponse{}

	err = json.Unmarshal(data, &varJobRunResponse)

	if err != nil {
		return err
	}

	*o = JobRunResponse(varJobRunResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "job_result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobRunResponse struct {
	value *JobRunResponse
	isSet bool
}

func (v NullableJobRunResponse) Get() *JobRunResponse {
	return v.value
}

func (v *NullableJobRunResponse) Set(val *JobRunResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRunResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRunResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRunResponse(val *JobRunResponse) *NullableJobRunResponse {
	return &NullableJobRunResponse{value: val, isSet: true}
}

func (v NullableJobRunResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRunResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

