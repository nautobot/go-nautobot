# CloudServiceNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CloudNetwork** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CloudService** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewCloudServiceNetworkAssignmentRequest

`func NewCloudServiceNetworkAssignmentRequest(cloudNetwork ApprovalWorkflowStageResponseApprovalWorkflowStage, cloudService ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *CloudServiceNetworkAssignmentRequest`

NewCloudServiceNetworkAssignmentRequest instantiates a new CloudServiceNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceNetworkAssignmentRequestWithDefaults

`func NewCloudServiceNetworkAssignmentRequestWithDefaults() *CloudServiceNetworkAssignmentRequest`

NewCloudServiceNetworkAssignmentRequestWithDefaults instantiates a new CloudServiceNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudServiceNetworkAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudServiceNetworkAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudServiceNetworkAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudServiceNetworkAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloudNetwork

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudNetwork() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudNetworkOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *CloudServiceNetworkAssignmentRequest) SetCloudNetwork(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetCloudService

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudService() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudServiceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *CloudServiceNetworkAssignmentRequest) SetCloudService(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCloudService sets CloudService field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


