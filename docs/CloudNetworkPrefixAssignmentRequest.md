# CloudNetworkPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CloudNetwork** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Prefix** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewCloudNetworkPrefixAssignmentRequest

`func NewCloudNetworkPrefixAssignmentRequest(cloudNetwork ApprovalWorkflowStageResponseApprovalWorkflowStage, prefix ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *CloudNetworkPrefixAssignmentRequest`

NewCloudNetworkPrefixAssignmentRequest instantiates a new CloudNetworkPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkPrefixAssignmentRequestWithDefaults

`func NewCloudNetworkPrefixAssignmentRequestWithDefaults() *CloudNetworkPrefixAssignmentRequest`

NewCloudNetworkPrefixAssignmentRequestWithDefaults instantiates a new CloudNetworkPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudNetworkPrefixAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudNetworkPrefixAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudNetworkPrefixAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudNetworkPrefixAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloudNetwork

`func (o *CloudNetworkPrefixAssignmentRequest) GetCloudNetwork() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *CloudNetworkPrefixAssignmentRequest) GetCloudNetworkOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *CloudNetworkPrefixAssignmentRequest) SetCloudNetwork(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetPrefix

`func (o *CloudNetworkPrefixAssignmentRequest) GetPrefix() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CloudNetworkPrefixAssignmentRequest) GetPrefixOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CloudNetworkPrefixAssignmentRequest) SetPrefix(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


