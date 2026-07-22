# RadioProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ChannelWidth** | Pointer to [**[]RadioProfileChannelWidthInner**](RadioProfileChannelWidthInner.md) |  | [optional] 
**AllowedChannelList** | Pointer to **[]int32** |  | [optional] 
**Name** | **string** |  | 
**Frequency** | Pointer to [**BulkWritableRadioProfileRequestFrequency**](BulkWritableRadioProfileRequestFrequency.md) |  | [optional] 
**TxPowerMin** | Pointer to **NullableInt32** |  | [optional] 
**TxPowerMax** | Pointer to **NullableInt32** |  | [optional] 
**RegulatoryDomain** | [**RegulatoryDomainEnum**](RegulatoryDomainEnum.md) |  | 
**RxPowerMin** | Pointer to **NullableInt32** |  | [optional] 
**SupportedDataRates** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewRadioProfile

`func NewRadioProfile(objectType string, display string, url string, naturalSlug string, name string, regulatoryDomain RegulatoryDomainEnum, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *RadioProfile`

NewRadioProfile instantiates a new RadioProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadioProfileWithDefaults

`func NewRadioProfileWithDefaults() *RadioProfile`

NewRadioProfileWithDefaults instantiates a new RadioProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RadioProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RadioProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RadioProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RadioProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *RadioProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RadioProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RadioProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *RadioProfile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RadioProfile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RadioProfile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *RadioProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RadioProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RadioProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *RadioProfile) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *RadioProfile) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *RadioProfile) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetChannelWidth

`func (o *RadioProfile) GetChannelWidth() []RadioProfileChannelWidthInner`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *RadioProfile) GetChannelWidthOk() (*[]RadioProfileChannelWidthInner, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *RadioProfile) SetChannelWidth(v []RadioProfileChannelWidthInner)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *RadioProfile) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetAllowedChannelList

`func (o *RadioProfile) GetAllowedChannelList() []int32`

GetAllowedChannelList returns the AllowedChannelList field if non-nil, zero value otherwise.

### GetAllowedChannelListOk

`func (o *RadioProfile) GetAllowedChannelListOk() (*[]int32, bool)`

GetAllowedChannelListOk returns a tuple with the AllowedChannelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChannelList

`func (o *RadioProfile) SetAllowedChannelList(v []int32)`

SetAllowedChannelList sets AllowedChannelList field to given value.

### HasAllowedChannelList

`func (o *RadioProfile) HasAllowedChannelList() bool`

HasAllowedChannelList returns a boolean if a field has been set.

### GetName

`func (o *RadioProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadioProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadioProfile) SetName(v string)`

SetName sets Name field to given value.


### GetFrequency

