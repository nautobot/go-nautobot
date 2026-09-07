# PatchedBulkWritableVLANLocationAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Vlan** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVLANLocationAssignmentRequest

`func NewPatchedBulkWritableVLANLocationAssignmentRequest(id string, ) *PatchedBulkWritableVLANLocationAssignmentRequest`

NewPatchedBulkWritableVLANLocationAssignmentRequest instantiates a new PatchedBulkWritableVLANLocationAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVLANLocationAssignmentRequestWithDefaults

`func NewPatchedBulkWritableVLANLocationAssignmentRequestWithDefaults() *PatchedBulkWritableVLANLocationAssignmentRequest`

NewPatchedBulkWritableVLANLocationAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableVLANLocationAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVlan

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) GetVlan() BulkWritableCableRequestStatus`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) GetVlanOk() (*BulkWritableCableRequestStatus, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) SetVlan(v BulkWritableCableRequestStatus)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkWritableVLANLocationAssignmentRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


