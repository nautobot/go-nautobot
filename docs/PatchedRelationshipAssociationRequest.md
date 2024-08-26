# PatchedRelationshipAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**DestinationId** | Pointer to **string** |  | [optional] 
**Relationship** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedRelationshipAssociationRequest

`func NewPatchedRelationshipAssociationRequest() *PatchedRelationshipAssociationRequest`

NewPatchedRelationshipAssociationRequest instantiates a new PatchedRelationshipAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRelationshipAssociationRequestWithDefaults

`func NewPatchedRelationshipAssociationRequestWithDefaults() *PatchedRelationshipAssociationRequest`

NewPatchedRelationshipAssociationRequestWithDefaults instantiates a new PatchedRelationshipAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *PatchedRelationshipAssociationRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PatchedRelationshipAssociationRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PatchedRelationshipAssociationRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PatchedRelationshipAssociationRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *PatchedRelationshipAssociationRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PatchedRelationshipAssociationRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PatchedRelationshipAssociationRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *PatchedRelationshipAssociationRequest) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetSourceId

`func (o *PatchedRelationshipAssociationRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PatchedRelationshipAssociationRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PatchedRelationshipAssociationRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *PatchedRelationshipAssociationRequest) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationId

`func (o *PatchedRelationshipAssociationRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *PatchedRelationshipAssociationRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *PatchedRelationshipAssociationRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *PatchedRelationshipAssociationRequest) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetRelationship

`func (o *PatchedRelationshipAssociationRequest) GetRelationship() BulkWritableCableRequestStatus`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *PatchedRelationshipAssociationRequest) GetRelationshipOk() (*BulkWritableCableRequestStatus, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *PatchedRelationshipAssociationRequest) SetRelationship(v BulkWritableCableRequestStatus)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *PatchedRelationshipAssociationRequest) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


