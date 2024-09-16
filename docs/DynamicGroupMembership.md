# DynamicGroupMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 
**Weight** | **int32** |  | 
**Group** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**ParentGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewDynamicGroupMembership

`func NewDynamicGroupMembership(id string, objectType string, display string, url string, naturalSlug string, operator OperatorEnum, weight int32, group BulkWritableCableRequestStatus, parentGroup BulkWritableCableRequestStatus, ) *DynamicGroupMembership`

NewDynamicGroupMembership instantiates a new DynamicGroupMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicGroupMembershipWithDefaults

`func NewDynamicGroupMembershipWithDefaults() *DynamicGroupMembership`

NewDynamicGroupMembershipWithDefaults instantiates a new DynamicGroupMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DynamicGroupMembership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicGroupMembership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicGroupMembership) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *DynamicGroupMembership) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DynamicGroupMembership) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DynamicGroupMembership) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *DynamicGroupMembership) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DynamicGroupMembership) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DynamicGroupMembership) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *DynamicGroupMembership) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DynamicGroupMembership) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DynamicGroupMembership) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *DynamicGroupMembership) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *DynamicGroupMembership) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *DynamicGroupMembership) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetOperator

`func (o *DynamicGroupMembership) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DynamicGroupMembership) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DynamicGroupMembership) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.


### GetWeight

`func (o *DynamicGroupMembership) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DynamicGroupMembership) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DynamicGroupMembership) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetGroup

`func (o *DynamicGroupMembership) GetGroup() BulkWritableCableRequestStatus`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DynamicGroupMembership) GetGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DynamicGroupMembership) SetGroup(v BulkWritableCableRequestStatus)`

SetGroup sets Group field to given value.


### GetParentGroup

`func (o *DynamicGroupMembership) GetParentGroup() BulkWritableCableRequestStatus`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *DynamicGroupMembership) GetParentGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *DynamicGroupMembership) SetParentGroup(v BulkWritableCableRequestStatus)`

SetParentGroup sets ParentGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


