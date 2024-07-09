# VirtualMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalConfigContextData** | Pointer to **map[string]interface{}** |  | [optional] 
**LocalConfigContextDataOwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Vcpus** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Disk** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalConfigContextSchema** | Pointer to [**NullableBulkWritableConfigContextRequestConfigContextSchema**](BulkWritableConfigContextRequestConfigContextSchema.md) |  | [optional] 
**LocalConfigContextDataOwnerContentType** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Cluster** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullablePrimaryIPv4**](PrimaryIPv4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullablePrimaryIPv6**](PrimaryIPv6.md) |  | [optional] 
**SoftwareVersion** | Pointer to [**NullableBulkWritableVirtualMachineRequestSoftwareVersion**](BulkWritableVirtualMachineRequestSoftwareVersion.md) |  | [optional] 
**SoftwareImageFiles** | Pointer to [**[]SoftwareImageFiles**](SoftwareImageFiles.md) | Override the software image files associated with the software version for this virtual machine | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewVirtualMachineRequest

`func NewVirtualMachineRequest(name string, cluster BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *VirtualMachineRequest`

NewVirtualMachineRequest instantiates a new VirtualMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineRequestWithDefaults

`func NewVirtualMachineRequestWithDefaults() *VirtualMachineRequest`

NewVirtualMachineRequestWithDefaults instantiates a new VirtualMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalConfigContextData

`func (o *VirtualMachineRequest) GetLocalConfigContextData() map[string]interface{}`

GetLocalConfigContextData returns the LocalConfigContextData field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOk

`func (o *VirtualMachineRequest) GetLocalConfigContextDataOk() (*map[string]interface{}, bool)`

GetLocalConfigContextDataOk returns a tuple with the LocalConfigContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextData

`func (o *VirtualMachineRequest) SetLocalConfigContextData(v map[string]interface{})`

SetLocalConfigContextData sets LocalConfigContextData field to given value.

### HasLocalConfigContextData

`func (o *VirtualMachineRequest) HasLocalConfigContextData() bool`

HasLocalConfigContextData returns a boolean if a field has been set.

### SetLocalConfigContextDataNil

`func (o *VirtualMachineRequest) SetLocalConfigContextDataNil(b bool)`

 SetLocalConfigContextDataNil sets the value for LocalConfigContextData to be an explicit nil

### UnsetLocalConfigContextData
`func (o *VirtualMachineRequest) UnsetLocalConfigContextData()`

UnsetLocalConfigContextData ensures that no value is present for LocalConfigContextData, not even an explicit nil
### GetLocalConfigContextDataOwnerObjectId

`func (o *VirtualMachineRequest) GetLocalConfigContextDataOwnerObjectId() string`

GetLocalConfigContextDataOwnerObjectId returns the LocalConfigContextDataOwnerObjectId field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerObjectIdOk

`func (o *VirtualMachineRequest) GetLocalConfigContextDataOwnerObjectIdOk() (*string, bool)`

GetLocalConfigContextDataOwnerObjectIdOk returns a tuple with the LocalConfigContextDataOwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerObjectId

`func (o *VirtualMachineRequest) SetLocalConfigContextDataOwnerObjectId(v string)`

SetLocalConfigContextDataOwnerObjectId sets LocalConfigContextDataOwnerObjectId field to given value.

### HasLocalConfigContextDataOwnerObjectId

`func (o *VirtualMachineRequest) HasLocalConfigContextDataOwnerObjectId() bool`

HasLocalConfigContextDataOwnerObjectId returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerObjectIdNil

`func (o *VirtualMachineRequest) SetLocalConfigContextDataOwnerObjectIdNil(b bool)`

 SetLocalConfigContextDataOwnerObjectIdNil sets the value for LocalConfigContextDataOwnerObjectId to be an explicit nil

### UnsetLocalConfigContextDataOwnerObjectId
`func (o *VirtualMachineRequest) UnsetLocalConfigContextDataOwnerObjectId()`

UnsetLocalConfigContextDataOwnerObjectId ensures that no value is present for LocalConfigContextDataOwnerObjectId, not even an explicit nil
### GetName

`func (o *VirtualMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVcpus

`func (o *VirtualMachineRequest) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VirtualMachineRequest) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VirtualMachineRequest) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VirtualMachineRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *VirtualMachineRequest) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *VirtualMachineRequest) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *VirtualMachineRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VirtualMachineRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VirtualMachineRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VirtualMachineRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *VirtualMachineRequest) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *VirtualMachineRequest) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *VirtualMachineRequest) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualMachineRequest) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualMachineRequest) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualMachineRequest) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *VirtualMachineRequest) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *VirtualMachineRequest) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetComments

`func (o *VirtualMachineRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualMachineRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualMachineRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualMachineRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalConfigContextSchema

`func (o *VirtualMachineRequest) GetLocalConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetLocalConfigContextSchema returns the LocalConfigContextSchema field if non-nil, zero value otherwise.

### GetLocalConfigContextSchemaOk

`func (o *VirtualMachineRequest) GetLocalConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetLocalConfigContextSchemaOk returns a tuple with the LocalConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextSchema

`func (o *VirtualMachineRequest) SetLocalConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetLocalConfigContextSchema sets LocalConfigContextSchema field to given value.

