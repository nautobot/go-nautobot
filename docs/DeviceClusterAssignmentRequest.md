# DeviceClusterAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Device** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Cluster** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewDeviceClusterAssignmentRequest

`func NewDeviceClusterAssignmentRequest(device BulkWritableCableRequestStatus, cluster BulkWritableCableRequestStatus, ) *DeviceClusterAssignmentRequest`

NewDeviceClusterAssignmentRequest instantiates a new DeviceClusterAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceClusterAssignmentRequestWithDefaults

`func NewDeviceClusterAssignmentRequestWithDefaults() *DeviceClusterAssignmentRequest`

NewDeviceClusterAssignmentRequestWithDefaults instantiates a new DeviceClusterAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceClusterAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceClusterAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceClusterAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceClusterAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDevice

`func (o *DeviceClusterAssignmentRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DeviceClusterAssignmentRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DeviceClusterAssignmentRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.


### GetCluster

`func (o *DeviceClusterAssignmentRequest) GetCluster() BulkWritableCableRequestStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeviceClusterAssignmentRequest) GetClusterOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeviceClusterAssignmentRequest) SetCluster(v BulkWritableCableRequestStatus)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


