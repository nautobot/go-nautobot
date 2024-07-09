# BulkWritableCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentTypes** | **[]string** |  | 
**Type** | [**CustomFieldTypeChoices**](CustomFieldTypeChoices.md) |  | 
**FilterLogic** | Pointer to [**FilterLogicEnum**](FilterLogicEnum.md) |  | [optional] 
**Label** | **string** |  | 
**Grouping** | Pointer to **string** | Human-readable grouping that this custom field belongs to. | [optional] 
**Key** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**Description** | Pointer to **string** | A helpful description for this field. | [optional] 
**Required** | Pointer to **bool** | If true, this field is required when creating new objects or editing an existing object. | [optional] 
**Default** | Pointer to **map[string]interface{}** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**Weight** | Pointer to **int32** | Fields with higher weights appear lower in a form. | [optional] 
**ValidationMinimum** | Pointer to **NullableInt64** | Minimum allowed value (for numeric fields) or length (for text fields). | [optional] 
**ValidationMaximum** | Pointer to **NullableInt64** | Maximum allowed value (for numeric fields) or length (for text fields). | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. Regular expression on select and multi-select will be applied at &lt;code&gt;Custom Field Choices&lt;/code&gt; definition. | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 

## Methods

### NewBulkWritableCustomFieldRequest

`func NewBulkWritableCustomFieldRequest(id string, contentTypes []string, type_ CustomFieldTypeChoices, label string, ) *BulkWritableCustomFieldRequest`

NewBulkWritableCustomFieldRequest instantiates a new BulkWritableCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCustomFieldRequestWithDefaults

`func NewBulkWritableCustomFieldRequestWithDefaults() *BulkWritableCustomFieldRequest`

NewBulkWritableCustomFieldRequestWithDefaults instantiates a new BulkWritableCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCustomFieldRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCustomFieldRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCustomFieldRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentTypes

`func (o *BulkWritableCustomFieldRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *BulkWritableCustomFieldRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *BulkWritableCustomFieldRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetType

`func (o *BulkWritableCustomFieldRequest) GetType() CustomFieldTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkWritableCustomFieldRequest) GetTypeOk() (*CustomFieldTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkWritableCustomFieldRequest) SetType(v CustomFieldTypeChoices)`

SetType sets Type field to given value.


### GetFilterLogic

`func (o *BulkWritableCustomFieldRequest) GetFilterLogic() FilterLogicEnum`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *BulkWritableCustomFieldRequest) GetFilterLogicOk() (*FilterLogicEnum, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *BulkWritableCustomFieldRequest) SetFilterLogic(v FilterLogicEnum)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *BulkWritableCustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetLabel

`func (o *BulkWritableCustomFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkWritableCustomFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkWritableCustomFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGrouping

`func (o *BulkWritableCustomFieldRequest) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BulkWritableCustomFieldRequest) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BulkWritableCustomFieldRequest) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BulkWritableCustomFieldRequest) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetKey

`func (o *BulkWritableCustomFieldRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BulkWritableCustomFieldRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BulkWritableCustomFieldRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BulkWritableCustomFieldRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritableCustomFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableCustomFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableCustomFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableCustomFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *BulkWritableCustomFieldRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BulkWritableCustomFieldRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BulkWritableCustomFieldRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BulkWritableCustomFieldRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefault

`func (o *BulkWritableCustomFieldRequest) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BulkWritableCustomFieldRequest) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BulkWritableCustomFieldRequest) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *BulkWritableCustomFieldRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *BulkWritableCustomFieldRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *BulkWritableCustomFieldRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetWeight

`func (o *BulkWritableCustomFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkWritableCustomFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkWritableCustomFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkWritableCustomFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *BulkWritableCustomFieldRequest) GetValidationMinimum() int64`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *BulkWritableCustomFieldRequest) GetValidationMinimumOk() (*int64, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *BulkWritableCustomFieldRequest) SetValidationMinimum(v int64)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *BulkWritableCustomFieldRequest) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *BulkWritableCustomFieldRequest) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *BulkWritableCustomFieldRequest) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *BulkWritableCustomFieldRequest) GetValidationMaximum() int64`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *BulkWritableCustomFieldRequest) GetValidationMaximumOk() (*int64, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *BulkWritableCustomFieldRequest) SetValidationMaximum(v int64)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *BulkWritableCustomFieldRequest) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *BulkWritableCustomFieldRequest) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *BulkWritableCustomFieldRequest) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *BulkWritableCustomFieldRequest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *BulkWritableCustomFieldRequest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *BulkWritableCustomFieldRequest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *BulkWritableCustomFieldRequest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetAdvancedUi

`func (o *BulkWritableCustomFieldRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *BulkWritableCustomFieldRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *BulkWritableCustomFieldRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *BulkWritableCustomFieldRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


