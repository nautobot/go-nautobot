# PatchedBulkWritableDeviceClusterAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Cluster** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableDeviceClusterAssignmentRequest

`func NewPatchedBulkWritableDeviceClusterAssignmentRequest(id string, ) *PatchedBulkWritableDeviceClusterAssignmentRequest`

NewPatchedBulkWritableDeviceClusterAssignmentRequest instantiates a new PatchedBulkWritableDeviceClusterAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableDeviceClusterAssignmentRequestWithDefaults

`func NewPatchedBulkWritableDeviceClusterAssignmentRequestWithDefaults() *PatchedBulkWritableDeviceClusterAssignmentRequest`

NewPatchedBulkWritableDeviceClusterAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableDeviceClusterAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) GetDevice() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) GetDeviceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) SetDevice(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetCluster

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) GetCluster() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) GetClusterOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) SetCluster(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedBulkWritableDeviceClusterAssignmentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


