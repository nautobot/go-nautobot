# BulkWritableGraphQLQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Query** | **string** |  | 

## Methods

### NewBulkWritableGraphQLQueryRequest

`func NewBulkWritableGraphQLQueryRequest(id string, name string, query string, ) *BulkWritableGraphQLQueryRequest`

NewBulkWritableGraphQLQueryRequest instantiates a new BulkWritableGraphQLQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableGraphQLQueryRequestWithDefaults

`func NewBulkWritableGraphQLQueryRequestWithDefaults() *BulkWritableGraphQLQueryRequest`

NewBulkWritableGraphQLQueryRequestWithDefaults instantiates a new BulkWritableGraphQLQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableGraphQLQueryRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableGraphQLQueryRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableGraphQLQueryRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerContentType

`func (o *BulkWritableGraphQLQueryRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *BulkWritableGraphQLQueryRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *BulkWritableGraphQLQueryRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *BulkWritableGraphQLQueryRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *BulkWritableGraphQLQueryRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *BulkWritableGraphQLQueryRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetOwnerObjectId

`func (o *BulkWritableGraphQLQueryRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *BulkWritableGraphQLQueryRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *BulkWritableGraphQLQueryRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *BulkWritableGraphQLQueryRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *BulkWritableGraphQLQueryRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *BulkWritableGraphQLQueryRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetName

`func (o *BulkWritableGraphQLQueryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableGraphQLQueryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableGraphQLQueryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *BulkWritableGraphQLQueryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkWritableGraphQLQueryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkWritableGraphQLQueryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


