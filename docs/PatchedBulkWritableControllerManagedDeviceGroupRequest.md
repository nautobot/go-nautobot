# PatchedBulkWritableControllerManagedDeviceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Capabilities** | Pointer to [**[]CapabilitiesEnum**](CapabilitiesEnum.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the controller device group | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Weight of the controller device group, used to sort the groups within its parent group | [optional] 
**Parent** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Controller** | Pointer to [**BulkWritableControllerManagedDeviceGroupRequestController**](BulkWritableControllerManagedDeviceGroupRequestController.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableControllerManagedDeviceGroupRequest

`func NewPatchedBulkWritableControllerManagedDeviceGroupRequest(id string, ) *PatchedBulkWritableControllerManagedDeviceGroupRequest`

NewPatchedBulkWritableControllerManagedDeviceGroupRequest instantiates a new PatchedBulkWritableControllerManagedDeviceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableControllerManagedDeviceGroupRequestWithDefaults

`func NewPatchedBulkWritableControllerManagedDeviceGroupRequestWithDefaults() *PatchedBulkWritableControllerManagedDeviceGroupRequest`

NewPatchedBulkWritableControllerManagedDeviceGroupRequestWithDefaults instantiates a new PatchedBulkWritableControllerManagedDeviceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCapabilities

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetCapabilities() []CapabilitiesEnum`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetCapabilitiesOk() (*[]CapabilitiesEnum, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetCapabilities(v []CapabilitiesEnum)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetName

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetParent() ApprovalWorkflowUser`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetParentOk() (*ApprovalWorkflowUser, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetParent(v ApprovalWorkflowUser)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetController

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetController() BulkWritableControllerManagedDeviceGroupRequestController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetControllerOk() (*BulkWritableControllerManagedDeviceGroupRequestController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetController(v BulkWritableControllerManagedDeviceGroupRequestController)`

SetController sets Controller field to given value.

### HasController

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasController() bool`

HasController returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableControllerManagedDeviceGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


