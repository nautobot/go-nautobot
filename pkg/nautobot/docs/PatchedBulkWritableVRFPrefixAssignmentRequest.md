# PatchedBulkWritableVRFPrefixAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Vrf** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Prefix** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVRFPrefixAssignmentRequest

`func NewPatchedBulkWritableVRFPrefixAssignmentRequest(id string, ) *PatchedBulkWritableVRFPrefixAssignmentRequest`

NewPatchedBulkWritableVRFPrefixAssignmentRequest instantiates a new PatchedBulkWritableVRFPrefixAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVRFPrefixAssignmentRequestWithDefaults

`func NewPatchedBulkWritableVRFPrefixAssignmentRequestWithDefaults() *PatchedBulkWritableVRFPrefixAssignmentRequest`

NewPatchedBulkWritableVRFPrefixAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableVRFPrefixAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVrf

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) GetVrf() BulkWritableCableRequestStatus`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) GetVrfOk() (*BulkWritableCableRequestStatus, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) SetVrf(v BulkWritableCableRequestStatus)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetPrefix

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedBulkWritableVRFPrefixAssignmentRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


