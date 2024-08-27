# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Face** | Pointer to [**DeviceFace**](DeviceFace.md) |  | [optional] 
**LocalConfigContextData** | Pointer to **interface{}** |  | [optional] 
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
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ParentBay** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewDevice

`func NewDevice(id string, objectType string, display string, url string, naturalSlug string, deviceType BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, role BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Device) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Device) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Device) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Device) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Device) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Device) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Device) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Device) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Device) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Device) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Device) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Device) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetFace

`func (o *Device) GetFace() DeviceFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *Device) GetFaceOk() (*DeviceFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *Device) SetFace(v DeviceFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *Device) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetLocalConfigContextData

`func (o *Device) GetLocalConfigContextData() interface{}`

GetLocalConfigContextData returns the LocalConfigContextData field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOk

`func (o *Device) GetLocalConfigContextDataOk() (*interface{}, bool)`

GetLocalConfigContextDataOk returns a tuple with the LocalConfigContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextData

`func (o *Device) SetLocalConfigContextData(v interface{})`

SetLocalConfigContextData sets LocalConfigContextData field to given value.

### HasLocalConfigContextData

`func (o *Device) HasLocalConfigContextData() bool`

HasLocalConfigContextData returns a boolean if a field has been set.

### SetLocalConfigContextDataNil

`func (o *Device) SetLocalConfigContextDataNil(b bool)`

 SetLocalConfigContextDataNil sets the value for LocalConfigContextData to be an explicit nil

### UnsetLocalConfigContextData
`func (o *Device) UnsetLocalConfigContextData()`

UnsetLocalConfigContextData ensures that no value is present for LocalConfigContextData, not even an explicit nil
### GetLocalConfigContextDataOwnerObjectId

`func (o *Device) GetLocalConfigContextDataOwnerObjectId() string`

GetLocalConfigContextDataOwnerObjectId returns the LocalConfigContextDataOwnerObjectId field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerObjectIdOk

`func (o *Device) GetLocalConfigContextDataOwnerObjectIdOk() (*string, bool)`

GetLocalConfigContextDataOwnerObjectIdOk returns a tuple with the LocalConfigContextDataOwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerObjectId

`func (o *Device) SetLocalConfigContextDataOwnerObjectId(v string)`

SetLocalConfigContextDataOwnerObjectId sets LocalConfigContextDataOwnerObjectId field to given value.

### HasLocalConfigContextDataOwnerObjectId

`func (o *Device) HasLocalConfigContextDataOwnerObjectId() bool`

HasLocalConfigContextDataOwnerObjectId returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerObjectIdNil

`func (o *Device) SetLocalConfigContextDataOwnerObjectIdNil(b bool)`

 SetLocalConfigContextDataOwnerObjectIdNil sets the value for LocalConfigContextDataOwnerObjectId to be an explicit nil

### UnsetLocalConfigContextDataOwnerObjectId
`func (o *Device) UnsetLocalConfigContextDataOwnerObjectId()`

UnsetLocalConfigContextDataOwnerObjectId ensures that no value is present for LocalConfigContextDataOwnerObjectId, not even an explicit nil
### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Device) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Device) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSerial

`func (o *Device) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Device) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Device) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Device) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *Device) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Device) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Device) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Device) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *Device) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *Device) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetPosition

`func (o *Device) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Device) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Device) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Device) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *Device) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *Device) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetDeviceRedundancyGroupPriority

`func (o *Device) GetDeviceRedundancyGroupPriority() int32`

GetDeviceRedundancyGroupPriority returns the DeviceRedundancyGroupPriority field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupPriorityOk

`func (o *Device) GetDeviceRedundancyGroupPriorityOk() (*int32, bool)`

GetDeviceRedundancyGroupPriorityOk returns a tuple with the DeviceRedundancyGroupPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroupPriority

`func (o *Device) SetDeviceRedundancyGroupPriority(v int32)`

SetDeviceRedundancyGroupPriority sets DeviceRedundancyGroupPriority field to given value.

### HasDeviceRedundancyGroupPriority

`func (o *Device) HasDeviceRedundancyGroupPriority() bool`

HasDeviceRedundancyGroupPriority returns a boolean if a field has been set.

### SetDeviceRedundancyGroupPriorityNil

`func (o *Device) SetDeviceRedundancyGroupPriorityNil(b bool)`

 SetDeviceRedundancyGroupPriorityNil sets the value for DeviceRedundancyGroupPriority to be an explicit nil

### UnsetDeviceRedundancyGroupPriority
`func (o *Device) UnsetDeviceRedundancyGroupPriority()`

