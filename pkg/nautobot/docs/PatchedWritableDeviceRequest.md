# PatchedWritableDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalConfigContextData** | Pointer to **map[string]interface{}** |  | [optional] 
**LocalConfigContextDataOwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Position** | Pointer to **NullableInt32** | The lowest-numbered unit occupied by the device | [optional] 
**Face** | Pointer to [**RackFace**](RackFace.md) |  | [optional] 
**DeviceRedundancyGroupPriority** | Pointer to **NullableInt32** | The priority the device has in the device redundancy group. | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalConfigContextSchema** | Pointer to [**NullableBulkWritableConfigContextRequestConfigContextSchema**](BulkWritableConfigContextRequestConfigContextSchema.md) |  | [optional] 
**LocalConfigContextDataOwnerContentType** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**DeviceType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Rack** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullablePrimaryIPv4**](PrimaryIPv4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullablePrimaryIPv6**](PrimaryIPv6.md) |  | [optional] 
**Cluster** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**DeviceRedundancyGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**SoftwareVersion** | Pointer to [**NullableBulkWritableDeviceRequestSoftwareVersion**](BulkWritableDeviceRequestSoftwareVersion.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ControllerManagedDeviceGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**SoftwareImageFiles** | Pointer to [**[]SoftwareImageFiles**](SoftwareImageFiles.md) | Override the software image files associated with the software version for this device | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**ParentBay** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedWritableDeviceRequest

`func NewPatchedWritableDeviceRequest() *PatchedWritableDeviceRequest`

NewPatchedWritableDeviceRequest instantiates a new PatchedWritableDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDeviceRequestWithDefaults

`func NewPatchedWritableDeviceRequestWithDefaults() *PatchedWritableDeviceRequest`

NewPatchedWritableDeviceRequestWithDefaults instantiates a new PatchedWritableDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalConfigContextData

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextData() map[string]interface{}`

GetLocalConfigContextData returns the LocalConfigContextData field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOk

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextDataOk() (*map[string]interface{}, bool)`

GetLocalConfigContextDataOk returns a tuple with the LocalConfigContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextData

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextData(v map[string]interface{})`

SetLocalConfigContextData sets LocalConfigContextData field to given value.

### HasLocalConfigContextData

`func (o *PatchedWritableDeviceRequest) HasLocalConfigContextData() bool`

HasLocalConfigContextData returns a boolean if a field has been set.

### SetLocalConfigContextDataNil

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextDataNil(b bool)`

 SetLocalConfigContextDataNil sets the value for LocalConfigContextData to be an explicit nil

### UnsetLocalConfigContextData
`func (o *PatchedWritableDeviceRequest) UnsetLocalConfigContextData()`

UnsetLocalConfigContextData ensures that no value is present for LocalConfigContextData, not even an explicit nil
### GetLocalConfigContextDataOwnerObjectId

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextDataOwnerObjectId() string`

GetLocalConfigContextDataOwnerObjectId returns the LocalConfigContextDataOwnerObjectId field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerObjectIdOk

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextDataOwnerObjectIdOk() (*string, bool)`

GetLocalConfigContextDataOwnerObjectIdOk returns a tuple with the LocalConfigContextDataOwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerObjectId

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextDataOwnerObjectId(v string)`

SetLocalConfigContextDataOwnerObjectId sets LocalConfigContextDataOwnerObjectId field to given value.

### HasLocalConfigContextDataOwnerObjectId

`func (o *PatchedWritableDeviceRequest) HasLocalConfigContextDataOwnerObjectId() bool`

HasLocalConfigContextDataOwnerObjectId returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerObjectIdNil

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextDataOwnerObjectIdNil(b bool)`

 SetLocalConfigContextDataOwnerObjectIdNil sets the value for LocalConfigContextDataOwnerObjectId to be an explicit nil

### UnsetLocalConfigContextDataOwnerObjectId
`func (o *PatchedWritableDeviceRequest) UnsetLocalConfigContextDataOwnerObjectId()`

UnsetLocalConfigContextDataOwnerObjectId ensures that no value is present for LocalConfigContextDataOwnerObjectId, not even an explicit nil
### GetName

`func (o *PatchedWritableDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedWritableDeviceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedWritableDeviceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSerial

`func (o *PatchedWritableDeviceRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableDeviceRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableDeviceRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableDeviceRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableDeviceRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableDeviceRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableDeviceRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableDeviceRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableDeviceRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableDeviceRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetPosition

`func (o *PatchedWritableDeviceRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedWritableDeviceRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedWritableDeviceRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedWritableDeviceRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PatchedWritableDeviceRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PatchedWritableDeviceRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *PatchedWritableDeviceRequest) GetFace() RackFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *PatchedWritableDeviceRequest) GetFaceOk() (*RackFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *PatchedWritableDeviceRequest) SetFace(v RackFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *PatchedWritableDeviceRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetDeviceRedundancyGroupPriority

`func (o *PatchedWritableDeviceRequest) GetDeviceRedundancyGroupPriority() int32`

GetDeviceRedundancyGroupPriority returns the DeviceRedundancyGroupPriority field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupPriorityOk

`func (o *PatchedWritableDeviceRequest) GetDeviceRedundancyGroupPriorityOk() (*int32, bool)`

GetDeviceRedundancyGroupPriorityOk returns a tuple with the DeviceRedundancyGroupPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroupPriority

`func (o *PatchedWritableDeviceRequest) SetDeviceRedundancyGroupPriority(v int32)`

SetDeviceRedundancyGroupPriority sets DeviceRedundancyGroupPriority field to given value.

### HasDeviceRedundancyGroupPriority

`func (o *PatchedWritableDeviceRequest) HasDeviceRedundancyGroupPriority() bool`

HasDeviceRedundancyGroupPriority returns a boolean if a field has been set.

### SetDeviceRedundancyGroupPriorityNil

`func (o *PatchedWritableDeviceRequest) SetDeviceRedundancyGroupPriorityNil(b bool)`

 SetDeviceRedundancyGroupPriorityNil sets the value for DeviceRedundancyGroupPriority to be an explicit nil

### UnsetDeviceRedundancyGroupPriority
`func (o *PatchedWritableDeviceRequest) UnsetDeviceRedundancyGroupPriority()`

UnsetDeviceRedundancyGroupPriority ensures that no value is present for DeviceRedundancyGroupPriority, not even an explicit nil
### GetVcPosition

`func (o *PatchedWritableDeviceRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *PatchedWritableDeviceRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *PatchedWritableDeviceRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *PatchedWritableDeviceRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *PatchedWritableDeviceRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *PatchedWritableDeviceRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *PatchedWritableDeviceRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *PatchedWritableDeviceRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *PatchedWritableDeviceRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *PatchedWritableDeviceRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *PatchedWritableDeviceRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *PatchedWritableDeviceRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *PatchedWritableDeviceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDeviceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDeviceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDeviceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalConfigContextSchema

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetLocalConfigContextSchema returns the LocalConfigContextSchema field if non-nil, zero value otherwise.

### GetLocalConfigContextSchemaOk

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetLocalConfigContextSchemaOk returns a tuple with the LocalConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextSchema

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetLocalConfigContextSchema sets LocalConfigContextSchema field to given value.

### HasLocalConfigContextSchema

`func (o *PatchedWritableDeviceRequest) HasLocalConfigContextSchema() bool`

HasLocalConfigContextSchema returns a boolean if a field has been set.

### SetLocalConfigContextSchemaNil

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextSchemaNil(b bool)`

 SetLocalConfigContextSchemaNil sets the value for LocalConfigContextSchema to be an explicit nil

### UnsetLocalConfigContextSchema
`func (o *PatchedWritableDeviceRequest) UnsetLocalConfigContextSchema()`

UnsetLocalConfigContextSchema ensures that no value is present for LocalConfigContextSchema, not even an explicit nil
### GetLocalConfigContextDataOwnerContentType

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextDataOwnerContentType() BulkWritableCircuitRequestTenant`

GetLocalConfigContextDataOwnerContentType returns the LocalConfigContextDataOwnerContentType field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerContentTypeOk

`func (o *PatchedWritableDeviceRequest) GetLocalConfigContextDataOwnerContentTypeOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocalConfigContextDataOwnerContentTypeOk returns a tuple with the LocalConfigContextDataOwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerContentType

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextDataOwnerContentType(v BulkWritableCircuitRequestTenant)`

SetLocalConfigContextDataOwnerContentType sets LocalConfigContextDataOwnerContentType field to given value.

### HasLocalConfigContextDataOwnerContentType

`func (o *PatchedWritableDeviceRequest) HasLocalConfigContextDataOwnerContentType() bool`

HasLocalConfigContextDataOwnerContentType returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerContentTypeNil

`func (o *PatchedWritableDeviceRequest) SetLocalConfigContextDataOwnerContentTypeNil(b bool)`

 SetLocalConfigContextDataOwnerContentTypeNil sets the value for LocalConfigContextDataOwnerContentType to be an explicit nil

### UnsetLocalConfigContextDataOwnerContentType
`func (o *PatchedWritableDeviceRequest) UnsetLocalConfigContextDataOwnerContentType()`

UnsetLocalConfigContextDataOwnerContentType ensures that no value is present for LocalConfigContextDataOwnerContentType, not even an explicit nil
### GetDeviceType

`func (o *PatchedWritableDeviceRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableDeviceRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableDeviceRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableDeviceRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableDeviceRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableDeviceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableDeviceRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableDeviceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableDeviceRequest) GetRole() BulkWritableCableRequestStatus`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableDeviceRequest) GetRoleOk() (*BulkWritableCableRequestStatus, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableDeviceRequest) SetRole(v BulkWritableCableRequestStatus)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableDeviceRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableDeviceRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableDeviceRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableDeviceRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableDeviceRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableDeviceRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableDeviceRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableDeviceRequest) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableDeviceRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableDeviceRequest) SetPlatform(v BulkWritableCircuitRequestTenant)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableDeviceRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableDeviceRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableDeviceRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLocation

