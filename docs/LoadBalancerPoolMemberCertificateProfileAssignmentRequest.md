# LoadBalancerPoolMemberCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LoadBalancerPoolMember** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CertificateProfile** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewLoadBalancerPoolMemberCertificateProfileAssignmentRequest

`func NewLoadBalancerPoolMemberCertificateProfileAssignmentRequest(loadBalancerPoolMember ApprovalWorkflowStageResponseApprovalWorkflowStage, certificateProfile ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *LoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewLoadBalancerPoolMemberCertificateProfileAssignmentRequest instantiates a new LoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults

`func NewLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults() *LoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults instantiates a new LoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancerPoolMember

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMember() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLoadBalancerPoolMember returns the LoadBalancerPoolMember field if non-nil, zero value otherwise.

### GetLoadBalancerPoolMemberOk

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMemberOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLoadBalancerPoolMemberOk returns a tuple with the LoadBalancerPoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPoolMember

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetLoadBalancerPoolMember(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLoadBalancerPoolMember sets LoadBalancerPoolMember field to given value.


### GetCertificateProfile

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *LoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetCertificateProfile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCertificateProfile sets CertificateProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


