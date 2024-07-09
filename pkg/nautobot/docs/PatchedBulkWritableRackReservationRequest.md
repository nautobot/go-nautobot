# PatchedBulkWritableRackReservationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Units** | Pointer to **map[string]interface{}** | List of rack unit numbers to reserve | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Rack** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**User** | Pointer to [**BulkWritableRackReservationRequestUser**](BulkWritableRackReservationRequestUser.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableRackReservationRequest

`func NewPatchedBulkWritableRackReservationRequest(id string, ) *PatchedBulkWritableRackReservationRequest`

NewPatchedBulkWritableRackReservationRequest instantiates a new PatchedBulkWritableRackReservationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableRackReservationRequestWithDefaults

`func NewPatchedBulkWritableRackReservationRequestWithDefaults() *PatchedBulkWritableRackReservationRequest`

NewPatchedBulkWritableRackReservationRequestWithDefaults instantiates a new PatchedBulkWritableRackReservationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableRackReservationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableRackReservationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableRackReservationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUnits

`func (o *PatchedBulkWritableRackReservationRequest) GetUnits() map[string]interface{}`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *PatchedBulkWritableRackReservationRequest) GetUnitsOk() (*map[string]interface{}, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *PatchedBulkWritableRackReservationRequest) SetUnits(v map[string]interface{})`

SetUnits sets Units field to given value.

### HasUnits

`func (o *PatchedBulkWritableRackReservationRequest) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableRackReservationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableRackReservationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableRackReservationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableRackReservationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRack

`func (o *PatchedBulkWritableRackReservationRequest) GetRack() BulkWritableCableRequestStatus`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedBulkWritableRackReservationRequest) GetRackOk() (*BulkWritableCableRequestStatus, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedBulkWritableRackReservationRequest) SetRack(v BulkWritableCableRequestStatus)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedBulkWritableRackReservationRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWritableRackReservationRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableRackReservationRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableRackReservationRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableRackReservationRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableRackReservationRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableRackReservationRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetUser

`func (o *PatchedBulkWritableRackReservationRequest) GetUser() BulkWritableRackReservationRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkWritableRackReservationRequest) GetUserOk() (*BulkWritableRackReservationRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkWritableRackReservationRequest) SetUser(v BulkWritableRackReservationRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkWritableRackReservationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableRackReservationRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableRackReservationRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableRackReservationRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableRackReservationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableRackReservationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableRackReservationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableRackReservationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableRackReservationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableRackReservationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableRackReservationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableRackReservationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableRackReservationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


