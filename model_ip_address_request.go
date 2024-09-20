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

// checks if the IPAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPAddressRequest{}

// IPAddressRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type IPAddressRequest struct {
	Address string `json:"address"`
	Namespace *BulkWritableIPAddressRequestNamespace `json:"namespace,omitempty"`
	Type *IPAddressTypeChoices `json:"type,omitempty"`
	// Hostname or FQDN (not case-sensitive)
	DnsName *string `json:"dns_name,omitempty" validate:"regexp=^[0-9A-Za-z._-]+$"`
	Description *string `json:"description,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Parent NullableBulkWritableIPAddressRequestParent `json:"parent,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	NatInside NullableNATInside `json:"nat_inside,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPAddressRequest IPAddressRequest

// NewIPAddressRequest instantiates a new IPAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAddressRequest(address string, status BulkWritableCableRequestStatus) *IPAddressRequest {
	this := IPAddressRequest{}
	this.Address = address
	this.Status = status
	return &this
}

// NewIPAddressRequestWithDefaults instantiates a new IPAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAddressRequestWithDefaults() *IPAddressRequest {
	this := IPAddressRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *IPAddressRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *IPAddressRequest) SetAddress(v string) {
	o.Address = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IPAddressRequest) GetNamespace() BulkWritableIPAddressRequestNamespace {
	if o == nil || IsNil(o.Namespace) {
		var ret BulkWritableIPAddressRequestNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetNamespaceOk() (*BulkWritableIPAddressRequestNamespace, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IPAddressRequest) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given BulkWritableIPAddressRequestNamespace and assigns it to the Namespace field.
func (o *IPAddressRequest) SetNamespace(v BulkWritableIPAddressRequestNamespace) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPAddressRequest) GetType() IPAddressTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret IPAddressTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetTypeOk() (*IPAddressTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPAddressRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IPAddressTypeChoices and assigns it to the Type field.
func (o *IPAddressRequest) SetType(v IPAddressTypeChoices) {
	o.Type = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *IPAddressRequest) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *IPAddressRequest) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *IPAddressRequest) SetDnsName(v string) {
	o.DnsName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPAddressRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPAddressRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPAddressRequest) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *IPAddressRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IPAddressRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddressRequest) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddressRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *IPAddressRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *IPAddressRequest) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *IPAddressRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *IPAddressRequest) UnsetRole() {
	o.Role.Unset()
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddressRequest) GetParent() BulkWritableIPAddressRequestParent {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritableIPAddressRequestParent
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddressRequest) GetParentOk() (*BulkWritableIPAddressRequestParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *IPAddressRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritableIPAddressRequestParent and assigns it to the Parent field.
func (o *IPAddressRequest) SetParent(v BulkWritableIPAddressRequestParent) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *IPAddressRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *IPAddressRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddressRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddressRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *IPAddressRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *IPAddressRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *IPAddressRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *IPAddressRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetNatInside returns the NatInside field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddressRequest) GetNatInside() NATInside {
	if o == nil || IsNil(o.NatInside.Get()) {
		var ret NATInside
		return ret
	}
	return *o.NatInside.Get()
}

// GetNatInsideOk returns a tuple with the NatInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddressRequest) GetNatInsideOk() (*NATInside, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatInside.Get(), o.NatInside.IsSet()
}

// HasNatInside returns a boolean if a field has been set.
func (o *IPAddressRequest) HasNatInside() bool {
	if o != nil && o.NatInside.IsSet() {
		return true
	}

	return false
}

// SetNatInside gets a reference to the given NullableNATInside and assigns it to the NatInside field.
func (o *IPAddressRequest) SetNatInside(v NATInside) {
	o.NatInside.Set(&v)
}
// SetNatInsideNil sets the value for NatInside to be an explicit nil
func (o *IPAddressRequest) SetNatInsideNil() {
	o.NatInside.Set(nil)
}

// UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
func (o *IPAddressRequest) UnsetNatInside() {
	o.NatInside.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPAddressRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPAddressRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *IPAddressRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPAddressRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPAddressRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IPAddressRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IPAddressRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IPAddressRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *IPAddressRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o IPAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DnsName) {
		toSerialize["dns_name"] = o.DnsName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.NatInside.IsSet() {
		toSerialize["nat_inside"] = o.NatInside.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPAddressRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
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

	varIPAddressRequest := _IPAddressRequest{}

	err = json.Unmarshal(data, &varIPAddressRequest)

	if err != nil {
		return err
	}

	*o = IPAddressRequest(varIPAddressRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "type")
		delete(additionalProperties, "dns_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "nat_inside")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPAddressRequest struct {
	value *IPAddressRequest
	isSet bool
}

func (v NullableIPAddressRequest) Get() *IPAddressRequest {
	return v.value
}

func (v *NullableIPAddressRequest) Set(val *IPAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressRequest(val *IPAddressRequest) *NullableIPAddressRequest {
	return &NullableIPAddressRequest{value: val, isSet: true}
}

func (v NullableIPAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


