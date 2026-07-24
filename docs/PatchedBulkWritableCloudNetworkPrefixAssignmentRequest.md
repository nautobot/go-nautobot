# PatchedBulkWritableCloudNetworkPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CloudNetwork** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Prefix** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableCloudNetworkPrefixAssignmentRequest

`func NewPatchedBulkWritableCloudNetworkPrefixAssignmentRequest(id string, ) *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest`

NewPatchedBulkWritableCloudNetworkPrefixAssignmentRequest instantiates a new PatchedBulkWritableCloudNetworkPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCloudNetworkPrefixAssignmentRequestWithDefaults

`func NewPatchedBulkWritableCloudNetworkPrefixAssignmentRequestWithDefaults() *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest`

NewPatchedBulkWritableCloudNetworkPrefixAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableCloudNetworkPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCloudNetwork

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.

### HasCloudNetwork

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) HasCloudNetwork() bool`

HasCloudNetwork returns a boolean if a field has been set.

### GetPrefix

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedBulkWritableCloudNetworkPrefixAssignmentRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


