/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ConsolePort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsolePort{}

// ConsolePort Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type ConsolePort struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	CablePeerType NullableString `json:"cable_peer_type"`
	CablePeer NullableCableTermination `json:"cable_peer"`
	ConnectedEndpointType NullableString `json:"connected_endpoint_type"`
	ConnectedEndpoint NullablePathEndpoint `json:"connected_endpoint"`
	ConnectedEndpointReachable NullableBool `json:"connected_endpoint_reachable"`
	Type *ConsolePortType `json:"type,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Device NullableBulkWritableCircuitRequestTenant `json:"device,omitempty"`
	Module NullableBulkWritableCircuitRequestTenant `json:"module,omitempty"`
	Cable NullableCircuitCircuitTerminationA `json:"cable"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsolePort ConsolePort

// NewConsolePort instantiates a new ConsolePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsolePort(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, connectedEndpointType NullableString, connectedEndpoint NullablePathEndpoint, connectedEndpointReachable NullableBool, name string, cable NullableCircuitCircuitTerminationA, created NullableTime, lastUpdated NullableTime, notesUrl string) *ConsolePort {
	this := ConsolePort{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.CablePeerType = cablePeerType
	this.CablePeer = cablePeer
	this.ConnectedEndpointType = connectedEndpointType
	this.ConnectedEndpoint = connectedEndpoint
	this.ConnectedEndpointReachable = connectedEndpointReachable
	this.Name = name
	this.Cable = cable
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewConsolePortWithDefaults instantiates a new ConsolePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsolePortWithDefaults() *ConsolePort {
	this := ConsolePort{}
	return &this
}

// GetId returns the Id field value
func (o *ConsolePort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsolePort) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *ConsolePort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConsolePort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *ConsolePort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConsolePort) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *ConsolePort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConsolePort) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *ConsolePort) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *ConsolePort) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetCablePeerType returns the CablePeerType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConsolePort) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// SetCablePeerType sets field value
func (o *ConsolePort) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}

// GetCablePeer returns the CablePeer field value
// If the value is explicit nil, the zero value for CableTermination will be returned
func (o *ConsolePort) GetCablePeer() CableTermination {
	if o == nil || o.CablePeer.Get() == nil {
		var ret CableTermination
		return ret
	}

	return *o.CablePeer.Get()
}

// GetCablePeerOk returns a tuple with the CablePeer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetCablePeerOk() (*CableTermination, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeer.Get(), o.CablePeer.IsSet()
}

// SetCablePeer sets field value
func (o *ConsolePort) SetCablePeer(v CableTermination) {
	o.CablePeer.Set(&v)
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConsolePort) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConnectedEndpointType.Get()
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointType.Get(), o.ConnectedEndpointType.IsSet()
}

// SetConnectedEndpointType sets field value
func (o *ConsolePort) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType.Set(&v)
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value
// If the value is explicit nil, the zero value for PathEndpoint will be returned
func (o *ConsolePort) GetConnectedEndpoint() PathEndpoint {
	if o == nil || o.ConnectedEndpoint.Get() == nil {
		var ret PathEndpoint
		return ret
	}

	return *o.ConnectedEndpoint.Get()
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetConnectedEndpointOk() (*PathEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpoint.Get(), o.ConnectedEndpoint.IsSet()
}

// SetConnectedEndpoint sets field value
func (o *ConsolePort) SetConnectedEndpoint(v PathEndpoint) {
	o.ConnectedEndpoint.Set(&v)
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ConsolePort) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// SetConnectedEndpointReachable sets field value
func (o *ConsolePort) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsolePort) GetType() ConsolePortType {
	if o == nil || IsNil(o.Type) {
		var ret ConsolePortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetTypeOk() (*ConsolePortType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsolePort) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortType and assigns it to the Type field.
func (o *ConsolePort) SetType(v ConsolePortType) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *ConsolePort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsolePort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsolePort) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsolePort) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsolePort) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsolePort) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsolePort) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsolePort) SetDescription(v string) {
	o.Description = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsolePort) GetDevice() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *ConsolePort) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Device field.
func (o *ConsolePort) SetDevice(v BulkWritableCircuitRequestTenant) {
	o.Device.Set(&v)
}
// SetDeviceNil sets the value for Device to be an explicit nil
func (o *ConsolePort) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *ConsolePort) UnsetDevice() {
	o.Device.Unset()
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsolePort) GetModule() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetModuleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *ConsolePort) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Module field.
func (o *ConsolePort) SetModule(v BulkWritableCircuitRequestTenant) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *ConsolePort) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *ConsolePort) UnsetModule() {
	o.Module.Unset()
}

// GetCable returns the Cable field value
// If the value is explicit nil, the zero value for CircuitCircuitTerminationA will be returned
func (o *ConsolePort) GetCable() CircuitCircuitTerminationA {
	if o == nil || o.Cable.Get() == nil {
		var ret CircuitCircuitTerminationA
		return ret
	}

	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetCableOk() (*CircuitCircuitTerminationA, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// SetCable sets field value
func (o *ConsolePort) SetCable(v CircuitCircuitTerminationA) {
	o.Cable.Set(&v)
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConsolePort) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ConsolePort) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConsolePort) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ConsolePort) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *ConsolePort) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *ConsolePort) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ConsolePort) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ConsolePort) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ConsolePort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConsolePort) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConsolePort) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *ConsolePort) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o ConsolePort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsolePort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["cable_peer_type"] = o.CablePeerType.Get()
	toSerialize["cable_peer"] = o.CablePeer.Get()
	toSerialize["connected_endpoint_type"] = o.ConnectedEndpointType.Get()
	toSerialize["connected_endpoint"] = o.ConnectedEndpoint.Get()
	toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable.Get()
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["cable"] = o.Cable.Get()
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConsolePort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"cable_peer_type",
		"cable_peer",
		"connected_endpoint_type",
		"connected_endpoint",
		"connected_endpoint_reachable",
		"name",
		"cable",
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

	varConsolePort := _ConsolePort{}

	err = json.Unmarshal(data, &varConsolePort)

	if err != nil {
		return err
	}

	*o = ConsolePort(varConsolePort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "cable_peer_type")
		delete(additionalProperties, "cable_peer")
		delete(additionalProperties, "connected_endpoint_type")
		delete(additionalProperties, "connected_endpoint")
		delete(additionalProperties, "connected_endpoint_reachable")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsolePort struct {
	value *ConsolePort
	isSet bool
}

func (v NullableConsolePort) Get() *ConsolePort {
	return v.value
}

func (v *NullableConsolePort) Set(val *ConsolePort) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePort) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePort(val *ConsolePort) *NullableConsolePort {
	return &NullableConsolePort{value: val, isSet: true}
}

func (v NullableConsolePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


