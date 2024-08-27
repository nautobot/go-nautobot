# MetadataChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**MetadataType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewMetadataChoiceRequest

`func NewMetadataChoiceRequest(value string, metadataType BulkWritableCableRequestStatus, ) *MetadataChoiceRequest`

NewMetadataChoiceRequest instantiates a new MetadataChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataChoiceRequestWithDefaults

`func NewMetadataChoiceRequestWithDefaults() *MetadataChoiceRequest`

NewMetadataChoiceRequestWithDefaults instantiates a new MetadataChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *MetadataChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetadataChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetadataChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *MetadataChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MetadataChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MetadataChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *MetadataChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMetadataType

`func (o *MetadataChoiceRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *MetadataChoiceRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *MetadataChoiceRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


