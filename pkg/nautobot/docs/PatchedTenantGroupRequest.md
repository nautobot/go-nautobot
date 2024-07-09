# PatchedTenantGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedTenantGroupRequest

`func NewPatchedTenantGroupRequest() *PatchedTenantGroupRequest`

NewPatchedTenantGroupRequest instantiates a new PatchedTenantGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTenantGroupRequestWithDefaults

`func NewPatchedTenantGroupRequestWithDefaults() *PatchedTenantGroupRequest`

NewPatchedTenantGroupRequestWithDefaults instantiates a new PatchedTenantGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedTenantGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTenantGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTenantGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTenantGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTenantGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTenantGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTenantGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTenantGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *PatchedTenantGroupRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedTenantGroupRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedTenantGroupRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedTenantGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedTenantGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedTenantGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetCustomFields

`func (o *PatchedTenantGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedTenantGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedTenantGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedTenantGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedTenantGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedTenantGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedTenantGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedTenantGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


