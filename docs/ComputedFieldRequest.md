# ComputedFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**Key** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**Label** | **string** | Name of the field as displayed to users | 
**Description** | Pointer to **string** |  | [optional] 
**Template** | **string** | Jinja2 template code for field value | 
**FallbackValue** | Pointer to **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 

## Methods

### NewComputedFieldRequest

`func NewComputedFieldRequest(contentType string, label string, template string, ) *ComputedFieldRequest`

NewComputedFieldRequest instantiates a new ComputedFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputedFieldRequestWithDefaults

`func NewComputedFieldRequestWithDefaults() *ComputedFieldRequest`

NewComputedFieldRequestWithDefaults instantiates a new ComputedFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ComputedFieldRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ComputedFieldRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ComputedFieldRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetKey

`func (o *ComputedFieldRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComputedFieldRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComputedFieldRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ComputedFieldRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *ComputedFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ComputedFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ComputedFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ComputedFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComputedFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComputedFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComputedFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplate

`func (o *ComputedFieldRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ComputedFieldRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ComputedFieldRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetFallbackValue

`func (o *ComputedFieldRequest) GetFallbackValue() string`

GetFallbackValue returns the FallbackValue field if non-nil, zero value otherwise.

### GetFallbackValueOk

`func (o *ComputedFieldRequest) GetFallbackValueOk() (*string, bool)`

GetFallbackValueOk returns a tuple with the FallbackValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackValue

`func (o *ComputedFieldRequest) SetFallbackValue(v string)`

SetFallbackValue sets FallbackValue field to given value.

### HasFallbackValue

`func (o *ComputedFieldRequest) HasFallbackValue() bool`

HasFallbackValue returns a boolean if a field has been set.

### GetWeight

`func (o *ComputedFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ComputedFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ComputedFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ComputedFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetAdvancedUi

`func (o *ComputedFieldRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *ComputedFieldRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *ComputedFieldRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *ComputedFieldRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


