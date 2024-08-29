/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the BulkWritableCloudAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableCloudAccountRequest{}

// BulkWritableCloudAccountRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableCloudAccountRequest struct {
	Id string `json:"id"`
	// The name of this Cloud Account.
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// The account identifier of this Cloud Account.
	AccountNumber string `json:"account_number"`
	Provider BulkWritableCloudAccountRequestProvider `json:"provider"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableCloudAccountRequest BulkWritableCloudAccountRequest

// NewBulkWritableCloudAccountRequest instantiates a new BulkWritableCloudAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableCloudAccountRequest(id string, name string, accountNumber string, provider BulkWritableCloudAccountRequestProvider) *BulkWritableCloudAccountRequest {
	this := BulkWritableCloudAccountRequest{}
	this.Id = id
	this.Name = name
	this.AccountNumber = accountNumber
	this.Provider = provider
	return &this
}

// NewBulkWritableCloudAccountRequestWithDefaults instantiates a new BulkWritableCloudAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableCloudAccountRequestWithDefaults() *BulkWritableCloudAccountRequest {
	this := BulkWritableCloudAccountRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableCloudAccountRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableCloudAccountRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BulkWritableCloudAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableCloudAccountRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableCloudAccountRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableCloudAccountRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableCloudAccountRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAccountNumber returns the AccountNumber field value
func (o *BulkWritableCloudAccountRequest) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *BulkWritableCloudAccountRequest) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetProvider returns the Provider field value
func (o *BulkWritableCloudAccountRequest) GetProvider() BulkWritableCloudAccountRequestProvider {
	if o == nil {
		var ret BulkWritableCloudAccountRequestProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *BulkWritableCloudAccountRequest) SetProvider(v BulkWritableCloudAccountRequestProvider) {
	o.Provider = v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableCloudAccountRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableCloudAccountRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *BulkWritableCloudAccountRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *BulkWritableCloudAccountRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *BulkWritableCloudAccountRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *BulkWritableCloudAccountRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableCloudAccountRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableCloudAccountRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableCloudAccountRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableCloudAccountRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableCloudAccountRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableCloudAccountRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableCloudAccountRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCloudAccountRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableCloudAccountRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableCloudAccountRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableCloudAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableCloudAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["provider"] = o.Provider
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

func (o *BulkWritableCloudAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"account_number",
		"provider",
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

	varBulkWritableCloudAccountRequest := _BulkWritableCloudAccountRequest{}

	err = json.Unmarshal(data, &varBulkWritableCloudAccountRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableCloudAccountRequest(varBulkWritableCloudAccountRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableCloudAccountRequest struct {
	value *BulkWritableCloudAccountRequest
	isSet bool
}

func (v NullableBulkWritableCloudAccountRequest) Get() *BulkWritableCloudAccountRequest {
	return v.value
}

func (v *NullableBulkWritableCloudAccountRequest) Set(val *BulkWritableCloudAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableCloudAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableCloudAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableCloudAccountRequest(val *BulkWritableCloudAccountRequest) *NullableBulkWritableCloudAccountRequest {
	return &NullableBulkWritableCloudAccountRequest{value: val, isSet: true}
}

func (v NullableBulkWritableCloudAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableCloudAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