UnsetDeviceRedundancyGroupPriority ensures that no value is present for DeviceRedundancyGroupPriority, not even an explicit nil
### GetVcPosition

`func (o *Device) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *Device) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *Device) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *Device) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *Device) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *Device) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *Device) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *Device) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *Device) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *Device) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *Device) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *Device) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *Device) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Device) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Device) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Device) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalConfigContextSchema

`func (o *Device) GetLocalConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema`

GetLocalConfigContextSchema returns the LocalConfigContextSchema field if non-nil, zero value otherwise.

### GetLocalConfigContextSchemaOk

`func (o *Device) GetLocalConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool)`

GetLocalConfigContextSchemaOk returns a tuple with the LocalConfigContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextSchema

`func (o *Device) SetLocalConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema)`

SetLocalConfigContextSchema sets LocalConfigContextSchema field to given value.

### HasLocalConfigContextSchema

`func (o *Device) HasLocalConfigContextSchema() bool`

HasLocalConfigContextSchema returns a boolean if a field has been set.

### SetLocalConfigContextSchemaNil

`func (o *Device) SetLocalConfigContextSchemaNil(b bool)`

 SetLocalConfigContextSchemaNil sets the value for LocalConfigContextSchema to be an explicit nil

### UnsetLocalConfigContextSchema
`func (o *Device) UnsetLocalConfigContextSchema()`

UnsetLocalConfigContextSchema ensures that no value is present for LocalConfigContextSchema, not even an explicit nil
### GetLocalConfigContextDataOwnerContentType

`func (o *Device) GetLocalConfigContextDataOwnerContentType() BulkWritableCircuitRequestTenant`

GetLocalConfigContextDataOwnerContentType returns the LocalConfigContextDataOwnerContentType field if non-nil, zero value otherwise.

### GetLocalConfigContextDataOwnerContentTypeOk

`func (o *Device) GetLocalConfigContextDataOwnerContentTypeOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocalConfigContextDataOwnerContentTypeOk returns a tuple with the LocalConfigContextDataOwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigContextDataOwnerContentType

`func (o *Device) SetLocalConfigContextDataOwnerContentType(v BulkWritableCircuitRequestTenant)`

SetLocalConfigContextDataOwnerContentType sets LocalConfigContextDataOwnerContentType field to given value.

### HasLocalConfigContextDataOwnerContentType

`func (o *Device) HasLocalConfigContextDataOwnerContentType() bool`

HasLocalConfigContextDataOwnerContentType returns a boolean if a field has been set.

### SetLocalConfigContextDataOwnerContentTypeNil

`func (o *Device) SetLocalConfigContextDataOwnerContentTypeNil(b bool)`

 SetLocalConfigContextDataOwnerContentTypeNil sets the value for LocalConfigContextDataOwnerContentType to be an explicit nil

### UnsetLocalConfigContextDataOwnerContentType
`func (o *Device) UnsetLocalConfigContextDataOwnerContentType()`

UnsetLocalConfigContextDataOwnerContentType ensures that no value is present for LocalConfigContextDataOwnerContentType, not even an explicit nil
### GetDeviceType

`func (o *Device) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Device) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Device) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.


### GetStatus

`func (o *Device) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *Device) GetRole() BulkWritableCableRequestStatus`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Device) GetRoleOk() (*BulkWritableCableRequestStatus, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Device) SetRole(v BulkWritableCableRequestStatus)`

SetRole sets Role field to given value.


### GetTenant

`func (o *Device) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Device) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Device) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Device) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Device) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Device) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *Device) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Device) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Device) SetPlatform(v BulkWritableCircuitRequestTenant)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Device) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Device) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Device) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLocation

`func (o *Device) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Device) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Device) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetRack

`func (o *Device) GetRack() BulkWritableCircuitRequestTenant`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *Device) GetRackOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *Device) SetRack(v BulkWritableCircuitRequestTenant)`

SetRack sets Rack field to given value.

### HasRack

`func (o *Device) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *Device) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *Device) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPrimaryIp4

`func (o *Device) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *Device) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *Device) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *Device) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *Device) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *Device) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *Device) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *Device) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *Device) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *Device) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *Device) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *Device) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetCluster

`func (o *Device) GetCluster() BulkWritableCircuitRequestTenant`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Device) GetClusterOk() (*BulkWritableCircuitRequestTenant, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Device) SetCluster(v BulkWritableCircuitRequestTenant)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Device) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *Device) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *Device) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *Device) GetVirtualChassis() BulkWritableCircuitRequestTenant`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *Device) GetVirtualChassisOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *Device) SetVirtualChassis(v BulkWritableCircuitRequestTenant)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *Device) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *Device) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *Device) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetDeviceRedundancyGroup

