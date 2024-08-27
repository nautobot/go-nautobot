# CloudServiceNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNetwork** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CloudService** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewCloudServiceNetworkAssignmentRequest

`func NewCloudServiceNetworkAssignmentRequest(cloudNetwork BulkWritableCableRequestStatus, cloudService BulkWritableCableRequestStatus, ) *CloudServiceNetworkAssignmentRequest`

NewCloudServiceNetworkAssignmentRequest instantiates a new CloudServiceNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceNetworkAssignmentRequestWithDefaults

`func NewCloudServiceNetworkAssignmentRequestWithDefaults() *CloudServiceNetworkAssignmentRequest`

NewCloudServiceNetworkAssignmentRequestWithDefaults instantiates a new CloudServiceNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNetwork

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *CloudServiceNetworkAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetCloudService

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudService() BulkWritableCableRequestStatus`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *CloudServiceNetworkAssignmentRequest) GetCloudServiceOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *CloudServiceNetworkAssignmentRequest) SetCloudService(v BulkWritableCableRequestStatus)`

SetCloudService sets CloudService field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


