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

// checks if the Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cluster{}

// Cluster Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type Cluster struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	VirtualmachineCount *int32 `json:"virtualmachine_count,omitempty"`
	Name string `json:"name"`
	Comments *string `json:"comments,omitempty"`
	ClusterType BulkWritableCableRequestStatus `json:"cluster_type"`
	ClusterGroup NullableBulkWritableCircuitRequestTenant `json:"cluster_group,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Location NullableBulkWritableCircuitRequestTenant `json:"location,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Cluster Cluster

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(id string, objectType string, display string, url string, naturalSlug string, name string, clusterType BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string) *Cluster {
	this := Cluster{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Name = name
	this.ClusterType = clusterType
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetId returns the Id field value
func (o *Cluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cluster) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *Cluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *Cluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *Cluster) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Cluster) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *Cluster) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Cluster) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *Cluster) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *Cluster) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *Cluster) GetDeviceCount() int32 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDeviceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *Cluster) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *Cluster) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *Cluster) GetVirtualmachineCount() int32 {
	if o == nil || IsNil(o.VirtualmachineCount) {
		var ret int32
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil || IsNil(o.VirtualmachineCount) {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *Cluster) HasVirtualmachineCount() bool {
	if o != nil && !IsNil(o.VirtualmachineCount) {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int32 and assigns it to the VirtualmachineCount field.
func (o *Cluster) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = &v
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Cluster) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Cluster) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Cluster) SetComments(v string) {
	o.Comments = &v
}

// GetClusterType returns the ClusterType field value
func (o *Cluster) GetClusterType() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetClusterTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterType, true
}

// SetClusterType sets field value
func (o *Cluster) SetClusterType(v BulkWritableCableRequestStatus) {
	o.ClusterType = v
}

// GetClusterGroup returns the ClusterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetClusterGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ClusterGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ClusterGroup.Get()
}

// GetClusterGroupOk returns a tuple with the ClusterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetClusterGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterGroup.Get(), o.ClusterGroup.IsSet()
}

// HasClusterGroup returns a boolean if a field has been set.
func (o *Cluster) HasClusterGroup() bool {
	if o != nil && o.ClusterGroup.IsSet() {
		return true
	}

	return false
}

// SetClusterGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ClusterGroup field.
func (o *Cluster) SetClusterGroup(v BulkWritableCircuitRequestTenant) {
	o.ClusterGroup.Set(&v)
}
// SetClusterGroupNil sets the value for ClusterGroup to be an explicit nil
func (o *Cluster) SetClusterGroupNil() {
	o.ClusterGroup.Set(nil)
}

// UnsetClusterGroup ensures that no value is present for ClusterGroup, not even an explicit nil
func (o *Cluster) UnsetClusterGroup() {
	o.ClusterGroup.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Cluster) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *Cluster) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Cluster) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Cluster) UnsetTenant() {
	o.Tenant.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetLocation() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Location.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *Cluster) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Location field.
func (o *Cluster) SetLocation(v BulkWritableCircuitRequestTenant) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *Cluster) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *Cluster) UnsetLocation() {
	o.Location.Unset()
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Cluster) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Cluster) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Cluster) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Cluster) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Cluster) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Cluster) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *Cluster) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetNotesUrl returns the NotesUrl field value
func (o *Cluster) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *Cluster) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Cluster) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Cluster) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Cluster) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	if !IsNil(o.DeviceCount) {
		toSerialize["device_count"] = o.DeviceCount
	}
	if !IsNil(o.VirtualmachineCount) {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["cluster_type"] = o.ClusterType
	if o.ClusterGroup.IsSet() {
		toSerialize["cluster_group"] = o.ClusterGroup.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cluster) UnmarshalJSON(data []byte) (err error) {
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
		"cluster_type",
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

	varCluster := _Cluster{}

	err = json.Unmarshal(data, &varCluster)

	if err != nil {
		return err
	}

	*o = Cluster(varCluster)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "virtualmachine_count")
		delete(additionalProperties, "name")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "cluster_type")
		delete(additionalProperties, "cluster_group")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "location")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


