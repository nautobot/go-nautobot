# VRFPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Vrf** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Prefix** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewVRFPrefixAssignmentRequest

`func NewVRFPrefixAssignmentRequest(vrf ApprovalWorkflowStageResponseApprovalWorkflowStage, prefix ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *VRFPrefixAssignmentRequest`

NewVRFPrefixAssignmentRequest instantiates a new VRFPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVRFPrefixAssignmentRequestWithDefaults

`func NewVRFPrefixAssignmentRequestWithDefaults() *VRFPrefixAssignmentRequest`

NewVRFPrefixAssignmentRequestWithDefaults instantiates a new VRFPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VRFPrefixAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VRFPrefixAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VRFPrefixAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VRFPrefixAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVrf

`func (o *VRFPrefixAssignmentRequest) GetVrf() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VRFPrefixAssignmentRequest) GetVrfOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VRFPrefixAssignmentRequest) SetVrf(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVrf sets Vrf field to given value.


### GetPrefix

`func (o *VRFPrefixAssignmentRequest) GetPrefix() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VRFPrefixAssignmentRequest) GetPrefixOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VRFPrefixAssignmentRequest) SetPrefix(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


