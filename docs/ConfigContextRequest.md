# ConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Data** | **interface{}** |  | 
**ConfigContextSchema** | Pointer to [**NullableBulkWritableConfigContextRequestConfigContextSchema**](BulkWritableConfigContextRequestConfigContextSchema.md) |  | [optional] 
**Locations** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Roles** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**DeviceTypes** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**DeviceRedundancyGroups** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Platforms** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**ClusterGroups** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Clusters** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**TenantGroups** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Tenants** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**DeviceFamilies** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewConfigContextRequest

`func NewConfigContextRequest(name string, data interface{}, ) *ConfigContextRequest`

NewConfigContextRequest instantiates a new ConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextRequestWithDefaults

`func NewConfigContextRequestWithDefaults() *ConfigContextRequest`

NewConfigContextRequestWithDefaults instantiates a new ConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContextRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContextRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContextRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigContextRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerContentType

`func (o *ConfigContextRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *ConfigContextRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *ConfigContextRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *ConfigContextRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *ConfigContextRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *ConfigContextRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetName

`func (o *ConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerObjectId

`func (o *ConfigContextRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *ConfigContextRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *ConfigContextRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *ConfigContextRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *ConfigContextRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *ConfigContextRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetWeight

`func (o *ConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *ConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetData

`func (o *ConfigContextRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigContextRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigContextRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *ConfigContextRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ConfigContextRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetConfigContextSchema

`func (o *ConfigContextRequest) GetConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetConfigContextSchema returns the ConfigContextSchema field if non-nil, zero value otherwise.

### GetConfigContextSchemaOk

`func (o *ConfigContextRequest) GetConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetConfigContextSchemaOk returns a tuple with the ConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContextSchema

`func (o *ConfigContextRequest) SetConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetConfigContextSchema sets ConfigContextSchema field to given value.

### HasConfigContextSchema

`func (o *ConfigContextRequest) HasConfigContextSchema() bool`

HasConfigContextSchema returns a boolean if a field has been set.

### SetConfigContextSchemaNil

`func (o *ConfigContextRequest) SetConfigContextSchemaNil(b bool)`

 SetConfigContextSchemaNil sets the value for ConfigContextSchema to be an explicit nil

### UnsetConfigContextSchema
`func (o *ConfigContextRequest) UnsetConfigContextSchema()`

UnsetConfigContextSchema ensures that no value is present for ConfigContextSchema, not even an explicit nil
### GetLocations

`func (o *ConfigContextRequest) GetLocations() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ConfigContextRequest) GetLocationsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ConfigContextRequest) SetLocations(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRoles

`func (o *ConfigContextRequest) GetRoles() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ConfigContextRequest) GetRolesOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ConfigContextRequest) SetRoles(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *ConfigContextRequest) GetDeviceTypes() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *ConfigContextRequest) GetDeviceTypesOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *ConfigContextRequest) SetDeviceTypes(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *ConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetDeviceRedundancyGroups

`func (o *ConfigContextRequest) GetDeviceRedundancyGroups() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceRedundancyGroups returns the DeviceRedundancyGroups field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupsOk

`func (o *ConfigContextRequest) GetDeviceRedundancyGroupsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceRedundancyGroupsOk returns a tuple with the DeviceRedundancyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroups

`func (o *ConfigContextRequest) SetDeviceRedundancyGroups(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceRedundancyGroups sets DeviceRedundancyGroups field to given value.

### HasDeviceRedundancyGroups

`func (o *ConfigContextRequest) HasDeviceRedundancyGroups() bool`

HasDeviceRedundancyGroups returns a boolean if a field has been set.

### GetPlatforms

`func (o *ConfigContextRequest) GetPlatforms() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ConfigContextRequest) GetPlatformsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ConfigContextRequest) SetPlatforms(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterGroups

`func (o *ConfigContextRequest) GetClusterGroups() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *ConfigContextRequest) GetClusterGroupsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *ConfigContextRequest) SetClusterGroups(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *ConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *ConfigContextRequest) GetClusters() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ConfigContextRequest) GetClustersOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ConfigContextRequest) SetClusters(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *ConfigContextRequest) GetTenantGroups() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *ConfigContextRequest) GetTenantGroupsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *ConfigContextRequest) SetTenantGroups(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *ConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *ConfigContextRequest) GetTenants() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ConfigContextRequest) GetTenantsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ConfigContextRequest) SetTenants(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetDeviceFamilies

`func (o *ConfigContextRequest) GetDeviceFamilies() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceFamilies returns the DeviceFamilies field if non-nil, zero value otherwise.

### GetDeviceFamiliesOk

`func (o *ConfigContextRequest) GetDeviceFamiliesOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceFamiliesOk returns a tuple with the DeviceFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFamilies

`func (o *ConfigContextRequest) SetDeviceFamilies(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceFamilies sets DeviceFamilies field to given value.

### HasDeviceFamilies

`func (o *ConfigContextRequest) HasDeviceFamilies() bool`

HasDeviceFamilies returns a boolean if a field has been set.

### GetTags

`func (o *ConfigContextRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigContextRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigContextRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