`func (o *PatchedWritableDeviceRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritableDeviceRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritableDeviceRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritableDeviceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRack

`func (o *PatchedWritableDeviceRequest) GetRack() BulkWritableCircuitRequestTenant`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritableDeviceRequest) GetRackOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritableDeviceRequest) SetRack(v BulkWritableCircuitRequestTenant)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritableDeviceRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritableDeviceRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritableDeviceRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPrimaryIp4

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableDeviceRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableDeviceRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableDeviceRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableDeviceRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetCluster

`func (o *PatchedWritableDeviceRequest) GetCluster() BulkWritableCircuitRequestTenant`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableDeviceRequest) GetClusterOk() (*BulkWritableCircuitRequestTenant, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableDeviceRequest) SetCluster(v BulkWritableCircuitRequestTenant)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableDeviceRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedWritableDeviceRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedWritableDeviceRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedWritableDeviceRequest) GetVirtualChassis() BulkWritableCircuitRequestTenant`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedWritableDeviceRequest) GetVirtualChassisOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedWritableDeviceRequest) SetVirtualChassis(v BulkWritableCircuitRequestTenant)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedWritableDeviceRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedWritableDeviceRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedWritableDeviceRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetDeviceRedundancyGroup

`func (o *PatchedWritableDeviceRequest) GetDeviceRedundancyGroup() BulkWritableCircuitRequestTenant`

