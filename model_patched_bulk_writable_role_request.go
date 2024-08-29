/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the PatchedBulkWritableRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableRoleRequest{}

// PatchedBulkWritableRoleRequest Serializer for `Role` objects.
type PatchedBulkWritableRoleRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types,omitempty"`
	Name *string `json:"name,omitempty"`
	// RGB color in hexadecimal (e.g. 00ff00)
	Color *string `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	Description *string `json:"description,omitempty"`
	Weight NullableInt32 `json:"weight,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableRoleRequest PatchedBulkWritableRoleRequest

// NewPatchedBulkWritableRoleRequest instantiates a new PatchedBulkWritableRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableRoleRequest(id string) *PatchedBulkWritableRoleRequest {
	this := PatchedBulkWritableRoleRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableRoleRequestWithDefaults instantiates a new PatchedBulkWritableRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableRoleRequestWithDefaults() *PatchedBulkWritableRoleRequest {
	this := PatchedBulkWritableRoleRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableRoleRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableRoleRequest) SetId(v string) {
	o.Id = v
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *PatchedBulkWritableRoleRequest) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *PatchedBulkWritableRoleRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableRoleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableRoleRequest) SetName(v string) {
	o.Name = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PatchedBulkWritableRoleRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PatchedBulkWritableRoleRequest) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableRoleRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret int32
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableRoleRequest) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableInt32 and assigns it to the Weight field.
func (o *PatchedBulkWritableRoleRequest) SetWeight(v int32) {
	o.Weight.Set(&v)
}
// SetWeightNil sets the value for Weight to be an explicit nil
func (o *PatchedBulkWritableRoleRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *PatchedBulkWritableRoleRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableRoleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableRoleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableRoleRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableRoleRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableRoleRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableRoleRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritableRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ContentTypes) {
		toSerialize["content_types"] = o.ContentTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
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

func (o *PatchedBulkWritableRoleRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableRoleRequest := _PatchedBulkWritableRoleRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableRoleRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableRoleRequest(varPatchedBulkWritableRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "color")
		delete(additionalProperties, "description")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableRoleRequest struct {
	value *PatchedBulkWritableRoleRequest
	isSet bool
}

func (v NullablePatchedBulkWritableRoleRequest) Get() *PatchedBulkWritableRoleRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableRoleRequest) Set(val *PatchedBulkWritableRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableRoleRequest(val *PatchedBulkWritableRoleRequest) *NullablePatchedBulkWritableRoleRequest {
	return &NullablePatchedBulkWritableRoleRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

