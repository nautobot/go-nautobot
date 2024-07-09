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

// checks if the BulkWritableDynamicGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableDynamicGroupRequest{}

// BulkWritableDynamicGroupRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableDynamicGroupRequest struct {
	Id string `json:"id"`
	ContentType string `json:"content_type"`
	// Dynamic Group name
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// A JSON-encoded dictionary of filter parameters for group membership
	Filter map[string]interface{} `json:"filter"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableDynamicGroupRequest BulkWritableDynamicGroupRequest

// NewBulkWritableDynamicGroupRequest instantiates a new BulkWritableDynamicGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableDynamicGroupRequest(id string, contentType string, name string, filter map[string]interface{}) *BulkWritableDynamicGroupRequest {
	this := BulkWritableDynamicGroupRequest{}
	this.Id = id
	this.ContentType = contentType
	this.Name = name
	this.Filter = filter
	return &this
}

// NewBulkWritableDynamicGroupRequestWithDefaults instantiates a new BulkWritableDynamicGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableDynamicGroupRequestWithDefaults() *BulkWritableDynamicGroupRequest {
	this := BulkWritableDynamicGroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableDynamicGroupRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableDynamicGroupRequest) SetId(v string) {
	o.Id = v
}

// GetContentType returns the ContentType field value
func (o *BulkWritableDynamicGroupRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *BulkWritableDynamicGroupRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetName returns the Name field value
func (o *BulkWritableDynamicGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableDynamicGroupRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableDynamicGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableDynamicGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableDynamicGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFilter returns the Filter field value
func (o *BulkWritableDynamicGroupRequest) GetFilter() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// SetFilter sets field value
func (o *BulkWritableDynamicGroupRequest) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableDynamicGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableDynamicGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableDynamicGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableDynamicGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDynamicGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableDynamicGroupRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableDynamicGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableDynamicGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableDynamicGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["content_type"] = o.ContentType
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["filter"] = o.Filter
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

func (o *BulkWritableDynamicGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"content_type",
		"name",
		"filter",
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

	varBulkWritableDynamicGroupRequest := _BulkWritableDynamicGroupRequest{}

	err = json.Unmarshal(data, &varBulkWritableDynamicGroupRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableDynamicGroupRequest(varBulkWritableDynamicGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableDynamicGroupRequest struct {
	value *BulkWritableDynamicGroupRequest
	isSet bool
}

func (v NullableBulkWritableDynamicGroupRequest) Get() *BulkWritableDynamicGroupRequest {
	return v.value
}

func (v *NullableBulkWritableDynamicGroupRequest) Set(val *BulkWritableDynamicGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableDynamicGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableDynamicGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableDynamicGroupRequest(val *BulkWritableDynamicGroupRequest) *NullableBulkWritableDynamicGroupRequest {
	return &NullableBulkWritableDynamicGroupRequest{value: val, isSet: true}
}

func (v NullableBulkWritableDynamicGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableDynamicGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

