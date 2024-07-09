# BulkWritableConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Data** | **map[string]interface{}** |  | 
**ConfigContextSchema** | Pointer to [**NullableBulkWritableConfigContextRequestConfigContextSchema**](BulkWritableConfigContextRequestConfigContextSchema.md) |  | [optional] 
**Locations** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Roles** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**DeviceTypes** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**DeviceRedundancyGroups** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Platforms** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ClusterGroups** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Clusters** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**TenantGroups** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tenants** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableConfigContextRequest

`func NewBulkWritableConfigContextRequest(id string, name string, data map[string]interface{}, ) *BulkWritableConfigContextRequest`

NewBulkWritableConfigContextRequest instantiates a new BulkWritableConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableConfigContextRequestWithDefaults

`func NewBulkWritableConfigContextRequestWithDefaults() *BulkWritableConfigContextRequest`

NewBulkWritableConfigContextRequestWithDefaults instantiates a new BulkWritableConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableConfigContextRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableConfigContextRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableConfigContextRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerContentType

`func (o *BulkWritableConfigContextRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *BulkWritableConfigContextRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *BulkWritableConfigContextRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *BulkWritableConfigContextRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *BulkWritableConfigContextRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *BulkWritableConfigContextRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetName

`func (o *BulkWritableConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerObjectId

`func (o *BulkWritableConfigContextRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *BulkWritableConfigContextRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *BulkWritableConfigContextRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *BulkWritableConfigContextRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *BulkWritableConfigContextRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *BulkWritableConfigContextRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetWeight

`func (o *BulkWritableConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkWritableConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkWritableConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkWritableConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritableConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *BulkWritableConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BulkWritableConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BulkWritableConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BulkWritableConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetData

`func (o *BulkWritableConfigContextRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkWritableConfigContextRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkWritableConfigContextRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetConfigContextSchema

`func (o *BulkWritableConfigContextRequest) GetConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetConfigContextSchema returns the ConfigContextSchema field if non-nil, zero value otherwise.

### GetConfigContextSchemaOk

`func (o *BulkWritableConfigContextRequest) GetConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetConfigContextSchemaOk returns a tuple with the ConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContextSchema

`func (o *BulkWritableConfigContextRequest) SetConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetConfigContextSchema sets ConfigContextSchema field to given value.

### HasConfigContextSchema

`func (o *BulkWritableConfigContextRequest) HasConfigContextSchema() bool`

HasConfigContextSchema returns a boolean if a field has been set.

### SetConfigContextSchemaNil

`func (o *BulkWritableConfigContextRequest) SetConfigContextSchemaNil(b bool)`

 SetConfigContextSchemaNil sets the value for ConfigContextSchema to be an explicit nil

### UnsetConfigContextSchema
`func (o *BulkWritableConfigContextRequest) UnsetConfigContextSchema()`

UnsetConfigContextSchema ensures that no value is present for ConfigContextSchema, not even an explicit nil
### GetLocations

`func (o *BulkWritableConfigContextRequest) GetLocations() []BulkWritableCableRequestStatus`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *BulkWritableConfigContextRequest) GetLocationsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *BulkWritableConfigContextRequest) SetLocations(v []BulkWritableCableRequestStatus)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *BulkWritableConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRoles

`func (o *BulkWritableConfigContextRequest) GetRoles() []BulkWritableCableRequestStatus`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BulkWritableConfigContextRequest) GetRolesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BulkWritableConfigContextRequest) SetRoles(v []BulkWritableCableRequestStatus)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BulkWritableConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *BulkWritableConfigContextRequest) GetDeviceTypes() []BulkWritableCableRequestStatus`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *BulkWritableConfigContextRequest) GetDeviceTypesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *BulkWritableConfigContextRequest) SetDeviceTypes(v []BulkWritableCableRequestStatus)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *BulkWritableConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetDeviceRedundancyGroups

`func (o *BulkWritableConfigContextRequest) GetDeviceRedundancyGroups() []BulkWritableCableRequestStatus`

GetDeviceRedundancyGroups returns the DeviceRedundancyGroups field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupsOk

`func (o *BulkWritableConfigContextRequest) GetDeviceRedundancyGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetDeviceRedundancyGroupsOk returns a tuple with the DeviceRedundancyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroups

`func (o *BulkWritableConfigContextRequest) SetDeviceRedundancyGroups(v []BulkWritableCableRequestStatus)`

SetDeviceRedundancyGroups sets DeviceRedundancyGroups field to given value.

### HasDeviceRedundancyGroups

`func (o *BulkWritableConfigContextRequest) HasDeviceRedundancyGroups() bool`

HasDeviceRedundancyGroups returns a boolean if a field has been set.

### GetPlatforms

`func (o *BulkWritableConfigContextRequest) GetPlatforms() []BulkWritableCableRequestStatus`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *BulkWritableConfigContextRequest) GetPlatformsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *BulkWritableConfigContextRequest) SetPlatforms(v []BulkWritableCableRequestStatus)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *BulkWritableConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterGroups

`func (o *BulkWritableConfigContextRequest) GetClusterGroups() []BulkWritableCableRequestStatus`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *BulkWritableConfigContextRequest) GetClusterGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *BulkWritableConfigContextRequest) SetClusterGroups(v []BulkWritableCableRequestStatus)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *BulkWritableConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *BulkWritableConfigContextRequest) GetClusters() []BulkWritableCableRequestStatus`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *BulkWritableConfigContextRequest) GetClustersOk() (*[]BulkWritableCableRequestStatus, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *BulkWritableConfigContextRequest) SetClusters(v []BulkWritableCableRequestStatus)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *BulkWritableConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *BulkWritableConfigContextRequest) GetTenantGroups() []BulkWritableCableRequestStatus`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *BulkWritableConfigContextRequest) GetTenantGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *BulkWritableConfigContextRequest) SetTenantGroups(v []BulkWritableCableRequestStatus)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *BulkWritableConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *BulkWritableConfigContextRequest) GetTenants() []BulkWritableCableRequestStatus`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *BulkWritableConfigContextRequest) GetTenantsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *BulkWritableConfigContextRequest) SetTenants(v []BulkWritableCableRequestStatus)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *BulkWritableConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableConfigContextRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableConfigContextRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableConfigContextRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


