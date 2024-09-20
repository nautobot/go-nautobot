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

// checks if the BulkWritableRelationshipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableRelationshipRequest{}

// BulkWritableRelationshipRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableRelationshipRequest struct {
	Id string `json:"id"`
	SourceType string `json:"source_type"`
	DestinationType string `json:"destination_type"`
	// Label of the relationship as displayed to users
	Label string `json:"label"`
	// Internal relationship key. Please use underscores rather than dashes in this key.
	Key *string `json:"key,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description *string `json:"description,omitempty"`
	// Cardinality of this relationship
	Type *RelationshipTypeChoices `json:"type,omitempty"`
	RequiredOn *BulkWritableRelationshipRequestRequiredOn `json:"required_on,omitempty"`
	// Label for related destination objects, as displayed on the source object.
	SourceLabel *string `json:"source_label,omitempty"`
	// Hide this relationship on the source object.
	SourceHidden *bool `json:"source_hidden,omitempty"`
	// Filterset filter matching the applicable source objects of the selected type
	SourceFilter interface{} `json:"source_filter,omitempty"`
	// Label for related source objects, as displayed on the destination object.
	DestinationLabel *string `json:"destination_label,omitempty"`
	// Hide this relationship on the destination object.
	DestinationHidden *bool `json:"destination_hidden,omitempty"`
	// Filterset filter matching the applicable destination objects of the selected type
	DestinationFilter interface{} `json:"destination_filter,omitempty"`
	// Hide this field from the object's primary information tab. It will appear in the \"Advanced\" tab instead.
	AdvancedUi *bool `json:"advanced_ui,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableRelationshipRequest BulkWritableRelationshipRequest

// NewBulkWritableRelationshipRequest instantiates a new BulkWritableRelationshipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableRelationshipRequest(id string, sourceType string, destinationType string, label string) *BulkWritableRelationshipRequest {
	this := BulkWritableRelationshipRequest{}
	this.Id = id
	this.SourceType = sourceType
	this.DestinationType = destinationType
	this.Label = label
	return &this
}

// NewBulkWritableRelationshipRequestWithDefaults instantiates a new BulkWritableRelationshipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableRelationshipRequestWithDefaults() *BulkWritableRelationshipRequest {
	this := BulkWritableRelationshipRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableRelationshipRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableRelationshipRequest) SetId(v string) {
	o.Id = v
}

// GetSourceType returns the SourceType field value
func (o *BulkWritableRelationshipRequest) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *BulkWritableRelationshipRequest) SetSourceType(v string) {
	o.SourceType = v
}

// GetDestinationType returns the DestinationType field value
func (o *BulkWritableRelationshipRequest) GetDestinationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetDestinationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *BulkWritableRelationshipRequest) SetDestinationType(v string) {
	o.DestinationType = v
}

// GetLabel returns the Label field value
func (o *BulkWritableRelationshipRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *BulkWritableRelationshipRequest) SetLabel(v string) {
	o.Label = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BulkWritableRelationshipRequest) SetKey(v string) {
	o.Key = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableRelationshipRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetType() RelationshipTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret RelationshipTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetTypeOk() (*RelationshipTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RelationshipTypeChoices and assigns it to the Type field.
func (o *BulkWritableRelationshipRequest) SetType(v RelationshipTypeChoices) {
	o.Type = &v
}

// GetRequiredOn returns the RequiredOn field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetRequiredOn() BulkWritableRelationshipRequestRequiredOn {
	if o == nil || IsNil(o.RequiredOn) {
		var ret BulkWritableRelationshipRequestRequiredOn
		return ret
	}
	return *o.RequiredOn
}

// GetRequiredOnOk returns a tuple with the RequiredOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetRequiredOnOk() (*BulkWritableRelationshipRequestRequiredOn, bool) {
	if o == nil || IsNil(o.RequiredOn) {
		return nil, false
	}
	return o.RequiredOn, true
}

// HasRequiredOn returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasRequiredOn() bool {
	if o != nil && !IsNil(o.RequiredOn) {
		return true
	}

	return false
}

