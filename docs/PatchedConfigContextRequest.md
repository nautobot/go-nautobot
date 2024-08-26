# PatchedConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
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

### NewPatchedConfigContextRequest

`func NewPatchedConfigContextRequest() *PatchedConfigContextRequest`

NewPatchedConfigContextRequest instantiates a new PatchedConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConfigContextRequestWithDefaults

`func NewPatchedConfigContextRequestWithDefaults() *PatchedConfigContextRequest`

NewPatchedConfigContextRequestWithDefaults instantiates a new PatchedConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerContentType

`func (o *PatchedConfigContextRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *PatchedConfigContextRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *PatchedConfigContextRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *PatchedConfigContextRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *PatchedConfigContextRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *PatchedConfigContextRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetName

`func (o *PatchedConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerObjectId

`func (o *PatchedConfigContextRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *PatchedConfigContextRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *PatchedConfigContextRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *PatchedConfigContextRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *PatchedConfigContextRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *PatchedConfigContextRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetWeight

`func (o *PatchedConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetData

`func (o *PatchedConfigContextRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchedConfigContextRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchedConfigContextRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PatchedConfigContextRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetConfigContextSchema

`func (o *PatchedConfigContextRequest) GetConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetConfigContextSchema returns the ConfigContextSchema field if non-nil, zero value otherwise.

### GetConfigContextSchemaOk

`func (o *PatchedConfigContextRequest) GetConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetConfigContextSchemaOk returns a tuple with the ConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContextSchema

`func (o *PatchedConfigContextRequest) SetConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetConfigContextSchema sets ConfigContextSchema field to given value.

### HasConfigContextSchema

`func (o *PatchedConfigContextRequest) HasConfigContextSchema() bool`

HasConfigContextSchema returns a boolean if a field has been set.

### SetConfigContextSchemaNil

`func (o *PatchedConfigContextRequest) SetConfigContextSchemaNil(b bool)`

 SetConfigContextSchemaNil sets the value for ConfigContextSchema to be an explicit nil

### UnsetConfigContextSchema
`func (o *PatchedConfigContextRequest) UnsetConfigContextSchema()`

UnsetConfigContextSchema ensures that no value is present for ConfigContextSchema, not even an explicit nil
### GetLocations

`func (o *PatchedConfigContextRequest) GetLocations() []BulkWritableCableRequestStatus`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *PatchedConfigContextRequest) GetLocationsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *PatchedConfigContextRequest) SetLocations(v []BulkWritableCableRequestStatus)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *PatchedConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRoles

`func (o *PatchedConfigContextRequest) GetRoles() []BulkWritableCableRequestStatus`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatchedConfigContextRequest) GetRolesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatchedConfigContextRequest) SetRoles(v []BulkWritableCableRequestStatus)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatchedConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *PatchedConfigContextRequest) GetDeviceTypes() []BulkWritableCableRequestStatus`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *PatchedConfigContextRequest) GetDeviceTypesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *PatchedConfigContextRequest) SetDeviceTypes(v []BulkWritableCableRequestStatus)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *PatchedConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetDeviceRedundancyGroups

`func (o *PatchedConfigContextRequest) GetDeviceRedundancyGroups() []BulkWritableCableRequestStatus`

GetDeviceRedundancyGroups returns the DeviceRedundancyGroups field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupsOk

`func (o *PatchedConfigContextRequest) GetDeviceRedundancyGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetDeviceRedundancyGroupsOk returns a tuple with the DeviceRedundancyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroups

`func (o *PatchedConfigContextRequest) SetDeviceRedundancyGroups(v []BulkWritableCableRequestStatus)`

SetDeviceRedundancyGroups sets DeviceRedundancyGroups field to given value.

### HasDeviceRedundancyGroups

`func (o *PatchedConfigContextRequest) HasDeviceRedundancyGroups() bool`

HasDeviceRedundancyGroups returns a boolean if a field has been set.

### GetPlatforms

`func (o *PatchedConfigContextRequest) GetPlatforms() []BulkWritableCableRequestStatus`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PatchedConfigContextRequest) GetPlatformsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PatchedConfigContextRequest) SetPlatforms(v []BulkWritableCableRequestStatus)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PatchedConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterGroups

`func (o *PatchedConfigContextRequest) GetClusterGroups() []BulkWritableCableRequestStatus`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *PatchedConfigContextRequest) GetClusterGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *PatchedConfigContextRequest) SetClusterGroups(v []BulkWritableCableRequestStatus)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *PatchedConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *PatchedConfigContextRequest) GetClusters() []BulkWritableCableRequestStatus`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *PatchedConfigContextRequest) GetClustersOk() (*[]BulkWritableCableRequestStatus, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *PatchedConfigContextRequest) SetClusters(v []BulkWritableCableRequestStatus)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *PatchedConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *PatchedConfigContextRequest) GetTenantGroups() []BulkWritableCableRequestStatus`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *PatchedConfigContextRequest) GetTenantGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *PatchedConfigContextRequest) SetTenantGroups(v []BulkWritableCableRequestStatus)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *PatchedConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *PatchedConfigContextRequest) GetTenants() []BulkWritableCableRequestStatus`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *PatchedConfigContextRequest) GetTenantsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *PatchedConfigContextRequest) SetTenants(v []BulkWritableCableRequestStatus)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *PatchedConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTags

`func (o *PatchedConfigContextRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedConfigContextRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedConfigContextRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


