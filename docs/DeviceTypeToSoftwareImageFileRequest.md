# DeviceTypeToSoftwareImageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeviceType** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**SoftwareImageFile** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewDeviceTypeToSoftwareImageFileRequest

`func NewDeviceTypeToSoftwareImageFileRequest(deviceType ApprovalWorkflowStageResponseApprovalWorkflowStage, softwareImageFile ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *DeviceTypeToSoftwareImageFileRequest`

NewDeviceTypeToSoftwareImageFileRequest instantiates a new DeviceTypeToSoftwareImageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTypeToSoftwareImageFileRequestWithDefaults

`func NewDeviceTypeToSoftwareImageFileRequestWithDefaults() *DeviceTypeToSoftwareImageFileRequest`

NewDeviceTypeToSoftwareImageFileRequestWithDefaults instantiates a new DeviceTypeToSoftwareImageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceTypeToSoftwareImageFileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceTypeToSoftwareImageFileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceTypeToSoftwareImageFileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceTypeToSoftwareImageFileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceTypeToSoftwareImageFileRequest) GetDeviceType() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceTypeToSoftwareImageFileRequest) GetDeviceTypeOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceTypeToSoftwareImageFileRequest) SetDeviceType(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDeviceType sets DeviceType field to given value.


### GetSoftwareImageFile

`func (o *DeviceTypeToSoftwareImageFileRequest) GetSoftwareImageFile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetSoftwareImageFile returns the SoftwareImageFile field if non-nil, zero value otherwise.

### GetSoftwareImageFileOk

`func (o *DeviceTypeToSoftwareImageFileRequest) GetSoftwareImageFileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetSoftwareImageFileOk returns a tuple with the SoftwareImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFile

`func (o *DeviceTypeToSoftwareImageFileRequest) SetSoftwareImageFile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetSoftwareImageFile sets SoftwareImageFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


