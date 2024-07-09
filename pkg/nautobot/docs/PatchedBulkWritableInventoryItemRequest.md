# PatchedBulkWritableInventoryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this item | [optional] 
**Discovered** | Pointer to **bool** | This item was automatically discovered | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Device** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**SoftwareVersion** | Pointer to [**NullableTheSoftwareVersionInstalledOnThisInventoryItem**](TheSoftwareVersionInstalledOnThisInventoryItem.md) |  | [optional] 
**SoftwareImageFiles** | Pointer to [**[]SoftwareImageFiles**](SoftwareImageFiles.md) | Override the software image files associated with the software version for this inventory item | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableInventoryItemRequest

`func NewPatchedBulkWritableInventoryItemRequest(id string, ) *PatchedBulkWritableInventoryItemRequest`

NewPatchedBulkWritableInventoryItemRequest instantiates a new PatchedBulkWritableInventoryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableInventoryItemRequestWithDefaults

`func NewPatchedBulkWritableInventoryItemRequestWithDefaults() *PatchedBulkWritableInventoryItemRequest`

NewPatchedBulkWritableInventoryItemRequestWithDefaults instantiates a new PatchedBulkWritableInventoryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableInventoryItemRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableInventoryItemRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableInventoryItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableInventoryItemRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableInventoryItemRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkWritableInventoryItemRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkWritableInventoryItemRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkWritableInventoryItemRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableInventoryItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableInventoryItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableInventoryItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPartId

`func (o *PatchedBulkWritableInventoryItemRequest) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *PatchedBulkWritableInventoryItemRequest) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *PatchedBulkWritableInventoryItemRequest) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetSerial

`func (o *PatchedBulkWritableInventoryItemRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedBulkWritableInventoryItemRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedBulkWritableInventoryItemRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedBulkWritableInventoryItemRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedBulkWritableInventoryItemRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedBulkWritableInventoryItemRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedBulkWritableInventoryItemRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedBulkWritableInventoryItemRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDiscovered

`func (o *PatchedBulkWritableInventoryItemRequest) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *PatchedBulkWritableInventoryItemRequest) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *PatchedBulkWritableInventoryItemRequest) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkWritableInventoryItemRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkWritableInventoryItemRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkWritableInventoryItemRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkWritableInventoryItemRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkWritableInventoryItemRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDevice

`func (o *PatchedBulkWritableInventoryItemRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkWritableInventoryItemRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkWritableInventoryItemRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetManufacturer

`func (o *PatchedBulkWritableInventoryItemRequest) GetManufacturer() BulkWritableCircuitRequestTenant`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetManufacturerOk() (*BulkWritableCircuitRequestTenant, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedBulkWritableInventoryItemRequest) SetManufacturer(v BulkWritableCircuitRequestTenant)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedBulkWritableInventoryItemRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *PatchedBulkWritableInventoryItemRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *PatchedBulkWritableInventoryItemRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetSoftwareVersion

`func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareVersion() TheSoftwareVersionInstalledOnThisInventoryItem`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareVersionOk() (*TheSoftwareVersionInstalledOnThisInventoryItem, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PatchedBulkWritableInventoryItemRequest) SetSoftwareVersion(v TheSoftwareVersionInstalledOnThisInventoryItem)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *PatchedBulkWritableInventoryItemRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *PatchedBulkWritableInventoryItemRequest) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *PatchedBulkWritableInventoryItemRequest) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetSoftwareImageFiles

`func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareImageFiles() []SoftwareImageFiles`

GetSoftwareImageFiles returns the SoftwareImageFiles field if non-nil, zero value otherwise.

### GetSoftwareImageFilesOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetSoftwareImageFilesOk() (*[]SoftwareImageFiles, bool)`

GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFiles

`func (o *PatchedBulkWritableInventoryItemRequest) SetSoftwareImageFiles(v []SoftwareImageFiles)`

SetSoftwareImageFiles sets SoftwareImageFiles field to given value.

### HasSoftwareImageFiles

`func (o *PatchedBulkWritableInventoryItemRequest) HasSoftwareImageFiles() bool`

HasSoftwareImageFiles returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableInventoryItemRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableInventoryItemRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableInventoryItemRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableInventoryItemRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableInventoryItemRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableInventoryItemRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableInventoryItemRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableInventoryItemRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableInventoryItemRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableInventoryItemRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


