# RadioProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ChannelWidth** | Pointer to [**[]ChannelWidthEnum**](ChannelWidthEnum.md) |  | [optional] 
**AllowedChannelList** | Pointer to **[]int32** |  | [optional] 
**Name** | **string** |  | 
**Frequency** | Pointer to [**BulkWritableRadioProfileRequestFrequency**](BulkWritableRadioProfileRequestFrequency.md) |  | [optional] 
**TxPowerMin** | Pointer to **NullableInt32** |  | [optional] 
**TxPowerMax** | Pointer to **NullableInt32** |  | [optional] 
**RegulatoryDomain** | [**RegulatoryDomainEnum**](RegulatoryDomainEnum.md) |  | 
**RxPowerMin** | Pointer to **NullableInt32** |  | [optional] 
**SupportedDataRates** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewRadioProfileRequest

`func NewRadioProfileRequest(name string, regulatoryDomain RegulatoryDomainEnum, ) *RadioProfileRequest`

NewRadioProfileRequest instantiates a new RadioProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadioProfileRequestWithDefaults

`func NewRadioProfileRequestWithDefaults() *RadioProfileRequest`

NewRadioProfileRequestWithDefaults instantiates a new RadioProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RadioProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RadioProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RadioProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RadioProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelWidth

`func (o *RadioProfileRequest) GetChannelWidth() []ChannelWidthEnum`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *RadioProfileRequest) GetChannelWidthOk() (*[]ChannelWidthEnum, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *RadioProfileRequest) SetChannelWidth(v []ChannelWidthEnum)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *RadioProfileRequest) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetAllowedChannelList

`func (o *RadioProfileRequest) GetAllowedChannelList() []int32`

GetAllowedChannelList returns the AllowedChannelList field if non-nil, zero value otherwise.

### GetAllowedChannelListOk

`func (o *RadioProfileRequest) GetAllowedChannelListOk() (*[]int32, bool)`

GetAllowedChannelListOk returns a tuple with the AllowedChannelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChannelList

`func (o *RadioProfileRequest) SetAllowedChannelList(v []int32)`

SetAllowedChannelList sets AllowedChannelList field to given value.

### HasAllowedChannelList

`func (o *RadioProfileRequest) HasAllowedChannelList() bool`

HasAllowedChannelList returns a boolean if a field has been set.

### GetName

`func (o *RadioProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadioProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadioProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFrequency

`func (o *RadioProfileRequest) GetFrequency() BulkWritableRadioProfileRequestFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RadioProfileRequest) GetFrequencyOk() (*BulkWritableRadioProfileRequestFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RadioProfileRequest) SetFrequency(v BulkWritableRadioProfileRequestFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RadioProfileRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetTxPowerMin

`func (o *RadioProfileRequest) GetTxPowerMin() int32`

GetTxPowerMin returns the TxPowerMin field if non-nil, zero value otherwise.

### GetTxPowerMinOk

`func (o *RadioProfileRequest) GetTxPowerMinOk() (*int32, bool)`

GetTxPowerMinOk returns a tuple with the TxPowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerMin

`func (o *RadioProfileRequest) SetTxPowerMin(v int32)`

SetTxPowerMin sets TxPowerMin field to given value.

### HasTxPowerMin

`func (o *RadioProfileRequest) HasTxPowerMin() bool`

HasTxPowerMin returns a boolean if a field has been set.

### SetTxPowerMinNil

`func (o *RadioProfileRequest) SetTxPowerMinNil(b bool)`

 SetTxPowerMinNil sets the value for TxPowerMin to be an explicit nil

### UnsetTxPowerMin
`func (o *RadioProfileRequest) UnsetTxPowerMin()`

UnsetTxPowerMin ensures that no value is present for TxPowerMin, not even an explicit nil
### GetTxPowerMax

`func (o *RadioProfileRequest) GetTxPowerMax() int32`

GetTxPowerMax returns the TxPowerMax field if non-nil, zero value otherwise.

### GetTxPowerMaxOk

`func (o *RadioProfileRequest) GetTxPowerMaxOk() (*int32, bool)`

GetTxPowerMaxOk returns a tuple with the TxPowerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerMax

`func (o *RadioProfileRequest) SetTxPowerMax(v int32)`

SetTxPowerMax sets TxPowerMax field to given value.

### HasTxPowerMax

`func (o *RadioProfileRequest) HasTxPowerMax() bool`

HasTxPowerMax returns a boolean if a field has been set.

### SetTxPowerMaxNil

`func (o *RadioProfileRequest) SetTxPowerMaxNil(b bool)`

 SetTxPowerMaxNil sets the value for TxPowerMax to be an explicit nil

### UnsetTxPowerMax
`func (o *RadioProfileRequest) UnsetTxPowerMax()`

UnsetTxPowerMax ensures that no value is present for TxPowerMax, not even an explicit nil
### GetRegulatoryDomain

`func (o *RadioProfileRequest) GetRegulatoryDomain() RegulatoryDomainEnum`

GetRegulatoryDomain returns the RegulatoryDomain field if non-nil, zero value otherwise.

### GetRegulatoryDomainOk

`func (o *RadioProfileRequest) GetRegulatoryDomainOk() (*RegulatoryDomainEnum, bool)`

GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDomain

`func (o *RadioProfileRequest) SetRegulatoryDomain(v RegulatoryDomainEnum)`

SetRegulatoryDomain sets RegulatoryDomain field to given value.


### GetRxPowerMin

`func (o *RadioProfileRequest) GetRxPowerMin() int32`

GetRxPowerMin returns the RxPowerMin field if non-nil, zero value otherwise.

### GetRxPowerMinOk

`func (o *RadioProfileRequest) GetRxPowerMinOk() (*int32, bool)`

GetRxPowerMinOk returns a tuple with the RxPowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPowerMin

`func (o *RadioProfileRequest) SetRxPowerMin(v int32)`

SetRxPowerMin sets RxPowerMin field to given value.

### HasRxPowerMin

`func (o *RadioProfileRequest) HasRxPowerMin() bool`

HasRxPowerMin returns a boolean if a field has been set.

### SetRxPowerMinNil

`func (o *RadioProfileRequest) SetRxPowerMinNil(b bool)`

 SetRxPowerMinNil sets the value for RxPowerMin to be an explicit nil

### UnsetRxPowerMin
`func (o *RadioProfileRequest) UnsetRxPowerMin()`

UnsetRxPowerMin ensures that no value is present for RxPowerMin, not even an explicit nil
### GetSupportedDataRates

`func (o *RadioProfileRequest) GetSupportedDataRates() []BulkWritableCableRequestStatus`

GetSupportedDataRates returns the SupportedDataRates field if non-nil, zero value otherwise.

### GetSupportedDataRatesOk

`func (o *RadioProfileRequest) GetSupportedDataRatesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetSupportedDataRatesOk returns a tuple with the SupportedDataRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDataRates

`func (o *RadioProfileRequest) SetSupportedDataRates(v []BulkWritableCableRequestStatus)`

SetSupportedDataRates sets SupportedDataRates field to given value.

### HasSupportedDataRates

`func (o *RadioProfileRequest) HasSupportedDataRates() bool`

HasSupportedDataRates returns a boolean if a field has been set.

### GetCustomFields

`func (o *RadioProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RadioProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RadioProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RadioProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *RadioProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RadioProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RadioProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RadioProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *RadioProfileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RadioProfileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RadioProfileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RadioProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


