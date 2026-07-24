# PatchedVRFDeviceAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Device** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VirtualDeviceContext** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 

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

### GetId

`func (o *PatchedVRFDeviceAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVRFDeviceAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVRFDeviceAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

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

`func (o *PatchedVRFDeviceAssignmentRequest) GetVrf() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetVrfOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedVRFDeviceAssignmentRequest) SetVrf(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedVRFDeviceAssignmentRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedVRFDeviceAssignmentRequest) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedVRFDeviceAssignmentRequest) SetDevice(v ApprovalWorkflowUser)`

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

`func (o *PatchedVRFDeviceAssignmentRequest) GetVirtualMachine() ApprovalWorkflowUser`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetVirtualMachineOk() (*ApprovalWorkflowUser, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedVRFDeviceAssignmentRequest) SetVirtualMachine(v ApprovalWorkflowUser)`

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
### GetVirtualDeviceContext

`func (o *PatchedVRFDeviceAssignmentRequest) GetVirtualDeviceContext() ApprovalWorkflowUser`

GetVirtualDeviceContext returns the VirtualDeviceContext field if non-nil, zero value otherwise.

### GetVirtualDeviceContextOk

`func (o *PatchedVRFDeviceAssignmentRequest) GetVirtualDeviceContextOk() (*ApprovalWorkflowUser, bool)`

GetVirtualDeviceContextOk returns a tuple with the VirtualDeviceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceContext

`func (o *PatchedVRFDeviceAssignmentRequest) SetVirtualDeviceContext(v ApprovalWorkflowUser)`

SetVirtualDeviceContext sets VirtualDeviceContext field to given value.

### HasVirtualDeviceContext

`func (o *PatchedVRFDeviceAssignmentRequest) HasVirtualDeviceContext() bool`

HasVirtualDeviceContext returns a boolean if a field has been set.

### SetVirtualDeviceContextNil

`func (o *PatchedVRFDeviceAssignmentRequest) SetVirtualDeviceContextNil(b bool)`

 SetVirtualDeviceContextNil sets the value for VirtualDeviceContext to be an explicit nil

### UnsetVirtualDeviceContext
`func (o *PatchedVRFDeviceAssignmentRequest) UnsetVirtualDeviceContext()`

UnsetVirtualDeviceContext ensures that no value is present for VirtualDeviceContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


