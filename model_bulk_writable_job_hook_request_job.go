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

// checks if the BulkWritableJobHookRequestJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableJobHookRequestJob{}

// BulkWritableJobHookRequestJob The job that this job hook will initiate
type BulkWritableJobHookRequestJob struct {
	Id *BulkWritableCableRequestStatusId `json:"id,omitempty"`
	ObjectType *string `json:"object_type,omitempty" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableJobHookRequestJob BulkWritableJobHookRequestJob

// NewBulkWritableJobHookRequestJob instantiates a new BulkWritableJobHookRequestJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableJobHookRequestJob() *BulkWritableJobHookRequestJob {
	this := BulkWritableJobHookRequestJob{}
	return &this
}

// NewBulkWritableJobHookRequestJobWithDefaults instantiates a new BulkWritableJobHookRequestJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableJobHookRequestJobWithDefaults() *BulkWritableJobHookRequestJob {
	this := BulkWritableJobHookRequestJob{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkWritableJobHookRequestJob) GetId() BulkWritableCableRequestStatusId {
	if o == nil || IsNil(o.Id) {
		var ret BulkWritableCableRequestStatusId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableJobHookRequestJob) GetIdOk() (*BulkWritableCableRequestStatusId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkWritableJobHookRequestJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given BulkWritableCableRequestStatusId and assigns it to the Id field.
func (o *BulkWritableJobHookRequestJob) SetId(v BulkWritableCableRequestStatusId) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BulkWritableJobHookRequestJob) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableJobHookRequestJob) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BulkWritableJobHookRequestJob) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *BulkWritableJobHookRequestJob) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BulkWritableJobHookRequestJob) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableJobHookRequestJob) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BulkWritableJobHookRequestJob) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BulkWritableJobHookRequestJob) SetUrl(v string) {
	o.Url = &v
}

func (o BulkWritableJobHookRequestJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableJobHookRequestJob) ToMap() (map[string]interface{}, error) {
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

func (o *BulkWritableJobHookRequestJob) UnmarshalJSON(data []byte) (err error) {
	varBulkWritableJobHookRequestJob := _BulkWritableJobHookRequestJob{}

	err = json.Unmarshal(data, &varBulkWritableJobHookRequestJob)

	if err != nil {
		return err
	}

	*o = BulkWritableJobHookRequestJob(varBulkWritableJobHookRequestJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableJobHookRequestJob struct {
	value *BulkWritableJobHookRequestJob
	isSet bool
}

func (v NullableBulkWritableJobHookRequestJob) Get() *BulkWritableJobHookRequestJob {
	return v.value
}

func (v *NullableBulkWritableJobHookRequestJob) Set(val *BulkWritableJobHookRequestJob) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableJobHookRequestJob) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableJobHookRequestJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableJobHookRequestJob(val *BulkWritableJobHookRequestJob) *NullableBulkWritableJobHookRequestJob {
	return &NullableBulkWritableJobHookRequestJob{value: val, isSet: true}
}

func (v NullableBulkWritableJobHookRequestJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableJobHookRequestJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


