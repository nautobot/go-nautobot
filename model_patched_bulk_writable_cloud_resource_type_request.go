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

// checks if the PatchedBulkWritableCloudResourceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableCloudResourceTypeRequest{}

// PatchedBulkWritableCloudResourceTypeRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableCloudResourceTypeRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types,omitempty"`
	// Type of cloud objects
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ConfigSchema interface{} `json:"config_schema,omitempty"`
	Provider *BulkWritableCloudAccountRequestProvider `json:"provider,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableCloudResourceTypeRequest PatchedBulkWritableCloudResourceTypeRequest

// NewPatchedBulkWritableCloudResourceTypeRequest instantiates a new PatchedBulkWritableCloudResourceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableCloudResourceTypeRequest(id string) *PatchedBulkWritableCloudResourceTypeRequest {
	this := PatchedBulkWritableCloudResourceTypeRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableCloudResourceTypeRequestWithDefaults instantiates a new PatchedBulkWritableCloudResourceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableCloudResourceTypeRequestWithDefaults() *PatchedBulkWritableCloudResourceTypeRequest {
	this := PatchedBulkWritableCloudResourceTypeRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetId(v string) {
	o.Id = v
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetConfigSchema returns the ConfigSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetConfigSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ConfigSchema
}

// GetConfigSchemaOk returns a tuple with the ConfigSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetConfigSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ConfigSchema) {
		return nil, false
	}
	return &o.ConfigSchema, true
}

// HasConfigSchema returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasConfigSchema() bool {
	if o != nil && !IsNil(o.ConfigSchema) {
		return true
	}

	return false
}

// SetConfigSchema gets a reference to the given interface{} and assigns it to the ConfigSchema field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetConfigSchema(v interface{}) {
	o.ConfigSchema = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetProvider() BulkWritableCloudAccountRequestProvider {
	if o == nil || IsNil(o.Provider) {
		var ret BulkWritableCloudAccountRequestProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given BulkWritableCloudAccountRequestProvider and assigns it to the Provider field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetProvider(v BulkWritableCloudAccountRequestProvider) {
	o.Provider = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudResourceTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableCloudResourceTypeRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableCloudResourceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableCloudResourceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ContentTypes) {
		toSerialize["content_types"] = o.ContentTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ConfigSchema != nil {
		toSerialize["config_schema"] = o.ConfigSchema
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
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

func (o *PatchedBulkWritableCloudResourceTypeRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableCloudResourceTypeRequest := _PatchedBulkWritableCloudResourceTypeRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableCloudResourceTypeRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableCloudResourceTypeRequest(varPatchedBulkWritableCloudResourceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "config_schema")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableCloudResourceTypeRequest struct {
	value *PatchedBulkWritableCloudResourceTypeRequest
	isSet bool
}

func (v NullablePatchedBulkWritableCloudResourceTypeRequest) Get() *PatchedBulkWritableCloudResourceTypeRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableCloudResourceTypeRequest) Set(val *PatchedBulkWritableCloudResourceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableCloudResourceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableCloudResourceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableCloudResourceTypeRequest(val *PatchedBulkWritableCloudResourceTypeRequest) *NullablePatchedBulkWritableCloudResourceTypeRequest {
	return &NullablePatchedBulkWritableCloudResourceTypeRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableCloudResourceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableCloudResourceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


