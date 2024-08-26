# DynamicGroupMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 
**Weight** | **int32** |  | 
**Group** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**ParentGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewDynamicGroupMembershipRequest

`func NewDynamicGroupMembershipRequest(operator OperatorEnum, weight int32, group BulkWritableCableRequestStatus, parentGroup BulkWritableCableRequestStatus, ) *DynamicGroupMembershipRequest`

NewDynamicGroupMembershipRequest instantiates a new DynamicGroupMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicGroupMembershipRequestWithDefaults

`func NewDynamicGroupMembershipRequestWithDefaults() *DynamicGroupMembershipRequest`

NewDynamicGroupMembershipRequestWithDefaults instantiates a new DynamicGroupMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *DynamicGroupMembershipRequest) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DynamicGroupMembershipRequest) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DynamicGroupMembershipRequest) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.


### GetWeight

`func (o *DynamicGroupMembershipRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DynamicGroupMembershipRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DynamicGroupMembershipRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetGroup

`func (o *DynamicGroupMembershipRequest) GetGroup() BulkWritableCableRequestStatus`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DynamicGroupMembershipRequest) GetGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DynamicGroupMembershipRequest) SetGroup(v BulkWritableCableRequestStatus)`

SetGroup sets Group field to given value.


### GetParentGroup

`func (o *DynamicGroupMembershipRequest) GetParentGroup() BulkWritableCableRequestStatus`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *DynamicGroupMembershipRequest) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *DynamicGroupMembershipRequest) SetParentGroup(v BulkWritableCableRequestStatus)`

SetParentGroup sets ParentGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


