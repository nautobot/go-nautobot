# PatchedBulkWritableDynamicGroupMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Operator** | Pointer to [**OperatorEnum**](OperatorEnum.md) |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Group** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ParentGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableDynamicGroupMembershipRequest

`func NewPatchedBulkWritableDynamicGroupMembershipRequest(id string, ) *PatchedBulkWritableDynamicGroupMembershipRequest`

NewPatchedBulkWritableDynamicGroupMembershipRequest instantiates a new PatchedBulkWritableDynamicGroupMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableDynamicGroupMembershipRequestWithDefaults

`func NewPatchedBulkWritableDynamicGroupMembershipRequestWithDefaults() *PatchedBulkWritableDynamicGroupMembershipRequest`

NewPatchedBulkWritableDynamicGroupMembershipRequestWithDefaults instantiates a new PatchedBulkWritableDynamicGroupMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOperator

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetGroup() BulkWritableCableRequestStatus`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetGroup(v BulkWritableCableRequestStatus)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetParentGroup

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetParentGroup() BulkWritableCableRequestStatus`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) SetParentGroup(v BulkWritableCableRequestStatus)`

SetParentGroup sets ParentGroup field to given value.

### HasParentGroup

`func (o *PatchedBulkWritableDynamicGroupMembershipRequest) HasParentGroup() bool`

HasParentGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


