# InterfaceRedundancyGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **int32** |  | 
**InterfaceRedundancyGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Interface** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewInterfaceRedundancyGroupAssociationRequest

`func NewInterfaceRedundancyGroupAssociationRequest(priority int32, interfaceRedundancyGroup BulkWritableCableRequestStatus, interface_ BulkWritableCableRequestStatus, ) *InterfaceRedundancyGroupAssociationRequest`

NewInterfaceRedundancyGroupAssociationRequest instantiates a new InterfaceRedundancyGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceRedundancyGroupAssociationRequestWithDefaults

`func NewInterfaceRedundancyGroupAssociationRequestWithDefaults() *InterfaceRedundancyGroupAssociationRequest`

NewInterfaceRedundancyGroupAssociationRequestWithDefaults instantiates a new InterfaceRedundancyGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *InterfaceRedundancyGroupAssociationRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InterfaceRedundancyGroupAssociationRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InterfaceRedundancyGroupAssociationRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetInterfaceRedundancyGroup

`func (o *InterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroup() BulkWritableCableRequestStatus`

GetInterfaceRedundancyGroup returns the InterfaceRedundancyGroup field if non-nil, zero value otherwise.

### GetInterfaceRedundancyGroupOk

`func (o *InterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceRedundancyGroupOk returns a tuple with the InterfaceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyGroup

`func (o *InterfaceRedundancyGroupAssociationRequest) SetInterfaceRedundancyGroup(v BulkWritableCableRequestStatus)`

SetInterfaceRedundancyGroup sets InterfaceRedundancyGroup field to given value.


### GetInterface

`func (o *InterfaceRedundancyGroupAssociationRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InterfaceRedundancyGroupAssociationRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InterfaceRedundancyGroupAssociationRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


