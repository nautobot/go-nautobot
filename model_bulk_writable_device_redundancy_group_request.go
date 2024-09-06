/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the BulkWritableDeviceRedundancyGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableDeviceRedundancyGroupRequest{}

// BulkWritableDeviceRedundancyGroupRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableDeviceRedundancyGroupRequest struct {
	Id string `json:"id"`
	FailoverStrategy *BulkWritableDeviceRedundancyGroupRequestFailoverStrategy `json:"failover_strategy,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableDeviceRedundancyGroupRequest BulkWritableDeviceRedundancyGroupRequest

// NewBulkWritableDeviceRedundancyGroupRequest instantiates a new BulkWritableDeviceRedundancyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableDeviceRedundancyGroupRequest(id string, name string, status BulkWritableCableRequestStatus) *BulkWritableDeviceRedundancyGroupRequest {
	this := BulkWritableDeviceRedundancyGroupRequest{}
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewBulkWritableDeviceRedundancyGroupRequestWithDefaults instantiates a new BulkWritableDeviceRedundancyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableDeviceRedundancyGroupRequestWithDefaults() *BulkWritableDeviceRedundancyGroupRequest {
	this := BulkWritableDeviceRedundancyGroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableDeviceRedundancyGroupRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableDeviceRedundancyGroupRequest) SetId(v string) {
	o.Id = v
}

// GetFailoverStrategy returns the FailoverStrategy field value if set, zero value otherwise.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetFailoverStrategy() BulkWritableDeviceRedundancyGroupRequestFailoverStrategy {
	if o == nil || IsNil(o.FailoverStrategy) {
		var ret BulkWritableDeviceRedundancyGroupRequestFailoverStrategy
		return ret
	}
	return *o.FailoverStrategy
}

// GetFailoverStrategyOk returns a tuple with the FailoverStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetFailoverStrategyOk() (*BulkWritableDeviceRedundancyGroupRequestFailoverStrategy, bool) {
	if o == nil || IsNil(o.FailoverStrategy) {
		return nil, false
	}
	return o.FailoverStrategy, true
}

// HasFailoverStrategy returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasFailoverStrategy() bool {
	if o != nil && !IsNil(o.FailoverStrategy) {
		return true
	}

	return false
}

// SetFailoverStrategy gets a reference to the given BulkWritableDeviceRedundancyGroupRequestFailoverStrategy and assigns it to the FailoverStrategy field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetFailoverStrategy(v BulkWritableDeviceRedundancyGroupRequestFailoverStrategy) {
	o.FailoverStrategy = &v
}

// GetName returns the Name field value
func (o *BulkWritableDeviceRedundancyGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableDeviceRedundancyGroupRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetComments(v string) {
	o.Comments = &v
}

// GetStatus returns the Status field value
func (o *BulkWritableDeviceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableDeviceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableDeviceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableDeviceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *BulkWritableDeviceRedundancyGroupRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *BulkWritableDeviceRedundancyGroupRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableDeviceRedundancyGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableDeviceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableDeviceRedundancyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableDeviceRedundancyGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.FailoverStrategy) {
		toSerialize["failover_strategy"] = o.FailoverStrategy
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["status"] = o.Status
	if o.SecretsGroup.IsSet() {
		toSerialize["secrets_group"] = o.SecretsGroup.Get()
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

func (o *BulkWritableDeviceRedundancyGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varBulkWritableDeviceRedundancyGroupRequest := _BulkWritableDeviceRedundancyGroupRequest{}

	err = json.Unmarshal(data, &varBulkWritableDeviceRedundancyGroupRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableDeviceRedundancyGroupRequest(varBulkWritableDeviceRedundancyGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "failover_strategy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "status")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableDeviceRedundancyGroupRequest struct {
	value *BulkWritableDeviceRedundancyGroupRequest
	isSet bool
}

func (v NullableBulkWritableDeviceRedundancyGroupRequest) Get() *BulkWritableDeviceRedundancyGroupRequest {
	return v.value
}

func (v *NullableBulkWritableDeviceRedundancyGroupRequest) Set(val *BulkWritableDeviceRedundancyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableDeviceRedundancyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableDeviceRedundancyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableDeviceRedundancyGroupRequest(val *BulkWritableDeviceRedundancyGroupRequest) *NullableBulkWritableDeviceRedundancyGroupRequest {
	return &NullableBulkWritableDeviceRedundancyGroupRequest{value: val, isSet: true}
}

func (v NullableBulkWritableDeviceRedundancyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableDeviceRedundancyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


