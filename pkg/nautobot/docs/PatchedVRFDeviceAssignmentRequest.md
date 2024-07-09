# PatchedVRFDeviceAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Device** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedVRFDeviceAssignmentRequest

`func NewPatchedVRFDeviceAssignmentRequest() *PatchedVRFDeviceAssignmentRequest`

NewPatchedVRFDeviceAssignmentRequest instantiates a new PatchedVRFDeviceAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVRFDeviceAssignmentRequestWithDefaults

`func NewPatchedVRFDeviceAssignmentRequestWithDefaults() *PatchedVRFDeviceAssignmentRequest`

NewPatchedVRFDeviceAssignmentRequestWithDefaults instantiates a new PatchedVRFDeviceAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRd

`func (o *PatchedVRFDeviceAssignmentRequest) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *PatchedVRFDeviceAssignmentRequest) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *PatchedVRFDeviceAssignmentRequest) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *PatchedVRFDeviceAssignmentRequest) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *PatchedVRFDeviceAssignmentRequest) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetName

`func (o *PatchedVRFDeviceAssignmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVRFDeviceAssignmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVRFDeviceAssignmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVrf

`func (o *PatchedVRFDeviceAssignmentRequest) GetVrf() BulkWritableCableRequestStatus`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetVrfOk() (*BulkWritableCableRequestStatus, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedVRFDeviceAssignmentRequest) SetVrf(v BulkWritableCableRequestStatus)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedVRFDeviceAssignmentRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedVRFDeviceAssignmentRequest) GetDevice() BulkWritableCircuitRequestTenant`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedVRFDeviceAssignmentRequest) SetDevice(v BulkWritableCircuitRequestTenant)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedVRFDeviceAssignmentRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedVRFDeviceAssignmentRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedVRFDeviceAssignmentRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVirtualMachine

`func (o *PatchedVRFDeviceAssignmentRequest) GetVirtualMachine() BulkWritableCircuitRequestTenant`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetVirtualMachineOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedVRFDeviceAssignmentRequest) SetVirtualMachine(v BulkWritableCircuitRequestTenant)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PatchedVRFDeviceAssignmentRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *PatchedVRFDeviceAssignmentRequest) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *PatchedVRFDeviceAssignmentRequest) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


