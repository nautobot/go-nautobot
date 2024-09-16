# PatchedDynamicGroupMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**OperatorEnum**](OperatorEnum.md) |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Group** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ParentGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedDynamicGroupMembershipRequest

`func NewPatchedDynamicGroupMembershipRequest() *PatchedDynamicGroupMembershipRequest`

NewPatchedDynamicGroupMembershipRequest instantiates a new PatchedDynamicGroupMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDynamicGroupMembershipRequestWithDefaults

`func NewPatchedDynamicGroupMembershipRequestWithDefaults() *PatchedDynamicGroupMembershipRequest`

NewPatchedDynamicGroupMembershipRequestWithDefaults instantiates a new PatchedDynamicGroupMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *PatchedDynamicGroupMembershipRequest) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PatchedDynamicGroupMembershipRequest) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PatchedDynamicGroupMembershipRequest) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PatchedDynamicGroupMembershipRequest) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedDynamicGroupMembershipRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedDynamicGroupMembershipRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedDynamicGroupMembershipRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedDynamicGroupMembershipRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedDynamicGroupMembershipRequest) GetGroup() BulkWritableCableRequestStatus`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedDynamicGroupMembershipRequest) GetGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedDynamicGroupMembershipRequest) SetGroup(v BulkWritableCableRequestStatus)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedDynamicGroupMembershipRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetParentGroup

`func (o *PatchedDynamicGroupMembershipRequest) GetParentGroup() BulkWritableCableRequestStatus`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *PatchedDynamicGroupMembershipRequest) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *PatchedDynamicGroupMembershipRequest) SetParentGroup(v BulkWritableCableRequestStatus)`

SetParentGroup sets ParentGroup field to given value.

### HasParentGroup

`func (o *PatchedDynamicGroupMembershipRequest) HasParentGroup() bool`

HasParentGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


