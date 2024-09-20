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

// checks if the PatchedBulkWritableInventoryItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableInventoryItemRequest{}

// PatchedBulkWritableInventoryItemRequest Add a `tree_depth` field to non-nested model serializers based on django-tree-queries.
type PatchedBulkWritableInventoryItemRequest struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	// Manufacturer-assigned part identifier
	PartId *string `json:"part_id,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this item
	AssetTag NullableString `json:"asset_tag,omitempty"`
	// This item was automatically discovered
	Discovered *bool `json:"discovered,omitempty"`
	Parent NullableBulkWritableCircuitRequestTenant `json:"parent,omitempty"`
	Device *BulkWritableCableRequestStatus `json:"device,omitempty"`
	Manufacturer NullableBulkWritableCircuitRequestTenant `json:"manufacturer,omitempty"`
	SoftwareVersion NullableBulkWritableInventoryItemRequestSoftwareVersion `json:"software_version,omitempty"`
	// Override the software image files associated with the software version for this inventory item
	SoftwareImageFiles []SoftwareImageFiles `json:"software_image_files,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableInventoryItemRequest PatchedBulkWritableInventoryItemRequest

// NewPatchedBulkWritableInventoryItemRequest instantiates a new PatchedBulkWritableInventoryItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableInventoryItemRequest(id string) *PatchedBulkWritableInventoryItemRequest {
	this := PatchedBulkWritableInventoryItemRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableInventoryItemRequestWithDefaults instantiates a new PatchedBulkWritableInventoryItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableInventoryItemRequestWithDefaults() *PatchedBulkWritableInventoryItemRequest {
	this := PatchedBulkWritableInventoryItemRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableInventoryItemRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableInventoryItemRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableInventoryItemRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedBulkWritableInventoryItemRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableInventoryItemRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetPartId() string {
	if o == nil || IsNil(o.PartId) {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetPartIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartId) {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasPartId() bool {
	if o != nil && !IsNil(o.PartId) {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *PatchedBulkWritableInventoryItemRequest) SetPartId(v string) {
	o.PartId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *PatchedBulkWritableInventoryItemRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableInventoryItemRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableInventoryItemRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *PatchedBulkWritableInventoryItemRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetDiscovered() bool {
	if o == nil || IsNil(o.Discovered) {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.Discovered) {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasDiscovered() bool {
	if o != nil && !IsNil(o.Discovered) {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *PatchedBulkWritableInventoryItemRequest) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableInventoryItemRequest) GetParent() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableInventoryItemRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Parent field.
func (o *PatchedBulkWritableInventoryItemRequest) SetParent(v BulkWritableCircuitRequestTenant) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetDevice() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Device) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Device field.
func (o *PatchedBulkWritableInventoryItemRequest) SetDevice(v BulkWritableCableRequestStatus) {
	o.Device = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableInventoryItemRequest) GetManufacturer() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableInventoryItemRequest) GetManufacturerOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Manufacturer field.
func (o *PatchedBulkWritableInventoryItemRequest) SetManufacturer(v BulkWritableCircuitRequestTenant) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareVersion() BulkWritableInventoryItemRequestSoftwareVersion {
	if o == nil || IsNil(o.SoftwareVersion.Get()) {
		var ret BulkWritableInventoryItemRequestSoftwareVersion
		return ret
	}
	return *o.SoftwareVersion.Get()
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareVersionOk() (*BulkWritableInventoryItemRequestSoftwareVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoftwareVersion.Get(), o.SoftwareVersion.IsSet()
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasSoftwareVersion() bool {
	if o != nil && o.SoftwareVersion.IsSet() {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given NullableBulkWritableInventoryItemRequestSoftwareVersion and assigns it to the SoftwareVersion field.
func (o *PatchedBulkWritableInventoryItemRequest) SetSoftwareVersion(v BulkWritableInventoryItemRequestSoftwareVersion) {
	o.SoftwareVersion.Set(&v)
}
// SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) SetSoftwareVersionNil() {
	o.SoftwareVersion.Set(nil)
}

// UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
func (o *PatchedBulkWritableInventoryItemRequest) UnsetSoftwareVersion() {
	o.SoftwareVersion.Unset()
}

// GetSoftwareImageFiles returns the SoftwareImageFiles field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareImageFiles() []SoftwareImageFiles {
	if o == nil || IsNil(o.SoftwareImageFiles) {
		var ret []SoftwareImageFiles
		return ret
	}
	return o.SoftwareImageFiles
}

// GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareImageFilesOk() ([]SoftwareImageFiles, bool) {
	if o == nil || IsNil(o.SoftwareImageFiles) {
		return nil, false
	}
	return o.SoftwareImageFiles, true
}

// HasSoftwareImageFiles returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasSoftwareImageFiles() bool {
	if o != nil && !IsNil(o.SoftwareImageFiles) {
		return true
	}

	return false
}

// SetSoftwareImageFiles gets a reference to the given []SoftwareImageFiles and assigns it to the SoftwareImageFiles field.
func (o *PatchedBulkWritableInventoryItemRequest) SetSoftwareImageFiles(v []SoftwareImageFiles) {
	o.SoftwareImageFiles = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableInventoryItemRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableInventoryItemRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableInventoryItemRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableInventoryItemRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableInventoryItemRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableInventoryItemRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableInventoryItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableInventoryItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PartId) {
		toSerialize["part_id"] = o.PartId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.Discovered) {
		toSerialize["discovered"] = o.Discovered
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.SoftwareVersion.IsSet() {
		toSerialize["software_version"] = o.SoftwareVersion.Get()
	}
	if !IsNil(o.SoftwareImageFiles) {
		toSerialize["software_image_files"] = o.SoftwareImageFiles
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

func (o *PatchedBulkWritableInventoryItemRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableInventoryItemRequest := _PatchedBulkWritableInventoryItemRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableInventoryItemRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableInventoryItemRequest(varPatchedBulkWritableInventoryItemRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "part_id")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "discovered")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "device")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "software_version")
		delete(additionalProperties, "software_image_files")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableInventoryItemRequest struct {
	value *PatchedBulkWritableInventoryItemRequest
	isSet bool
}

func (v NullablePatchedBulkWritableInventoryItemRequest) Get() *PatchedBulkWritableInventoryItemRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableInventoryItemRequest) Set(val *PatchedBulkWritableInventoryItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableInventoryItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableInventoryItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableInventoryItemRequest(val *PatchedBulkWritableInventoryItemRequest) *NullablePatchedBulkWritableInventoryItemRequest {
	return &NullablePatchedBulkWritableInventoryItemRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableInventoryItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableInventoryItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