### HasLocalConfigContextSchema

`func (o *VirtualMachineRequest) HasLocalConfigContextSchema() bool`

HasLocalConfigContextSchema returns a boolean if a field has been set.

### SetLocalConfigContextSchemaNil

`func (o *VirtualMachineRequest) SetLocalConfigContextSchemaNil(b bool)`

 SetLocalConfigContextSchemaNil sets the value for LocalConfigContextSchema to be an explicit nil

### UnsetLocalConfigContextSchema
`func (o *VirtualMachineRequest) UnsetLocalConfigContextSchema()`

UnsetLocalConfigContextSchema ensures that no value is present for LocalConfigContextSchema, not even an explicit nil
### GetLocalConfigContextDataOwnerContentType

`func (o *VirtualMachineRequest) GetLocalConfigContextDataOwnerContentType() BulkWritableCircuitRequestTenant`

GetLocalConfigContextDataOwnerContentType returns the LocalConfigContextDataOwnerContentType field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerContentTypeOk

`func (o *VirtualMachineRequest) GetLocalConfigContextDataOwnerContentTypeOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocalConfigContextDataOwnerContentTypeOk returns a tuple with the LocalConfigContextDataOwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerContentType

`func (o *VirtualMachineRequest) SetLocalConfigContextDataOwnerContentType(v BulkWritableCircuitRequestTenant)`

SetLocalConfigContextDataOwnerContentType sets LocalConfigContextDataOwnerContentType field to given value.

### HasLocalConfigContextDataOwnerContentType

`func (o *VirtualMachineRequest) HasLocalConfigContextDataOwnerContentType() bool`

HasLocalConfigContextDataOwnerContentType returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerContentTypeNil

`func (o *VirtualMachineRequest) SetLocalConfigContextDataOwnerContentTypeNil(b bool)`

 SetLocalConfigContextDataOwnerContentTypeNil sets the value for LocalConfigContextDataOwnerContentType to be an explicit nil

### UnsetLocalConfigContextDataOwnerContentType
`func (o *VirtualMachineRequest) UnsetLocalConfigContextDataOwnerContentType()`

UnsetLocalConfigContextDataOwnerContentType ensures that no value is present for LocalConfigContextDataOwnerContentType, not even an explicit nil
### GetCluster

`func (o *VirtualMachineRequest) GetCluster() BulkWritableCableRequestStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualMachineRequest) GetClusterOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualMachineRequest) SetCluster(v BulkWritableCableRequestStatus)`

SetCluster sets Cluster field to given value.


### GetTenant

`func (o *VirtualMachineRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualMachineRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualMachineRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualMachineRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualMachineRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualMachineRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *VirtualMachineRequest) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VirtualMachineRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VirtualMachineRequest) SetPlatform(v BulkWritableCircuitRequestTenant)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VirtualMachineRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VirtualMachineRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VirtualMachineRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetStatus

`func (o *VirtualMachineRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualMachineRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualMachineRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *VirtualMachineRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualMachineRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualMachineRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualMachineRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VirtualMachineRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VirtualMachineRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPrimaryIp4

`func (o *VirtualMachineRequest) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *VirtualMachineRequest) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *VirtualMachineRequest) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *VirtualMachineRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *VirtualMachineRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *VirtualMachineRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *VirtualMachineRequest) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *VirtualMachineRequest) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *VirtualMachineRequest) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *VirtualMachineRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *VirtualMachineRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *VirtualMachineRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetSoftwareVersion

`func (o *VirtualMachineRequest) GetSoftwareVersion() BulkWritableVirtualMachineRequestSoftwareVersion`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *VirtualMachineRequest) GetSoftwareVersionOk() (*BulkWritableVirtualMachineRequestSoftwareVersion, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *VirtualMachineRequest) SetSoftwareVersion(v BulkWritableVirtualMachineRequestSoftwareVersion)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *VirtualMachineRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *VirtualMachineRequest) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *VirtualMachineRequest) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetSoftwareImageFiles

`func (o *VirtualMachineRequest) GetSoftwareImageFiles() []SoftwareImageFiles`

GetSoftwareImageFiles returns the SoftwareImageFiles field if non-nil, zero value otherwise.

### GetSoftwareImageFilesOk

`func (o *VirtualMachineRequest) GetSoftwareImageFilesOk() (*[]SoftwareImageFiles, bool)`

GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFiles

`func (o *VirtualMachineRequest) SetSoftwareImageFiles(v []SoftwareImageFiles)`

SetSoftwareImageFiles sets SoftwareImageFiles field to given value.

### HasSoftwareImageFiles

`func (o *VirtualMachineRequest) HasSoftwareImageFiles() bool`

HasSoftwareImageFiles returns a boolean if a field has been set.

### GetTags

`func (o *VirtualMachineRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualMachineRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualMachineRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualMachineRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualMachineRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualMachineRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualMachineRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualMachineRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VirtualMachineRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VirtualMachineRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VirtualMachineRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VirtualMachineRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


