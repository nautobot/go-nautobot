# PatchedBulkWritableCloudServiceNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CloudNetwork** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CloudService** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableCloudServiceNetworkAssignmentRequest

`func NewPatchedBulkWritableCloudServiceNetworkAssignmentRequest(id string, ) *PatchedBulkWritableCloudServiceNetworkAssignmentRequest`

NewPatchedBulkWritableCloudServiceNetworkAssignmentRequest instantiates a new PatchedBulkWritableCloudServiceNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCloudServiceNetworkAssignmentRequestWithDefaults

`func NewPatchedBulkWritableCloudServiceNetworkAssignmentRequestWithDefaults() *PatchedBulkWritableCloudServiceNetworkAssignmentRequest`

NewPatchedBulkWritableCloudServiceNetworkAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableCloudServiceNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCloudNetwork

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.

### HasCloudNetwork

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) HasCloudNetwork() bool`

HasCloudNetwork returns a boolean if a field has been set.

### GetCloudService

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) GetCloudService() BulkWritableCableRequestStatus`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) GetCloudServiceOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) SetCloudService(v BulkWritableCableRequestStatus)`

SetCloudService sets CloudService field to given value.

### HasCloudService

`func (o *PatchedBulkWritableCloudServiceNetworkAssignmentRequest) HasCloudService() bool`

HasCloudService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


