# PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Priority** | Pointer to **int32** |  | [optional] 
**InterfaceRedundancyGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Interface** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableInterfaceRedundancyGroupAssociationRequest

`func NewPatchedBulkWritableInterfaceRedundancyGroupAssociationRequest(id string, ) *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest`

NewPatchedBulkWritableInterfaceRedundancyGroupAssociationRequest instantiates a new PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableInterfaceRedundancyGroupAssociationRequestWithDefaults

`func NewPatchedBulkWritableInterfaceRedundancyGroupAssociationRequestWithDefaults() *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest`

NewPatchedBulkWritableInterfaceRedundancyGroupAssociationRequestWithDefaults instantiates a new PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPriority

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetInterfaceRedundancyGroup

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroup() BulkWritableCableRequestStatus`

GetInterfaceRedundancyGroup returns the InterfaceRedundancyGroup field if non-nil, zero value otherwise.

### GetInterfaceRedundancyGroupOk

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterfaceRedundancyGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceRedundancyGroupOk returns a tuple with the InterfaceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyGroup

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) SetInterfaceRedundancyGroup(v BulkWritableCableRequestStatus)`

SetInterfaceRedundancyGroup sets InterfaceRedundancyGroup field to given value.

### HasInterfaceRedundancyGroup

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) HasInterfaceRedundancyGroup() bool`

HasInterfaceRedundancyGroup returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedBulkWritableInterfaceRedundancyGroupAssociationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


