/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedNoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedNoteRequest{}

// PatchedNoteRequest This base serializer implements common fields and logic for all ModelSerializers.  Namely, it:  - defines the `display` field which exposes a human friendly value for the given object. - ensures that `id` field is always present on the serializer as well. - ensures that `created` and `last_updated` fields are always present if applicable to this model and serializer. - ensures that `object_type` field is always present on the serializer which represents the content-type of this   serializer's associated model (e.g. \"dcim.device\"). This is required as the OpenAPI schema, using the   PolymorphicProxySerializer class defined below, relies upon this field as a way to identify to the client   which of several possible serializers are in use for a given attribute. - supports `?depth` query parameter. It is passed in as `nested_depth` to the `build_nested_field()` function   to enable the dynamic generation of nested serializers.
type PatchedNoteRequest struct {
	AssignedObjectType *string `json:"assigned_object_type,omitempty"`
	AssignedObjectId *string `json:"assigned_object_id,omitempty"`
	Note *string `json:"note,omitempty"`
	User NullableBulkWritableCircuitRequestTenant `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedNoteRequest PatchedNoteRequest

// NewPatchedNoteRequest instantiates a new PatchedNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNoteRequest() *PatchedNoteRequest {
	this := PatchedNoteRequest{}
	return &this
}

// NewPatchedNoteRequestWithDefaults instantiates a new PatchedNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNoteRequestWithDefaults() *PatchedNoteRequest {
	this := PatchedNoteRequest{}
	return &this
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise.
func (o *PatchedNoteRequest) GetAssignedObjectType() string {
	if o == nil || IsNil(o.AssignedObjectType) {
		var ret string
		return ret
	}
	return *o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNoteRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedObjectType) {
		return nil, false
	}
	return o.AssignedObjectType, true
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasAssignedObjectType() bool {
	if o != nil && !IsNil(o.AssignedObjectType) {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given string and assigns it to the AssignedObjectType field.
func (o *PatchedNoteRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = &v
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise.
func (o *PatchedNoteRequest) GetAssignedObjectId() string {
	if o == nil || IsNil(o.AssignedObjectId) {
		var ret string
		return ret
	}
	return *o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNoteRequest) GetAssignedObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedObjectId) {
		return nil, false
	}
	return o.AssignedObjectId, true
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasAssignedObjectId() bool {
	if o != nil && !IsNil(o.AssignedObjectId) {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given string and assigns it to the AssignedObjectId field.
func (o *PatchedNoteRequest) SetAssignedObjectId(v string) {
	o.AssignedObjectId = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *PatchedNoteRequest) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNoteRequest) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *PatchedNoteRequest) SetNote(v string) {
	o.Note = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedNoteRequest) GetUser() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.User.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedNoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the User field.
func (o *PatchedNoteRequest) SetUser(v BulkWritableCircuitRequestTenant) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *PatchedNoteRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PatchedNoteRequest) UnsetUser() {
	o.User.Unset()
}

func (o PatchedNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignedObjectType) {
		toSerialize["assigned_object_type"] = o.AssignedObjectType
	}
	if !IsNil(o.AssignedObjectId) {
		toSerialize["assigned_object_id"] = o.AssignedObjectId
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedNoteRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedNoteRequest := _PatchedNoteRequest{}

	err = json.Unmarshal(data, &varPatchedNoteRequest)

	if err != nil {
		return err
	}

	*o = PatchedNoteRequest(varPatchedNoteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "note")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedNoteRequest struct {
	value *PatchedNoteRequest
	isSet bool
}

func (v NullablePatchedNoteRequest) Get() *PatchedNoteRequest {
	return v.value
}

func (v *NullablePatchedNoteRequest) Set(val *PatchedNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNoteRequest(val *PatchedNoteRequest) *NullablePatchedNoteRequest {
	return &NullablePatchedNoteRequest{value: val, isSet: true}
}

func (v NullablePatchedNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

