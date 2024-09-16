# PatchedVLANLocationAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedVLANLocationAssignmentRequest

`func NewPatchedVLANLocationAssignmentRequest() *PatchedVLANLocationAssignmentRequest`

NewPatchedVLANLocationAssignmentRequest instantiates a new PatchedVLANLocationAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVLANLocationAssignmentRequestWithDefaults

`func NewPatchedVLANLocationAssignmentRequestWithDefaults() *PatchedVLANLocationAssignmentRequest`

NewPatchedVLANLocationAssignmentRequestWithDefaults instantiates a new PatchedVLANLocationAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *PatchedVLANLocationAssignmentRequest) GetVlan() BulkWritableCableRequestStatus`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedVLANLocationAssignmentRequest) GetVlanOk() (*BulkWritableCableRequestStatus, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedVLANLocationAssignmentRequest) SetVlan(v BulkWritableCableRequestStatus)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedVLANLocationAssignmentRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedVLANLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedVLANLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedVLANLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedVLANLocationAssignmentRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


