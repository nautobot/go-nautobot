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

// checks if the PatchedBulkWritablePowerPanelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritablePowerPanelRequest{}

// PatchedBulkWritablePowerPanelRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritablePowerPanelRequest struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Location *BulkWritableCableRequestStatus `json:"location,omitempty"`
	RackGroup NullableBulkWritableCircuitRequestTenant `json:"rack_group,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritablePowerPanelRequest PatchedBulkWritablePowerPanelRequest

// NewPatchedBulkWritablePowerPanelRequest instantiates a new PatchedBulkWritablePowerPanelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritablePowerPanelRequest(id string) *PatchedBulkWritablePowerPanelRequest {
	this := PatchedBulkWritablePowerPanelRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritablePowerPanelRequestWithDefaults instantiates a new PatchedBulkWritablePowerPanelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritablePowerPanelRequestWithDefaults() *PatchedBulkWritablePowerPanelRequest {
	this := PatchedBulkWritablePowerPanelRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritablePowerPanelRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerPanelRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritablePowerPanelRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerPanelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerPanelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerPanelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritablePowerPanelRequest) SetName(v string) {
	o.Name = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerPanelRequest) GetLocation() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Location) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerPanelRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerPanelRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Location field.
func (o *PatchedBulkWritablePowerPanelRequest) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = &v
}

// GetRackGroup returns the RackGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePowerPanelRequest) GetRackGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.RackGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.RackGroup.Get()
}

// GetRackGroupOk returns a tuple with the RackGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePowerPanelRequest) GetRackGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.RackGroup.Get(), o.RackGroup.IsSet()
}

// HasRackGroup returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerPanelRequest) HasRackGroup() bool {
	if o != nil && o.RackGroup.IsSet() {
		return true
	}

	return false
}

// SetRackGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the RackGroup field.
func (o *PatchedBulkWritablePowerPanelRequest) SetRackGroup(v BulkWritableCircuitRequestTenant) {
	o.RackGroup.Set(&v)
}
// SetRackGroupNil sets the value for RackGroup to be an explicit nil
func (o *PatchedBulkWritablePowerPanelRequest) SetRackGroupNil() {
	o.RackGroup.Set(nil)
}

// UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
func (o *PatchedBulkWritablePowerPanelRequest) UnsetRackGroup() {
	o.RackGroup.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerPanelRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerPanelRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerPanelRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritablePowerPanelRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerPanelRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerPanelRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerPanelRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritablePowerPanelRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerPanelRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerPanelRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerPanelRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritablePowerPanelRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritablePowerPanelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritablePowerPanelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if o.RackGroup.IsSet() {
		toSerialize["rack_group"] = o.RackGroup.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritablePowerPanelRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritablePowerPanelRequest := _PatchedBulkWritablePowerPanelRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritablePowerPanelRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritablePowerPanelRequest(varPatchedBulkWritablePowerPanelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		delete(additionalProperties, "rack_group")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritablePowerPanelRequest struct {
	value *PatchedBulkWritablePowerPanelRequest
	isSet bool
}

func (v NullablePatchedBulkWritablePowerPanelRequest) Get() *PatchedBulkWritablePowerPanelRequest {
	return v.value
}

func (v *NullablePatchedBulkWritablePowerPanelRequest) Set(val *PatchedBulkWritablePowerPanelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritablePowerPanelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritablePowerPanelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritablePowerPanelRequest(val *PatchedBulkWritablePowerPanelRequest) *NullablePatchedBulkWritablePowerPanelRequest {
	return &NullablePatchedBulkWritablePowerPanelRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritablePowerPanelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritablePowerPanelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

