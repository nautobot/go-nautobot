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

// checks if the PatchedConfigContextSchemaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedConfigContextSchemaRequest{}

// PatchedConfigContextSchemaRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedConfigContextSchemaRequest struct {
	OwnerContentType NullableString `json:"owner_content_type,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// A JSON Schema document which is used to validate a config context object.
	DataSchema interface{} `json:"data_schema,omitempty"`
	OwnerObjectId NullableString `json:"owner_object_id,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedConfigContextSchemaRequest PatchedConfigContextSchemaRequest

// NewPatchedConfigContextSchemaRequest instantiates a new PatchedConfigContextSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedConfigContextSchemaRequest() *PatchedConfigContextSchemaRequest {
	this := PatchedConfigContextSchemaRequest{}
	return &this
}

// NewPatchedConfigContextSchemaRequestWithDefaults instantiates a new PatchedConfigContextSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedConfigContextSchemaRequestWithDefaults() *PatchedConfigContextSchemaRequest {
	this := PatchedConfigContextSchemaRequest{}
	return &this
}

// GetOwnerContentType returns the OwnerContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedConfigContextSchemaRequest) GetOwnerContentType() string {
	if o == nil || IsNil(o.OwnerContentType.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerContentType.Get()
}

// GetOwnerContentTypeOk returns a tuple with the OwnerContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedConfigContextSchemaRequest) GetOwnerContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerContentType.Get(), o.OwnerContentType.IsSet()
}

// HasOwnerContentType returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasOwnerContentType() bool {
	if o != nil && o.OwnerContentType.IsSet() {
		return true
	}

	return false
}

// SetOwnerContentType gets a reference to the given NullableString and assigns it to the OwnerContentType field.
func (o *PatchedConfigContextSchemaRequest) SetOwnerContentType(v string) {
	o.OwnerContentType.Set(&v)
}
// SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil
func (o *PatchedConfigContextSchemaRequest) SetOwnerContentTypeNil() {
	o.OwnerContentType.Set(nil)
}

// UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
func (o *PatchedConfigContextSchemaRequest) UnsetOwnerContentType() {
	o.OwnerContentType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedConfigContextSchemaRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigContextSchemaRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedConfigContextSchemaRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedConfigContextSchemaRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigContextSchemaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedConfigContextSchemaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDataSchema returns the DataSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedConfigContextSchemaRequest) GetDataSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DataSchema
}

// GetDataSchemaOk returns a tuple with the DataSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedConfigContextSchemaRequest) GetDataSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DataSchema) {
		return nil, false
	}
	return &o.DataSchema, true
}

// HasDataSchema returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasDataSchema() bool {
	if o != nil && !IsNil(o.DataSchema) {
		return true
	}

	return false
}

// SetDataSchema gets a reference to the given interface{} and assigns it to the DataSchema field.
func (o *PatchedConfigContextSchemaRequest) SetDataSchema(v interface{}) {
	o.DataSchema = v
}

// GetOwnerObjectId returns the OwnerObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedConfigContextSchemaRequest) GetOwnerObjectId() string {
	if o == nil || IsNil(o.OwnerObjectId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerObjectId.Get()
}

// GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedConfigContextSchemaRequest) GetOwnerObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerObjectId.Get(), o.OwnerObjectId.IsSet()
}

// HasOwnerObjectId returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasOwnerObjectId() bool {
	if o != nil && o.OwnerObjectId.IsSet() {
		return true
	}

	return false
}

// SetOwnerObjectId gets a reference to the given NullableString and assigns it to the OwnerObjectId field.
func (o *PatchedConfigContextSchemaRequest) SetOwnerObjectId(v string) {
	o.OwnerObjectId.Set(&v)
}
// SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil
func (o *PatchedConfigContextSchemaRequest) SetOwnerObjectIdNil() {
	o.OwnerObjectId.Set(nil)
}

// UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
func (o *PatchedConfigContextSchemaRequest) UnsetOwnerObjectId() {
	o.OwnerObjectId.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedConfigContextSchemaRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigContextSchemaRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedConfigContextSchemaRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedConfigContextSchemaRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigContextSchemaRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedConfigContextSchemaRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedConfigContextSchemaRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedConfigContextSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedConfigContextSchemaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerContentType.IsSet() {
		toSerialize["owner_content_type"] = o.OwnerContentType.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.DataSchema != nil {
		toSerialize["data_schema"] = o.DataSchema
	}
	if o.OwnerObjectId.IsSet() {
		toSerialize["owner_object_id"] = o.OwnerObjectId.Get()
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

func (o *PatchedConfigContextSchemaRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedConfigContextSchemaRequest := _PatchedConfigContextSchemaRequest{}

	err = json.Unmarshal(data, &varPatchedConfigContextSchemaRequest)

	if err != nil {
		return err
	}

	*o = PatchedConfigContextSchemaRequest(varPatchedConfigContextSchemaRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "owner_content_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "data_schema")
		delete(additionalProperties, "owner_object_id")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedConfigContextSchemaRequest struct {
	value *PatchedConfigContextSchemaRequest
	isSet bool
}

func (v NullablePatchedConfigContextSchemaRequest) Get() *PatchedConfigContextSchemaRequest {
	return v.value
}

func (v *NullablePatchedConfigContextSchemaRequest) Set(val *PatchedConfigContextSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedConfigContextSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedConfigContextSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedConfigContextSchemaRequest(val *PatchedConfigContextSchemaRequest) *NullablePatchedConfigContextSchemaRequest {
	return &NullablePatchedConfigContextSchemaRequest{value: val, isSet: true}
}

func (v NullablePatchedConfigContextSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedConfigContextSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


