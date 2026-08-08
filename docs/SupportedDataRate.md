# SupportedDataRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Standard** | [**StandardEnum**](StandardEnum.md) |  | 
**Rate** | **int32** |  | 
**McsIndex** | Pointer to **NullableInt32** | The Modulation and Coding Scheme (MCS) index is a value used in wireless communications to define the modulation type, coding rate, and number of spatial streams used in a transmission. | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewSupportedDataRate

`func NewSupportedDataRate(objectType string, display string, url string, naturalSlug string, standard StandardEnum, rate int32, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *SupportedDataRate`

NewSupportedDataRate instantiates a new SupportedDataRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedDataRateWithDefaults

`func NewSupportedDataRateWithDefaults() *SupportedDataRate`

NewSupportedDataRateWithDefaults instantiates a new SupportedDataRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportedDataRate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedDataRate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedDataRate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportedDataRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *SupportedDataRate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SupportedDataRate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SupportedDataRate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *SupportedDataRate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SupportedDataRate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SupportedDataRate) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *SupportedDataRate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupportedDataRate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupportedDataRate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *SupportedDataRate) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *SupportedDataRate) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *SupportedDataRate) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetStandard

`func (o *SupportedDataRate) GetStandard() StandardEnum`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *SupportedDataRate) GetStandardOk() (*StandardEnum, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *SupportedDataRate) SetStandard(v StandardEnum)`

SetStandard sets Standard field to given value.


### GetRate

`func (o *SupportedDataRate) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SupportedDataRate) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SupportedDataRate) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetMcsIndex

`func (o *SupportedDataRate) GetMcsIndex() int32`

GetMcsIndex returns the McsIndex field if non-nil, zero value otherwise.

### GetMcsIndexOk

`func (o *SupportedDataRate) GetMcsIndexOk() (*int32, bool)`

GetMcsIndexOk returns a tuple with the McsIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsIndex

`func (o *SupportedDataRate) SetMcsIndex(v int32)`

SetMcsIndex sets McsIndex field to given value.

### HasMcsIndex

`func (o *SupportedDataRate) HasMcsIndex() bool`

HasMcsIndex returns a boolean if a field has been set.

### SetMcsIndexNil

`func (o *SupportedDataRate) SetMcsIndexNil(b bool)`

 SetMcsIndexNil sets the value for McsIndex to be an explicit nil

### UnsetMcsIndex
`func (o *SupportedDataRate) UnsetMcsIndex()`

UnsetMcsIndex ensures that no value is present for McsIndex, not even an explicit nil
### GetCreated

`func (o *SupportedDataRate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SupportedDataRate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SupportedDataRate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *SupportedDataRate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SupportedDataRate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *SupportedDataRate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SupportedDataRate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SupportedDataRate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *SupportedDataRate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *SupportedDataRate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *SupportedDataRate) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *SupportedDataRate) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *SupportedDataRate) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *SupportedDataRate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SupportedDataRate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SupportedDataRate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SupportedDataRate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *SupportedDataRate) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SupportedDataRate) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SupportedDataRate) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SupportedDataRate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


