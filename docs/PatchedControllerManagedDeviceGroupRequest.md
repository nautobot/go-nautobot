# PatchedControllerManagedDeviceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Capabilities** | Pointer to [**[]CapabilitiesEnum**](CapabilitiesEnum.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the controller device group | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Weight of the controller device group, used to sort the groups within its parent group | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Controller** | Pointer to [**BulkWritableControllerManagedDeviceGroupRequestController**](BulkWritableControllerManagedDeviceGroupRequestController.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedControllerManagedDeviceGroupRequest

`func NewPatchedControllerManagedDeviceGroupRequest() *PatchedControllerManagedDeviceGroupRequest`

NewPatchedControllerManagedDeviceGroupRequest instantiates a new PatchedControllerManagedDeviceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedControllerManagedDeviceGroupRequestWithDefaults

`func NewPatchedControllerManagedDeviceGroupRequestWithDefaults() *PatchedControllerManagedDeviceGroupRequest`

NewPatchedControllerManagedDeviceGroupRequestWithDefaults instantiates a new PatchedControllerManagedDeviceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedControllerManagedDeviceGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedControllerManagedDeviceGroupRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedControllerManagedDeviceGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilities

`func (o *PatchedControllerManagedDeviceGroupRequest) GetCapabilities() []CapabilitiesEnum`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetCapabilitiesOk() (*[]CapabilitiesEnum, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PatchedControllerManagedDeviceGroupRequest) SetCapabilities(v []CapabilitiesEnum)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PatchedControllerManagedDeviceGroupRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *PatchedControllerManagedDeviceGroupRequest) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *PatchedControllerManagedDeviceGroupRequest) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetName

`func (o *PatchedControllerManagedDeviceGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedControllerManagedDeviceGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedControllerManagedDeviceGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedControllerManagedDeviceGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedControllerManagedDeviceGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedControllerManagedDeviceGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedControllerManagedDeviceGroupRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedControllerManagedDeviceGroupRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedControllerManagedDeviceGroupRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetParent

`func (o *PatchedControllerManagedDeviceGroupRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedControllerManagedDeviceGroupRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedControllerManagedDeviceGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedControllerManagedDeviceGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedControllerManagedDeviceGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetController

`func (o *PatchedControllerManagedDeviceGroupRequest) GetController() BulkWritableControllerManagedDeviceGroupRequestController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetControllerOk() (*BulkWritableControllerManagedDeviceGroupRequestController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *PatchedControllerManagedDeviceGroupRequest) SetController(v BulkWritableControllerManagedDeviceGroupRequestController)`

SetController sets Controller field to given value.

### HasController

`func (o *PatchedControllerManagedDeviceGroupRequest) HasController() bool`

HasController returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedControllerManagedDeviceGroupRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedControllerManagedDeviceGroupRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedControllerManagedDeviceGroupRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedControllerManagedDeviceGroupRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedControllerManagedDeviceGroupRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedControllerManagedDeviceGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedControllerManagedDeviceGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedControllerManagedDeviceGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedControllerManagedDeviceGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedControllerManagedDeviceGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedControllerManagedDeviceGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedControllerManagedDeviceGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedControllerManagedDeviceGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedControllerManagedDeviceGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedControllerManagedDeviceGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