// SetRequiredOn gets a reference to the given BulkWritableRelationshipRequestRequiredOn and assigns it to the RequiredOn field.
func (o *BulkWritableRelationshipRequest) SetRequiredOn(v BulkWritableRelationshipRequestRequiredOn) {
	o.RequiredOn = &v
}

// GetSourceLabel returns the SourceLabel field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetSourceLabel() string {
	if o == nil || IsNil(o.SourceLabel) {
		var ret string
		return ret
	}
	return *o.SourceLabel
}

// GetSourceLabelOk returns a tuple with the SourceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetSourceLabelOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLabel) {
		return nil, false
	}
	return o.SourceLabel, true
}

// HasSourceLabel returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasSourceLabel() bool {
	if o != nil && !IsNil(o.SourceLabel) {
		return true
	}

	return false
}

// SetSourceLabel gets a reference to the given string and assigns it to the SourceLabel field.
func (o *BulkWritableRelationshipRequest) SetSourceLabel(v string) {
	o.SourceLabel = &v
}

// GetSourceHidden returns the SourceHidden field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetSourceHidden() bool {
	if o == nil || IsNil(o.SourceHidden) {
		var ret bool
		return ret
	}
	return *o.SourceHidden
}

// GetSourceHiddenOk returns a tuple with the SourceHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetSourceHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.SourceHidden) {
		return nil, false
	}
	return o.SourceHidden, true
}

// HasSourceHidden returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasSourceHidden() bool {
	if o != nil && !IsNil(o.SourceHidden) {
		return true
	}

	return false
}

// SetSourceHidden gets a reference to the given bool and assigns it to the SourceHidden field.
func (o *BulkWritableRelationshipRequest) SetSourceHidden(v bool) {
	o.SourceHidden = &v
}

// GetSourceFilter returns the SourceFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRelationshipRequest) GetSourceFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SourceFilter
}

// GetSourceFilterOk returns a tuple with the SourceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRelationshipRequest) GetSourceFilterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SourceFilter) {
		return nil, false
	}
	return &o.SourceFilter, true
}

// HasSourceFilter returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasSourceFilter() bool {
	if o != nil && !IsNil(o.SourceFilter) {
		return true
	}

	return false
}

// SetSourceFilter gets a reference to the given interface{} and assigns it to the SourceFilter field.
func (o *BulkWritableRelationshipRequest) SetSourceFilter(v interface{}) {
	o.SourceFilter = v
}

// GetDestinationLabel returns the DestinationLabel field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetDestinationLabel() string {
	if o == nil || IsNil(o.DestinationLabel) {
		var ret string
		return ret
	}
	return *o.DestinationLabel
}

// GetDestinationLabelOk returns a tuple with the DestinationLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetDestinationLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationLabel) {
		return nil, false
	}
	return o.DestinationLabel, true
}

// HasDestinationLabel returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasDestinationLabel() bool {
	if o != nil && !IsNil(o.DestinationLabel) {
		return true
	}

	return false
}

// SetDestinationLabel gets a reference to the given string and assigns it to the DestinationLabel field.
func (o *BulkWritableRelationshipRequest) SetDestinationLabel(v string) {
	o.DestinationLabel = &v
}

// GetDestinationHidden returns the DestinationHidden field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetDestinationHidden() bool {
	if o == nil || IsNil(o.DestinationHidden) {
		var ret bool
		return ret
	}
	return *o.DestinationHidden
}

// GetDestinationHiddenOk returns a tuple with the DestinationHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetDestinationHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.DestinationHidden) {
		return nil, false
	}
	return o.DestinationHidden, true
}

// HasDestinationHidden returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasDestinationHidden() bool {
	if o != nil && !IsNil(o.DestinationHidden) {
		return true
	}

	return false
}

// SetDestinationHidden gets a reference to the given bool and assigns it to the DestinationHidden field.
func (o *BulkWritableRelationshipRequest) SetDestinationHidden(v bool) {
	o.DestinationHidden = &v
}

