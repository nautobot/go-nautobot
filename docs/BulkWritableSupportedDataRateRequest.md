# BulkWritableSupportedDataRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Standard** | [**StandardEnum**](StandardEnum.md) |  | 
**Rate** | **int32** |  | 
**McsIndex** | Pointer to **NullableInt32** | The Modulation and Coding Scheme (MCS) index is a value used in wireless communications to define the modulation type, coding rate, and number of spatial streams used in a transmission. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableSupportedDataRateRequest

`func NewBulkWritableSupportedDataRateRequest(id string, standard StandardEnum, rate int32, ) *BulkWritableSupportedDataRateRequest`

NewBulkWritableSupportedDataRateRequest instantiates a new BulkWritableSupportedDataRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableSupportedDataRateRequestWithDefaults

`func NewBulkWritableSupportedDataRateRequestWithDefaults() *BulkWritableSupportedDataRateRequest`

NewBulkWritableSupportedDataRateRequestWithDefaults instantiates a new BulkWritableSupportedDataRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableSupportedDataRateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableSupportedDataRateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableSupportedDataRateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStandard

`func (o *BulkWritableSupportedDataRateRequest) GetStandard() StandardEnum`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *BulkWritableSupportedDataRateRequest) GetStandardOk() (*StandardEnum, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *BulkWritableSupportedDataRateRequest) SetStandard(v StandardEnum)`

SetStandard sets Standard field to given value.


### GetRate

`func (o *BulkWritableSupportedDataRateRequest) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *BulkWritableSupportedDataRateRequest) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *BulkWritableSupportedDataRateRequest) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetMcsIndex

`func (o *BulkWritableSupportedDataRateRequest) GetMcsIndex() int32`

GetMcsIndex returns the McsIndex field if non-nil, zero value otherwise.

### GetMcsIndexOk

`func (o *BulkWritableSupportedDataRateRequest) GetMcsIndexOk() (*int32, bool)`

GetMcsIndexOk returns a tuple with the McsIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsIndex

`func (o *BulkWritableSupportedDataRateRequest) SetMcsIndex(v int32)`

SetMcsIndex sets McsIndex field to given value.

### HasMcsIndex

`func (o *BulkWritableSupportedDataRateRequest) HasMcsIndex() bool`

HasMcsIndex returns a boolean if a field has been set.

### SetMcsIndexNil

`func (o *BulkWritableSupportedDataRateRequest) SetMcsIndexNil(b bool)`

 SetMcsIndexNil sets the value for McsIndex to be an explicit nil

### UnsetMcsIndex
`func (o *BulkWritableSupportedDataRateRequest) UnsetMcsIndex()`

UnsetMcsIndex ensures that no value is present for McsIndex, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableSupportedDataRateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableSupportedDataRateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableSupportedDataRateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableSupportedDataRateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableSupportedDataRateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableSupportedDataRateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableSupportedDataRateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableSupportedDataRateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableSupportedDataRateRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableSupportedDataRateRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableSupportedDataRateRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableSupportedDataRateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


