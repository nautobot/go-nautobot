# BulkWritableDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Face** | Pointer to [**FaceEnum**](FaceEnum.md) |  | [optional] 
**LocalConfigContextData** | Pointer to **map[string]interface{}** |  | [optional] 
**LocalConfigContextDataOwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Position** | Pointer to **NullableInt32** | The lowest-numbered unit occupied by the device | [optional] 
**DeviceRedundancyGroupPriority** | Pointer to **NullableInt32** | The priority the device has in the device redundancy group. | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalConfigContextSchema** | Pointer to [**NullableBulkWritableConfigContextRequestConfigContextSchema**](BulkWritableConfigContextRequestConfigContextSchema.md) |  | [optional] 
**LocalConfigContextDataOwnerContentType** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**DeviceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
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

### NewBulkWritableDeviceRequest

`func NewBulkWritableDeviceRequest(id string, deviceType BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, role BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, ) *BulkWritableDeviceRequest`

NewBulkWritableDeviceRequest instantiates a new BulkWritableDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDeviceRequestWithDefaults

`func NewBulkWritableDeviceRequestWithDefaults() *BulkWritableDeviceRequest`

NewBulkWritableDeviceRequestWithDefaults instantiates a new BulkWritableDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDeviceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDeviceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDeviceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFace

`func (o *BulkWritableDeviceRequest) GetFace() FaceEnum`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *BulkWritableDeviceRequest) GetFaceOk() (*FaceEnum, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *BulkWritableDeviceRequest) SetFace(v FaceEnum)`

SetFace sets Face field to given value.

### HasFace

`func (o *BulkWritableDeviceRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetLocalConfigContextData

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextData() map[string]interface{}`

GetLocalConfigContextData returns the LocalConfigContextData field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOk

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextDataOk() (*map[string]interface{}, bool)`

GetLocalConfigContextDataOk returns a tuple with the LocalConfigContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextData

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextData(v map[string]interface{})`

SetLocalConfigContextData sets LocalConfigContextData field to given value.

### HasLocalConfigContextData

`func (o *BulkWritableDeviceRequest) HasLocalConfigContextData() bool`

HasLocalConfigContextData returns a boolean if a field has been set.

### SetLocalConfigContextDataNil

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextDataNil(b bool)`

 SetLocalConfigContextDataNil sets the value for LocalConfigContextData to be an explicit nil

### UnsetLocalConfigContextData
`func (o *BulkWritableDeviceRequest) UnsetLocalConfigContextData()`

UnsetLocalConfigContextData ensures that no value is present for LocalConfigContextData, not even an explicit nil
### GetLocalConfigContextDataOwnerObjectId

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextDataOwnerObjectId() string`

GetLocalConfigContextDataOwnerObjectId returns the LocalConfigContextDataOwnerObjectId field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerObjectIdOk

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextDataOwnerObjectIdOk() (*string, bool)`

GetLocalConfigContextDataOwnerObjectIdOk returns a tuple with the LocalConfigContextDataOwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerObjectId

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextDataOwnerObjectId(v string)`

SetLocalConfigContextDataOwnerObjectId sets LocalConfigContextDataOwnerObjectId field to given value.

### HasLocalConfigContextDataOwnerObjectId

`func (o *BulkWritableDeviceRequest) HasLocalConfigContextDataOwnerObjectId() bool`

HasLocalConfigContextDataOwnerObjectId returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerObjectIdNil

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextDataOwnerObjectIdNil(b bool)`

 SetLocalConfigContextDataOwnerObjectIdNil sets the value for LocalConfigContextDataOwnerObjectId to be an explicit nil

### UnsetLocalConfigContextDataOwnerObjectId
`func (o *BulkWritableDeviceRequest) UnsetLocalConfigContextDataOwnerObjectId()`

UnsetLocalConfigContextDataOwnerObjectId ensures that no value is present for LocalConfigContextDataOwnerObjectId, not even an explicit nil
### GetName

`func (o *BulkWritableDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkWritableDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BulkWritableDeviceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BulkWritableDeviceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSerial

`func (o *BulkWritableDeviceRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BulkWritableDeviceRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BulkWritableDeviceRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BulkWritableDeviceRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *BulkWritableDeviceRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *BulkWritableDeviceRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *BulkWritableDeviceRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *BulkWritableDeviceRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *BulkWritableDeviceRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *BulkWritableDeviceRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetPosition

`func (o *BulkWritableDeviceRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BulkWritableDeviceRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BulkWritableDeviceRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BulkWritableDeviceRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *BulkWritableDeviceRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *BulkWritableDeviceRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetDeviceRedundancyGroupPriority

`func (o *BulkWritableDeviceRequest) GetDeviceRedundancyGroupPriority() int32`

GetDeviceRedundancyGroupPriority returns the DeviceRedundancyGroupPriority field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupPriorityOk

`func (o *BulkWritableDeviceRequest) GetDeviceRedundancyGroupPriorityOk() (*int32, bool)`

GetDeviceRedundancyGroupPriorityOk returns a tuple with the DeviceRedundancyGroupPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroupPriority

`func (o *BulkWritableDeviceRequest) SetDeviceRedundancyGroupPriority(v int32)`

SetDeviceRedundancyGroupPriority sets DeviceRedundancyGroupPriority field to given value.

### HasDeviceRedundancyGroupPriority

`func (o *BulkWritableDeviceRequest) HasDeviceRedundancyGroupPriority() bool`

HasDeviceRedundancyGroupPriority returns a boolean if a field has been set.

### SetDeviceRedundancyGroupPriorityNil

`func (o *BulkWritableDeviceRequest) SetDeviceRedundancyGroupPriorityNil(b bool)`

 SetDeviceRedundancyGroupPriorityNil sets the value for DeviceRedundancyGroupPriority to be an explicit nil

### UnsetDeviceRedundancyGroupPriority
`func (o *BulkWritableDeviceRequest) UnsetDeviceRedundancyGroupPriority()`

UnsetDeviceRedundancyGroupPriority ensures that no value is present for DeviceRedundancyGroupPriority, not even an explicit nil
### GetVcPosition

`func (o *BulkWritableDeviceRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *BulkWritableDeviceRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *BulkWritableDeviceRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *BulkWritableDeviceRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *BulkWritableDeviceRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *BulkWritableDeviceRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *BulkWritableDeviceRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *BulkWritableDeviceRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *BulkWritableDeviceRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *BulkWritableDeviceRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *BulkWritableDeviceRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *BulkWritableDeviceRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *BulkWritableDeviceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableDeviceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableDeviceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableDeviceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalConfigContextSchema

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetLocalConfigContextSchema returns the LocalConfigContextSchema field if non-nil, zero value otherwise.

### GetLocalConfigContextSchemaOk

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetLocalConfigContextSchemaOk returns a tuple with the LocalConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextSchema

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetLocalConfigContextSchema sets LocalConfigContextSchema field to given value.

### HasLocalConfigContextSchema

`func (o *BulkWritableDeviceRequest) HasLocalConfigContextSchema() bool`

HasLocalConfigContextSchema returns a boolean if a field has been set.

### SetLocalConfigContextSchemaNil

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextSchemaNil(b bool)`

 SetLocalConfigContextSchemaNil sets the value for LocalConfigContextSchema to be an explicit nil

### UnsetLocalConfigContextSchema
`func (o *BulkWritableDeviceRequest) UnsetLocalConfigContextSchema()`

UnsetLocalConfigContextSchema ensures that no value is present for LocalConfigContextSchema, not even an explicit nil
### GetLocalConfigContextDataOwnerContentType

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextDataOwnerContentType() BulkWritableCircuitRequestTenant`

GetLocalConfigContextDataOwnerContentType returns the LocalConfigContextDataOwnerContentType field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerContentTypeOk

`func (o *BulkWritableDeviceRequest) GetLocalConfigContextDataOwnerContentTypeOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocalConfigContextDataOwnerContentTypeOk returns a tuple with the LocalConfigContextDataOwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerContentType

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextDataOwnerContentType(v BulkWritableCircuitRequestTenant)`

SetLocalConfigContextDataOwnerContentType sets LocalConfigContextDataOwnerContentType field to given value.

### HasLocalConfigContextDataOwnerContentType

`func (o *BulkWritableDeviceRequest) HasLocalConfigContextDataOwnerContentType() bool`

HasLocalConfigContextDataOwnerContentType returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerContentTypeNil

`func (o *BulkWritableDeviceRequest) SetLocalConfigContextDataOwnerContentTypeNil(b bool)`

 SetLocalConfigContextDataOwnerContentTypeNil sets the value for LocalConfigContextDataOwnerContentType to be an explicit nil

### UnsetLocalConfigContextDataOwnerContentType
`func (o *BulkWritableDeviceRequest) UnsetLocalConfigContextDataOwnerContentType()`

UnsetLocalConfigContextDataOwnerContentType ensures that no value is present for LocalConfigContextDataOwnerContentType, not even an explicit nil
### GetDeviceType

`func (o *BulkWritableDeviceRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkWritableDeviceRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkWritableDeviceRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.


### GetStatus

`func (o *BulkWritableDeviceRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableDeviceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableDeviceRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *BulkWritableDeviceRequest) GetRole() BulkWritableCableRequestStatus`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkWritableDeviceRequest) GetRoleOk() (*BulkWritableCableRequestStatus, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkWritableDeviceRequest) SetRole(v BulkWritableCableRequestStatus)`

SetRole sets Role field to given value.


### GetTenant

`func (o *BulkWritableDeviceRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableDeviceRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableDeviceRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableDeviceRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableDeviceRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableDeviceRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *BulkWritableDeviceRequest) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BulkWritableDeviceRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BulkWritableDeviceRequest) SetPlatform(v BulkWritableCircuitRequestTenant)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *BulkWritableDeviceRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *BulkWritableDeviceRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *BulkWritableDeviceRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLocation

`func (o *BulkWritableDeviceRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BulkWritableDeviceRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BulkWritableDeviceRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetRack

`func (o *BulkWritableDeviceRequest) GetRack() BulkWritableCircuitRequestTenant`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *BulkWritableDeviceRequest) GetRackOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *BulkWritableDeviceRequest) SetRack(v BulkWritableCircuitRequestTenant)`

SetRack sets Rack field to given value.

### HasRack

`func (o *BulkWritableDeviceRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *BulkWritableDeviceRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *BulkWritableDeviceRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPrimaryIp4

`func (o *BulkWritableDeviceRequest) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *BulkWritableDeviceRequest) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *BulkWritableDeviceRequest) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *BulkWritableDeviceRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *BulkWritableDeviceRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *BulkWritableDeviceRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *BulkWritableDeviceRequest) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *BulkWritableDeviceRequest) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *BulkWritableDeviceRequest) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *BulkWritableDeviceRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *BulkWritableDeviceRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *BulkWritableDeviceRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetCluster

`func (o *BulkWritableDeviceRequest) GetCluster() BulkWritableCircuitRequestTenant`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BulkWritableDeviceRequest) GetClusterOk() (*BulkWritableCircuitRequestTenant, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BulkWritableDeviceRequest) SetCluster(v BulkWritableCircuitRequestTenant)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *BulkWritableDeviceRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *BulkWritableDeviceRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *BulkWritableDeviceRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *BulkWritableDeviceRequest) GetVirtualChassis() BulkWritableCircuitRequestTenant`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *BulkWritableDeviceRequest) GetVirtualChassisOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *BulkWritableDeviceRequest) SetVirtualChassis(v BulkWritableCircuitRequestTenant)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *BulkWritableDeviceRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *BulkWritableDeviceRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *BulkWritableDeviceRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetDeviceRedundancyGroup

`func (o *BulkWritableDeviceRequest) GetDeviceRedundancyGroup() BulkWritableCircuitRequestTenant`

GetDeviceRedundancyGroup returns the DeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupOk

`func (o *BulkWritableDeviceRequest) GetDeviceRedundancyGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceRedundancyGroupOk returns a tuple with the DeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroup

`func (o *BulkWritableDeviceRequest) SetDeviceRedundancyGroup(v BulkWritableCircuitRequestTenant)`

SetDeviceRedundancyGroup sets DeviceRedundancyGroup field to given value.

### HasDeviceRedundancyGroup

`func (o *BulkWritableDeviceRequest) HasDeviceRedundancyGroup() bool`

HasDeviceRedundancyGroup returns a boolean if a field has been set.

### SetDeviceRedundancyGroupNil

`func (o *BulkWritableDeviceRequest) SetDeviceRedundancyGroupNil(b bool)`

 SetDeviceRedundancyGroupNil sets the value for DeviceRedundancyGroup to be an explicit nil

### UnsetDeviceRedundancyGroup
`func (o *BulkWritableDeviceRequest) UnsetDeviceRedundancyGroup()`

UnsetDeviceRedundancyGroup ensures that no value is present for DeviceRedundancyGroup, not even an explicit nil
### GetSoftwareVersion

`func (o *BulkWritableDeviceRequest) GetSoftwareVersion() BulkWritableDeviceRequestSoftwareVersion`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *BulkWritableDeviceRequest) GetSoftwareVersionOk() (*BulkWritableDeviceRequestSoftwareVersion, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *BulkWritableDeviceRequest) SetSoftwareVersion(v BulkWritableDeviceRequestSoftwareVersion)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *BulkWritableDeviceRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *BulkWritableDeviceRequest) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *BulkWritableDeviceRequest) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetSecretsGroup

`func (o *BulkWritableDeviceRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableDeviceRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableDeviceRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableDeviceRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableDeviceRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableDeviceRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetControllerManagedDeviceGroup

`func (o *BulkWritableDeviceRequest) GetControllerManagedDeviceGroup() BulkWritableCircuitRequestTenant`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *BulkWritableDeviceRequest) GetControllerManagedDeviceGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *BulkWritableDeviceRequest) SetControllerManagedDeviceGroup(v BulkWritableCircuitRequestTenant)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.

