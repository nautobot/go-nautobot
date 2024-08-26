# BulkWritableDynamicGroupMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 
**Weight** | **int32** |  | 
**Group** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**ParentGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableDynamicGroupMembershipRequest

`func NewBulkWritableDynamicGroupMembershipRequest(id string, operator OperatorEnum, weight int32, group BulkWritableCableRequestStatus, parentGroup BulkWritableCableRequestStatus, ) *BulkWritableDynamicGroupMembershipRequest`

NewBulkWritableDynamicGroupMembershipRequest instantiates a new BulkWritableDynamicGroupMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDynamicGroupMembershipRequestWithDefaults

`func NewBulkWritableDynamicGroupMembershipRequestWithDefaults() *BulkWritableDynamicGroupMembershipRequest`

NewBulkWritableDynamicGroupMembershipRequestWithDefaults instantiates a new BulkWritableDynamicGroupMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDynamicGroupMembershipRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDynamicGroupMembershipRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDynamicGroupMembershipRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOperator

`func (o *BulkWritableDynamicGroupMembershipRequest) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BulkWritableDynamicGroupMembershipRequest) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BulkWritableDynamicGroupMembershipRequest) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.


### GetWeight

`func (o *BulkWritableDynamicGroupMembershipRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkWritableDynamicGroupMembershipRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkWritableDynamicGroupMembershipRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetGroup

`func (o *BulkWritableDynamicGroupMembershipRequest) GetGroup() BulkWritableCableRequestStatus`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkWritableDynamicGroupMembershipRequest) GetGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkWritableDynamicGroupMembershipRequest) SetGroup(v BulkWritableCableRequestStatus)`

SetGroup sets Group field to given value.


### GetParentGroup

`func (o *BulkWritableDynamicGroupMembershipRequest) GetParentGroup() BulkWritableCableRequestStatus`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *BulkWritableDynamicGroupMembershipRequest) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *BulkWritableDynamicGroupMembershipRequest) SetParentGroup(v BulkWritableCableRequestStatus)`

SetParentGroup sets ParentGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


