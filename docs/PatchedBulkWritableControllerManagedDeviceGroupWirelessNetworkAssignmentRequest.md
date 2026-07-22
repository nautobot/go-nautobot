# PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ControllerManagedDeviceGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**WirelessNetwork** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest

`func NewPatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest(id string, ) *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest`

NewPatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest instantiates a new PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults

`func NewPatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults() *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest`

NewPatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetControllerManagedDeviceGroup

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetControllerManagedDeviceGroup() BulkWritableCableRequestStatus`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetControllerManagedDeviceGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetControllerManagedDeviceGroup(v BulkWritableCableRequestStatus)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.

### HasControllerManagedDeviceGroup

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasControllerManagedDeviceGroup() bool`

HasControllerManagedDeviceGroup returns a boolean if a field has been set.

### GetWirelessNetwork

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetWirelessNetwork() BulkWritableCableRequestStatus`

GetWirelessNetwork returns the WirelessNetwork field if non-nil, zero value otherwise.

### GetWirelessNetworkOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetWirelessNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetWirelessNetworkOk returns a tuple with the WirelessNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetwork

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetWirelessNetwork(v BulkWritableCableRequestStatus)`

SetWirelessNetwork sets WirelessNetwork field to given value.

### HasWirelessNetwork

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasWirelessNetwork() bool`

HasWirelessNetwork returns a boolean if a field has been set.

### GetVlan

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetVlan() BulkWritableCircuitRequestTenant`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetVlan(v BulkWritableCircuitRequestTenant)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PatchedBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


