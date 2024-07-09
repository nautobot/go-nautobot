/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedWritableFrontPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableFrontPortRequest{}

// PatchedWritableFrontPortRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedWritableFrontPortRequest struct {
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *PortTypeChoices `json:"type,omitempty"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	Device *BulkWritableCableRequestStatus `json:"device,omitempty"`
	RearPort *BulkWritableCableRequestStatus `json:"rear_port,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableFrontPortRequest PatchedWritableFrontPortRequest

// NewPatchedWritableFrontPortRequest instantiates a new PatchedWritableFrontPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableFrontPortRequest() *PatchedWritableFrontPortRequest {
	this := PatchedWritableFrontPortRequest{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// NewPatchedWritableFrontPortRequestWithDefaults instantiates a new PatchedWritableFrontPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableFrontPortRequestWithDefaults() *PatchedWritableFrontPortRequest {
	this := PatchedWritableFrontPortRequest{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableFrontPortRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableFrontPortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableFrontPortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetType() PortTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PortTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetTypeOk() (*PortTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PortTypeChoices and assigns it to the Type field.
func (o *PatchedWritableFrontPortRequest) SetType(v PortTypeChoices) {
	o.Type = &v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *PatchedWritableFrontPortRequest) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetDevice() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Device) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Device field.
func (o *PatchedWritableFrontPortRequest) SetDevice(v BulkWritableCableRequestStatus) {
	o.Device = &v
}

// GetRearPort returns the RearPort field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetRearPort() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.RearPort) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetRearPortOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.RearPort) {
		return nil, false
	}
	return o.RearPort, true
}

// HasRearPort returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasRearPort() bool {
	if o != nil && !IsNil(o.RearPort) {
		return true
	}

	return false
}

// SetRearPort gets a reference to the given BulkWritableCableRequestStatus and assigns it to the RearPort field.
func (o *PatchedWritableFrontPortRequest) SetRearPort(v BulkWritableCableRequestStatus) {
	o.RearPort = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedWritableFrontPortRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableFrontPortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedWritableFrontPortRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedWritableFrontPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableFrontPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RearPortPosition) {
		toSerialize["rear_port_position"] = o.RearPortPosition
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.RearPort) {
		toSerialize["rear_port"] = o.RearPort
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

func (o *PatchedWritableFrontPortRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableFrontPortRequest := _PatchedWritableFrontPortRequest{}

	err = json.Unmarshal(data, &varPatchedWritableFrontPortRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableFrontPortRequest(varPatchedWritableFrontPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "device")
		delete(additionalProperties, "rear_port")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableFrontPortRequest struct {
	value *PatchedWritableFrontPortRequest
	isSet bool
}

func (v NullablePatchedWritableFrontPortRequest) Get() *PatchedWritableFrontPortRequest {
	return v.value
}

func (v *NullablePatchedWritableFrontPortRequest) Set(val *PatchedWritableFrontPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableFrontPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableFrontPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableFrontPortRequest(val *PatchedWritableFrontPortRequest) *NullablePatchedWritableFrontPortRequest {
	return &NullablePatchedWritableFrontPortRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableFrontPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableFrontPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

