# PatchedMetadataChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**MetadataType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedMetadataChoiceRequest

`func NewPatchedMetadataChoiceRequest() *PatchedMetadataChoiceRequest`

NewPatchedMetadataChoiceRequest instantiates a new PatchedMetadataChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedMetadataChoiceRequestWithDefaults

`func NewPatchedMetadataChoiceRequestWithDefaults() *PatchedMetadataChoiceRequest`

NewPatchedMetadataChoiceRequestWithDefaults instantiates a new PatchedMetadataChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PatchedMetadataChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedMetadataChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedMetadataChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedMetadataChoiceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedMetadataChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedMetadataChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedMetadataChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedMetadataChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMetadataType

`func (o *PatchedMetadataChoiceRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *PatchedMetadataChoiceRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *PatchedMetadataChoiceRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *PatchedMetadataChoiceRequest) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


