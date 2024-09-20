/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the BulkWritableDeviceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableDeviceTypeRequest{}

// BulkWritableDeviceTypeRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableDeviceTypeRequest struct {
	Id string `json:"id"`
	SubdeviceRole *SubdeviceRoleEnum `json:"subdevice_role,omitempty"`
	FrontImage **os.File `json:"front_image,omitempty"`
	RearImage **os.File `json:"rear_image,omitempty"`
	Model string `json:"model"`
	// Discrete part number (optional)
	PartNumber *string `json:"part_number,omitempty"`
	UHeight *int32 `json:"u_height,omitempty"`
	// Device consumes both front and rear rack faces
	IsFullDepth *bool `json:"is_full_depth,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Manufacturer BulkWritableCableRequestStatus `json:"manufacturer"`
	DeviceFamily NullableBulkWritableCircuitRequestTenant `json:"device_family,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableDeviceTypeRequest BulkWritableDeviceTypeRequest

// NewBulkWritableDeviceTypeRequest instantiates a new BulkWritableDeviceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableDeviceTypeRequest(id string, model string, manufacturer BulkWritableCableRequestStatus) *BulkWritableDeviceTypeRequest {
	this := BulkWritableDeviceTypeRequest{}
	this.Id = id
	this.Model = model
	this.Manufacturer = manufacturer
	return &this
}

// NewBulkWritableDeviceTypeRequestWithDefaults instantiates a new BulkWritableDeviceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableDeviceTypeRequestWithDefaults() *BulkWritableDeviceTypeRequest {
	this := BulkWritableDeviceTypeRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableDeviceTypeRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableDeviceTypeRequest) SetId(v string) {
	o.Id = v
}

// GetSubdeviceRole returns the SubdeviceRole field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetSubdeviceRole() SubdeviceRoleEnum {
	if o == nil || IsNil(o.SubdeviceRole) {
		var ret SubdeviceRoleEnum
		return ret
	}
	return *o.SubdeviceRole
}

// GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetSubdeviceRoleOk() (*SubdeviceRoleEnum, bool) {
	if o == nil || IsNil(o.SubdeviceRole) {
		return nil, false
	}
	return o.SubdeviceRole, true
}

// HasSubdeviceRole returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasSubdeviceRole() bool {
	if o != nil && !IsNil(o.SubdeviceRole) {
		return true
	}

	return false
}

// SetSubdeviceRole gets a reference to the given SubdeviceRoleEnum and assigns it to the SubdeviceRole field.
func (o *BulkWritableDeviceTypeRequest) SetSubdeviceRole(v SubdeviceRoleEnum) {
	o.SubdeviceRole = &v
}

// GetFrontImage returns the FrontImage field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetFrontImage() *os.File {
	if o == nil || IsNil(o.FrontImage) {
		var ret *os.File
		return ret
	}
	return *o.FrontImage
}

// GetFrontImageOk returns a tuple with the FrontImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetFrontImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.FrontImage) {
		return nil, false
	}
	return o.FrontImage, true
}

// HasFrontImage returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasFrontImage() bool {
	if o != nil && !IsNil(o.FrontImage) {
		return true
	}

	return false
}

// SetFrontImage gets a reference to the given *os.File and assigns it to the FrontImage field.
func (o *BulkWritableDeviceTypeRequest) SetFrontImage(v *os.File) {
	o.FrontImage = &v
}

// GetRearImage returns the RearImage field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetRearImage() *os.File {
	if o == nil || IsNil(o.RearImage) {
		var ret *os.File
		return ret
	}
	return *o.RearImage
}

// GetRearImageOk returns a tuple with the RearImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetRearImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.RearImage) {
		return nil, false
	}
	return o.RearImage, true
}

// HasRearImage returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasRearImage() bool {
	if o != nil && !IsNil(o.RearImage) {
		return true
	}

	return false
}

// SetRearImage gets a reference to the given *os.File and assigns it to the RearImage field.
func (o *BulkWritableDeviceTypeRequest) SetRearImage(v *os.File) {
	o.RearImage = &v
}

