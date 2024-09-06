/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the DeviceFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceFamily{}

// DeviceFamily Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type DeviceFamily struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	DeviceTypeCount *int32 `json:"device_type_count,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceFamily DeviceFamily

// NewDeviceFamily instantiates a new DeviceFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceFamily(id string, objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string) *DeviceFamily {
	this := DeviceFamily{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewDeviceFamilyWithDefaults instantiates a new DeviceFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceFamilyWithDefaults() *DeviceFamily {
	this := DeviceFamily{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceFamily) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceFamily) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *DeviceFamily) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DeviceFamily) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *DeviceFamily) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceFamily) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *DeviceFamily) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceFamily) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *DeviceFamily) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *DeviceFamily) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetDeviceTypeCount returns the DeviceTypeCount field value if set, zero value otherwise.
func (o *DeviceFamily) GetDeviceTypeCount() int32 {
	if o == nil || IsNil(o.DeviceTypeCount) {
		var ret int32
		return ret
	}
	return *o.DeviceTypeCount
}

// GetDeviceTypeCountOk returns a tuple with the DeviceTypeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetDeviceTypeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceTypeCount) {
		return nil, false
	}
	return o.DeviceTypeCount, true
}

// HasDeviceTypeCount returns a boolean if a field has been set.
func (o *DeviceFamily) HasDeviceTypeCount() bool {
	if o != nil && !IsNil(o.DeviceTypeCount) {
		return true
	}

	return false
}

// SetDeviceTypeCount gets a reference to the given int32 and assigns it to the DeviceTypeCount field.
func (o *DeviceFamily) SetDeviceTypeCount(v int32) {
	o.DeviceTypeCount = &v
}

// GetName returns the Name field value
func (o *DeviceFamily) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceFamily) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceFamily) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceFamily) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceFamily) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceFamily) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceFamily) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *DeviceFamily) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceFamily) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceFamily) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *DeviceFamily) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *DeviceFamily) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *DeviceFamily) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceFamily) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceFamily) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceFamily) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceFamily) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o DeviceFamily) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	if !IsNil(o.DeviceTypeCount) {
		toSerialize["device_type_count"] = o.DeviceTypeCount
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceFamily) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"name",
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

	varDeviceFamily := _DeviceFamily{}

	err = json.Unmarshal(data, &varDeviceFamily)

	if err != nil {
		return err
	}

	*o = DeviceFamily(varDeviceFamily)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "device_type_count")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceFamily struct {
	value *DeviceFamily
	isSet bool
}

func (v NullableDeviceFamily) Get() *DeviceFamily {
	return v.value
}

func (v *NullableDeviceFamily) Set(val *DeviceFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceFamily(val *DeviceFamily) *NullableDeviceFamily {
	return &NullableDeviceFamily{value: val, isSet: true}
}

func (v NullableDeviceFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


