# BulkWritableVirtualMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableVirtualMachineRequest

`func NewBulkWritableVirtualMachineRequest(id string, name string, cluster BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *BulkWritableVirtualMachineRequest`

NewBulkWritableVirtualMachineRequest instantiates a new BulkWritableVirtualMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVirtualMachineRequestWithDefaults

`func NewBulkWritableVirtualMachineRequestWithDefaults() *BulkWritableVirtualMachineRequest`

NewBulkWritableVirtualMachineRequestWithDefaults instantiates a new BulkWritableVirtualMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVirtualMachineRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVirtualMachineRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVirtualMachineRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLocalConfigContextData

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextData() map[string]interface{}`

GetLocalConfigContextData returns the LocalConfigContextData field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOk

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextDataOk() (*map[string]interface{}, bool)`

GetLocalConfigContextDataOk returns a tuple with the LocalConfigContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextData

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextData(v map[string]interface{})`

SetLocalConfigContextData sets LocalConfigContextData field to given value.

### HasLocalConfigContextData

`func (o *BulkWritableVirtualMachineRequest) HasLocalConfigContextData() bool`

HasLocalConfigContextData returns a boolean if a field has been set.

### SetLocalConfigContextDataNil

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextDataNil(b bool)`

 SetLocalConfigContextDataNil sets the value for LocalConfigContextData to be an explicit nil

### UnsetLocalConfigContextData
`func (o *BulkWritableVirtualMachineRequest) UnsetLocalConfigContextData()`

UnsetLocalConfigContextData ensures that no value is present for LocalConfigContextData, not even an explicit nil
### GetLocalConfigContextDataOwnerObjectId

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerObjectId() string`

GetLocalConfigContextDataOwnerObjectId returns the LocalConfigContextDataOwnerObjectId field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerObjectIdOk

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerObjectIdOk() (*string, bool)`

GetLocalConfigContextDataOwnerObjectIdOk returns a tuple with the LocalConfigContextDataOwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerObjectId

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerObjectId(v string)`

SetLocalConfigContextDataOwnerObjectId sets LocalConfigContextDataOwnerObjectId field to given value.

### HasLocalConfigContextDataOwnerObjectId

`func (o *BulkWritableVirtualMachineRequest) HasLocalConfigContextDataOwnerObjectId() bool`

HasLocalConfigContextDataOwnerObjectId returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerObjectIdNil

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerObjectIdNil(b bool)`

 SetLocalConfigContextDataOwnerObjectIdNil sets the value for LocalConfigContextDataOwnerObjectId to be an explicit nil

### UnsetLocalConfigContextDataOwnerObjectId
`func (o *BulkWritableVirtualMachineRequest) UnsetLocalConfigContextDataOwnerObjectId()`

UnsetLocalConfigContextDataOwnerObjectId ensures that no value is present for LocalConfigContextDataOwnerObjectId, not even an explicit nil
### GetName

`func (o *BulkWritableVirtualMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVirtualMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVirtualMachineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVcpus

`func (o *BulkWritableVirtualMachineRequest) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *BulkWritableVirtualMachineRequest) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *BulkWritableVirtualMachineRequest) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *BulkWritableVirtualMachineRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *BulkWritableVirtualMachineRequest) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *BulkWritableVirtualMachineRequest) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *BulkWritableVirtualMachineRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *BulkWritableVirtualMachineRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *BulkWritableVirtualMachineRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *BulkWritableVirtualMachineRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *BulkWritableVirtualMachineRequest) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *BulkWritableVirtualMachineRequest) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *BulkWritableVirtualMachineRequest) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *BulkWritableVirtualMachineRequest) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *BulkWritableVirtualMachineRequest) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *BulkWritableVirtualMachineRequest) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *BulkWritableVirtualMachineRequest) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *BulkWritableVirtualMachineRequest) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetComments

`func (o *BulkWritableVirtualMachineRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableVirtualMachineRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableVirtualMachineRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableVirtualMachineRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalConfigContextSchema

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetLocalConfigContextSchema returns the LocalConfigContextSchema field if non-nil, zero value otherwise.

### GetLocalConfigContextSchemaOk

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetLocalConfigContextSchemaOk returns a tuple with the LocalConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextSchema

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetLocalConfigContextSchema sets LocalConfigContextSchema field to given value.

### HasLocalConfigContextSchema

`func (o *BulkWritableVirtualMachineRequest) HasLocalConfigContextSchema() bool`

HasLocalConfigContextSchema returns a boolean if a field has been set.

### SetLocalConfigContextSchemaNil

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextSchemaNil(b bool)`

 SetLocalConfigContextSchemaNil sets the value for LocalConfigContextSchema to be an explicit nil

### UnsetLocalConfigContextSchema
`func (o *BulkWritableVirtualMachineRequest) UnsetLocalConfigContextSchema()`

