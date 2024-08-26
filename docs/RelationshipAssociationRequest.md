# RelationshipAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** |  | 
**DestinationType** | **string** |  | 
**SourceId** | **string** |  | 
**DestinationId** | **string** |  | 
**Relationship** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewRelationshipAssociationRequest

`func NewRelationshipAssociationRequest(sourceType string, destinationType string, sourceId string, destinationId string, relationship BulkWritableCableRequestStatus, ) *RelationshipAssociationRequest`

NewRelationshipAssociationRequest instantiates a new RelationshipAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipAssociationRequestWithDefaults

`func NewRelationshipAssociationRequestWithDefaults() *RelationshipAssociationRequest`

NewRelationshipAssociationRequestWithDefaults instantiates a new RelationshipAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *RelationshipAssociationRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RelationshipAssociationRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RelationshipAssociationRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetDestinationType

`func (o *RelationshipAssociationRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RelationshipAssociationRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RelationshipAssociationRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetSourceId

`func (o *RelationshipAssociationRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RelationshipAssociationRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RelationshipAssociationRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *RelationshipAssociationRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *RelationshipAssociationRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *RelationshipAssociationRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetRelationship

`func (o *RelationshipAssociationRequest) GetRelationship() BulkWritableCableRequestStatus`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *RelationshipAssociationRequest) GetRelationshipOk() (*BulkWritableCableRequestStatus, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *RelationshipAssociationRequest) SetRelationship(v BulkWritableCableRequestStatus)`

SetRelationship sets Relationship field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


