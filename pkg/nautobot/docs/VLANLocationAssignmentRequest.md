# VLANLocationAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewVLANLocationAssignmentRequest

`func NewVLANLocationAssignmentRequest(vlan BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, ) *VLANLocationAssignmentRequest`

NewVLANLocationAssignmentRequest instantiates a new VLANLocationAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANLocationAssignmentRequestWithDefaults

`func NewVLANLocationAssignmentRequestWithDefaults() *VLANLocationAssignmentRequest`

NewVLANLocationAssignmentRequestWithDefaults instantiates a new VLANLocationAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *VLANLocationAssignmentRequest) GetVlan() BulkWritableCableRequestStatus`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VLANLocationAssignmentRequest) GetVlanOk() (*BulkWritableCableRequestStatus, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VLANLocationAssignmentRequest) SetVlan(v BulkWritableCableRequestStatus)`

SetVlan sets Vlan field to given value.


### GetLocation

`func (o *VLANLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VLANLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VLANLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