GetDeviceRedundancyGroup returns the DeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupOk

`func (o *PatchedWritableDeviceRequest) GetDeviceRedundancyGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceRedundancyGroupOk returns a tuple with the DeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroup

`func (o *PatchedWritableDeviceRequest) SetDeviceRedundancyGroup(v BulkWritableCircuitRequestTenant)`

SetDeviceRedundancyGroup sets DeviceRedundancyGroup field to given value.

### HasDeviceRedundancyGroup

`func (o *PatchedWritableDeviceRequest) HasDeviceRedundancyGroup() bool`

HasDeviceRedundancyGroup returns a boolean if a field has been set.

### SetDeviceRedundancyGroupNil

`func (o *PatchedWritableDeviceRequest) SetDeviceRedundancyGroupNil(b bool)`

 SetDeviceRedundancyGroupNil sets the value for DeviceRedundancyGroup to be an explicit nil

### UnsetDeviceRedundancyGroup
`func (o *PatchedWritableDeviceRequest) UnsetDeviceRedundancyGroup()`

UnsetDeviceRedundancyGroup ensures that no value is present for DeviceRedundancyGroup, not even an explicit nil
### GetSoftwareVersion

`func (o *PatchedWritableDeviceRequest) GetSoftwareVersion() BulkWritableDeviceRequestSoftwareVersion`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PatchedWritableDeviceRequest) GetSoftwareVersionOk() (*BulkWritableDeviceRequestSoftwareVersion, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PatchedWritableDeviceRequest) SetSoftwareVersion(v BulkWritableDeviceRequestSoftwareVersion)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *PatchedWritableDeviceRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *PatchedWritableDeviceRequest) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *PatchedWritableDeviceRequest) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetSecretsGroup

`func (o *PatchedWritableDeviceRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedWritableDeviceRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedWritableDeviceRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedWritableDeviceRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedWritableDeviceRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedWritableDeviceRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetControllerManagedDeviceGroup

`func (o *PatchedWritableDeviceRequest) GetControllerManagedDeviceGroup() BulkWritableCircuitRequestTenant`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *PatchedWritableDeviceRequest) GetControllerManagedDeviceGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *PatchedWritableDeviceRequest) SetControllerManagedDeviceGroup(v BulkWritableCircuitRequestTenant)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.

