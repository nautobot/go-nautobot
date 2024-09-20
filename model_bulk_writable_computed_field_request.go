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

// checks if the BulkWritableComputedFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableComputedFieldRequest{}

// BulkWritableComputedFieldRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableComputedFieldRequest struct {
	Id string `json:"id"`
	ContentType string `json:"content_type"`
	// Internal field name. Please use underscores rather than dashes in this key.
	Key *string `json:"key,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// Human-readable grouping that this computed field belongs to.
	Grouping *string `json:"grouping,omitempty"`
	// Name of the field as displayed to users
	Label string `json:"label"`
	Description *string `json:"description,omitempty"`
	// Jinja2 template code for field value
	Template string `json:"template"`
	// Fallback value (if any) to be output for the field in the case of a template rendering error.
	FallbackValue *string `json:"fallback_value,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	// Hide this field from the object's primary information tab. It will appear in the \"Advanced\" tab instead.
	AdvancedUi *bool `json:"advanced_ui,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableComputedFieldRequest BulkWritableComputedFieldRequest

// NewBulkWritableComputedFieldRequest instantiates a new BulkWritableComputedFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableComputedFieldRequest(id string, contentType string, label string, template string) *BulkWritableComputedFieldRequest {
	this := BulkWritableComputedFieldRequest{}
	this.Id = id
	this.ContentType = contentType
	this.Label = label
	this.Template = template
	return &this
}

// NewBulkWritableComputedFieldRequestWithDefaults instantiates a new BulkWritableComputedFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableComputedFieldRequestWithDefaults() *BulkWritableComputedFieldRequest {
	this := BulkWritableComputedFieldRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableComputedFieldRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableComputedFieldRequest) SetId(v string) {
	o.Id = v
}

// GetContentType returns the ContentType field value
func (o *BulkWritableComputedFieldRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *BulkWritableComputedFieldRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BulkWritableComputedFieldRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BulkWritableComputedFieldRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BulkWritableComputedFieldRequest) SetKey(v string) {
	o.Key = &v
}

// GetGrouping returns the Grouping field value if set, zero value otherwise.
func (o *BulkWritableComputedFieldRequest) GetGrouping() string {
	if o == nil || IsNil(o.Grouping) {
		var ret string
		return ret
	}
	return *o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetGroupingOk() (*string, bool) {
	if o == nil || IsNil(o.Grouping) {
		return nil, false
	}
	return o.Grouping, true
}

// HasGrouping returns a boolean if a field has been set.
func (o *BulkWritableComputedFieldRequest) HasGrouping() bool {
	if o != nil && !IsNil(o.Grouping) {
		return true
	}

	return false
}

// SetGrouping gets a reference to the given string and assigns it to the Grouping field.
func (o *BulkWritableComputedFieldRequest) SetGrouping(v string) {
	o.Grouping = &v
}

// GetLabel returns the Label field value
func (o *BulkWritableComputedFieldRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *BulkWritableComputedFieldRequest) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableComputedFieldRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableComputedFieldRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableComputedFieldRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTemplate returns the Template field value
func (o *BulkWritableComputedFieldRequest) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *BulkWritableComputedFieldRequest) SetTemplate(v string) {
	o.Template = v
}

// GetFallbackValue returns the FallbackValue field value if set, zero value otherwise.
func (o *BulkWritableComputedFieldRequest) GetFallbackValue() string {
	if o == nil || IsNil(o.FallbackValue) {
		var ret string
		return ret
	}
	return *o.FallbackValue
}

// GetFallbackValueOk returns a tuple with the FallbackValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetFallbackValueOk() (*string, bool) {
	if o == nil || IsNil(o.FallbackValue) {
		return nil, false
	}
	return o.FallbackValue, true
}

// HasFallbackValue returns a boolean if a field has been set.
func (o *BulkWritableComputedFieldRequest) HasFallbackValue() bool {
	if o != nil && !IsNil(o.FallbackValue) {
		return true
	}

	return false
}

// SetFallbackValue gets a reference to the given string and assigns it to the FallbackValue field.
func (o *BulkWritableComputedFieldRequest) SetFallbackValue(v string) {
	o.FallbackValue = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *BulkWritableComputedFieldRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *BulkWritableComputedFieldRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *BulkWritableComputedFieldRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetAdvancedUi returns the AdvancedUi field value if set, zero value otherwise.
func (o *BulkWritableComputedFieldRequest) GetAdvancedUi() bool {
	if o == nil || IsNil(o.AdvancedUi) {
		var ret bool
		return ret
	}
	return *o.AdvancedUi
}

// GetAdvancedUiOk returns a tuple with the AdvancedUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableComputedFieldRequest) GetAdvancedUiOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvancedUi) {
		return nil, false
	}
	return o.AdvancedUi, true
}

// HasAdvancedUi returns a boolean if a field has been set.
func (o *BulkWritableComputedFieldRequest) HasAdvancedUi() bool {
	if o != nil && !IsNil(o.AdvancedUi) {
		return true
	}

	return false
}

// SetAdvancedUi gets a reference to the given bool and assigns it to the AdvancedUi field.
func (o *BulkWritableComputedFieldRequest) SetAdvancedUi(v bool) {
	o.AdvancedUi = &v
}

func (o BulkWritableComputedFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableComputedFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["content_type"] = o.ContentType
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Grouping) {
		toSerialize["grouping"] = o.Grouping
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["template"] = o.Template
	if !IsNil(o.FallbackValue) {
		toSerialize["fallback_value"] = o.FallbackValue
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.AdvancedUi) {
		toSerialize["advanced_ui"] = o.AdvancedUi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableComputedFieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"content_type",
		"label",
		"template",
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

	varBulkWritableComputedFieldRequest := _BulkWritableComputedFieldRequest{}

	err = json.Unmarshal(data, &varBulkWritableComputedFieldRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableComputedFieldRequest(varBulkWritableComputedFieldRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "key")
		delete(additionalProperties, "grouping")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "template")
		delete(additionalProperties, "fallback_value")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "advanced_ui")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableComputedFieldRequest struct {
	value *BulkWritableComputedFieldRequest
	isSet bool
}

func (v NullableBulkWritableComputedFieldRequest) Get() *BulkWritableComputedFieldRequest {
	return v.value
}

func (v *NullableBulkWritableComputedFieldRequest) Set(val *BulkWritableComputedFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableComputedFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableComputedFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableComputedFieldRequest(val *BulkWritableComputedFieldRequest) *NullableBulkWritableComputedFieldRequest {
	return &NullableBulkWritableComputedFieldRequest{value: val, isSet: true}
}

func (v NullableBulkWritableComputedFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableComputedFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


