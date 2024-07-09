# VirtualChassisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Domain** | Pointer to **string** |  | [optional] 
**Master** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewVirtualChassisRequest

`func NewVirtualChassisRequest(name string, ) *VirtualChassisRequest`

NewVirtualChassisRequest instantiates a new VirtualChassisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisRequestWithDefaults

`func NewVirtualChassisRequestWithDefaults() *VirtualChassisRequest`

NewVirtualChassisRequestWithDefaults instantiates a new VirtualChassisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualChassisRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualChassisRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualChassisRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *VirtualChassisRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VirtualChassisRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VirtualChassisRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *VirtualChassisRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMaster

`func (o *VirtualChassisRequest) GetMaster() BulkWritableCircuitRequestTenant`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *VirtualChassisRequest) GetMasterOk() (*BulkWritableCircuitRequestTenant, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *VirtualChassisRequest) SetMaster(v BulkWritableCircuitRequestTenant)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *VirtualChassisRequest) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *VirtualChassisRequest) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *VirtualChassisRequest) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetTags

`func (o *VirtualChassisRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualChassisRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualChassisRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualChassisRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualChassisRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualChassisRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualChassisRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualChassisRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VirtualChassisRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VirtualChassisRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VirtualChassisRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VirtualChassisRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