// GetModel returns the Model field value
func (o *BulkWritableDeviceTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BulkWritableDeviceTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BulkWritableDeviceTypeRequest) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *BulkWritableDeviceTypeRequest) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetIsFullDepth returns the IsFullDepth field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetIsFullDepth() bool {
	if o == nil || IsNil(o.IsFullDepth) {
		var ret bool
		return ret
	}
	return *o.IsFullDepth
}

// GetIsFullDepthOk returns a tuple with the IsFullDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetIsFullDepthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullDepth) {
		return nil, false
	}
	return o.IsFullDepth, true
}

// HasIsFullDepth returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasIsFullDepth() bool {
	if o != nil && !IsNil(o.IsFullDepth) {
		return true
	}

	return false
}

// SetIsFullDepth gets a reference to the given bool and assigns it to the IsFullDepth field.
func (o *BulkWritableDeviceTypeRequest) SetIsFullDepth(v bool) {
	o.IsFullDepth = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *BulkWritableDeviceTypeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetManufacturer returns the Manufacturer field value
func (o *BulkWritableDeviceTypeRequest) GetManufacturer() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetManufacturerOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *BulkWritableDeviceTypeRequest) SetManufacturer(v BulkWritableCableRequestStatus) {
	o.Manufacturer = v
}

// GetDeviceFamily returns the DeviceFamily field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableDeviceTypeRequest) GetDeviceFamily() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.DeviceFamily.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.DeviceFamily.Get()
}

// GetDeviceFamilyOk returns a tuple with the DeviceFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableDeviceTypeRequest) GetDeviceFamilyOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceFamily.Get(), o.DeviceFamily.IsSet()
}

// HasDeviceFamily returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasDeviceFamily() bool {
	if o != nil && o.DeviceFamily.IsSet() {
		return true
	}

	return false
}

// SetDeviceFamily gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the DeviceFamily field.
func (o *BulkWritableDeviceTypeRequest) SetDeviceFamily(v BulkWritableCircuitRequestTenant) {
	o.DeviceFamily.Set(&v)
}
// SetDeviceFamilyNil sets the value for DeviceFamily to be an explicit nil
func (o *BulkWritableDeviceTypeRequest) SetDeviceFamilyNil() {
	o.DeviceFamily.Set(nil)
}

// UnsetDeviceFamily ensures that no value is present for DeviceFamily, not even an explicit nil
func (o *BulkWritableDeviceTypeRequest) UnsetDeviceFamily() {
	o.DeviceFamily.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableDeviceTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableDeviceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableDeviceTypeRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceTypeRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableDeviceTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableDeviceTypeRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableDeviceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.SubdeviceRole) {
		toSerialize["subdevice_role"] = o.SubdeviceRole
	}
	if !IsNil(o.FrontImage) {
		toSerialize["front_image"] = o.FrontImage
	}
	if !IsNil(o.RearImage) {
		toSerialize["rear_image"] = o.RearImage
	}
	toSerialize["model"] = o.Model
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.IsFullDepth) {
		toSerialize["is_full_depth"] = o.IsFullDepth
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["manufacturer"] = o.Manufacturer
	if o.DeviceFamily.IsSet() {
		toSerialize["device_family"] = o.DeviceFamily.Get()
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

func (o *BulkWritableDeviceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"model",
		"manufacturer",
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

	varBulkWritableDeviceTypeRequest := _BulkWritableDeviceTypeRequest{}

	err = json.Unmarshal(data, &varBulkWritableDeviceTypeRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableDeviceTypeRequest(varBulkWritableDeviceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "subdevice_role")
		delete(additionalProperties, "front_image")
		delete(additionalProperties, "rear_image")
		delete(additionalProperties, "model")
		delete(additionalProperties, "part_number")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "is_full_depth")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "device_family")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableDeviceTypeRequest struct {
	value *BulkWritableDeviceTypeRequest
	isSet bool
}

func (v NullableBulkWritableDeviceTypeRequest) Get() *BulkWritableDeviceTypeRequest {
	return v.value
}

func (v *NullableBulkWritableDeviceTypeRequest) Set(val *BulkWritableDeviceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableDeviceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableDeviceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableDeviceTypeRequest(val *BulkWritableDeviceTypeRequest) *NullableBulkWritableDeviceTypeRequest {
	return &NullableBulkWritableDeviceTypeRequest{value: val, isSet: true}
}

func (v NullableBulkWritableDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableDeviceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


