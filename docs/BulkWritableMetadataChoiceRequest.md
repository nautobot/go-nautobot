# BulkWritableMetadataChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**MetadataType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableMetadataChoiceRequest

`func NewBulkWritableMetadataChoiceRequest(id string, value string, metadataType BulkWritableCableRequestStatus, ) *BulkWritableMetadataChoiceRequest`

NewBulkWritableMetadataChoiceRequest instantiates a new BulkWritableMetadataChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableMetadataChoiceRequestWithDefaults

`func NewBulkWritableMetadataChoiceRequestWithDefaults() *BulkWritableMetadataChoiceRequest`

NewBulkWritableMetadataChoiceRequestWithDefaults instantiates a new BulkWritableMetadataChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableMetadataChoiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableMetadataChoiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableMetadataChoiceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *BulkWritableMetadataChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkWritableMetadataChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkWritableMetadataChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *BulkWritableMetadataChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkWritableMetadataChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkWritableMetadataChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkWritableMetadataChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMetadataType

`func (o *BulkWritableMetadataChoiceRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *BulkWritableMetadataChoiceRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *BulkWritableMetadataChoiceRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


