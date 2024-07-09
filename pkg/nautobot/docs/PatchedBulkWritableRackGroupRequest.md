# PatchedBulkWritableRackGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableRackGroupRequest

`func NewPatchedBulkWritableRackGroupRequest(id string, ) *PatchedBulkWritableRackGroupRequest`

NewPatchedBulkWritableRackGroupRequest instantiates a new PatchedBulkWritableRackGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableRackGroupRequestWithDefaults

`func NewPatchedBulkWritableRackGroupRequestWithDefaults() *PatchedBulkWritableRackGroupRequest`

NewPatchedBulkWritableRackGroupRequestWithDefaults instantiates a new PatchedBulkWritableRackGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableRackGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableRackGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableRackGroupRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableRackGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableRackGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableRackGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableRackGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableRackGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableRackGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableRackGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableRackGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkWritableRackGroupRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkWritableRackGroupRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkWritableRackGroupRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkWritableRackGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkWritableRackGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkWritableRackGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetLocation

`func (o *PatchedBulkWritableRackGroupRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkWritableRackGroupRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkWritableRackGroupRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkWritableRackGroupRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableRackGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableRackGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableRackGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableRackGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableRackGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableRackGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableRackGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableRackGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


