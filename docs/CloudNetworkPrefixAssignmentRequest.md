# CloudNetworkPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNetwork** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Prefix** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewCloudNetworkPrefixAssignmentRequest

`func NewCloudNetworkPrefixAssignmentRequest(cloudNetwork BulkWritableCableRequestStatus, prefix BulkWritableCableRequestStatus, ) *CloudNetworkPrefixAssignmentRequest`

NewCloudNetworkPrefixAssignmentRequest instantiates a new CloudNetworkPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkPrefixAssignmentRequestWithDefaults

`func NewCloudNetworkPrefixAssignmentRequestWithDefaults() *CloudNetworkPrefixAssignmentRequest`

NewCloudNetworkPrefixAssignmentRequestWithDefaults instantiates a new CloudNetworkPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNetwork

`func (o *CloudNetworkPrefixAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *CloudNetworkPrefixAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *CloudNetworkPrefixAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetPrefix

`func (o *CloudNetworkPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CloudNetworkPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CloudNetworkPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


