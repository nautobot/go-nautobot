# PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ControllerManagedDeviceGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**WirelessNetwork** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest

`func NewPatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest() *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest`

NewPatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest instantiates a new PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults

`func NewPatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults() *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest`

NewPatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults instantiates a new PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetControllerManagedDeviceGroup

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetControllerManagedDeviceGroup() BulkWritableCableRequestStatus`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetControllerManagedDeviceGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetControllerManagedDeviceGroup(v BulkWritableCableRequestStatus)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.

### HasControllerManagedDeviceGroup

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasControllerManagedDeviceGroup() bool`

HasControllerManagedDeviceGroup returns a boolean if a field has been set.

### GetWirelessNetwork

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetWirelessNetwork() BulkWritableCableRequestStatus`

GetWirelessNetwork returns the WirelessNetwork field if non-nil, zero value otherwise.

### GetWirelessNetworkOk

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetWirelessNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetWirelessNetworkOk returns a tuple with the WirelessNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetwork

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetWirelessNetwork(v BulkWritableCableRequestStatus)`

SetWirelessNetwork sets WirelessNetwork field to given value.

### HasWirelessNetwork

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasWirelessNetwork() bool`

HasWirelessNetwork returns a boolean if a field has been set.

### GetVlan

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetVlan() BulkWritableCircuitRequestTenant`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetVlan(v BulkWritableCircuitRequestTenant)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PatchedControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