`func (o *RadioProfile) GetFrequency() BulkWritableRadioProfileRequestFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RadioProfile) GetFrequencyOk() (*BulkWritableRadioProfileRequestFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RadioProfile) SetFrequency(v BulkWritableRadioProfileRequestFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RadioProfile) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetTxPowerMin

`func (o *RadioProfile) GetTxPowerMin() int32`

GetTxPowerMin returns the TxPowerMin field if non-nil, zero value otherwise.

### GetTxPowerMinOk

`func (o *RadioProfile) GetTxPowerMinOk() (*int32, bool)`

GetTxPowerMinOk returns a tuple with the TxPowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerMin

`func (o *RadioProfile) SetTxPowerMin(v int32)`

SetTxPowerMin sets TxPowerMin field to given value.

### HasTxPowerMin

`func (o *RadioProfile) HasTxPowerMin() bool`

HasTxPowerMin returns a boolean if a field has been set.

### SetTxPowerMinNil

`func (o *RadioProfile) SetTxPowerMinNil(b bool)`

 SetTxPowerMinNil sets the value for TxPowerMin to be an explicit nil

### UnsetTxPowerMin
`func (o *RadioProfile) UnsetTxPowerMin()`

UnsetTxPowerMin ensures that no value is present for TxPowerMin, not even an explicit nil
### GetTxPowerMax

`func (o *RadioProfile) GetTxPowerMax() int32`

GetTxPowerMax returns the TxPowerMax field if non-nil, zero value otherwise.

### GetTxPowerMaxOk

`func (o *RadioProfile) GetTxPowerMaxOk() (*int32, bool)`

GetTxPowerMaxOk returns a tuple with the TxPowerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerMax

`func (o *RadioProfile) SetTxPowerMax(v int32)`

SetTxPowerMax sets TxPowerMax field to given value.

### HasTxPowerMax

`func (o *RadioProfile) HasTxPowerMax() bool`

HasTxPowerMax returns a boolean if a field has been set.

### SetTxPowerMaxNil

`func (o *RadioProfile) SetTxPowerMaxNil(b bool)`

 SetTxPowerMaxNil sets the value for TxPowerMax to be an explicit nil

### UnsetTxPowerMax
`func (o *RadioProfile) UnsetTxPowerMax()`

UnsetTxPowerMax ensures that no value is present for TxPowerMax, not even an explicit nil
### GetRegulatoryDomain

`func (o *RadioProfile) GetRegulatoryDomain() RegulatoryDomainEnum`

GetRegulatoryDomain returns the RegulatoryDomain field if non-nil, zero value otherwise.

### GetRegulatoryDomainOk

`func (o *RadioProfile) GetRegulatoryDomainOk() (*RegulatoryDomainEnum, bool)`

GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDomain

`func (o *RadioProfile) SetRegulatoryDomain(v RegulatoryDomainEnum)`

SetRegulatoryDomain sets RegulatoryDomain field to given value.


### GetRxPowerMin

`func (o *RadioProfile) GetRxPowerMin() int32`

GetRxPowerMin returns the RxPowerMin field if non-nil, zero value otherwise.

### GetRxPowerMinOk

`func (o *RadioProfile) GetRxPowerMinOk() (*int32, bool)`

GetRxPowerMinOk returns a tuple with the RxPowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPowerMin

`func (o *RadioProfile) SetRxPowerMin(v int32)`

SetRxPowerMin sets RxPowerMin field to given value.

### HasRxPowerMin

`func (o *RadioProfile) HasRxPowerMin() bool`

HasRxPowerMin returns a boolean if a field has been set.

### SetRxPowerMinNil

`func (o *RadioProfile) SetRxPowerMinNil(b bool)`

 SetRxPowerMinNil sets the value for RxPowerMin to be an explicit nil

### UnsetRxPowerMin
`func (o *RadioProfile) UnsetRxPowerMin()`

UnsetRxPowerMin ensures that no value is present for RxPowerMin, not even an explicit nil
### GetSupportedDataRates

`func (o *RadioProfile) GetSupportedDataRates() []BulkWritableCableRequestStatus`

GetSupportedDataRates returns the SupportedDataRates field if non-nil, zero value otherwise.

### GetSupportedDataRatesOk

`func (o *RadioProfile) GetSupportedDataRatesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetSupportedDataRatesOk returns a tuple with the SupportedDataRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDataRates

`func (o *RadioProfile) SetSupportedDataRates(v []BulkWritableCableRequestStatus)`

SetSupportedDataRates sets SupportedDataRates field to given value.

### HasSupportedDataRates

`func (o *RadioProfile) HasSupportedDataRates() bool`

HasSupportedDataRates returns a boolean if a field has been set.

### GetCreated

`func (o *RadioProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RadioProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RadioProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *RadioProfile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *RadioProfile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *RadioProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RadioProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RadioProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *RadioProfile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *RadioProfile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *RadioProfile) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *RadioProfile) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *RadioProfile) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *RadioProfile) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RadioProfile) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RadioProfile) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RadioProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *RadioProfile) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RadioProfile) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RadioProfile) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RadioProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


