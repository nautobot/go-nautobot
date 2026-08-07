# SupportedDataRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Standard** | [**StandardEnum**](StandardEnum.md) |  | 
**Rate** | **int32** |  | 
**McsIndex** | Pointer to **NullableInt32** | The Modulation and Coding Scheme (MCS) index is a value used in wireless communications to define the modulation type, coding rate, and number of spatial streams used in a transmission. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewSupportedDataRateRequest

`func NewSupportedDataRateRequest(standard StandardEnum, rate int32, ) *SupportedDataRateRequest`

NewSupportedDataRateRequest instantiates a new SupportedDataRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedDataRateRequestWithDefaults

`func NewSupportedDataRateRequestWithDefaults() *SupportedDataRateRequest`

NewSupportedDataRateRequestWithDefaults instantiates a new SupportedDataRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportedDataRateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedDataRateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedDataRateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportedDataRateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStandard

`func (o *SupportedDataRateRequest) GetStandard() StandardEnum`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *SupportedDataRateRequest) GetStandardOk() (*StandardEnum, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *SupportedDataRateRequest) SetStandard(v StandardEnum)`

SetStandard sets Standard field to given value.


### GetRate

`func (o *SupportedDataRateRequest) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SupportedDataRateRequest) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SupportedDataRateRequest) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetMcsIndex

`func (o *SupportedDataRateRequest) GetMcsIndex() int32`

GetMcsIndex returns the McsIndex field if non-nil, zero value otherwise.

### GetMcsIndexOk

`func (o *SupportedDataRateRequest) GetMcsIndexOk() (*int32, bool)`

GetMcsIndexOk returns a tuple with the McsIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsIndex

`func (o *SupportedDataRateRequest) SetMcsIndex(v int32)`

SetMcsIndex sets McsIndex field to given value.

### HasMcsIndex

`func (o *SupportedDataRateRequest) HasMcsIndex() bool`

HasMcsIndex returns a boolean if a field has been set.

### SetMcsIndexNil

`func (o *SupportedDataRateRequest) SetMcsIndexNil(b bool)`

 SetMcsIndexNil sets the value for McsIndex to be an explicit nil

### UnsetMcsIndex
`func (o *SupportedDataRateRequest) UnsetMcsIndex()`

UnsetMcsIndex ensures that no value is present for McsIndex, not even an explicit nil
### GetCustomFields

`func (o *SupportedDataRateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SupportedDataRateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SupportedDataRateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SupportedDataRateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *SupportedDataRateRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SupportedDataRateRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SupportedDataRateRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SupportedDataRateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *SupportedDataRateRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SupportedDataRateRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SupportedDataRateRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SupportedDataRateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


