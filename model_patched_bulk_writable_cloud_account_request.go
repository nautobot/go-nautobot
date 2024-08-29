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

// checks if the PatchedBulkWritableCloudAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableCloudAccountRequest{}

// PatchedBulkWritableCloudAccountRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableCloudAccountRequest struct {
	Id string `json:"id"`
	// The name of this Cloud Account.
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// The account identifier of this Cloud Account.
	AccountNumber *string `json:"account_number,omitempty"`
	Provider *BulkWritableCloudAccountRequestProvider `json:"provider,omitempty"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableCloudAccountRequest PatchedBulkWritableCloudAccountRequest

// NewPatchedBulkWritableCloudAccountRequest instantiates a new PatchedBulkWritableCloudAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableCloudAccountRequest(id string) *PatchedBulkWritableCloudAccountRequest {
	this := PatchedBulkWritableCloudAccountRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableCloudAccountRequestWithDefaults instantiates a new PatchedBulkWritableCloudAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableCloudAccountRequestWithDefaults() *PatchedBulkWritableCloudAccountRequest {
	this := PatchedBulkWritableCloudAccountRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableCloudAccountRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableCloudAccountRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableCloudAccountRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableCloudAccountRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PatchedBulkWritableCloudAccountRequest) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetProvider() BulkWritableCloudAccountRequestProvider {
	if o == nil || IsNil(o.Provider) {
		var ret BulkWritableCloudAccountRequestProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given BulkWritableCloudAccountRequestProvider and assigns it to the Provider field.
func (o *PatchedBulkWritableCloudAccountRequest) SetProvider(v BulkWritableCloudAccountRequestProvider) {
	o.Provider = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCloudAccountRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCloudAccountRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *PatchedBulkWritableCloudAccountRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *PatchedBulkWritableCloudAccountRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *PatchedBulkWritableCloudAccountRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableCloudAccountRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableCloudAccountRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudAccountRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudAccountRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudAccountRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableCloudAccountRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableCloudAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableCloudAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
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

func (o *PatchedBulkWritableCloudAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPatchedBulkWritableCloudAccountRequest := _PatchedBulkWritableCloudAccountRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableCloudAccountRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableCloudAccountRequest(varPatchedBulkWritableCloudAccountRequest)

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

type NullablePatchedBulkWritableCloudAccountRequest struct {
	value *PatchedBulkWritableCloudAccountRequest
	isSet bool
}

func (v NullablePatchedBulkWritableCloudAccountRequest) Get() *PatchedBulkWritableCloudAccountRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableCloudAccountRequest) Set(val *PatchedBulkWritableCloudAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableCloudAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableCloudAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableCloudAccountRequest(val *PatchedBulkWritableCloudAccountRequest) *NullablePatchedBulkWritableCloudAccountRequest {
	return &NullablePatchedBulkWritableCloudAccountRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableCloudAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableCloudAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

