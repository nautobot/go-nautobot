# PatchedInterfaceRedundancyGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** |  | [optional] 
**InterfaceRedundancyGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Interface** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedInterfaceRedundancyGroupAssociationRequest

`func NewPatchedInterfaceRedundancyGroupAssociationRequest() *PatchedInterfaceRedundancyGroupAssociationRequest`

NewPatchedInterfaceRedundancyGroupAssociationRequest instantiates a new PatchedInterfaceRedundancyGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInterfaceRedundancyGroupAssociationRequestWithDefaults

`func NewPatchedInterfaceRedundancyGroupAssociationRequestWithDefaults() *PatchedInterfaceRedundancyGroupAssociationRequest`

NewPatchedInterfaceRedundancyGroupAssociationRequestWithDefaults instantiates a new PatchedInterfaceRedundancyGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetInterfaceRedundancyGroup

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroup() BulkWritableCableRequestStatus`

GetInterfaceRedundancyGroup returns the InterfaceRedundancyGroup field if non-nil, zero value otherwise.

### GetInterfaceRedundancyGroupOk

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceRedundancyGroupOk returns a tuple with the InterfaceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyGroup

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) SetInterfaceRedundancyGroup(v BulkWritableCableRequestStatus)`

SetInterfaceRedundancyGroup sets InterfaceRedundancyGroup field to given value.

### HasInterfaceRedundancyGroup

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) HasInterfaceRedundancyGroup() bool`

HasInterfaceRedundancyGroup returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedInterfaceRedundancyGroupAssociationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


