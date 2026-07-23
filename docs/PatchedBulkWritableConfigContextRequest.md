# PatchedBulkWritableConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
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

### NewPatchedBulkWritableConfigContextRequest

`func NewPatchedBulkWritableConfigContextRequest(id string, ) *PatchedBulkWritableConfigContextRequest`

NewPatchedBulkWritableConfigContextRequest instantiates a new PatchedBulkWritableConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableConfigContextRequestWithDefaults

`func NewPatchedBulkWritableConfigContextRequestWithDefaults() *PatchedBulkWritableConfigContextRequest`

NewPatchedBulkWritableConfigContextRequestWithDefaults instantiates a new PatchedBulkWritableConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableConfigContextRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableConfigContextRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableConfigContextRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerContentType

`func (o *PatchedBulkWritableConfigContextRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *PatchedBulkWritableConfigContextRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *PatchedBulkWritableConfigContextRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *PatchedBulkWritableConfigContextRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *PatchedBulkWritableConfigContextRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *PatchedBulkWritableConfigContextRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetName

`func (o *PatchedBulkWritableConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerObjectId

`func (o *PatchedBulkWritableConfigContextRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *PatchedBulkWritableConfigContextRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *PatchedBulkWritableConfigContextRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *PatchedBulkWritableConfigContextRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *PatchedBulkWritableConfigContextRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *PatchedBulkWritableConfigContextRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetWeight

`func (o *PatchedBulkWritableConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedBulkWritableConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedBulkWritableConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedBulkWritableConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedBulkWritableConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetData

`func (o *PatchedBulkWritableConfigContextRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchedBulkWritableConfigContextRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchedBulkWritableConfigContextRequest) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PatchedBulkWritableConfigContextRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PatchedBulkWritableConfigContextRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PatchedBulkWritableConfigContextRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetConfigContextSchema

`func (o *PatchedBulkWritableConfigContextRequest) GetConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetConfigContextSchema returns the ConfigContextSchema field if non-nil, zero value otherwise.

### GetConfigContextSchemaOk

`func (o *PatchedBulkWritableConfigContextRequest) GetConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetConfigContextSchemaOk returns a tuple with the ConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContextSchema

`func (o *PatchedBulkWritableConfigContextRequest) SetConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetConfigContextSchema sets ConfigContextSchema field to given value.

### HasConfigContextSchema

`func (o *PatchedBulkWritableConfigContextRequest) HasConfigContextSchema() bool`

HasConfigContextSchema returns a boolean if a field has been set.

### SetConfigContextSchemaNil

`func (o *PatchedBulkWritableConfigContextRequest) SetConfigContextSchemaNil(b bool)`

 SetConfigContextSchemaNil sets the value for ConfigContextSchema to be an explicit nil

### UnsetConfigContextSchema
`func (o *PatchedBulkWritableConfigContextRequest) UnsetConfigContextSchema()`

UnsetConfigContextSchema ensures that no value is present for ConfigContextSchema, not even an explicit nil
### GetLocations

`func (o *PatchedBulkWritableConfigContextRequest) GetLocations() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetLocationsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *PatchedBulkWritableConfigContextRequest) SetLocations(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *PatchedBulkWritableConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRoles

`func (o *PatchedBulkWritableConfigContextRequest) GetRoles() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatchedBulkWritableConfigContextRequest) GetRolesOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatchedBulkWritableConfigContextRequest) SetRoles(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatchedBulkWritableConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *PatchedBulkWritableConfigContextRequest) GetDeviceTypes() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *PatchedBulkWritableConfigContextRequest) GetDeviceTypesOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *PatchedBulkWritableConfigContextRequest) SetDeviceTypes(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *PatchedBulkWritableConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetDeviceRedundancyGroups

`func (o *PatchedBulkWritableConfigContextRequest) GetDeviceRedundancyGroups() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceRedundancyGroups returns the DeviceRedundancyGroups field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetDeviceRedundancyGroupsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceRedundancyGroupsOk returns a tuple with the DeviceRedundancyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroups

`func (o *PatchedBulkWritableConfigContextRequest) SetDeviceRedundancyGroups(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceRedundancyGroups sets DeviceRedundancyGroups field to given value.

### HasDeviceRedundancyGroups

`func (o *PatchedBulkWritableConfigContextRequest) HasDeviceRedundancyGroups() bool`

HasDeviceRedundancyGroups returns a boolean if a field has been set.

### GetPlatforms

`func (o *PatchedBulkWritableConfigContextRequest) GetPlatforms() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetPlatformsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PatchedBulkWritableConfigContextRequest) SetPlatforms(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PatchedBulkWritableConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterGroups

`func (o *PatchedBulkWritableConfigContextRequest) GetClusterGroups() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetClusterGroupsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *PatchedBulkWritableConfigContextRequest) SetClusterGroups(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *PatchedBulkWritableConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *PatchedBulkWritableConfigContextRequest) GetClusters() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *PatchedBulkWritableConfigContextRequest) GetClustersOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *PatchedBulkWritableConfigContextRequest) SetClusters(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *PatchedBulkWritableConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *PatchedBulkWritableConfigContextRequest) GetTenantGroups() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetTenantGroupsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *PatchedBulkWritableConfigContextRequest) SetTenantGroups(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *PatchedBulkWritableConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *PatchedBulkWritableConfigContextRequest) GetTenants() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetTenantsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *PatchedBulkWritableConfigContextRequest) SetTenants(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *PatchedBulkWritableConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetDeviceFamilies

`func (o *PatchedBulkWritableConfigContextRequest) GetDeviceFamilies() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceFamilies returns the DeviceFamilies field if non-nil, zero value otherwise.

### GetDeviceFamiliesOk

`func (o *PatchedBulkWritableConfigContextRequest) GetDeviceFamiliesOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceFamiliesOk returns a tuple with the DeviceFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFamilies

`func (o *PatchedBulkWritableConfigContextRequest) SetDeviceFamilies(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceFamilies sets DeviceFamilies field to given value.

### HasDeviceFamilies

`func (o *PatchedBulkWritableConfigContextRequest) HasDeviceFamilies() bool`

HasDeviceFamilies returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableConfigContextRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableConfigContextRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableConfigContextRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


