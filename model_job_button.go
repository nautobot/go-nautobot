/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the JobButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobButton{}

// JobButton Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type JobButton struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	ContentTypes []string `json:"content_types"`
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	// Jinja2 template code for button text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>. Buttons which render as empty text will not be displayed.
	Text string `json:"text"`
	Weight *int32 `json:"weight,omitempty"`
	// Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group.
	GroupName *string `json:"group_name,omitempty"`
	ButtonClass *ButtonClassEnum `json:"button_class,omitempty"`
	// Enable confirmation pop-up box. <span class='text-danger'>WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!</span>
	Confirmation *bool `json:"confirmation,omitempty"`
	Job BulkWritableJobButtonRequestJob `json:"job"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	AdditionalProperties map[string]interface{}
}

type _JobButton JobButton

// NewJobButton instantiates a new JobButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobButton(id string, objectType string, display string, url string, naturalSlug string, contentTypes []string, name string, text string, job BulkWritableJobButtonRequestJob, created NullableTime, lastUpdated NullableTime, notesUrl string) *JobButton {
	this := JobButton{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.ContentTypes = contentTypes
	this.Name = name
	this.Text = text
	this.Job = job
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewJobButtonWithDefaults instantiates a new JobButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobButtonWithDefaults() *JobButton {
	this := JobButton{}
	return &this
}

// GetId returns the Id field value
func (o *JobButton) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobButton) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *JobButton) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *JobButton) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *JobButton) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *JobButton) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *JobButton) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JobButton) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *JobButton) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *JobButton) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetContentTypes returns the ContentTypes field value
func (o *JobButton) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *JobButton) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *JobButton) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobButton) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *JobButton) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobButton) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *JobButton) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *JobButton) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetText returns the Text field value
func (o *JobButton) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *JobButton) SetText(v string) {
	o.Text = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *JobButton) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobButton) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *JobButton) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *JobButton) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *JobButton) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobButton) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *JobButton) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *JobButton) SetGroupName(v string) {
	o.GroupName = &v
}

// GetButtonClass returns the ButtonClass field value if set, zero value otherwise.
func (o *JobButton) GetButtonClass() ButtonClassEnum {
	if o == nil || IsNil(o.ButtonClass) {
		var ret ButtonClassEnum
		return ret
	}
	return *o.ButtonClass
}

// GetButtonClassOk returns a tuple with the ButtonClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobButton) GetButtonClassOk() (*ButtonClassEnum, bool) {
	if o == nil || IsNil(o.ButtonClass) {
		return nil, false
	}
	return o.ButtonClass, true
}

// HasButtonClass returns a boolean if a field has been set.
func (o *JobButton) HasButtonClass() bool {
	if o != nil && !IsNil(o.ButtonClass) {
		return true
	}

	return false
}

// SetButtonClass gets a reference to the given ButtonClassEnum and assigns it to the ButtonClass field.
func (o *JobButton) SetButtonClass(v ButtonClassEnum) {
	o.ButtonClass = &v
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise.
func (o *JobButton) GetConfirmation() bool {
	if o == nil || IsNil(o.Confirmation) {
		var ret bool
		return ret
	}
	return *o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobButton) GetConfirmationOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmation) {
		return nil, false
	}
	return o.Confirmation, true
}

// HasConfirmation returns a boolean if a field has been set.
func (o *JobButton) HasConfirmation() bool {
	if o != nil && !IsNil(o.Confirmation) {
		return true
	}

	return false
}

// SetConfirmation gets a reference to the given bool and assigns it to the Confirmation field.
func (o *JobButton) SetConfirmation(v bool) {
	o.Confirmation = &v
}

// GetJob returns the Job field value
func (o *JobButton) GetJob() BulkWritableJobButtonRequestJob {
	if o == nil {
		var ret BulkWritableJobButtonRequestJob
		return ret
	}

	return o.Job
}

// GetJobOk returns a tuple with the Job field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetJobOk() (*BulkWritableJobButtonRequestJob, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Job, true
}

// SetJob sets field value
func (o *JobButton) SetJob(v BulkWritableJobButtonRequestJob) {
	o.Job = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *JobButton) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobButton) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *JobButton) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *JobButton) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobButton) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *JobButton) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *JobButton) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *JobButton) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *JobButton) SetNotesUrl(v string) {
	o.NotesUrl = v
}

func (o JobButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["text"] = o.Text
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.ButtonClass) {
		toSerialize["button_class"] = o.ButtonClass
	}
	if !IsNil(o.Confirmation) {
		toSerialize["confirmation"] = o.Confirmation
	}
	toSerialize["job"] = o.Job
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["notes_url"] = o.NotesUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobButton) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"content_types",
		"name",
		"text",
		"job",
		"created",
		"last_updated",
		"notes_url",
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

	varJobButton := _JobButton{}

	err = json.Unmarshal(data, &varJobButton)

	if err != nil {
		return err
	}

	*o = JobButton(varJobButton)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "text")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "button_class")
		delete(additionalProperties, "confirmation")
		delete(additionalProperties, "job")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobButton struct {
	value *JobButton
	isSet bool
}

func (v NullableJobButton) Get() *JobButton {
	return v.value
}

func (v *NullableJobButton) Set(val *JobButton) {
	v.value = val
	v.isSet = true
}

func (v NullableJobButton) IsSet() bool {
	return v.isSet
}

func (v *NullableJobButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobButton(val *JobButton) *NullableJobButton {
	return &NullableJobButton{value: val, isSet: true}
}

func (v NullableJobButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

