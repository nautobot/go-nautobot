# InterfaceVDCAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VirtualDeviceContext** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Interface** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewInterfaceVDCAssignmentRequest

`func NewInterfaceVDCAssignmentRequest(virtualDeviceContext ApprovalWorkflowStageResponseApprovalWorkflowStage, interface_ ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *InterfaceVDCAssignmentRequest`

NewInterfaceVDCAssignmentRequest instantiates a new InterfaceVDCAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceVDCAssignmentRequestWithDefaults

`func NewInterfaceVDCAssignmentRequestWithDefaults() *InterfaceVDCAssignmentRequest`

NewInterfaceVDCAssignmentRequestWithDefaults instantiates a new InterfaceVDCAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceVDCAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceVDCAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceVDCAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InterfaceVDCAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVirtualDeviceContext

`func (o *InterfaceVDCAssignmentRequest) GetVirtualDeviceContext() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVirtualDeviceContext returns the VirtualDeviceContext field if non-nil, zero value otherwise.

### GetVirtualDeviceContextOk

`func (o *InterfaceVDCAssignmentRequest) GetVirtualDeviceContextOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVirtualDeviceContextOk returns a tuple with the VirtualDeviceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceContext

`func (o *InterfaceVDCAssignmentRequest) SetVirtualDeviceContext(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVirtualDeviceContext sets VirtualDeviceContext field to given value.


### GetInterface

`func (o *InterfaceVDCAssignmentRequest) GetInterface() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InterfaceVDCAssignmentRequest) GetInterfaceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InterfaceVDCAssignmentRequest) SetInterface(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetInterface sets Interface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


