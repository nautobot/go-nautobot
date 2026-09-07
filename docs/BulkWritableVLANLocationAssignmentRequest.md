# BulkWritableVLANLocationAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Vlan** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableVLANLocationAssignmentRequest

`func NewBulkWritableVLANLocationAssignmentRequest(id string, vlan BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, ) *BulkWritableVLANLocationAssignmentRequest`

NewBulkWritableVLANLocationAssignmentRequest instantiates a new BulkWritableVLANLocationAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVLANLocationAssignmentRequestWithDefaults

`func NewBulkWritableVLANLocationAssignmentRequestWithDefaults() *BulkWritableVLANLocationAssignmentRequest`

NewBulkWritableVLANLocationAssignmentRequestWithDefaults instantiates a new BulkWritableVLANLocationAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVLANLocationAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVLANLocationAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVLANLocationAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVlan

`func (o *BulkWritableVLANLocationAssignmentRequest) GetVlan() BulkWritableCableRequestStatus`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *BulkWritableVLANLocationAssignmentRequest) GetVlanOk() (*BulkWritableCableRequestStatus, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *BulkWritableVLANLocationAssignmentRequest) SetVlan(v BulkWritableCableRequestStatus)`

SetVlan sets Vlan field to given value.


### GetLocation

`func (o *BulkWritableVLANLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BulkWritableVLANLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BulkWritableVLANLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


