# BulkWritableDeviceClusterAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Cluster** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewBulkWritableDeviceClusterAssignmentRequest

`func NewBulkWritableDeviceClusterAssignmentRequest(id string, device ApprovalWorkflowStageResponseApprovalWorkflowStage, cluster ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableDeviceClusterAssignmentRequest`

NewBulkWritableDeviceClusterAssignmentRequest instantiates a new BulkWritableDeviceClusterAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDeviceClusterAssignmentRequestWithDefaults

`func NewBulkWritableDeviceClusterAssignmentRequestWithDefaults() *BulkWritableDeviceClusterAssignmentRequest`

NewBulkWritableDeviceClusterAssignmentRequestWithDefaults instantiates a new BulkWritableDeviceClusterAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDeviceClusterAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetDevice() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetDeviceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkWritableDeviceClusterAssignmentRequest) SetDevice(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDevice sets Device field to given value.


### GetCluster

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetCluster() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetClusterOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BulkWritableDeviceClusterAssignmentRequest) SetCluster(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