### HasControllerManagedDeviceGroup

`func (o *BulkWritableDeviceRequest) HasControllerManagedDeviceGroup() bool`

HasControllerManagedDeviceGroup returns a boolean if a field has been set.

### SetControllerManagedDeviceGroupNil

`func (o *BulkWritableDeviceRequest) SetControllerManagedDeviceGroupNil(b bool)`

 SetControllerManagedDeviceGroupNil sets the value for ControllerManagedDeviceGroup to be an explicit nil

### UnsetControllerManagedDeviceGroup
`func (o *BulkWritableDeviceRequest) UnsetControllerManagedDeviceGroup()`

UnsetControllerManagedDeviceGroup ensures that no value is present for ControllerManagedDeviceGroup, not even an explicit nil
### GetSoftwareImageFiles

`func (o *BulkWritableDeviceRequest) GetSoftwareImageFiles() []SoftwareImageFiles`

GetSoftwareImageFiles returns the SoftwareImageFiles field if non-nil, zero value otherwise.

### GetSoftwareImageFilesOk

`func (o *BulkWritableDeviceRequest) GetSoftwareImageFilesOk() (*[]SoftwareImageFiles, bool)`

GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFiles

`func (o *BulkWritableDeviceRequest) SetSoftwareImageFiles(v []SoftwareImageFiles)`

