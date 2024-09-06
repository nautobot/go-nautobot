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

// checks if the Rack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rack{}

// Rack Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type Rack struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Type *RackType `json:"type,omitempty"`
	Width *RackWidth `json:"width,omitempty"`
	OuterUnit *RackOuterUnit `json:"outer_unit,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	PowerFeedCount *int32 `json:"power_feed_count,omitempty"`
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
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Rack Rack

// NewRack instantiates a new Rack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRack(id string, objectType string, display string, url string, naturalSlug string, name string, status BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string) *Rack {
	this := Rack{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Name = name
	this.Status = status
	this.Location = location
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewRackWithDefaults instantiates a new Rack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackWithDefaults() *Rack {
	this := Rack{}
	return &this
}

// GetId returns the Id field value
func (o *Rack) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Rack) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Rack) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *Rack) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *Rack) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *Rack) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *Rack) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Rack) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Rack) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *Rack) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Rack) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Rack) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *Rack) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *Rack) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *Rack) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Rack) GetType() RackType {
	if o == nil || IsNil(o.Type) {
		var ret RackType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetTypeOk() (*RackType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Rack) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RackType and assigns it to the Type field.
func (o *Rack) SetType(v RackType) {
	o.Type = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Rack) GetWidth() RackWidth {
	if o == nil || IsNil(o.Width) {
		var ret RackWidth
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetWidthOk() (*RackWidth, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Rack) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given RackWidth and assigns it to the Width field.
func (o *Rack) SetWidth(v RackWidth) {
	o.Width = &v
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise.
func (o *Rack) GetOuterUnit() RackOuterUnit {
	if o == nil || IsNil(o.OuterUnit) {
		var ret RackOuterUnit
		return ret
	}
	return *o.OuterUnit
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetOuterUnitOk() (*RackOuterUnit, bool) {
	if o == nil || IsNil(o.OuterUnit) {
		return nil, false
	}
	return o.OuterUnit, true
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *Rack) HasOuterUnit() bool {
	if o != nil && !IsNil(o.OuterUnit) {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given RackOuterUnit and assigns it to the OuterUnit field.
func (o *Rack) SetOuterUnit(v RackOuterUnit) {
	o.OuterUnit = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *Rack) GetDeviceCount() int32 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetDeviceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *Rack) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *Rack) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetPowerFeedCount returns the PowerFeedCount field value if set, zero value otherwise.
func (o *Rack) GetPowerFeedCount() int32 {
	if o == nil || IsNil(o.PowerFeedCount) {
		var ret int32
		return ret
	}
	return *o.PowerFeedCount
}

// GetPowerFeedCountOk returns a tuple with the PowerFeedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetPowerFeedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PowerFeedCount) {
		return nil, false
	}
	return o.PowerFeedCount, true
}

// HasPowerFeedCount returns a boolean if a field has been set.
func (o *Rack) HasPowerFeedCount() bool {
	if o != nil && !IsNil(o.PowerFeedCount) {
		return true
	}

	return false
}

// SetPowerFeedCount gets a reference to the given int32 and assigns it to the PowerFeedCount field.
func (o *Rack) SetPowerFeedCount(v int32) {
	o.PowerFeedCount = &v
}

// GetName returns the Name field value
func (o *Rack) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Rack) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Rack) SetName(v string) {
	o.Name = v
}

// GetFacilityId returns the FacilityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetFacilityId() string {
	if o == nil || IsNil(o.FacilityId.Get()) {
		var ret string
		return ret
	}
	return *o.FacilityId.Get()
}

// GetFacilityIdOk returns a tuple with the FacilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetFacilityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacilityId.Get(), o.FacilityId.IsSet()
}

// HasFacilityId returns a boolean if a field has been set.
func (o *Rack) HasFacilityId() bool {
	if o != nil && o.FacilityId.IsSet() {
		return true
	}

	return false
}

// SetFacilityId gets a reference to the given NullableString and assigns it to the FacilityId field.
func (o *Rack) SetFacilityId(v string) {
	o.FacilityId.Set(&v)
}
// SetFacilityIdNil sets the value for FacilityId to be an explicit nil
func (o *Rack) SetFacilityIdNil() {
	o.FacilityId.Set(nil)
}

// UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
func (o *Rack) UnsetFacilityId() {
	o.FacilityId.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *Rack) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *Rack) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *Rack) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *Rack) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *Rack) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *Rack) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *Rack) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *Rack) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *Rack) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *Rack) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *Rack) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *Rack) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *Rack) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *Rack) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *Rack) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}
// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *Rack) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *Rack) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *Rack) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *Rack) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}
// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *Rack) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *Rack) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Rack) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Rack) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Rack) SetComments(v string) {
	o.Comments = &v
}

// GetStatus returns the Status field value
func (o *Rack) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Rack) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Rack) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *Rack) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *Rack) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *Rack) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *Rack) UnsetRole() {
	o.Role.Unset()
}

// GetLocation returns the Location field value
func (o *Rack) GetLocation() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Rack) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Rack) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = v
}

// GetRackGroup returns the RackGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetRackGroup() BulkWritableRackRequestRackGroup {
	if o == nil || IsNil(o.RackGroup.Get()) {
		var ret BulkWritableRackRequestRackGroup
		return ret
	}
	return *o.RackGroup.Get()
}

// GetRackGroupOk returns a tuple with the RackGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetRackGroupOk() (*BulkWritableRackRequestRackGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.RackGroup.Get(), o.RackGroup.IsSet()
}

// HasRackGroup returns a boolean if a field has been set.
func (o *Rack) HasRackGroup() bool {
	if o != nil && o.RackGroup.IsSet() {
		return true
	}

	return false
}

// SetRackGroup gets a reference to the given NullableBulkWritableRackRequestRackGroup and assigns it to the RackGroup field.
func (o *Rack) SetRackGroup(v BulkWritableRackRequestRackGroup) {
	o.RackGroup.Set(&v)
}
// SetRackGroupNil sets the value for RackGroup to be an explicit nil
func (o *Rack) SetRackGroupNil() {
	o.RackGroup.Set(nil)
}

// UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
func (o *Rack) UnsetRackGroup() {
	o.RackGroup.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Rack) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Rack) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *Rack) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Rack) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Rack) UnsetTenant() {
	o.Tenant.Unset()
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Rack) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Rack) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Rack) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Rack) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Rack) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *Rack) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *Rack) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *Rack) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Rack) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Rack) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Rack) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Rack) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rack) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Rack) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *Rack) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o Rack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.OuterUnit) {
		toSerialize["outer_unit"] = o.OuterUnit
	}
	if !IsNil(o.DeviceCount) {
		toSerialize["device_count"] = o.DeviceCount
	}
	if !IsNil(o.PowerFeedCount) {
		toSerialize["power_feed_count"] = o.PowerFeedCount
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

func (o *Rack) UnmarshalJSON(data []byte) (err error) {
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
		"status",
		"location",
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

	varRack := _Rack{}

	err = json.Unmarshal(data, &varRack)

	if err != nil {
		return err
	}

	*o = Rack(varRack)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "type")
		delete(additionalProperties, "width")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "power_feed_count")
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
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRack struct {
	value *Rack
	isSet bool
}

func (v NullableRack) Get() *Rack {
	return v.value
}

func (v *NullableRack) Set(val *Rack) {
	v.value = val
	v.isSet = true
}

func (v NullableRack) IsSet() bool {
	return v.isSet
}

func (v *NullableRack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRack(val *Rack) *NullableRack {
	return &NullableRack{value: val, isSet: true}
}

func (v NullableRack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


