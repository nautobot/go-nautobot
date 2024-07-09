# TenantGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewTenantGroupRequest

`func NewTenantGroupRequest(name string, ) *TenantGroupRequest`

NewTenantGroupRequest instantiates a new TenantGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantGroupRequestWithDefaults

`func NewTenantGroupRequestWithDefaults() *TenantGroupRequest`

NewTenantGroupRequestWithDefaults instantiates a new TenantGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TenantGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TenantGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *TenantGroupRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TenantGroupRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TenantGroupRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TenantGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *TenantGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *TenantGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetCustomFields

`func (o *TenantGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TenantGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TenantGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TenantGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *TenantGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TenantGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TenantGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TenantGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


