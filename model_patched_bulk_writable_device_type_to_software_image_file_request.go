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

// checks if the PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest{}

// PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest struct {
	Id string `json:"id"`
	DeviceType *BulkWritableCableRequestStatus `json:"device_type,omitempty"`
	SoftwareImageFile *BulkWritableCableRequestStatus `json:"software_image_file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest

// NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest instantiates a new PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest(id string) *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest {
	this := PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequestWithDefaults instantiates a new PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequestWithDefaults() *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest {
	this := PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) SetId(v string) {
	o.Id = v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetDeviceType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.DeviceType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the DeviceType field.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) SetDeviceType(v BulkWritableCableRequestStatus) {
	o.DeviceType = &v
}

// GetSoftwareImageFile returns the SoftwareImageFile field value if set, zero value otherwise.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetSoftwareImageFile() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.SoftwareImageFile) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.SoftwareImageFile
}

// GetSoftwareImageFileOk returns a tuple with the SoftwareImageFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetSoftwareImageFileOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.SoftwareImageFile) {
		return nil, false
	}
	return o.SoftwareImageFile, true
}

// HasSoftwareImageFile returns a boolean if a field has been set.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) HasSoftwareImageFile() bool {
	if o != nil && !IsNil(o.SoftwareImageFile) {
		return true
	}

	return false
}

// SetSoftwareImageFile gets a reference to the given BulkWritableCableRequestStatus and assigns it to the SoftwareImageFile field.
func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) SetSoftwareImageFile(v BulkWritableCableRequestStatus) {
	o.SoftwareImageFile = &v
}

func (o PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.SoftwareImageFile) {
		toSerialize["software_image_file"] = o.SoftwareImageFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest := _PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest(varPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "software_image_file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest struct {
	value *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest
	isSet bool
}

func (v NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) Get() *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) Set(val *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest(val *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) *NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest {
	return &NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


