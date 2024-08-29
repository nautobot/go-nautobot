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

// checks if the BulkWritableRackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableRackRequest{}

// BulkWritableRackRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableRackRequest struct {
	Id string `json:"id"`
	Type *RackTypeChoices `json:"type,omitempty"`
	// Rail-to-rail width (in inches)
	Width *WidthEnum `json:"width,omitempty"`
	OuterUnit *OuterUnitEnum `json:"outer_unit,omitempty"`
	Name string `json:"name"`
	// Locally-assigned identifier
	FacilityId NullableString `json:"facility_id,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this rack
	AssetTag NullableString `json:"asset_tag,omitempty"`
	// Height in rack units
	UHeight *int32 `json:"u_height,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth NullableInt32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth NullableInt32 `json:"outer_depth,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Location BulkWritableCableRequestStatus `json:"location"`
	RackGroup NullableBulkWritableRackRequestRackGroup `json:"rack_group,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableRackRequest BulkWritableRackRequest

// NewBulkWritableRackRequest instantiates a new BulkWritableRackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableRackRequest(id string, name string, status BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus) *BulkWritableRackRequest {
	this := BulkWritableRackRequest{}
	this.Id = id
	this.Name = name
	this.Status = status
	this.Location = location
	return &this
}

// NewBulkWritableRackRequestWithDefaults instantiates a new BulkWritableRackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableRackRequestWithDefaults() *BulkWritableRackRequest {
	this := BulkWritableRackRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableRackRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableRackRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetType() RackTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret RackTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetTypeOk() (*RackTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RackTypeChoices and assigns it to the Type field.
func (o *BulkWritableRackRequest) SetType(v RackTypeChoices) {
	o.Type = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetWidth() WidthEnum {
	if o == nil || IsNil(o.Width) {
		var ret WidthEnum
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetWidthOk() (*WidthEnum, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given WidthEnum and assigns it to the Width field.
func (o *BulkWritableRackRequest) SetWidth(v WidthEnum) {
	o.Width = &v
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetOuterUnit() OuterUnitEnum {
	if o == nil || IsNil(o.OuterUnit) {
		var ret OuterUnitEnum
		return ret
	}
	return *o.OuterUnit
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetOuterUnitOk() (*OuterUnitEnum, bool) {
	if o == nil || IsNil(o.OuterUnit) {
		return nil, false
	}
	return o.OuterUnit, true
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasOuterUnit() bool {
	if o != nil && !IsNil(o.OuterUnit) {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given OuterUnitEnum and assigns it to the OuterUnit field.
func (o *BulkWritableRackRequest) SetOuterUnit(v OuterUnitEnum) {
	o.OuterUnit = &v
}

// GetName returns the Name field value
func (o *BulkWritableRackRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableRackRequest) SetName(v string) {
	o.Name = v
}

// GetFacilityId returns the FacilityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetFacilityId() string {
	if o == nil || IsNil(o.FacilityId.Get()) {
		var ret string
		return ret
	}
	return *o.FacilityId.Get()
}

// GetFacilityIdOk returns a tuple with the FacilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetFacilityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacilityId.Get(), o.FacilityId.IsSet()
}

// HasFacilityId returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasFacilityId() bool {
	if o != nil && o.FacilityId.IsSet() {
		return true
	}

	return false
}

// SetFacilityId gets a reference to the given NullableString and assigns it to the FacilityId field.
func (o *BulkWritableRackRequest) SetFacilityId(v string) {
	o.FacilityId.Set(&v)
}
// SetFacilityIdNil sets the value for FacilityId to be an explicit nil
func (o *BulkWritableRackRequest) SetFacilityIdNil() {
	o.FacilityId.Set(nil)
}

// UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetFacilityId() {
	o.FacilityId.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *BulkWritableRackRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *BulkWritableRackRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *BulkWritableRackRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *BulkWritableRackRequest) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *BulkWritableRackRequest) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *BulkWritableRackRequest) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}
// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *BulkWritableRackRequest) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *BulkWritableRackRequest) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}
// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *BulkWritableRackRequest) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *BulkWritableRackRequest) SetComments(v string) {
	o.Comments = &v
}

// GetStatus returns the Status field value
func (o *BulkWritableRackRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableRackRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *BulkWritableRackRequest) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *BulkWritableRackRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetRole() {
	o.Role.Unset()
}

// GetLocation returns the Location field value
func (o *BulkWritableRackRequest) GetLocation() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *BulkWritableRackRequest) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = v
}

// GetRackGroup returns the RackGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetRackGroup() BulkWritableRackRequestRackGroup {
	if o == nil || IsNil(o.RackGroup.Get()) {
		var ret BulkWritableRackRequestRackGroup
		return ret
	}
	return *o.RackGroup.Get()
}

// GetRackGroupOk returns a tuple with the RackGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetRackGroupOk() (*BulkWritableRackRequestRackGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.RackGroup.Get(), o.RackGroup.IsSet()
}

// HasRackGroup returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasRackGroup() bool {
	if o != nil && o.RackGroup.IsSet() {
		return true
	}

	return false
}

// SetRackGroup gets a reference to the given NullableBulkWritableRackRequestRackGroup and assigns it to the RackGroup field.
func (o *BulkWritableRackRequest) SetRackGroup(v BulkWritableRackRequestRackGroup) {
	o.RackGroup.Set(&v)
}
// SetRackGroupNil sets the value for RackGroup to be an explicit nil
func (o *BulkWritableRackRequest) SetRackGroupNil() {
	o.RackGroup.Set(nil)
}

// UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetRackGroup() {
	o.RackGroup.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *BulkWritableRackRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *BulkWritableRackRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *BulkWritableRackRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableRackRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableRackRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableRackRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableRackRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableRackRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableRackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableRackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.OuterUnit) {
		toSerialize["outer_unit"] = o.OuterUnit
	}
	toSerialize["name"] = o.Name
	if o.FacilityId.IsSet() {
		toSerialize["facility_id"] = o.FacilityId.Get()
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.DescUnits) {
		toSerialize["desc_units"] = o.DescUnits
	}
	if o.OuterWidth.IsSet() {
		toSerialize["outer_width"] = o.OuterWidth.Get()
	}
	if o.OuterDepth.IsSet() {
		toSerialize["outer_depth"] = o.OuterDepth.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["status"] = o.Status
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	toSerialize["location"] = o.Location
	if o.RackGroup.IsSet() {
		toSerialize["rack_group"] = o.RackGroup.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

func (o *BulkWritableRackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"status",
		"location",
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

	varBulkWritableRackRequest := _BulkWritableRackRequest{}

	err = json.Unmarshal(data, &varBulkWritableRackRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableRackRequest(varBulkWritableRackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "width")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "name")
		delete(additionalProperties, "facility_id")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "desc_units")
		delete(additionalProperties, "outer_width")
		delete(additionalProperties, "outer_depth")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "location")
		delete(additionalProperties, "rack_group")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableRackRequest struct {
	value *BulkWritableRackRequest
	isSet bool
}

func (v NullableBulkWritableRackRequest) Get() *BulkWritableRackRequest {
	return v.value
}

func (v *NullableBulkWritableRackRequest) Set(val *BulkWritableRackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableRackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableRackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableRackRequest(val *BulkWritableRackRequest) *NullableBulkWritableRackRequest {
	return &NullableBulkWritableRackRequest{value: val, isSet: true}
}

func (v NullableBulkWritableRackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableRackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

