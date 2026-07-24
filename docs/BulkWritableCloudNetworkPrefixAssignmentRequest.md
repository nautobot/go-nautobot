# BulkWritableCloudNetworkPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CloudNetwork** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Prefix** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableCloudNetworkPrefixAssignmentRequest

`func NewBulkWritableCloudNetworkPrefixAssignmentRequest(id string, cloudNetwork BulkWritableCableRequestStatus, prefix BulkWritableCableRequestStatus, ) *BulkWritableCloudNetworkPrefixAssignmentRequest`

NewBulkWritableCloudNetworkPrefixAssignmentRequest instantiates a new BulkWritableCloudNetworkPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCloudNetworkPrefixAssignmentRequestWithDefaults

`func NewBulkWritableCloudNetworkPrefixAssignmentRequestWithDefaults() *BulkWritableCloudNetworkPrefixAssignmentRequest`

NewBulkWritableCloudNetworkPrefixAssignmentRequestWithDefaults instantiates a new BulkWritableCloudNetworkPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCloudNetwork

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetPrefix

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BulkWritableCloudNetworkPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


