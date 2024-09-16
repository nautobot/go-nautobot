# BulkWritableCloudServiceNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CloudNetwork** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CloudService** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableCloudServiceNetworkAssignmentRequest

`func NewBulkWritableCloudServiceNetworkAssignmentRequest(id string, cloudNetwork BulkWritableCableRequestStatus, cloudService BulkWritableCableRequestStatus, ) *BulkWritableCloudServiceNetworkAssignmentRequest`

NewBulkWritableCloudServiceNetworkAssignmentRequest instantiates a new BulkWritableCloudServiceNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCloudServiceNetworkAssignmentRequestWithDefaults

`func NewBulkWritableCloudServiceNetworkAssignmentRequestWithDefaults() *BulkWritableCloudServiceNetworkAssignmentRequest`

NewBulkWritableCloudServiceNetworkAssignmentRequestWithDefaults instantiates a new BulkWritableCloudServiceNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCloudNetwork

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetCloudService

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) GetCloudService() BulkWritableCableRequestStatus`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) GetCloudServiceOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *BulkWritableCloudServiceNetworkAssignmentRequest) SetCloudService(v BulkWritableCableRequestStatus)`

SetCloudService sets CloudService field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