`func (o *Device) GetDeviceRedundancyGroup() BulkWritableCircuitRequestTenant`

GetDeviceRedundancyGroup returns the DeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupOk

`func (o *Device) GetDeviceRedundancyGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceRedundancyGroupOk returns a tuple with the DeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroup

`func (o *Device) SetDeviceRedundancyGroup(v BulkWritableCircuitRequestTenant)`

SetDeviceRedundancyGroup sets DeviceRedundancyGroup field to given value.

### HasDeviceRedundancyGroup

`func (o *Device) HasDeviceRedundancyGroup() bool`

HasDeviceRedundancyGroup returns a boolean if a field has been set.

### SetDeviceRedundancyGroupNil

`func (o *Device) SetDeviceRedundancyGroupNil(b bool)`

 SetDeviceRedundancyGroupNil sets the value for DeviceRedundancyGroup to be an explicit nil

### UnsetDeviceRedundancyGroup
`func (o *Device) UnsetDeviceRedundancyGroup()`

UnsetDeviceRedundancyGroup ensures that no value is present for DeviceRedundancyGroup, not even an explicit nil
### GetSoftwareVersion

`func (o *Device) GetSoftwareVersion() BulkWritableDeviceRequestSoftwareVersion`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *Device) GetSoftwareVersionOk() (*BulkWritableDeviceRequestSoftwareVersion, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *Device) SetSoftwareVersion(v BulkWritableDeviceRequestSoftwareVersion)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *Device) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *Device) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *Device) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetSecretsGroup

`func (o *Device) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *Device) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *Device) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *Device) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *Device) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *Device) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetControllerManagedDeviceGroup

`func (o *Device) GetControllerManagedDeviceGroup() BulkWritableCircuitRequestTenant`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *Device) GetControllerManagedDeviceGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *Device) SetControllerManagedDeviceGroup(v BulkWritableCircuitRequestTenant)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.

### HasControllerManagedDeviceGroup

`func (o *Device) HasControllerManagedDeviceGroup() bool`

HasControllerManagedDeviceGroup returns a boolean if a field has been set.

### SetControllerManagedDeviceGroupNil

`func (o *Device) SetControllerManagedDeviceGroupNil(b bool)`

 SetControllerManagedDeviceGroupNil sets the value for ControllerManagedDeviceGroup to be an explicit nil

### UnsetControllerManagedDeviceGroup
`func (o *Device) UnsetControllerManagedDeviceGroup()`

UnsetControllerManagedDeviceGroup ensures that no value is present for ControllerManagedDeviceGroup, not even an explicit nil
### GetSoftwareImageFiles

`func (o *Device) GetSoftwareImageFiles() []SoftwareImageFiles`

GetSoftwareImageFiles returns the SoftwareImageFiles field if non-nil, zero value otherwise.

### GetSoftwareImageFilesOk

`func (o *Device) GetSoftwareImageFilesOk() (*[]SoftwareImageFiles, bool)`

GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFiles

`func (o *Device) SetSoftwareImageFiles(v []SoftwareImageFiles)`

SetSoftwareImageFiles sets SoftwareImageFiles field to given value.

### HasSoftwareImageFiles

`func (o *Device) HasSoftwareImageFiles() bool`

HasSoftwareImageFiles returns a boolean if a field has been set.

### GetCreated

`func (o *Device) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Device) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Device) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Device) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Device) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Device) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Device) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Device) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Device) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Device) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Device) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Device) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Device) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Device) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Device) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Device) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Device) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *Device) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParentBay

`func (o *Device) GetParentBay() BulkWritableCircuitRequestTenant`

GetParentBay returns the ParentBay field if non-nil, zero value otherwise.

### GetParentBayOk

`func (o *Device) GetParentBayOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentBayOk returns a tuple with the ParentBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBay

`func (o *Device) SetParentBay(v BulkWritableCircuitRequestTenant)`

SetParentBay sets ParentBay field to given value.

### HasParentBay

`func (o *Device) HasParentBay() bool`

HasParentBay returns a boolean if a field has been set.

### SetParentBayNil

`func (o *Device) SetParentBayNil(b bool)`

 SetParentBayNil sets the value for ParentBay to be an explicit nil

### UnsetParentBay
`func (o *Device) UnsetParentBay()`

UnsetParentBay ensures that no value is present for ParentBay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


