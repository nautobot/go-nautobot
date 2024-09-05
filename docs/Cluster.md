# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**VirtualmachineCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Comments** | Pointer to **string** |  | [optional] 
**ClusterType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**ClusterGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, objectType string, display string, url string, naturalSlug string, name string, clusterType BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Cluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Cluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Cluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Cluster) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Cluster) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Cluster) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Cluster) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Cluster) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Cluster) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Cluster) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Cluster) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Cluster) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetDeviceCount

`func (o *Cluster) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Cluster) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Cluster) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *Cluster) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetVirtualmachineCount

`func (o *Cluster) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *Cluster) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *Cluster) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.

### HasVirtualmachineCount

`func (o *Cluster) HasVirtualmachineCount() bool`

HasVirtualmachineCount returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetComments

`func (o *Cluster) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Cluster) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Cluster) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Cluster) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetClusterType

`func (o *Cluster) GetClusterType() BulkWritableCableRequestStatus`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *Cluster) GetClusterTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *Cluster) SetClusterType(v BulkWritableCableRequestStatus)`

SetClusterType sets ClusterType field to given value.


### GetClusterGroup

`func (o *Cluster) GetClusterGroup() BulkWritableCircuitRequestTenant`

GetClusterGroup returns the ClusterGroup field if non-nil, zero value otherwise.

### GetClusterGroupOk

`func (o *Cluster) GetClusterGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetClusterGroupOk returns a tuple with the ClusterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroup

`func (o *Cluster) SetClusterGroup(v BulkWritableCircuitRequestTenant)`

SetClusterGroup sets ClusterGroup field to given value.

### HasClusterGroup

`func (o *Cluster) HasClusterGroup() bool`

HasClusterGroup returns a boolean if a field has been set.

### SetClusterGroupNil

`func (o *Cluster) SetClusterGroupNil(b bool)`

 SetClusterGroupNil sets the value for ClusterGroup to be an explicit nil

### UnsetClusterGroup
`func (o *Cluster) UnsetClusterGroup()`

UnsetClusterGroup ensures that no value is present for ClusterGroup, not even an explicit nil
### GetTenant

`func (o *Cluster) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Cluster) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Cluster) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Cluster) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Cluster) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Cluster) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLocation

`func (o *Cluster) GetLocation() BulkWritableCircuitRequestTenant`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Cluster) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Cluster) SetLocation(v BulkWritableCircuitRequestTenant)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Cluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Cluster) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Cluster) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetCreated

`func (o *Cluster) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Cluster) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Cluster) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Cluster) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Cluster) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Cluster) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Cluster) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Cluster) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Cluster) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Cluster) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *Cluster) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Cluster) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Cluster) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Cluster) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *Cluster) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Cluster) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Cluster) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Cluster) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Cluster) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Cluster) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Cluster) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


