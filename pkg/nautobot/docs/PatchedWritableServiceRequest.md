# PatchedWritableServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to **[]int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**ServiceProtocolChoices**](ServiceProtocolChoices.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**NullableBulkWritableServiceRequestDevice**](BulkWritableServiceRequestDevice.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableBulkWritableServiceRequestVirtualMachine**](BulkWritableServiceRequestVirtualMachine.md) |  | [optional] 
**IpAddresses** | Pointer to [**[]IPAddresses**](IPAddresses.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritableServiceRequest

`func NewPatchedWritableServiceRequest() *PatchedWritableServiceRequest`

NewPatchedWritableServiceRequest instantiates a new PatchedWritableServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableServiceRequestWithDefaults

`func NewPatchedWritableServiceRequestWithDefaults() *PatchedWritableServiceRequest`

NewPatchedWritableServiceRequestWithDefaults instantiates a new PatchedWritableServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *PatchedWritableServiceRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PatchedWritableServiceRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PatchedWritableServiceRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PatchedWritableServiceRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedWritableServiceRequest) GetProtocol() ServiceProtocolChoices`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedWritableServiceRequest) GetProtocolOk() (*ServiceProtocolChoices, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedWritableServiceRequest) SetProtocol(v ServiceProtocolChoices)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedWritableServiceRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedWritableServiceRequest) GetDevice() BulkWritableServiceRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableServiceRequest) GetDeviceOk() (*BulkWritableServiceRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableServiceRequest) SetDevice(v BulkWritableServiceRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableServiceRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedWritableServiceRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedWritableServiceRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVirtualMachine

`func (o *PatchedWritableServiceRequest) GetVirtualMachine() BulkWritableServiceRequestVirtualMachine`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedWritableServiceRequest) GetVirtualMachineOk() (*BulkWritableServiceRequestVirtualMachine, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedWritableServiceRequest) SetVirtualMachine(v BulkWritableServiceRequestVirtualMachine)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PatchedWritableServiceRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *PatchedWritableServiceRequest) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *PatchedWritableServiceRequest) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil
### GetIpAddresses

`func (o *PatchedWritableServiceRequest) GetIpAddresses() []IPAddresses`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *PatchedWritableServiceRequest) GetIpAddressesOk() (*[]IPAddresses, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *PatchedWritableServiceRequest) SetIpAddresses(v []IPAddresses)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *PatchedWritableServiceRequest) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableServiceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableServiceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableServiceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableServiceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableServiceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableServiceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableServiceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


