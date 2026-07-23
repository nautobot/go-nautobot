# VirtualServerCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VirtualServer** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CertificateProfile** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewVirtualServerCertificateProfileAssignmentRequest

`func NewVirtualServerCertificateProfileAssignmentRequest(virtualServer ApprovalWorkflowStageResponseApprovalWorkflowStage, certificateProfile ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *VirtualServerCertificateProfileAssignmentRequest`

NewVirtualServerCertificateProfileAssignmentRequest instantiates a new VirtualServerCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualServerCertificateProfileAssignmentRequestWithDefaults

`func NewVirtualServerCertificateProfileAssignmentRequestWithDefaults() *VirtualServerCertificateProfileAssignmentRequest`

NewVirtualServerCertificateProfileAssignmentRequestWithDefaults instantiates a new VirtualServerCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualServerCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualServerCertificateProfileAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVirtualServer

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetVirtualServer() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVirtualServer returns the VirtualServer field if non-nil, zero value otherwise.

### GetVirtualServerOk

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetVirtualServerOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVirtualServerOk returns a tuple with the VirtualServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServer

`func (o *VirtualServerCertificateProfileAssignmentRequest) SetVirtualServer(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVirtualServer sets VirtualServer field to given value.


### GetCertificateProfile

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetCertificateProfile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *VirtualServerCertificateProfileAssignmentRequest) SetCertificateProfile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCertificateProfile sets CertificateProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


