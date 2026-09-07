# WritableComputedFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**Key** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**OutputType** | Pointer to [**OutputTypeEnum**](OutputTypeEnum.md) | How the rendered result is to be output, by default it is as Plain Text | [optional] 
**Grouping** | Pointer to **string** | Human-readable grouping that this computed field belongs to. | [optional] 
**Label** | **string** | Name of the field as displayed to users | 
**Description** | Pointer to **string** |  | [optional] 
**Template** | **string** | Jinja2 template code for field value | 
**FallbackValue** | Pointer to **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 

## Methods

### NewWritableComputedFieldRequest

`func NewWritableComputedFieldRequest(contentType string, label string, template string, ) *WritableComputedFieldRequest`

NewWritableComputedFieldRequest instantiates a new WritableComputedFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableComputedFieldRequestWithDefaults

`func NewWritableComputedFieldRequestWithDefaults() *WritableComputedFieldRequest`

NewWritableComputedFieldRequestWithDefaults instantiates a new WritableComputedFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableComputedFieldRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableComputedFieldRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableComputedFieldRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WritableComputedFieldRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *WritableComputedFieldRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WritableComputedFieldRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WritableComputedFieldRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetKey

`func (o *WritableComputedFieldRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WritableComputedFieldRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WritableComputedFieldRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WritableComputedFieldRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOutputType

`func (o *WritableComputedFieldRequest) GetOutputType() OutputTypeEnum`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *WritableComputedFieldRequest) GetOutputTypeOk() (*OutputTypeEnum, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *WritableComputedFieldRequest) SetOutputType(v OutputTypeEnum)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *WritableComputedFieldRequest) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetGrouping

`func (o *WritableComputedFieldRequest) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *WritableComputedFieldRequest) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *WritableComputedFieldRequest) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *WritableComputedFieldRequest) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetLabel

`func (o *WritableComputedFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableComputedFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableComputedFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *WritableComputedFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableComputedFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableComputedFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableComputedFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplate

`func (o *WritableComputedFieldRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WritableComputedFieldRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WritableComputedFieldRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetFallbackValue

`func (o *WritableComputedFieldRequest) GetFallbackValue() string`

GetFallbackValue returns the FallbackValue field if non-nil, zero value otherwise.

### GetFallbackValueOk

`func (o *WritableComputedFieldRequest) GetFallbackValueOk() (*string, bool)`

GetFallbackValueOk returns a tuple with the FallbackValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackValue

`func (o *WritableComputedFieldRequest) SetFallbackValue(v string)`

SetFallbackValue sets FallbackValue field to given value.

### HasFallbackValue

`func (o *WritableComputedFieldRequest) HasFallbackValue() bool`

HasFallbackValue returns a boolean if a field has been set.

### GetWeight

`func (o *WritableComputedFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WritableComputedFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WritableComputedFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WritableComputedFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetAdvancedUi

`func (o *WritableComputedFieldRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *WritableComputedFieldRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *WritableComputedFieldRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *WritableComputedFieldRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


