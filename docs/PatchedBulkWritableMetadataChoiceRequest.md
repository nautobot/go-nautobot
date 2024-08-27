# PatchedBulkWritableMetadataChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**MetadataType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableMetadataChoiceRequest

`func NewPatchedBulkWritableMetadataChoiceRequest(id string, ) *PatchedBulkWritableMetadataChoiceRequest`

NewPatchedBulkWritableMetadataChoiceRequest instantiates a new PatchedBulkWritableMetadataChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableMetadataChoiceRequestWithDefaults

`func NewPatchedBulkWritableMetadataChoiceRequestWithDefaults() *PatchedBulkWritableMetadataChoiceRequest`

NewPatchedBulkWritableMetadataChoiceRequestWithDefaults instantiates a new PatchedBulkWritableMetadataChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableMetadataChoiceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedBulkWritableMetadataChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedBulkWritableMetadataChoiceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableMetadataChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableMetadataChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMetadataType

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *PatchedBulkWritableMetadataChoiceRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *PatchedBulkWritableMetadataChoiceRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *PatchedBulkWritableMetadataChoiceRequest) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


