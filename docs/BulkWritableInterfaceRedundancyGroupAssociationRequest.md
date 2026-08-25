# BulkWritableInterfaceRedundancyGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Priority** | **int32** |  | 
**InterfaceRedundancyGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Interface** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableInterfaceRedundancyGroupAssociationRequest

`func NewBulkWritableInterfaceRedundancyGroupAssociationRequest(id string, priority int32, interfaceRedundancyGroup BulkWritableCableRequestStatus, interface_ BulkWritableCableRequestStatus, ) *BulkWritableInterfaceRedundancyGroupAssociationRequest`

NewBulkWritableInterfaceRedundancyGroupAssociationRequest instantiates a new BulkWritableInterfaceRedundancyGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableInterfaceRedundancyGroupAssociationRequestWithDefaults

`func NewBulkWritableInterfaceRedundancyGroupAssociationRequestWithDefaults() *BulkWritableInterfaceRedundancyGroupAssociationRequest`

NewBulkWritableInterfaceRedundancyGroupAssociationRequestWithDefaults instantiates a new BulkWritableInterfaceRedundancyGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPriority

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetInterfaceRedundancyGroup

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroup() BulkWritableCableRequestStatus`

GetInterfaceRedundancyGroup returns the InterfaceRedundancyGroup field if non-nil, zero value otherwise.

### GetInterfaceRedundancyGroupOk

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceRedundancyGroupOk returns a tuple with the InterfaceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyGroup

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) SetInterfaceRedundancyGroup(v BulkWritableCableRequestStatus)`

SetInterfaceRedundancyGroup sets InterfaceRedundancyGroup field to given value.


### GetInterface

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BulkWritableInterfaceRedundancyGroupAssociationRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


