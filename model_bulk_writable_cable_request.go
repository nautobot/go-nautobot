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

// checks if the BulkWritableCableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableCableRequest{}

// BulkWritableCableRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableCableRequest struct {
	Id string `json:"id"`
	TerminationAType string `json:"termination_a_type"`
	TerminationBType string `json:"termination_b_type"`
	LengthUnit *LengthUnitEnum `json:"length_unit,omitempty"`
	Type *CableTypeChoices `json:"type,omitempty"`
	TerminationAId string `json:"termination_a_id"`
	TerminationBId string `json:"termination_b_id"`
	Label *string `json:"label,omitempty"`
	// RGB color in hexadecimal (e.g. 00ff00)
	Color *string `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	Length NullableInt32 `json:"length,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableCableRequest BulkWritableCableRequest

// NewBulkWritableCableRequest instantiates a new BulkWritableCableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableCableRequest(id string, terminationAType string, terminationBType string, terminationAId string, terminationBId string, status BulkWritableCableRequestStatus) *BulkWritableCableRequest {
	this := BulkWritableCableRequest{}
	this.Id = id
	this.TerminationAType = terminationAType
	this.TerminationBType = terminationBType
	this.TerminationAId = terminationAId
	this.TerminationBId = terminationBId
	this.Status = status
	return &this
}

// NewBulkWritableCableRequestWithDefaults instantiates a new BulkWritableCableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableCableRequestWithDefaults() *BulkWritableCableRequest {
	this := BulkWritableCableRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableCableRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableCableRequest) SetId(v string) {
	o.Id = v
}

// GetTerminationAType returns the TerminationAType field value
func (o *BulkWritableCableRequest) GetTerminationAType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationAType
}

// GetTerminationATypeOk returns a tuple with the TerminationAType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetTerminationATypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationAType, true
}

// SetTerminationAType sets field value
func (o *BulkWritableCableRequest) SetTerminationAType(v string) {
	o.TerminationAType = v
}

// GetTerminationBType returns the TerminationBType field value
func (o *BulkWritableCableRequest) GetTerminationBType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationBType
}

// GetTerminationBTypeOk returns a tuple with the TerminationBType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetTerminationBTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationBType, true
}

// SetTerminationBType sets field value
func (o *BulkWritableCableRequest) SetTerminationBType(v string) {
	o.TerminationBType = v
}

// GetLengthUnit returns the LengthUnit field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetLengthUnit() LengthUnitEnum {
	if o == nil || IsNil(o.LengthUnit) {
		var ret LengthUnitEnum
		return ret
	}
	return *o.LengthUnit
}

// GetLengthUnitOk returns a tuple with the LengthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetLengthUnitOk() (*LengthUnitEnum, bool) {
	if o == nil || IsNil(o.LengthUnit) {
		return nil, false
	}
	return o.LengthUnit, true
}

// HasLengthUnit returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasLengthUnit() bool {
	if o != nil && !IsNil(o.LengthUnit) {
		return true
	}

	return false
}

// SetLengthUnit gets a reference to the given LengthUnitEnum and assigns it to the LengthUnit field.
func (o *BulkWritableCableRequest) SetLengthUnit(v LengthUnitEnum) {
	o.LengthUnit = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetType() CableTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret CableTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetTypeOk() (*CableTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CableTypeChoices and assigns it to the Type field.
func (o *BulkWritableCableRequest) SetType(v CableTypeChoices) {
	o.Type = &v
}

// GetTerminationAId returns the TerminationAId field value
func (o *BulkWritableCableRequest) GetTerminationAId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationAId
}

// GetTerminationAIdOk returns a tuple with the TerminationAId field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetTerminationAIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationAId, true
}

// SetTerminationAId sets field value
func (o *BulkWritableCableRequest) SetTerminationAId(v string) {
	o.TerminationAId = v
}

// GetTerminationBId returns the TerminationBId field value
func (o *BulkWritableCableRequest) GetTerminationBId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationBId
}

// GetTerminationBIdOk returns a tuple with the TerminationBId field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetTerminationBIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationBId, true
}

// SetTerminationBId sets field value
func (o *BulkWritableCableRequest) SetTerminationBId(v string) {
	o.TerminationBId = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BulkWritableCableRequest) SetLabel(v string) {
	o.Label = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *BulkWritableCableRequest) SetColor(v string) {
	o.Color = &v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableCableRequest) GetLength() int32 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret int32
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableCableRequest) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableInt32 and assigns it to the Length field.
func (o *BulkWritableCableRequest) SetLength(v int32) {
	o.Length.Set(&v)
}
// SetLengthNil sets the value for Length to be an explicit nil
func (o *BulkWritableCableRequest) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *BulkWritableCableRequest) UnsetLength() {
	o.Length.Unset()
}

// GetStatus returns the Status field value
func (o *BulkWritableCableRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableCableRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableCableRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableCableRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableCableRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCableRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableCableRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableCableRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableCableRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableCableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["termination_a_type"] = o.TerminationAType
	toSerialize["termination_b_type"] = o.TerminationBType
	if !IsNil(o.LengthUnit) {
		toSerialize["length_unit"] = o.LengthUnit
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["termination_a_id"] = o.TerminationAId
	toSerialize["termination_b_id"] = o.TerminationBId
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	toSerialize["status"] = o.Status
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

func (o *BulkWritableCableRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"termination_a_type",
		"termination_b_type",
		"termination_a_id",
		"termination_b_id",
		"status",
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

	varBulkWritableCableRequest := _BulkWritableCableRequest{}

	err = json.Unmarshal(data, &varBulkWritableCableRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableCableRequest(varBulkWritableCableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "termination_a_type")
		delete(additionalProperties, "termination_b_type")
		delete(additionalProperties, "length_unit")
		delete(additionalProperties, "type")
		delete(additionalProperties, "termination_a_id")
		delete(additionalProperties, "termination_b_id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "color")
		delete(additionalProperties, "length")
		delete(additionalProperties, "status")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableCableRequest struct {
	value *BulkWritableCableRequest
	isSet bool
}

func (v NullableBulkWritableCableRequest) Get() *BulkWritableCableRequest {
	return v.value
}

func (v *NullableBulkWritableCableRequest) Set(val *BulkWritableCableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableCableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableCableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableCableRequest(val *BulkWritableCableRequest) *NullableBulkWritableCableRequest {
	return &NullableBulkWritableCableRequest{value: val, isSet: true}
}

func (v NullableBulkWritableCableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableCableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

