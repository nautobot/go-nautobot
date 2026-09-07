# BulkWritableRelationshipAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceType** | **string** |  | 
**DestinationType** | **string** |  | 
**SourceId** | **string** |  | 
**DestinationId** | **string** |  | 
**Relationship** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableRelationshipAssociationRequest

`func NewBulkWritableRelationshipAssociationRequest(id string, sourceType string, destinationType string, sourceId string, destinationId string, relationship BulkWritableCableRequestStatus, ) *BulkWritableRelationshipAssociationRequest`

NewBulkWritableRelationshipAssociationRequest instantiates a new BulkWritableRelationshipAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableRelationshipAssociationRequestWithDefaults

`func NewBulkWritableRelationshipAssociationRequestWithDefaults() *BulkWritableRelationshipAssociationRequest`

NewBulkWritableRelationshipAssociationRequestWithDefaults instantiates a new BulkWritableRelationshipAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableRelationshipAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableRelationshipAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableRelationshipAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSourceType

`func (o *BulkWritableRelationshipAssociationRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BulkWritableRelationshipAssociationRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BulkWritableRelationshipAssociationRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetDestinationType

`func (o *BulkWritableRelationshipAssociationRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *BulkWritableRelationshipAssociationRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *BulkWritableRelationshipAssociationRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetSourceId

`func (o *BulkWritableRelationshipAssociationRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BulkWritableRelationshipAssociationRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BulkWritableRelationshipAssociationRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *BulkWritableRelationshipAssociationRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *BulkWritableRelationshipAssociationRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *BulkWritableRelationshipAssociationRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetRelationship

`func (o *BulkWritableRelationshipAssociationRequest) GetRelationship() BulkWritableCableRequestStatus`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *BulkWritableRelationshipAssociationRequest) GetRelationshipOk() (*BulkWritableCableRequestStatus, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *BulkWritableRelationshipAssociationRequest) SetRelationship(v BulkWritableCableRequestStatus)`

SetRelationship sets Relationship field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