// GetDestinationFilter returns the DestinationFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRelationshipRequest) GetDestinationFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DestinationFilter
}

// GetDestinationFilterOk returns a tuple with the DestinationFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRelationshipRequest) GetDestinationFilterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DestinationFilter) {
		return nil, false
	}
	return &o.DestinationFilter, true
}

// HasDestinationFilter returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasDestinationFilter() bool {
	if o != nil && !IsNil(o.DestinationFilter) {
		return true
	}

	return false
}

// SetDestinationFilter gets a reference to the given interface{} and assigns it to the DestinationFilter field.
func (o *BulkWritableRelationshipRequest) SetDestinationFilter(v interface{}) {
	o.DestinationFilter = v
}

// GetAdvancedUi returns the AdvancedUi field value if set, zero value otherwise.
func (o *BulkWritableRelationshipRequest) GetAdvancedUi() bool {
	if o == nil || IsNil(o.AdvancedUi) {
		var ret bool
		return ret
	}
	return *o.AdvancedUi
}

// GetAdvancedUiOk returns a tuple with the AdvancedUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRelationshipRequest) GetAdvancedUiOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvancedUi) {
		return nil, false
	}
	return o.AdvancedUi, true
}

// HasAdvancedUi returns a boolean if a field has been set.
func (o *BulkWritableRelationshipRequest) HasAdvancedUi() bool {
	if o != nil && !IsNil(o.AdvancedUi) {
		return true
	}

	return false
}

// SetAdvancedUi gets a reference to the given bool and assigns it to the AdvancedUi field.
func (o *BulkWritableRelationshipRequest) SetAdvancedUi(v bool) {
	o.AdvancedUi = &v
}

func (o BulkWritableRelationshipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableRelationshipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["source_type"] = o.SourceType
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["label"] = o.Label
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RequiredOn) {
		toSerialize["required_on"] = o.RequiredOn
	}
	if !IsNil(o.SourceLabel) {
		toSerialize["source_label"] = o.SourceLabel
	}
	if !IsNil(o.SourceHidden) {
		toSerialize["source_hidden"] = o.SourceHidden
	}
	if o.SourceFilter != nil {
		toSerialize["source_filter"] = o.SourceFilter
	}
	if !IsNil(o.DestinationLabel) {
		toSerialize["destination_label"] = o.DestinationLabel
	}
	if !IsNil(o.DestinationHidden) {
		toSerialize["destination_hidden"] = o.DestinationHidden
	}
	if o.DestinationFilter != nil {
		toSerialize["destination_filter"] = o.DestinationFilter
	}
	if !IsNil(o.AdvancedUi) {
		toSerialize["advanced_ui"] = o.AdvancedUi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableRelationshipRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"source_type",
		"destination_type",
		"label",
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

	varBulkWritableRelationshipRequest := _BulkWritableRelationshipRequest{}

	err = json.Unmarshal(data, &varBulkWritableRelationshipRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableRelationshipRequest(varBulkWritableRelationshipRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "source_type")
		delete(additionalProperties, "destination_type")
		delete(additionalProperties, "label")
		delete(additionalProperties, "key")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "required_on")
		delete(additionalProperties, "source_label")
		delete(additionalProperties, "source_hidden")
		delete(additionalProperties, "source_filter")
		delete(additionalProperties, "destination_label")
		delete(additionalProperties, "destination_hidden")
		delete(additionalProperties, "destination_filter")
		delete(additionalProperties, "advanced_ui")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableRelationshipRequest struct {
	value *BulkWritableRelationshipRequest
	isSet bool
}

func (v NullableBulkWritableRelationshipRequest) Get() *BulkWritableRelationshipRequest {
	return v.value
}

func (v *NullableBulkWritableRelationshipRequest) Set(val *BulkWritableRelationshipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableRelationshipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableRelationshipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableRelationshipRequest(val *BulkWritableRelationshipRequest) *NullableBulkWritableRelationshipRequest {
	return &NullableBulkWritableRelationshipRequest{value: val, isSet: true}
}

func (v NullableBulkWritableRelationshipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableRelationshipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


