# PatchedBulkWritableVRFDeviceAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Device** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVRFDeviceAssignmentRequest

`func NewPatchedBulkWritableVRFDeviceAssignmentRequest(id string, ) *PatchedBulkWritableVRFDeviceAssignmentRequest`

NewPatchedBulkWritableVRFDeviceAssignmentRequest instantiates a new PatchedBulkWritableVRFDeviceAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVRFDeviceAssignmentRequestWithDefaults

`func NewPatchedBulkWritableVRFDeviceAssignmentRequestWithDefaults() *PatchedBulkWritableVRFDeviceAssignmentRequest`

NewPatchedBulkWritableVRFDeviceAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableVRFDeviceAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRd

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetName

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVrf

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetVrf() BulkWritableCableRequestStatus`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetVrfOk() (*BulkWritableCableRequestStatus, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetVrf(v BulkWritableCableRequestStatus)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetDevice() BulkWritableCircuitRequestTenant`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetDevice(v BulkWritableCircuitRequestTenant)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVirtualMachine

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetVirtualMachine() BulkWritableCircuitRequestTenant`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) GetVirtualMachineOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetVirtualMachine(v BulkWritableCircuitRequestTenant)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *PatchedBulkWritableVRFDeviceAssignmentRequest) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