### HasControllerManagedDeviceGroup

`func (o *PatchedWritableDeviceRequest) HasControllerManagedDeviceGroup() bool`

HasControllerManagedDeviceGroup returns a boolean if a field has been set.

### SetControllerManagedDeviceGroupNil

`func (o *PatchedWritableDeviceRequest) SetControllerManagedDeviceGroupNil(b bool)`

 SetControllerManagedDeviceGroupNil sets the value for ControllerManagedDeviceGroup to be an explicit nil

### UnsetControllerManagedDeviceGroup
`func (o *PatchedWritableDeviceRequest) UnsetControllerManagedDeviceGroup()`

UnsetControllerManagedDeviceGroup ensures that no value is present for ControllerManagedDeviceGroup, not even an explicit nil
### GetSoftwareImageFiles

`func (o *PatchedWritableDeviceRequest) GetSoftwareImageFiles() []SoftwareImageFiles`

GetSoftwareImageFiles returns the SoftwareImageFiles field if non-nil, zero value otherwise.

### GetSoftwareImageFilesOk

`func (o *PatchedWritableDeviceRequest) GetSoftwareImageFilesOk() (*[]SoftwareImageFiles, bool)`

GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFiles

`func (o *PatchedWritableDeviceRequest) SetSoftwareImageFiles(v []SoftwareImageFiles)`

SetSoftwareImageFiles sets SoftwareImageFiles field to given value.

### HasSoftwareImageFiles

`func (o *PatchedWritableDeviceRequest) HasSoftwareImageFiles() bool`

HasSoftwareImageFiles returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableDeviceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableDeviceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableDeviceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDeviceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDeviceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDeviceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDeviceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableDeviceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableDeviceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableDeviceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableDeviceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetParentBay

`func (o *PatchedWritableDeviceRequest) GetParentBay() BulkWritableCircuitRequestTenant`

GetParentBay returns the ParentBay field if non-nil, zero value otherwise.

### GetParentBayOk

`func (o *PatchedWritableDeviceRequest) GetParentBayOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentBayOk returns a tuple with the ParentBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBay

`func (o *PatchedWritableDeviceRequest) SetParentBay(v BulkWritableCircuitRequestTenant)`

SetParentBay sets ParentBay field to given value.

### HasParentBay

`func (o *PatchedWritableDeviceRequest) HasParentBay() bool`

HasParentBay returns a boolean if a field has been set.

### SetParentBayNil

`func (o *PatchedWritableDeviceRequest) SetParentBayNil(b bool)`

 SetParentBayNil sets the value for ParentBay to be an explicit nil

### UnsetParentBay
`func (o *PatchedWritableDeviceRequest) UnsetParentBay()`

UnsetParentBay ensures that no value is present for ParentBay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

