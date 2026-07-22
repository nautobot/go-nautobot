# PatchedInterfaceVDCAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VirtualDeviceContext** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Interface** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedInterfaceVDCAssignmentRequest

`func NewPatchedInterfaceVDCAssignmentRequest() *PatchedInterfaceVDCAssignmentRequest`

NewPatchedInterfaceVDCAssignmentRequest instantiates a new PatchedInterfaceVDCAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInterfaceVDCAssignmentRequestWithDefaults

`func NewPatchedInterfaceVDCAssignmentRequestWithDefaults() *PatchedInterfaceVDCAssignmentRequest`

NewPatchedInterfaceVDCAssignmentRequestWithDefaults instantiates a new PatchedInterfaceVDCAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedInterfaceVDCAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedInterfaceVDCAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedInterfaceVDCAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedInterfaceVDCAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVirtualDeviceContext

`func (o *PatchedInterfaceVDCAssignmentRequest) GetVirtualDeviceContext() BulkWritableCableRequestStatus`

GetVirtualDeviceContext returns the VirtualDeviceContext field if non-nil, zero value otherwise.

### GetVirtualDeviceContextOk

`func (o *PatchedInterfaceVDCAssignmentRequest) GetVirtualDeviceContextOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualDeviceContextOk returns a tuple with the VirtualDeviceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceContext

`func (o *PatchedInterfaceVDCAssignmentRequest) SetVirtualDeviceContext(v BulkWritableCableRequestStatus)`

SetVirtualDeviceContext sets VirtualDeviceContext field to given value.

### HasVirtualDeviceContext

`func (o *PatchedInterfaceVDCAssignmentRequest) HasVirtualDeviceContext() bool`

HasVirtualDeviceContext returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedInterfaceVDCAssignmentRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedInterfaceVDCAssignmentRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedInterfaceVDCAssignmentRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedInterfaceVDCAssignmentRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


