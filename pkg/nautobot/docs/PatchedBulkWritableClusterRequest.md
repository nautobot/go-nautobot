# PatchedBulkWritableClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ClusterType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ClusterGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableClusterRequest

`func NewPatchedBulkWritableClusterRequest(id string, ) *PatchedBulkWritableClusterRequest`

NewPatchedBulkWritableClusterRequest instantiates a new PatchedBulkWritableClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableClusterRequestWithDefaults

`func NewPatchedBulkWritableClusterRequestWithDefaults() *PatchedBulkWritableClusterRequest`

NewPatchedBulkWritableClusterRequestWithDefaults instantiates a new PatchedBulkWritableClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBulkWritableClusterRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWritableClusterRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWritableClusterRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWritableClusterRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetClusterType

`func (o *PatchedBulkWritableClusterRequest) GetClusterType() BulkWritableCableRequestStatus`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *PatchedBulkWritableClusterRequest) GetClusterTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *PatchedBulkWritableClusterRequest) SetClusterType(v BulkWritableCableRequestStatus)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *PatchedBulkWritableClusterRequest) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetClusterGroup

`func (o *PatchedBulkWritableClusterRequest) GetClusterGroup() BulkWritableCircuitRequestTenant`

GetClusterGroup returns the ClusterGroup field if non-nil, zero value otherwise.

### GetClusterGroupOk

`func (o *PatchedBulkWritableClusterRequest) GetClusterGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetClusterGroupOk returns a tuple with the ClusterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroup

`func (o *PatchedBulkWritableClusterRequest) SetClusterGroup(v BulkWritableCircuitRequestTenant)`

SetClusterGroup sets ClusterGroup field to given value.

### HasClusterGroup

`func (o *PatchedBulkWritableClusterRequest) HasClusterGroup() bool`

HasClusterGroup returns a boolean if a field has been set.

### SetClusterGroupNil

`func (o *PatchedBulkWritableClusterRequest) SetClusterGroupNil(b bool)`

 SetClusterGroupNil sets the value for ClusterGroup to be an explicit nil

### UnsetClusterGroup
`func (o *PatchedBulkWritableClusterRequest) UnsetClusterGroup()`

UnsetClusterGroup ensures that no value is present for ClusterGroup, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableClusterRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableClusterRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableClusterRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableClusterRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableClusterRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableClusterRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLocation

`func (o *PatchedBulkWritableClusterRequest) GetLocation() BulkWritableCircuitRequestTenant`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkWritableClusterRequest) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkWritableClusterRequest) SetLocation(v BulkWritableCircuitRequestTenant)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkWritableClusterRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedBulkWritableClusterRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedBulkWritableClusterRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetTags

`func (o *PatchedBulkWritableClusterRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableClusterRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableClusterRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableClusterRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableClusterRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableClusterRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableClusterRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableClusterRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableClusterRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableClusterRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableClusterRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableClusterRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


