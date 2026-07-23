# PrefixLocationAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Prefix** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Location** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewPrefixLocationAssignmentRequest

`func NewPrefixLocationAssignmentRequest(prefix ApprovalWorkflowStageResponseApprovalWorkflowStage, location ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *PrefixLocationAssignmentRequest`

NewPrefixLocationAssignmentRequest instantiates a new PrefixLocationAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixLocationAssignmentRequestWithDefaults

`func NewPrefixLocationAssignmentRequestWithDefaults() *PrefixLocationAssignmentRequest`

NewPrefixLocationAssignmentRequestWithDefaults instantiates a new PrefixLocationAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrefixLocationAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrefixLocationAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrefixLocationAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrefixLocationAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrefix

`func (o *PrefixLocationAssignmentRequest) GetPrefix() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixLocationAssignmentRequest) GetPrefixOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixLocationAssignmentRequest) SetPrefix(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPrefix sets Prefix field to given value.


### GetLocation

`func (o *PrefixLocationAssignmentRequest) GetLocation() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrefixLocationAssignmentRequest) GetLocationOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrefixLocationAssignmentRequest) SetLocation(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


