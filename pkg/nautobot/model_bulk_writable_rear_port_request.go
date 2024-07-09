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

// checks if the BulkWritableRearPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableRearPortRequest{}

// BulkWritableRearPortRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableRearPortRequest struct {
	Id string `json:"id"`
	Type PortTypeChoices `json:"type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Positions *int32 `json:"positions,omitempty"`
	Device BulkWritableCableRequestStatus `json:"device"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableRearPortRequest BulkWritableRearPortRequest

// NewBulkWritableRearPortRequest instantiates a new BulkWritableRearPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableRearPortRequest(id string, type_ PortTypeChoices, name string, device BulkWritableCableRequestStatus) *BulkWritableRearPortRequest {
	this := BulkWritableRearPortRequest{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Device = device
	return &this
}

// NewBulkWritableRearPortRequestWithDefaults instantiates a new BulkWritableRearPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableRearPortRequestWithDefaults() *BulkWritableRearPortRequest {
	this := BulkWritableRearPortRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableRearPortRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableRearPortRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *BulkWritableRearPortRequest) GetType() PortTypeChoices {
	if o == nil {
		var ret PortTypeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetTypeOk() (*PortTypeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BulkWritableRearPortRequest) SetType(v PortTypeChoices) {
	o.Type = v
}

// GetName returns the Name field value
func (o *BulkWritableRearPortRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableRearPortRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritableRearPortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritableRearPortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BulkWritableRearPortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableRearPortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableRearPortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableRearPortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *BulkWritableRearPortRequest) GetPositions() int32 {
	if o == nil || IsNil(o.Positions) {
		var ret int32
		return ret
	}
	return *o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetPositionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *BulkWritableRearPortRequest) HasPositions() bool {
	if o != nil && !IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given int32 and assigns it to the Positions field.
func (o *BulkWritableRearPortRequest) SetPositions(v int32) {
	o.Positions = &v
}

// GetDevice returns the Device field value
func (o *BulkWritableRearPortRequest) GetDevice() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *BulkWritableRearPortRequest) SetDevice(v BulkWritableCableRequestStatus) {
	o.Device = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableRearPortRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableRearPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableRearPortRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableRearPortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableRearPortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableRearPortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableRearPortRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRearPortRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableRearPortRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableRearPortRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableRearPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableRearPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	toSerialize["device"] = o.Device
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

func (o *BulkWritableRearPortRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"name",
		"device",
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

	varBulkWritableRearPortRequest := _BulkWritableRearPortRequest{}

	err = json.Unmarshal(data, &varBulkWritableRearPortRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableRearPortRequest(varBulkWritableRearPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "positions")
		delete(additionalProperties, "device")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableRearPortRequest struct {
	value *BulkWritableRearPortRequest
	isSet bool
}

func (v NullableBulkWritableRearPortRequest) Get() *BulkWritableRearPortRequest {
	return v.value
}

func (v *NullableBulkWritableRearPortRequest) Set(val *BulkWritableRearPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableRearPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableRearPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableRearPortRequest(val *BulkWritableRearPortRequest) *NullableBulkWritableRearPortRequest {
	return &NullableBulkWritableRearPortRequest{value: val, isSet: true}
}

func (v NullableBulkWritableRearPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableRearPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

