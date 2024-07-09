# PatchedClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPatchedClusterRequest

`func NewPatchedClusterRequest() *PatchedClusterRequest`

NewPatchedClusterRequest instantiates a new PatchedClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedClusterRequestWithDefaults

`func NewPatchedClusterRequestWithDefaults() *PatchedClusterRequest`

NewPatchedClusterRequestWithDefaults instantiates a new PatchedClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *PatchedClusterRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedClusterRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedClusterRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedClusterRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetClusterType

`func (o *PatchedClusterRequest) GetClusterType() BulkWritableCableRequestStatus`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *PatchedClusterRequest) GetClusterTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *PatchedClusterRequest) SetClusterType(v BulkWritableCableRequestStatus)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *PatchedClusterRequest) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetClusterGroup

`func (o *PatchedClusterRequest) GetClusterGroup() BulkWritableCircuitRequestTenant`

GetClusterGroup returns the ClusterGroup field if non-nil, zero value otherwise.

### GetClusterGroupOk

`func (o *PatchedClusterRequest) GetClusterGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetClusterGroupOk returns a tuple with the ClusterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroup

`func (o *PatchedClusterRequest) SetClusterGroup(v BulkWritableCircuitRequestTenant)`

SetClusterGroup sets ClusterGroup field to given value.

### HasClusterGroup

`func (o *PatchedClusterRequest) HasClusterGroup() bool`

HasClusterGroup returns a boolean if a field has been set.

### SetClusterGroupNil

`func (o *PatchedClusterRequest) SetClusterGroupNil(b bool)`

 SetClusterGroupNil sets the value for ClusterGroup to be an explicit nil

### UnsetClusterGroup
`func (o *PatchedClusterRequest) UnsetClusterGroup()`

UnsetClusterGroup ensures that no value is present for ClusterGroup, not even an explicit nil
### GetTenant

`func (o *PatchedClusterRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedClusterRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedClusterRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedClusterRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedClusterRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedClusterRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLocation

`func (o *PatchedClusterRequest) GetLocation() BulkWritableCircuitRequestTenant`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedClusterRequest) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedClusterRequest) SetLocation(v BulkWritableCircuitRequestTenant)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedClusterRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedClusterRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedClusterRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetTags

`func (o *PatchedClusterRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedClusterRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedClusterRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedClusterRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedClusterRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedClusterRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedClusterRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedClusterRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedClusterRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedClusterRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedClusterRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedClusterRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


