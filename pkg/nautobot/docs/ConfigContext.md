# ConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Owner** | [**NullableConfigContextOwner**](ConfigContextOwner.md) |  | [readonly] 
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
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewConfigContext

`func NewConfigContext(id string, objectType string, display string, url string, naturalSlug string, owner NullableConfigContextOwner, name string, data map[string]interface{}, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ConfigContext`

NewConfigContext instantiates a new ConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextWithDefaults

`func NewConfigContextWithDefaults() *ConfigContext`

NewConfigContextWithDefaults instantiates a new ConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContext) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *ConfigContext) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConfigContext) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConfigContext) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ConfigContext) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ConfigContext) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ConfigContext) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetOwnerContentType

`func (o *ConfigContext) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *ConfigContext) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *ConfigContext) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *ConfigContext) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *ConfigContext) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *ConfigContext) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetOwner

`func (o *ConfigContext) GetOwner() ConfigContextOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConfigContext) GetOwnerOk() (*ConfigContextOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConfigContext) SetOwner(v ConfigContextOwner)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *ConfigContext) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ConfigContext) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetName

`func (o *ConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContext) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerObjectId

`func (o *ConfigContext) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *ConfigContext) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *ConfigContext) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *ConfigContext) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *ConfigContext) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *ConfigContext) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetWeight

`func (o *ConfigContext) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ConfigContext) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ConfigContext) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ConfigContext) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *ConfigContext) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConfigContext) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConfigContext) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConfigContext) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetData

`func (o *ConfigContext) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigContext) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigContext) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetConfigContextSchema

`func (o *ConfigContext) GetConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetConfigContextSchema returns the ConfigContextSchema field if non-nil, zero value otherwise.

### GetConfigContextSchemaOk

`func (o *ConfigContext) GetConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetConfigContextSchemaOk returns a tuple with the ConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContextSchema

`func (o *ConfigContext) SetConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetConfigContextSchema sets ConfigContextSchema field to given value.

### HasConfigContextSchema

`func (o *ConfigContext) HasConfigContextSchema() bool`

HasConfigContextSchema returns a boolean if a field has been set.

### SetConfigContextSchemaNil

`func (o *ConfigContext) SetConfigContextSchemaNil(b bool)`

 SetConfigContextSchemaNil sets the value for ConfigContextSchema to be an explicit nil

### UnsetConfigContextSchema
`func (o *ConfigContext) UnsetConfigContextSchema()`

UnsetConfigContextSchema ensures that no value is present for ConfigContextSchema, not even an explicit nil
### GetLocations

`func (o *ConfigContext) GetLocations() []BulkWritableCableRequestStatus`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ConfigContext) GetLocationsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ConfigContext) SetLocations(v []BulkWritableCableRequestStatus)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ConfigContext) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRoles

`func (o *ConfigContext) GetRoles() []BulkWritableCableRequestStatus`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ConfigContext) GetRolesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ConfigContext) SetRoles(v []BulkWritableCableRequestStatus)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ConfigContext) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *ConfigContext) GetDeviceTypes() []BulkWritableCableRequestStatus`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *ConfigContext) GetDeviceTypesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *ConfigContext) SetDeviceTypes(v []BulkWritableCableRequestStatus)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *ConfigContext) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetDeviceRedundancyGroups

`func (o *ConfigContext) GetDeviceRedundancyGroups() []BulkWritableCableRequestStatus`

GetDeviceRedundancyGroups returns the DeviceRedundancyGroups field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupsOk

`func (o *ConfigContext) GetDeviceRedundancyGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetDeviceRedundancyGroupsOk returns a tuple with the DeviceRedundancyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroups

`func (o *ConfigContext) SetDeviceRedundancyGroups(v []BulkWritableCableRequestStatus)`

SetDeviceRedundancyGroups sets DeviceRedundancyGroups field to given value.

### HasDeviceRedundancyGroups

`func (o *ConfigContext) HasDeviceRedundancyGroups() bool`

HasDeviceRedundancyGroups returns a boolean if a field has been set.

### GetPlatforms

`func (o *ConfigContext) GetPlatforms() []BulkWritableCableRequestStatus`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ConfigContext) GetPlatformsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ConfigContext) SetPlatforms(v []BulkWritableCableRequestStatus)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ConfigContext) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterGroups

`func (o *ConfigContext) GetClusterGroups() []BulkWritableCableRequestStatus`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *ConfigContext) GetClusterGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *ConfigContext) SetClusterGroups(v []BulkWritableCableRequestStatus)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *ConfigContext) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *ConfigContext) GetClusters() []BulkWritableCableRequestStatus`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ConfigContext) GetClustersOk() (*[]BulkWritableCableRequestStatus, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ConfigContext) SetClusters(v []BulkWritableCableRequestStatus)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ConfigContext) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *ConfigContext) GetTenantGroups() []BulkWritableCableRequestStatus`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *ConfigContext) GetTenantGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *ConfigContext) SetTenantGroups(v []BulkWritableCableRequestStatus)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *ConfigContext) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *ConfigContext) GetTenants() []BulkWritableCableRequestStatus`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ConfigContext) GetTenantsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ConfigContext) SetTenants(v []BulkWritableCableRequestStatus)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ConfigContext) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetCreated

`func (o *ConfigContext) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConfigContext) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConfigContext) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ConfigContext) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ConfigContext) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ConfigContext) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ConfigContext) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *ConfigContext) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ConfigContext) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ConfigContext) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetTags

`func (o *ConfigContext) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigContext) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigContext) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


