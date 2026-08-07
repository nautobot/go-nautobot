# BulkWritableVRFPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Vrf** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Prefix** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableVRFPrefixAssignmentRequest

`func NewBulkWritableVRFPrefixAssignmentRequest(id string, vrf BulkWritableCableRequestStatus, prefix BulkWritableCableRequestStatus, ) *BulkWritableVRFPrefixAssignmentRequest`

NewBulkWritableVRFPrefixAssignmentRequest instantiates a new BulkWritableVRFPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVRFPrefixAssignmentRequestWithDefaults

`func NewBulkWritableVRFPrefixAssignmentRequestWithDefaults() *BulkWritableVRFPrefixAssignmentRequest`

NewBulkWritableVRFPrefixAssignmentRequestWithDefaults instantiates a new BulkWritableVRFPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVRFPrefixAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVRFPrefixAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVRFPrefixAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVrf

`func (o *BulkWritableVRFPrefixAssignmentRequest) GetVrf() BulkWritableCableRequestStatus`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *BulkWritableVRFPrefixAssignmentRequest) GetVrfOk() (*BulkWritableCableRequestStatus, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *BulkWritableVRFPrefixAssignmentRequest) SetVrf(v BulkWritableCableRequestStatus)`

SetVrf sets Vrf field to given value.


### GetPrefix

`func (o *BulkWritableVRFPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BulkWritableVRFPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BulkWritableVRFPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


