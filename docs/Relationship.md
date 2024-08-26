# Relationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**SourceType** | **string** |  | 
**DestinationType** | **string** |  | 
**Label** | **string** | Label of the relationship as displayed to users | 
**Key** | Pointer to **string** | Internal relationship key. Please use underscores rather than dashes in this key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**RelationshipTypeChoices**](RelationshipTypeChoices.md) | Cardinality of this relationship | [optional] 
**RequiredOn** | Pointer to [**BulkWritableRelationshipRequestRequiredOn**](BulkWritableRelationshipRequestRequiredOn.md) |  | [optional] 
**SourceLabel** | Pointer to **string** | Label for related destination objects, as displayed on the source object. | [optional] 
**SourceHidden** | Pointer to **bool** | Hide this relationship on the source object. | [optional] 
**SourceFilter** | Pointer to **map[string]interface{}** | Filterset filter matching the applicable source objects of the selected type | [optional] 
**DestinationLabel** | Pointer to **string** | Label for related source objects, as displayed on the destination object. | [optional] 
**DestinationHidden** | Pointer to **bool** | Hide this relationship on the destination object. | [optional] 
**DestinationFilter** | Pointer to **map[string]interface{}** | Filterset filter matching the applicable destination objects of the selected type | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 

## Methods

### NewRelationship

`func NewRelationship(id string, objectType string, display string, url string, naturalSlug string, sourceType string, destinationType string, label string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Relationship`

NewRelationship instantiates a new Relationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipWithDefaults

`func NewRelationshipWithDefaults() *Relationship`

NewRelationshipWithDefaults instantiates a new Relationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Relationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Relationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Relationship) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Relationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Relationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Relationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Relationship) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Relationship) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Relationship) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Relationship) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Relationship) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Relationship) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Relationship) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Relationship) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Relationship) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetSourceType

`func (o *Relationship) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Relationship) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Relationship) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetDestinationType

`func (o *Relationship) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *Relationship) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *Relationship) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetLabel

`func (o *Relationship) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Relationship) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Relationship) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetKey

`func (o *Relationship) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Relationship) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Relationship) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Relationship) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *Relationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Relationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Relationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Relationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Relationship) GetType() RelationshipTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Relationship) GetTypeOk() (*RelationshipTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Relationship) SetType(v RelationshipTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *Relationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequiredOn

`func (o *Relationship) GetRequiredOn() BulkWritableRelationshipRequestRequiredOn`

GetRequiredOn returns the RequiredOn field if non-nil, zero value otherwise.

### GetRequiredOnOk

`func (o *Relationship) GetRequiredOnOk() (*BulkWritableRelationshipRequestRequiredOn, bool)`

GetRequiredOnOk returns a tuple with the RequiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOn

`func (o *Relationship) SetRequiredOn(v BulkWritableRelationshipRequestRequiredOn)`

SetRequiredOn sets RequiredOn field to given value.

### HasRequiredOn

`func (o *Relationship) HasRequiredOn() bool`

HasRequiredOn returns a boolean if a field has been set.

### GetSourceLabel

`func (o *Relationship) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *Relationship) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *Relationship) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *Relationship) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### GetSourceHidden

`func (o *Relationship) GetSourceHidden() bool`

GetSourceHidden returns the SourceHidden field if non-nil, zero value otherwise.

### GetSourceHiddenOk

`func (o *Relationship) GetSourceHiddenOk() (*bool, bool)`

GetSourceHiddenOk returns a tuple with the SourceHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHidden

`func (o *Relationship) SetSourceHidden(v bool)`

SetSourceHidden sets SourceHidden field to given value.

### HasSourceHidden

`func (o *Relationship) HasSourceHidden() bool`

HasSourceHidden returns a boolean if a field has been set.

### GetSourceFilter

`func (o *Relationship) GetSourceFilter() map[string]interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *Relationship) GetSourceFilterOk() (*map[string]interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *Relationship) SetSourceFilter(v map[string]interface{})`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *Relationship) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### SetSourceFilterNil

`func (o *Relationship) SetSourceFilterNil(b bool)`

 SetSourceFilterNil sets the value for SourceFilter to be an explicit nil

### UnsetSourceFilter
`func (o *Relationship) UnsetSourceFilter()`

UnsetSourceFilter ensures that no value is present for SourceFilter, not even an explicit nil
### GetDestinationLabel

`func (o *Relationship) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *Relationship) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *Relationship) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.

### HasDestinationLabel

`func (o *Relationship) HasDestinationLabel() bool`

HasDestinationLabel returns a boolean if a field has been set.

### GetDestinationHidden

`func (o *Relationship) GetDestinationHidden() bool`

GetDestinationHidden returns the DestinationHidden field if non-nil, zero value otherwise.

### GetDestinationHiddenOk

`func (o *Relationship) GetDestinationHiddenOk() (*bool, bool)`

GetDestinationHiddenOk returns a tuple with the DestinationHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHidden

`func (o *Relationship) SetDestinationHidden(v bool)`

SetDestinationHidden sets DestinationHidden field to given value.

### HasDestinationHidden

`func (o *Relationship) HasDestinationHidden() bool`

HasDestinationHidden returns a boolean if a field has been set.

### GetDestinationFilter

`func (o *Relationship) GetDestinationFilter() map[string]interface{}`

GetDestinationFilter returns the DestinationFilter field if non-nil, zero value otherwise.

### GetDestinationFilterOk

`func (o *Relationship) GetDestinationFilterOk() (*map[string]interface{}, bool)`

GetDestinationFilterOk returns a tuple with the DestinationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFilter

`func (o *Relationship) SetDestinationFilter(v map[string]interface{})`

SetDestinationFilter sets DestinationFilter field to given value.

### HasDestinationFilter

`func (o *Relationship) HasDestinationFilter() bool`

HasDestinationFilter returns a boolean if a field has been set.

### SetDestinationFilterNil

`func (o *Relationship) SetDestinationFilterNil(b bool)`

 SetDestinationFilterNil sets the value for DestinationFilter to be an explicit nil

### UnsetDestinationFilter
`func (o *Relationship) UnsetDestinationFilter()`

UnsetDestinationFilter ensures that no value is present for DestinationFilter, not even an explicit nil
### GetAdvancedUi

`func (o *Relationship) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *Relationship) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *Relationship) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *Relationship) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.

### GetCreated

`func (o *Relationship) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Relationship) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Relationship) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Relationship) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Relationship) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Relationship) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Relationship) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Relationship) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Relationship) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Relationship) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Relationship) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Relationship) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Relationship) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


