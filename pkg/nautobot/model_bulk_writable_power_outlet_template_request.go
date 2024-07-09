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

// checks if the BulkWritablePowerOutletTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritablePowerOutletTemplateRequest{}

// BulkWritablePowerOutletTemplateRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritablePowerOutletTemplateRequest struct {
	Id string `json:"id"`
	Type *PowerOutletTypeChoices `json:"type,omitempty"`
	FeedLeg *FeedLegEnum `json:"feed_leg,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	DeviceType BulkWritableCableRequestStatus `json:"device_type"`
	PowerPortTemplate NullableBulkWritableCircuitRequestTenant `json:"power_port_template,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritablePowerOutletTemplateRequest BulkWritablePowerOutletTemplateRequest

// NewBulkWritablePowerOutletTemplateRequest instantiates a new BulkWritablePowerOutletTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritablePowerOutletTemplateRequest(id string, name string, deviceType BulkWritableCableRequestStatus) *BulkWritablePowerOutletTemplateRequest {
	this := BulkWritablePowerOutletTemplateRequest{}
	this.Id = id
	this.Name = name
	this.DeviceType = deviceType
	return &this
}

// NewBulkWritablePowerOutletTemplateRequestWithDefaults instantiates a new BulkWritablePowerOutletTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritablePowerOutletTemplateRequestWithDefaults() *BulkWritablePowerOutletTemplateRequest {
	this := BulkWritablePowerOutletTemplateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritablePowerOutletTemplateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritablePowerOutletTemplateRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BulkWritablePowerOutletTemplateRequest) GetType() PowerOutletTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PowerOutletTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetTypeOk() (*PowerOutletTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerOutletTypeChoices and assigns it to the Type field.
func (o *BulkWritablePowerOutletTemplateRequest) SetType(v PowerOutletTypeChoices) {
	o.Type = &v
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *BulkWritablePowerOutletTemplateRequest) GetFeedLeg() FeedLegEnum {
	if o == nil || IsNil(o.FeedLeg) {
		var ret FeedLegEnum
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetFeedLegOk() (*FeedLegEnum, bool) {
	if o == nil || IsNil(o.FeedLeg) {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasFeedLeg() bool {
	if o != nil && !IsNil(o.FeedLeg) {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given FeedLegEnum and assigns it to the FeedLeg field.
func (o *BulkWritablePowerOutletTemplateRequest) SetFeedLeg(v FeedLegEnum) {
	o.FeedLeg = &v
}

// GetName returns the Name field value
func (o *BulkWritablePowerOutletTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritablePowerOutletTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritablePowerOutletTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BulkWritablePowerOutletTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritablePowerOutletTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritablePowerOutletTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceType returns the DeviceType field value
func (o *BulkWritablePowerOutletTemplateRequest) GetDeviceType() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *BulkWritablePowerOutletTemplateRequest) SetDeviceType(v BulkWritableCableRequestStatus) {
	o.DeviceType = v
}

// GetPowerPortTemplate returns the PowerPortTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritablePowerOutletTemplateRequest) GetPowerPortTemplate() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.PowerPortTemplate.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.PowerPortTemplate.Get()
}

// GetPowerPortTemplateOk returns a tuple with the PowerPortTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritablePowerOutletTemplateRequest) GetPowerPortTemplateOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPortTemplate.Get(), o.PowerPortTemplate.IsSet()
}

// HasPowerPortTemplate returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasPowerPortTemplate() bool {
	if o != nil && o.PowerPortTemplate.IsSet() {
		return true
	}

	return false
}

// SetPowerPortTemplate gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the PowerPortTemplate field.
func (o *BulkWritablePowerOutletTemplateRequest) SetPowerPortTemplate(v BulkWritableCircuitRequestTenant) {
	o.PowerPortTemplate.Set(&v)
}
// SetPowerPortTemplateNil sets the value for PowerPortTemplate to be an explicit nil
func (o *BulkWritablePowerOutletTemplateRequest) SetPowerPortTemplateNil() {
	o.PowerPortTemplate.Set(nil)
}

// UnsetPowerPortTemplate ensures that no value is present for PowerPortTemplate, not even an explicit nil
func (o *BulkWritablePowerOutletTemplateRequest) UnsetPowerPortTemplate() {
	o.PowerPortTemplate.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritablePowerOutletTemplateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritablePowerOutletTemplateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritablePowerOutletTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerOutletTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritablePowerOutletTemplateRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritablePowerOutletTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritablePowerOutletTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritablePowerOutletTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FeedLeg) {
		toSerialize["feed_leg"] = o.FeedLeg
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["device_type"] = o.DeviceType
	if o.PowerPortTemplate.IsSet() {
		toSerialize["power_port_template"] = o.PowerPortTemplate.Get()
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

func (o *BulkWritablePowerOutletTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"device_type",
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

	varBulkWritablePowerOutletTemplateRequest := _BulkWritablePowerOutletTemplateRequest{}

	err = json.Unmarshal(data, &varBulkWritablePowerOutletTemplateRequest)

	if err != nil {
		return err
	}

	*o = BulkWritablePowerOutletTemplateRequest(varBulkWritablePowerOutletTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "power_port_template")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritablePowerOutletTemplateRequest struct {
	value *BulkWritablePowerOutletTemplateRequest
	isSet bool
}

func (v NullableBulkWritablePowerOutletTemplateRequest) Get() *BulkWritablePowerOutletTemplateRequest {
	return v.value
}

func (v *NullableBulkWritablePowerOutletTemplateRequest) Set(val *BulkWritablePowerOutletTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritablePowerOutletTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritablePowerOutletTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritablePowerOutletTemplateRequest(val *BulkWritablePowerOutletTemplateRequest) *NullableBulkWritablePowerOutletTemplateRequest {
	return &NullableBulkWritablePowerOutletTemplateRequest{value: val, isSet: true}
}

func (v NullableBulkWritablePowerOutletTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritablePowerOutletTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

