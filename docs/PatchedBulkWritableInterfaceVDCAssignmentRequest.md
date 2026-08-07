# PatchedBulkWritableInterfaceVDCAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VirtualDeviceContext** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Interface** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableInterfaceVDCAssignmentRequest

`func NewPatchedBulkWritableInterfaceVDCAssignmentRequest(id string, ) *PatchedBulkWritableInterfaceVDCAssignmentRequest`

NewPatchedBulkWritableInterfaceVDCAssignmentRequest instantiates a new PatchedBulkWritableInterfaceVDCAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableInterfaceVDCAssignmentRequestWithDefaults

`func NewPatchedBulkWritableInterfaceVDCAssignmentRequestWithDefaults() *PatchedBulkWritableInterfaceVDCAssignmentRequest`

NewPatchedBulkWritableInterfaceVDCAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableInterfaceVDCAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVirtualDeviceContext

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) GetVirtualDeviceContext() BulkWritableCableRequestStatus`

GetVirtualDeviceContext returns the VirtualDeviceContext field if non-nil, zero value otherwise.

### GetVirtualDeviceContextOk

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) GetVirtualDeviceContextOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualDeviceContextOk returns a tuple with the VirtualDeviceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceContext

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) SetVirtualDeviceContext(v BulkWritableCableRequestStatus)`

SetVirtualDeviceContext sets VirtualDeviceContext field to given value.

### HasVirtualDeviceContext

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) HasVirtualDeviceContext() bool`

HasVirtualDeviceContext returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedBulkWritableInterfaceVDCAssignmentRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


