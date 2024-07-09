# PatchedBulkWritableComputedFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**Label** | Pointer to **string** | Name of the field as displayed to users | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** | Jinja2 template code for field value | [optional] 
**FallbackValue** | Pointer to **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 

## Methods

### NewPatchedBulkWritableComputedFieldRequest

`func NewPatchedBulkWritableComputedFieldRequest(id string, ) *PatchedBulkWritableComputedFieldRequest`

NewPatchedBulkWritableComputedFieldRequest instantiates a new PatchedBulkWritableComputedFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableComputedFieldRequestWithDefaults

`func NewPatchedBulkWritableComputedFieldRequestWithDefaults() *PatchedBulkWritableComputedFieldRequest`

NewPatchedBulkWritableComputedFieldRequestWithDefaults instantiates a new PatchedBulkWritableComputedFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableComputedFieldRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableComputedFieldRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *PatchedBulkWritableComputedFieldRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedBulkWritableComputedFieldRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedBulkWritableComputedFieldRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetKey

`func (o *PatchedBulkWritableComputedFieldRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedBulkWritableComputedFieldRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedBulkWritableComputedFieldRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkWritableComputedFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkWritableComputedFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkWritableComputedFieldRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableComputedFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableComputedFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableComputedFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplate

`func (o *PatchedBulkWritableComputedFieldRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PatchedBulkWritableComputedFieldRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PatchedBulkWritableComputedFieldRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetFallbackValue

`func (o *PatchedBulkWritableComputedFieldRequest) GetFallbackValue() string`

GetFallbackValue returns the FallbackValue field if non-nil, zero value otherwise.

### GetFallbackValueOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetFallbackValueOk() (*string, bool)`

GetFallbackValueOk returns a tuple with the FallbackValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackValue

`func (o *PatchedBulkWritableComputedFieldRequest) SetFallbackValue(v string)`

SetFallbackValue sets FallbackValue field to given value.

### HasFallbackValue

`func (o *PatchedBulkWritableComputedFieldRequest) HasFallbackValue() bool`

HasFallbackValue returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkWritableComputedFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableComputedFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableComputedFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetAdvancedUi

`func (o *PatchedBulkWritableComputedFieldRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *PatchedBulkWritableComputedFieldRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *PatchedBulkWritableComputedFieldRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *PatchedBulkWritableComputedFieldRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


