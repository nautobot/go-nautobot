/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedUserSavedViewAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUserSavedViewAssociationRequest{}

// PatchedUserSavedViewAssociationRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedUserSavedViewAssociationRequest struct {
	ViewName *string `json:"view_name,omitempty"`
	SavedView *BulkWritableCableRequestStatus `json:"saved_view,omitempty"`
	User *BulkWritableCableRequestStatus `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedUserSavedViewAssociationRequest PatchedUserSavedViewAssociationRequest

// NewPatchedUserSavedViewAssociationRequest instantiates a new PatchedUserSavedViewAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserSavedViewAssociationRequest() *PatchedUserSavedViewAssociationRequest {
	this := PatchedUserSavedViewAssociationRequest{}
	return &this
}

// NewPatchedUserSavedViewAssociationRequestWithDefaults instantiates a new PatchedUserSavedViewAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserSavedViewAssociationRequestWithDefaults() *PatchedUserSavedViewAssociationRequest {
	this := PatchedUserSavedViewAssociationRequest{}
	return &this
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *PatchedUserSavedViewAssociationRequest) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSavedViewAssociationRequest) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *PatchedUserSavedViewAssociationRequest) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *PatchedUserSavedViewAssociationRequest) SetViewName(v string) {
	o.ViewName = &v
}

// GetSavedView returns the SavedView field value if set, zero value otherwise.
func (o *PatchedUserSavedViewAssociationRequest) GetSavedView() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.SavedView) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.SavedView
}

// GetSavedViewOk returns a tuple with the SavedView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSavedViewAssociationRequest) GetSavedViewOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.SavedView) {
		return nil, false
	}
	return o.SavedView, true
}

// HasSavedView returns a boolean if a field has been set.
func (o *PatchedUserSavedViewAssociationRequest) HasSavedView() bool {
	if o != nil && !IsNil(o.SavedView) {
		return true
	}

	return false
}

// SetSavedView gets a reference to the given BulkWritableCableRequestStatus and assigns it to the SavedView field.
func (o *PatchedUserSavedViewAssociationRequest) SetSavedView(v BulkWritableCableRequestStatus) {
	o.SavedView = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedUserSavedViewAssociationRequest) GetUser() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.User) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSavedViewAssociationRequest) GetUserOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedUserSavedViewAssociationRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given BulkWritableCableRequestStatus and assigns it to the User field.
func (o *PatchedUserSavedViewAssociationRequest) SetUser(v BulkWritableCableRequestStatus) {
	o.User = &v
}

func (o PatchedUserSavedViewAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUserSavedViewAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViewName) {
		toSerialize["view_name"] = o.ViewName
	}
	if !IsNil(o.SavedView) {
		toSerialize["saved_view"] = o.SavedView
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedUserSavedViewAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedUserSavedViewAssociationRequest := _PatchedUserSavedViewAssociationRequest{}

	err = json.Unmarshal(data, &varPatchedUserSavedViewAssociationRequest)

	if err != nil {
		return err
	}

	*o = PatchedUserSavedViewAssociationRequest(varPatchedUserSavedViewAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "view_name")
		delete(additionalProperties, "saved_view")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedUserSavedViewAssociationRequest struct {
	value *PatchedUserSavedViewAssociationRequest
	isSet bool
}

func (v NullablePatchedUserSavedViewAssociationRequest) Get() *PatchedUserSavedViewAssociationRequest {
	return v.value
}

func (v *NullablePatchedUserSavedViewAssociationRequest) Set(val *PatchedUserSavedViewAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserSavedViewAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserSavedViewAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserSavedViewAssociationRequest(val *PatchedUserSavedViewAssociationRequest) *NullablePatchedUserSavedViewAssociationRequest {
	return &NullablePatchedUserSavedViewAssociationRequest{value: val, isSet: true}
}

func (v NullablePatchedUserSavedViewAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserSavedViewAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


