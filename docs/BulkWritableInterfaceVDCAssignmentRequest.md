# BulkWritableInterfaceVDCAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VirtualDeviceContext** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Interface** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableInterfaceVDCAssignmentRequest

`func NewBulkWritableInterfaceVDCAssignmentRequest(id string, virtualDeviceContext BulkWritableCableRequestStatus, interface_ BulkWritableCableRequestStatus, ) *BulkWritableInterfaceVDCAssignmentRequest`

NewBulkWritableInterfaceVDCAssignmentRequest instantiates a new BulkWritableInterfaceVDCAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableInterfaceVDCAssignmentRequestWithDefaults

`func NewBulkWritableInterfaceVDCAssignmentRequestWithDefaults() *BulkWritableInterfaceVDCAssignmentRequest`

NewBulkWritableInterfaceVDCAssignmentRequestWithDefaults instantiates a new BulkWritableInterfaceVDCAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableInterfaceVDCAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableInterfaceVDCAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableInterfaceVDCAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVirtualDeviceContext

`func (o *BulkWritableInterfaceVDCAssignmentRequest) GetVirtualDeviceContext() BulkWritableCableRequestStatus`

GetVirtualDeviceContext returns the VirtualDeviceContext field if non-nil, zero value otherwise.

### GetVirtualDeviceContextOk

`func (o *BulkWritableInterfaceVDCAssignmentRequest) GetVirtualDeviceContextOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualDeviceContextOk returns a tuple with the VirtualDeviceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceContext

`func (o *BulkWritableInterfaceVDCAssignmentRequest) SetVirtualDeviceContext(v BulkWritableCableRequestStatus)`

SetVirtualDeviceContext sets VirtualDeviceContext field to given value.


### GetInterface

`func (o *BulkWritableInterfaceVDCAssignmentRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BulkWritableInterfaceVDCAssignmentRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BulkWritableInterfaceVDCAssignmentRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


