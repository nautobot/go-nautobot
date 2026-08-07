# PatchedDeviceClusterAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Cluster** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedDeviceClusterAssignmentRequest

`func NewPatchedDeviceClusterAssignmentRequest() *PatchedDeviceClusterAssignmentRequest`

NewPatchedDeviceClusterAssignmentRequest instantiates a new PatchedDeviceClusterAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDeviceClusterAssignmentRequestWithDefaults

`func NewPatchedDeviceClusterAssignmentRequestWithDefaults() *PatchedDeviceClusterAssignmentRequest`

NewPatchedDeviceClusterAssignmentRequestWithDefaults instantiates a new PatchedDeviceClusterAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDeviceClusterAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDeviceClusterAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDeviceClusterAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDeviceClusterAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedDeviceClusterAssignmentRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedDeviceClusterAssignmentRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedDeviceClusterAssignmentRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedDeviceClusterAssignmentRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetCluster

`func (o *PatchedDeviceClusterAssignmentRequest) GetCluster() BulkWritableCableRequestStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedDeviceClusterAssignmentRequest) GetClusterOk() (*BulkWritableCableRequestStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedDeviceClusterAssignmentRequest) SetCluster(v BulkWritableCableRequestStatus)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedDeviceClusterAssignmentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


