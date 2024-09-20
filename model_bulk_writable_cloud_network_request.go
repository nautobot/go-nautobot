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

// checks if the BulkWritableCloudNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableCloudNetworkRequest{}

// BulkWritableCloudNetworkRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableCloudNetworkRequest struct {
	Id string `json:"id"`
	ExtraConfig interface{} `json:"extra_config,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	CloudResourceType BulkWritableCableRequestStatus `json:"cloud_resource_type"`
	CloudAccount BulkWritableCableRequestStatus `json:"cloud_account"`
	Parent NullableBulkWritableCircuitRequestTenant `json:"parent,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableCloudNetworkRequest BulkWritableCloudNetworkRequest

// NewBulkWritableCloudNetworkRequest instantiates a new BulkWritableCloudNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableCloudNetworkRequest(id string, name string, cloudResourceType BulkWritableCableRequestStatus, cloudAccount BulkWritableCableRequestStatus) *BulkWritableCloudNetworkRequest {
	this := BulkWritableCloudNetworkRequest{}
	this.Id = id
	this.Name = name
	this.CloudResourceType = cloudResourceType
	this.CloudAccount = cloudAccount
	return &this
}

// NewBulkWritableCloudNetworkRequestWithDefaults instantiates a new BulkWritableCloudNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableCloudNetworkRequestWithDefaults() *BulkWritableCloudNetworkRequest {
	this := BulkWritableCloudNetworkRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableCloudNetworkRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableCloudNetworkRequest) SetId(v string) {
	o.Id = v
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableCloudNetworkRequest) GetExtraConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableCloudNetworkRequest) GetExtraConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExtraConfig) {
		return nil, false
	}
	return &o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *BulkWritableCloudNetworkRequest) HasExtraConfig() bool {
	if o != nil && !IsNil(o.ExtraConfig) {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given interface{} and assigns it to the ExtraConfig field.
func (o *BulkWritableCloudNetworkRequest) SetExtraConfig(v interface{}) {
	o.ExtraConfig = v
}

// GetName returns the Name field value
func (o *BulkWritableCloudNetworkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableCloudNetworkRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableCloudNetworkRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableCloudNetworkRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableCloudNetworkRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCloudResourceType returns the CloudResourceType field value
func (o *BulkWritableCloudNetworkRequest) GetCloudResourceType() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.CloudResourceType
}

// GetCloudResourceTypeOk returns a tuple with the CloudResourceType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetCloudResourceTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudResourceType, true
}

// SetCloudResourceType sets field value
func (o *BulkWritableCloudNetworkRequest) SetCloudResourceType(v BulkWritableCableRequestStatus) {
	o.CloudResourceType = v
}

// GetCloudAccount returns the CloudAccount field value
func (o *BulkWritableCloudNetworkRequest) GetCloudAccount() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.CloudAccount
}

// GetCloudAccountOk returns a tuple with the CloudAccount field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetCloudAccountOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudAccount, true
}

// SetCloudAccount sets field value
func (o *BulkWritableCloudNetworkRequest) SetCloudAccount(v BulkWritableCableRequestStatus) {
	o.CloudAccount = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableCloudNetworkRequest) GetParent() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableCloudNetworkRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *BulkWritableCloudNetworkRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Parent field.
func (o *BulkWritableCloudNetworkRequest) SetParent(v BulkWritableCircuitRequestTenant) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *BulkWritableCloudNetworkRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *BulkWritableCloudNetworkRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableCloudNetworkRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableCloudNetworkRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableCloudNetworkRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableCloudNetworkRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableCloudNetworkRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableCloudNetworkRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableCloudNetworkRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudNetworkRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableCloudNetworkRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableCloudNetworkRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableCloudNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableCloudNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.ExtraConfig != nil {
		toSerialize["extra_config"] = o.ExtraConfig
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["cloud_resource_type"] = o.CloudResourceType
	toSerialize["cloud_account"] = o.CloudAccount
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
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

func (o *BulkWritableCloudNetworkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"cloud_resource_type",
		"cloud_account",
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

	varBulkWritableCloudNetworkRequest := _BulkWritableCloudNetworkRequest{}

	err = json.Unmarshal(data, &varBulkWritableCloudNetworkRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableCloudNetworkRequest(varBulkWritableCloudNetworkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "extra_config")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cloud_resource_type")
		delete(additionalProperties, "cloud_account")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableCloudNetworkRequest struct {
	value *BulkWritableCloudNetworkRequest
	isSet bool
}

func (v NullableBulkWritableCloudNetworkRequest) Get() *BulkWritableCloudNetworkRequest {
	return v.value
}

func (v *NullableBulkWritableCloudNetworkRequest) Set(val *BulkWritableCloudNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableCloudNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableCloudNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableCloudNetworkRequest(val *BulkWritableCloudNetworkRequest) *NullableBulkWritableCloudNetworkRequest {
	return &NullableBulkWritableCloudNetworkRequest{value: val, isSet: true}
}

func (v NullableBulkWritableCloudNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableCloudNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


