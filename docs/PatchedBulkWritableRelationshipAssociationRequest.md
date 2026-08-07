# PatchedBulkWritableRelationshipAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**DestinationId** | Pointer to **string** |  | [optional] 
**Relationship** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableRelationshipAssociationRequest

`func NewPatchedBulkWritableRelationshipAssociationRequest(id string, ) *PatchedBulkWritableRelationshipAssociationRequest`

NewPatchedBulkWritableRelationshipAssociationRequest instantiates a new PatchedBulkWritableRelationshipAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableRelationshipAssociationRequestWithDefaults

`func NewPatchedBulkWritableRelationshipAssociationRequestWithDefaults() *PatchedBulkWritableRelationshipAssociationRequest`

NewPatchedBulkWritableRelationshipAssociationRequestWithDefaults instantiates a new PatchedBulkWritableRelationshipAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSourceType

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PatchedBulkWritableRelationshipAssociationRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PatchedBulkWritableRelationshipAssociationRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PatchedBulkWritableRelationshipAssociationRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *PatchedBulkWritableRelationshipAssociationRequest) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetSourceId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *PatchedBulkWritableRelationshipAssociationRequest) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetRelationship

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetRelationship() BulkWritableCableRequestStatus`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *PatchedBulkWritableRelationshipAssociationRequest) GetRelationshipOk() (*BulkWritableCableRequestStatus, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *PatchedBulkWritableRelationshipAssociationRequest) SetRelationship(v BulkWritableCableRequestStatus)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *PatchedBulkWritableRelationshipAssociationRequest) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


