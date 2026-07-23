# BulkWritableDeviceClusterAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Cluster** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableDeviceClusterAssignmentRequest

`func NewBulkWritableDeviceClusterAssignmentRequest(id string, device BulkWritableCableRequestStatus, cluster BulkWritableCableRequestStatus, ) *BulkWritableDeviceClusterAssignmentRequest`

NewBulkWritableDeviceClusterAssignmentRequest instantiates a new BulkWritableDeviceClusterAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDeviceClusterAssignmentRequestWithDefaults

`func NewBulkWritableDeviceClusterAssignmentRequestWithDefaults() *BulkWritableDeviceClusterAssignmentRequest`

NewBulkWritableDeviceClusterAssignmentRequestWithDefaults instantiates a new BulkWritableDeviceClusterAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDeviceClusterAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkWritableDeviceClusterAssignmentRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.


### GetCluster

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetCluster() BulkWritableCableRequestStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BulkWritableDeviceClusterAssignmentRequest) GetClusterOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BulkWritableDeviceClusterAssignmentRequest) SetCluster(v BulkWritableCableRequestStatus)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


