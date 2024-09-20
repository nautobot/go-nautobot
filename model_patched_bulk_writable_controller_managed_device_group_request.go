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

// checks if the PatchedBulkWritableControllerManagedDeviceGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableControllerManagedDeviceGroupRequest{}

// PatchedBulkWritableControllerManagedDeviceGroupRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableControllerManagedDeviceGroupRequest struct {
	Id string `json:"id"`
	// Name of the controller device group
	Name *string `json:"name,omitempty"`
	// Weight of the controller device group, used to sort the groups within its parent group
	Weight *int32 `json:"weight,omitempty"`
	Parent NullableBulkWritableCircuitRequestTenant `json:"parent,omitempty"`
	Controller *BulkWritableControllerManagedDeviceGroupRequestController `json:"controller,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableControllerManagedDeviceGroupRequest PatchedBulkWritableControllerManagedDeviceGroupRequest

// NewPatchedBulkWritableControllerManagedDeviceGroupRequest instantiates a new PatchedBulkWritableControllerManagedDeviceGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableControllerManagedDeviceGroupRequest(id string) *PatchedBulkWritableControllerManagedDeviceGroupRequest {
	this := PatchedBulkWritableControllerManagedDeviceGroupRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableControllerManagedDeviceGroupRequestWithDefaults instantiates a new PatchedBulkWritableControllerManagedDeviceGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableControllerManagedDeviceGroupRequestWithDefaults() *PatchedBulkWritableControllerManagedDeviceGroupRequest {
	this := PatchedBulkWritableControllerManagedDeviceGroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetParent() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Parent field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetParent(v BulkWritableCircuitRequestTenant) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetController() BulkWritableControllerManagedDeviceGroupRequestController {
	if o == nil || IsNil(o.Controller) {
		var ret BulkWritableControllerManagedDeviceGroupRequestController
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetControllerOk() (*BulkWritableControllerManagedDeviceGroupRequestController, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given BulkWritableControllerManagedDeviceGroupRequestController and assigns it to the Controller field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetController(v BulkWritableControllerManagedDeviceGroupRequestController) {
	o.Controller = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableControllerManagedDeviceGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableControllerManagedDeviceGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableControllerManagedDeviceGroupRequest := _PatchedBulkWritableControllerManagedDeviceGroupRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableControllerManagedDeviceGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableControllerManagedDeviceGroupRequest(varPatchedBulkWritableControllerManagedDeviceGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "controller")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableControllerManagedDeviceGroupRequest struct {
	value *PatchedBulkWritableControllerManagedDeviceGroupRequest
	isSet bool
}

func (v NullablePatchedBulkWritableControllerManagedDeviceGroupRequest) Get() *PatchedBulkWritableControllerManagedDeviceGroupRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableControllerManagedDeviceGroupRequest) Set(val *PatchedBulkWritableControllerManagedDeviceGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableControllerManagedDeviceGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableControllerManagedDeviceGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableControllerManagedDeviceGroupRequest(val *PatchedBulkWritableControllerManagedDeviceGroupRequest) *NullablePatchedBulkWritableControllerManagedDeviceGroupRequest {
	return &NullablePatchedBulkWritableControllerManagedDeviceGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableControllerManagedDeviceGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableControllerManagedDeviceGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