UnsetLocalConfigContextSchema ensures that no value is present for LocalConfigContextSchema, not even an explicit nil
### GetLocalConfigContextDataOwnerContentType

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerContentType() BulkWritableCircuitRequestTenant`

GetLocalConfigContextDataOwnerContentType returns the LocalConfigContextDataOwnerContentType field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerContentTypeOk

`func (o *BulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerContentTypeOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocalConfigContextDataOwnerContentTypeOk returns a tuple with the LocalConfigContextDataOwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerContentType

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerContentType(v BulkWritableCircuitRequestTenant)`

SetLocalConfigContextDataOwnerContentType sets LocalConfigContextDataOwnerContentType field to given value.

### HasLocalConfigContextDataOwnerContentType

`func (o *BulkWritableVirtualMachineRequest) HasLocalConfigContextDataOwnerContentType() bool`

HasLocalConfigContextDataOwnerContentType returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerContentTypeNil

`func (o *BulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerContentTypeNil(b bool)`

 SetLocalConfigContextDataOwnerContentTypeNil sets the value for LocalConfigContextDataOwnerContentType to be an explicit nil

### UnsetLocalConfigContextDataOwnerContentType
`func (o *BulkWritableVirtualMachineRequest) UnsetLocalConfigContextDataOwnerContentType()`

UnsetLocalConfigContextDataOwnerContentType ensures that no value is present for LocalConfigContextDataOwnerContentType, not even an explicit nil
### GetCluster

`func (o *BulkWritableVirtualMachineRequest) GetCluster() BulkWritableCableRequestStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BulkWritableVirtualMachineRequest) GetClusterOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BulkWritableVirtualMachineRequest) SetCluster(v BulkWritableCableRequestStatus)`

SetCluster sets Cluster field to given value.


### GetTenant

`func (o *BulkWritableVirtualMachineRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVirtualMachineRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVirtualMachineRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVirtualMachineRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVirtualMachineRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVirtualMachineRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *BulkWritableVirtualMachineRequest) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BulkWritableVirtualMachineRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BulkWritableVirtualMachineRequest) SetPlatform(v BulkWritableCircuitRequestTenant)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *BulkWritableVirtualMachineRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *BulkWritableVirtualMachineRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *BulkWritableVirtualMachineRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetStatus

`func (o *BulkWritableVirtualMachineRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableVirtualMachineRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableVirtualMachineRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *BulkWritableVirtualMachineRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkWritableVirtualMachineRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkWritableVirtualMachineRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkWritableVirtualMachineRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkWritableVirtualMachineRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkWritableVirtualMachineRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPrimaryIp4

`func (o *BulkWritableVirtualMachineRequest) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *BulkWritableVirtualMachineRequest) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *BulkWritableVirtualMachineRequest) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *BulkWritableVirtualMachineRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *BulkWritableVirtualMachineRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *BulkWritableVirtualMachineRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *BulkWritableVirtualMachineRequest) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *BulkWritableVirtualMachineRequest) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *BulkWritableVirtualMachineRequest) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *BulkWritableVirtualMachineRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *BulkWritableVirtualMachineRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *BulkWritableVirtualMachineRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetSoftwareVersion

`func (o *BulkWritableVirtualMachineRequest) GetSoftwareVersion() BulkWritableVirtualMachineRequestSoftwareVersion`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *BulkWritableVirtualMachineRequest) GetSoftwareVersionOk() (*BulkWritableVirtualMachineRequestSoftwareVersion, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *BulkWritableVirtualMachineRequest) SetSoftwareVersion(v BulkWritableVirtualMachineRequestSoftwareVersion)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *BulkWritableVirtualMachineRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *BulkWritableVirtualMachineRequest) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *BulkWritableVirtualMachineRequest) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetSoftwareImageFiles

`func (o *BulkWritableVirtualMachineRequest) GetSoftwareImageFiles() []SoftwareImageFiles`

GetSoftwareImageFiles returns the SoftwareImageFiles field if non-nil, zero value otherwise.

### GetSoftwareImageFilesOk

`func (o *BulkWritableVirtualMachineRequest) GetSoftwareImageFilesOk() (*[]SoftwareImageFiles, bool)`

GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFiles

`func (o *BulkWritableVirtualMachineRequest) SetSoftwareImageFiles(v []SoftwareImageFiles)`

SetSoftwareImageFiles sets SoftwareImageFiles field to given value.

### HasSoftwareImageFiles

`func (o *BulkWritableVirtualMachineRequest) HasSoftwareImageFiles() bool`

HasSoftwareImageFiles returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVirtualMachineRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVirtualMachineRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVirtualMachineRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVirtualMachineRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableVirtualMachineRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVirtualMachineRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVirtualMachineRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVirtualMachineRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVirtualMachineRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVirtualMachineRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVirtualMachineRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVirtualMachineRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


