# PatchedSupportedDataRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Standard** | Pointer to [**StandardEnum**](StandardEnum.md) |  | [optional] 
**Rate** | Pointer to **int32** |  | [optional] 
**McsIndex** | Pointer to **NullableInt32** | The Modulation and Coding Scheme (MCS) index is a value used in wireless communications to define the modulation type, coding rate, and number of spatial streams used in a transmission. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedSupportedDataRateRequest

`func NewPatchedSupportedDataRateRequest() *PatchedSupportedDataRateRequest`

NewPatchedSupportedDataRateRequest instantiates a new PatchedSupportedDataRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSupportedDataRateRequestWithDefaults

`func NewPatchedSupportedDataRateRequestWithDefaults() *PatchedSupportedDataRateRequest`

NewPatchedSupportedDataRateRequestWithDefaults instantiates a new PatchedSupportedDataRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSupportedDataRateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSupportedDataRateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSupportedDataRateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSupportedDataRateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStandard

`func (o *PatchedSupportedDataRateRequest) GetStandard() StandardEnum`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *PatchedSupportedDataRateRequest) GetStandardOk() (*StandardEnum, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *PatchedSupportedDataRateRequest) SetStandard(v StandardEnum)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *PatchedSupportedDataRateRequest) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetRate

`func (o *PatchedSupportedDataRateRequest) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PatchedSupportedDataRateRequest) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PatchedSupportedDataRateRequest) SetRate(v int32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *PatchedSupportedDataRateRequest) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetMcsIndex

`func (o *PatchedSupportedDataRateRequest) GetMcsIndex() int32`

GetMcsIndex returns the McsIndex field if non-nil, zero value otherwise.

### GetMcsIndexOk

`func (o *PatchedSupportedDataRateRequest) GetMcsIndexOk() (*int32, bool)`

GetMcsIndexOk returns a tuple with the McsIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsIndex

`func (o *PatchedSupportedDataRateRequest) SetMcsIndex(v int32)`

SetMcsIndex sets McsIndex field to given value.

### HasMcsIndex

`func (o *PatchedSupportedDataRateRequest) HasMcsIndex() bool`

HasMcsIndex returns a boolean if a field has been set.

### SetMcsIndexNil

`func (o *PatchedSupportedDataRateRequest) SetMcsIndexNil(b bool)`

 SetMcsIndexNil sets the value for McsIndex to be an explicit nil

### UnsetMcsIndex
`func (o *PatchedSupportedDataRateRequest) UnsetMcsIndex()`

UnsetMcsIndex ensures that no value is present for McsIndex, not even an explicit nil
### GetCustomFields

`func (o *PatchedSupportedDataRateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedSupportedDataRateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedSupportedDataRateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedSupportedDataRateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedSupportedDataRateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedSupportedDataRateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedSupportedDataRateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedSupportedDataRateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedSupportedDataRateRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedSupportedDataRateRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedSupportedDataRateRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedSupportedDataRateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


