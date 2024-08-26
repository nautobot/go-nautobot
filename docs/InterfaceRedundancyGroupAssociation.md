# InterfaceRedundancyGroupAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Priority** | **int32** |  | 
**InterfaceRedundancyGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Interface** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewInterfaceRedundancyGroupAssociation

`func NewInterfaceRedundancyGroupAssociation(id string, objectType string, display string, url string, naturalSlug string, priority int32, interfaceRedundancyGroup BulkWritableCableRequestStatus, interface_ BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, ) *InterfaceRedundancyGroupAssociation`

NewInterfaceRedundancyGroupAssociation instantiates a new InterfaceRedundancyGroupAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceRedundancyGroupAssociationWithDefaults

`func NewInterfaceRedundancyGroupAssociationWithDefaults() *InterfaceRedundancyGroupAssociation`

NewInterfaceRedundancyGroupAssociationWithDefaults instantiates a new InterfaceRedundancyGroupAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceRedundancyGroupAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceRedundancyGroupAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceRedundancyGroupAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *InterfaceRedundancyGroupAssociation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InterfaceRedundancyGroupAssociation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InterfaceRedundancyGroupAssociation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *InterfaceRedundancyGroupAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceRedundancyGroupAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceRedundancyGroupAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *InterfaceRedundancyGroupAssociation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InterfaceRedundancyGroupAssociation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InterfaceRedundancyGroupAssociation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *InterfaceRedundancyGroupAssociation) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *InterfaceRedundancyGroupAssociation) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *InterfaceRedundancyGroupAssociation) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetPriority

`func (o *InterfaceRedundancyGroupAssociation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InterfaceRedundancyGroupAssociation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InterfaceRedundancyGroupAssociation) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetInterfaceRedundancyGroup

`func (o *InterfaceRedundancyGroupAssociation) GetInterfaceRedundancyGroup() BulkWritableCableRequestStatus`

GetInterfaceRedundancyGroup returns the InterfaceRedundancyGroup field if non-nil, zero value otherwise.

### GetInterfaceRedundancyGroupOk

`func (o *InterfaceRedundancyGroupAssociation) GetInterfaceRedundancyGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceRedundancyGroupOk returns a tuple with the InterfaceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyGroup

`func (o *InterfaceRedundancyGroupAssociation) SetInterfaceRedundancyGroup(v BulkWritableCableRequestStatus)`

SetInterfaceRedundancyGroup sets InterfaceRedundancyGroup field to given value.


### GetInterface

`func (o *InterfaceRedundancyGroupAssociation) GetInterface() BulkWritableCableRequestStatus`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InterfaceRedundancyGroupAssociation) GetInterfaceOk() (*BulkWritableCableRequestStatus, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InterfaceRedundancyGroupAssociation) SetInterface(v BulkWritableCableRequestStatus)`

SetInterface sets Interface field to given value.


### GetCreated

`func (o *InterfaceRedundancyGroupAssociation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InterfaceRedundancyGroupAssociation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InterfaceRedundancyGroupAssociation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *InterfaceRedundancyGroupAssociation) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *InterfaceRedundancyGroupAssociation) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *InterfaceRedundancyGroupAssociation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InterfaceRedundancyGroupAssociation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InterfaceRedundancyGroupAssociation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *InterfaceRedundancyGroupAssociation) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *InterfaceRedundancyGroupAssociation) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


