# BulkWritableRadioProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableRadioProfileRequest

`func NewBulkWritableRadioProfileRequest(id string, name string, regulatoryDomain RegulatoryDomainEnum, ) *BulkWritableRadioProfileRequest`

NewBulkWritableRadioProfileRequest instantiates a new BulkWritableRadioProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableRadioProfileRequestWithDefaults

`func NewBulkWritableRadioProfileRequestWithDefaults() *BulkWritableRadioProfileRequest`

NewBulkWritableRadioProfileRequestWithDefaults instantiates a new BulkWritableRadioProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableRadioProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableRadioProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableRadioProfileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetChannelWidth

`func (o *BulkWritableRadioProfileRequest) GetChannelWidth() []ChannelWidthEnum`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *BulkWritableRadioProfileRequest) GetChannelWidthOk() (*[]ChannelWidthEnum, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *BulkWritableRadioProfileRequest) SetChannelWidth(v []ChannelWidthEnum)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *BulkWritableRadioProfileRequest) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetAllowedChannelList

`func (o *BulkWritableRadioProfileRequest) GetAllowedChannelList() []int32`

GetAllowedChannelList returns the AllowedChannelList field if non-nil, zero value otherwise.

### GetAllowedChannelListOk

`func (o *BulkWritableRadioProfileRequest) GetAllowedChannelListOk() (*[]int32, bool)`

GetAllowedChannelListOk returns a tuple with the AllowedChannelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChannelList

`func (o *BulkWritableRadioProfileRequest) SetAllowedChannelList(v []int32)`

SetAllowedChannelList sets AllowedChannelList field to given value.

### HasAllowedChannelList

`func (o *BulkWritableRadioProfileRequest) HasAllowedChannelList() bool`

HasAllowedChannelList returns a boolean if a field has been set.

### GetName

`func (o *BulkWritableRadioProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableRadioProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableRadioProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFrequency

`func (o *BulkWritableRadioProfileRequest) GetFrequency() BulkWritableRadioProfileRequestFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BulkWritableRadioProfileRequest) GetFrequencyOk() (*BulkWritableRadioProfileRequestFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BulkWritableRadioProfileRequest) SetFrequency(v BulkWritableRadioProfileRequestFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *BulkWritableRadioProfileRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetTxPowerMin

`func (o *BulkWritableRadioProfileRequest) GetTxPowerMin() int32`

GetTxPowerMin returns the TxPowerMin field if non-nil, zero value otherwise.

### GetTxPowerMinOk

`func (o *BulkWritableRadioProfileRequest) GetTxPowerMinOk() (*int32, bool)`

GetTxPowerMinOk returns a tuple with the TxPowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerMin

`func (o *BulkWritableRadioProfileRequest) SetTxPowerMin(v int32)`

SetTxPowerMin sets TxPowerMin field to given value.

### HasTxPowerMin

`func (o *BulkWritableRadioProfileRequest) HasTxPowerMin() bool`

HasTxPowerMin returns a boolean if a field has been set.

### SetTxPowerMinNil

`func (o *BulkWritableRadioProfileRequest) SetTxPowerMinNil(b bool)`

 SetTxPowerMinNil sets the value for TxPowerMin to be an explicit nil

### UnsetTxPowerMin
`func (o *BulkWritableRadioProfileRequest) UnsetTxPowerMin()`

UnsetTxPowerMin ensures that no value is present for TxPowerMin, not even an explicit nil
### GetTxPowerMax

`func (o *BulkWritableRadioProfileRequest) GetTxPowerMax() int32`

GetTxPowerMax returns the TxPowerMax field if non-nil, zero value otherwise.

### GetTxPowerMaxOk

`func (o *BulkWritableRadioProfileRequest) GetTxPowerMaxOk() (*int32, bool)`

GetTxPowerMaxOk returns a tuple with the TxPowerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerMax

`func (o *BulkWritableRadioProfileRequest) SetTxPowerMax(v int32)`

SetTxPowerMax sets TxPowerMax field to given value.

### HasTxPowerMax

`func (o *BulkWritableRadioProfileRequest) HasTxPowerMax() bool`

HasTxPowerMax returns a boolean if a field has been set.

### SetTxPowerMaxNil

`func (o *BulkWritableRadioProfileRequest) SetTxPowerMaxNil(b bool)`

 SetTxPowerMaxNil sets the value for TxPowerMax to be an explicit nil

### UnsetTxPowerMax
`func (o *BulkWritableRadioProfileRequest) UnsetTxPowerMax()`

UnsetTxPowerMax ensures that no value is present for TxPowerMax, not even an explicit nil
### GetRegulatoryDomain

`func (o *BulkWritableRadioProfileRequest) GetRegulatoryDomain() RegulatoryDomainEnum`

GetRegulatoryDomain returns the RegulatoryDomain field if non-nil, zero value otherwise.

### GetRegulatoryDomainOk

`func (o *BulkWritableRadioProfileRequest) GetRegulatoryDomainOk() (*RegulatoryDomainEnum, bool)`

GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDomain

`func (o *BulkWritableRadioProfileRequest) SetRegulatoryDomain(v RegulatoryDomainEnum)`

SetRegulatoryDomain sets RegulatoryDomain field to given value.


### GetRxPowerMin

`func (o *BulkWritableRadioProfileRequest) GetRxPowerMin() int32`

GetRxPowerMin returns the RxPowerMin field if non-nil, zero value otherwise.

### GetRxPowerMinOk

`func (o *BulkWritableRadioProfileRequest) GetRxPowerMinOk() (*int32, bool)`

GetRxPowerMinOk returns a tuple with the RxPowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPowerMin

`func (o *BulkWritableRadioProfileRequest) SetRxPowerMin(v int32)`

SetRxPowerMin sets RxPowerMin field to given value.

### HasRxPowerMin

`func (o *BulkWritableRadioProfileRequest) HasRxPowerMin() bool`

HasRxPowerMin returns a boolean if a field has been set.

### SetRxPowerMinNil

`func (o *BulkWritableRadioProfileRequest) SetRxPowerMinNil(b bool)`

 SetRxPowerMinNil sets the value for RxPowerMin to be an explicit nil

### UnsetRxPowerMin
`func (o *BulkWritableRadioProfileRequest) UnsetRxPowerMin()`

UnsetRxPowerMin ensures that no value is present for RxPowerMin, not even an explicit nil
### GetSupportedDataRates

`func (o *BulkWritableRadioProfileRequest) GetSupportedDataRates() []BulkWritableCableRequestStatus`

GetSupportedDataRates returns the SupportedDataRates field if non-nil, zero value otherwise.

### GetSupportedDataRatesOk

`func (o *BulkWritableRadioProfileRequest) GetSupportedDataRatesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetSupportedDataRatesOk returns a tuple with the SupportedDataRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDataRates

`func (o *BulkWritableRadioProfileRequest) SetSupportedDataRates(v []BulkWritableCableRequestStatus)`

SetSupportedDataRates sets SupportedDataRates field to given value.

### HasSupportedDataRates

`func (o *BulkWritableRadioProfileRequest) HasSupportedDataRates() bool`

HasSupportedDataRates returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableRadioProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableRadioProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableRadioProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableRadioProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableRadioProfileRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableRadioProfileRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableRadioProfileRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableRadioProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableRadioProfileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableRadioProfileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableRadioProfileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableRadioProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


