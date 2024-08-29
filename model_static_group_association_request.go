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

// checks if the StaticGroupAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticGroupAssociationRequest{}

// StaticGroupAssociationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type StaticGroupAssociationRequest struct {
	AssociatedObjectType string `json:"associated_object_type"`
	AssociatedObjectId string `json:"associated_object_id"`
	DynamicGroup BulkWritableCableRequestStatus `json:"dynamic_group"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StaticGroupAssociationRequest StaticGroupAssociationRequest

// NewStaticGroupAssociationRequest instantiates a new StaticGroupAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticGroupAssociationRequest(associatedObjectType string, associatedObjectId string, dynamicGroup BulkWritableCableRequestStatus) *StaticGroupAssociationRequest {
	this := StaticGroupAssociationRequest{}
	this.AssociatedObjectType = associatedObjectType
	this.AssociatedObjectId = associatedObjectId
	this.DynamicGroup = dynamicGroup
	return &this
}

// NewStaticGroupAssociationRequestWithDefaults instantiates a new StaticGroupAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticGroupAssociationRequestWithDefaults() *StaticGroupAssociationRequest {
	this := StaticGroupAssociationRequest{}
	return &this
}

// GetAssociatedObjectType returns the AssociatedObjectType field value
func (o *StaticGroupAssociationRequest) GetAssociatedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociatedObjectType
}

// GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field value
// and a boolean to check if the value has been set.
func (o *StaticGroupAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociatedObjectType, true
}

// SetAssociatedObjectType sets field value
func (o *StaticGroupAssociationRequest) SetAssociatedObjectType(v string) {
	o.AssociatedObjectType = v
}

// GetAssociatedObjectId returns the AssociatedObjectId field value
func (o *StaticGroupAssociationRequest) GetAssociatedObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociatedObjectId
}

// GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field value
// and a boolean to check if the value has been set.
func (o *StaticGroupAssociationRequest) GetAssociatedObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociatedObjectId, true
}

// SetAssociatedObjectId sets field value
func (o *StaticGroupAssociationRequest) SetAssociatedObjectId(v string) {
	o.AssociatedObjectId = v
}

// GetDynamicGroup returns the DynamicGroup field value
func (o *StaticGroupAssociationRequest) GetDynamicGroup() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.DynamicGroup
}

// GetDynamicGroupOk returns a tuple with the DynamicGroup field value
// and a boolean to check if the value has been set.
func (o *StaticGroupAssociationRequest) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DynamicGroup, true
}

// SetDynamicGroup sets field value
func (o *StaticGroupAssociationRequest) SetDynamicGroup(v BulkWritableCableRequestStatus) {
	o.DynamicGroup = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *StaticGroupAssociationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticGroupAssociationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *StaticGroupAssociationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *StaticGroupAssociationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StaticGroupAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticGroupAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StaticGroupAssociationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *StaticGroupAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o StaticGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticGroupAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associated_object_type"] = o.AssociatedObjectType
	toSerialize["associated_object_id"] = o.AssociatedObjectId
	toSerialize["dynamic_group"] = o.DynamicGroup
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

func (o *StaticGroupAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"associated_object_type",
		"associated_object_id",
		"dynamic_group",
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

	varStaticGroupAssociationRequest := _StaticGroupAssociationRequest{}

	err = json.Unmarshal(data, &varStaticGroupAssociationRequest)

	if err != nil {
		return err
	}

	*o = StaticGroupAssociationRequest(varStaticGroupAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associated_object_type")
		delete(additionalProperties, "associated_object_id")
		delete(additionalProperties, "dynamic_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStaticGroupAssociationRequest struct {
	value *StaticGroupAssociationRequest
	isSet bool
}

func (v NullableStaticGroupAssociationRequest) Get() *StaticGroupAssociationRequest {
	return v.value
}

func (v *NullableStaticGroupAssociationRequest) Set(val *StaticGroupAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticGroupAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticGroupAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticGroupAssociationRequest(val *StaticGroupAssociationRequest) *NullableStaticGroupAssociationRequest {
	return &NullableStaticGroupAssociationRequest{value: val, isSet: true}
}

func (v NullableStaticGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticGroupAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

