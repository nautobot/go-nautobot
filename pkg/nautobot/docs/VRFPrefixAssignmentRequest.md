# VRFPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vrf** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Prefix** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewVRFPrefixAssignmentRequest

`func NewVRFPrefixAssignmentRequest(vrf BulkWritableCableRequestStatus, prefix BulkWritableCableRequestStatus, ) *VRFPrefixAssignmentRequest`

NewVRFPrefixAssignmentRequest instantiates a new VRFPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVRFPrefixAssignmentRequestWithDefaults

`func NewVRFPrefixAssignmentRequestWithDefaults() *VRFPrefixAssignmentRequest`

NewVRFPrefixAssignmentRequestWithDefaults instantiates a new VRFPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVrf

`func (o *VRFPrefixAssignmentRequest) GetVrf() BulkWritableCableRequestStatus`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VRFPrefixAssignmentRequest) GetVrfOk() (*BulkWritableCableRequestStatus, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VRFPrefixAssignmentRequest) SetVrf(v BulkWritableCableRequestStatus)`

SetVrf sets Vrf field to given value.


### GetPrefix

`func (o *VRFPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VRFPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VRFPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


