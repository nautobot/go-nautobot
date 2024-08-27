# ComputedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ContentType** | **string** |  | 
**Key** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**Grouping** | Pointer to **string** | Human-readable grouping that this computed field belongs to. | [optional] 
**Label** | **string** | Name of the field as displayed to users | 
**Description** | Pointer to **string** |  | [optional] 
**Template** | **string** | Jinja2 template code for field value | 
**FallbackValue** | Pointer to **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 

## Methods

### NewComputedField

`func NewComputedField(id string, objectType string, display string, url string, naturalSlug string, contentType string, label string, template string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ComputedField`

NewComputedField instantiates a new ComputedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputedFieldWithDefaults

`func NewComputedFieldWithDefaults() *ComputedField`

NewComputedFieldWithDefaults instantiates a new ComputedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputedField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputedField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputedField) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *ComputedField) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputedField) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputedField) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ComputedField) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ComputedField) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ComputedField) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ComputedField) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ComputedField) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ComputedField) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ComputedField) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ComputedField) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ComputedField) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetContentType

`func (o *ComputedField) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ComputedField) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ComputedField) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetKey

`func (o *ComputedField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComputedField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComputedField) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ComputedField) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetGrouping

`func (o *ComputedField) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *ComputedField) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *ComputedField) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *ComputedField) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetLabel

`func (o *ComputedField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ComputedField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ComputedField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ComputedField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComputedField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComputedField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComputedField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplate

`func (o *ComputedField) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ComputedField) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ComputedField) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetFallbackValue

`func (o *ComputedField) GetFallbackValue() string`

GetFallbackValue returns the FallbackValue field if non-nil, zero value otherwise.

### GetFallbackValueOk

`func (o *ComputedField) GetFallbackValueOk() (*string, bool)`

GetFallbackValueOk returns a tuple with the FallbackValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackValue

`func (o *ComputedField) SetFallbackValue(v string)`

SetFallbackValue sets FallbackValue field to given value.

### HasFallbackValue

`func (o *ComputedField) HasFallbackValue() bool`

HasFallbackValue returns a boolean if a field has been set.

### GetWeight

`func (o *ComputedField) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ComputedField) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ComputedField) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ComputedField) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetAdvancedUi

`func (o *ComputedField) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *ComputedField) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *ComputedField) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *ComputedField) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.

### GetCreated

`func (o *ComputedField) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ComputedField) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ComputedField) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ComputedField) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ComputedField) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ComputedField) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ComputedField) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ComputedField) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ComputedField) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ComputedField) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *ComputedField) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ComputedField) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ComputedField) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