SetSoftwareImageFiles sets SoftwareImageFiles field to given value.

### HasSoftwareImageFiles

`func (o *BulkWritableDeviceRequest) HasSoftwareImageFiles() bool`

HasSoftwareImageFiles returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableDeviceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableDeviceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableDeviceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableDeviceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableDeviceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableDeviceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableDeviceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableDeviceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableDeviceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableDeviceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableDeviceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetParentBay

`func (o *BulkWritableDeviceRequest) GetParentBay() BulkWritableCircuitRequestTenant`

GetParentBay returns the ParentBay field if non-nil, zero value otherwise.

### GetParentBayOk

`func (o *BulkWritableDeviceRequest) GetParentBayOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentBayOk returns a tuple with the ParentBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBay

`func (o *BulkWritableDeviceRequest) SetParentBay(v BulkWritableCircuitRequestTenant)`

SetParentBay sets ParentBay field to given value.

### HasParentBay

`func (o *BulkWritableDeviceRequest) HasParentBay() bool`

HasParentBay returns a boolean if a field has been set.

### SetParentBayNil

`func (o *BulkWritableDeviceRequest) SetParentBayNil(b bool)`

 SetParentBayNil sets the value for ParentBay to be an explicit nil

### UnsetParentBay
`func (o *BulkWritableDeviceRequest) UnsetParentBay()`

UnsetParentBay ensures that no value is present for ParentBay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


