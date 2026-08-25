# PatchedBulkWritablePrefixLocationAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Prefix** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritablePrefixLocationAssignmentRequest

`func NewPatchedBulkWritablePrefixLocationAssignmentRequest(id string, ) *PatchedBulkWritablePrefixLocationAssignmentRequest`

NewPatchedBulkWritablePrefixLocationAssignmentRequest instantiates a new PatchedBulkWritablePrefixLocationAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritablePrefixLocationAssignmentRequestWithDefaults

`func NewPatchedBulkWritablePrefixLocationAssignmentRequestWithDefaults() *PatchedBulkWritablePrefixLocationAssignmentRequest`

NewPatchedBulkWritablePrefixLocationAssignmentRequestWithDefaults instantiates a new PatchedBulkWritablePrefixLocationAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPrefix

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


