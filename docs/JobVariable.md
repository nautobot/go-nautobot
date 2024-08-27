# JobVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**Type** | **string** |  | [readonly] 
**Label** | **string** |  | [readonly] 
**HelpText** | **string** |  | [readonly] 
**Default** | **interface{}** |  | [readonly] 
**Required** | **bool** |  | [readonly] 
**MinLength** | **int32** |  | [readonly] 
**MaxLength** | **int32** |  | [readonly] 
**MinValue** | **int32** |  | [readonly] 
**MaxValue** | **int32** |  | [readonly] 
**Choices** | **interface{}** |  | [readonly] 
**Model** | **string** |  | [readonly] 

## Methods

### NewJobVariable

`func NewJobVariable(name string, type_ string, label string, helpText string, default_ interface{}, required bool, minLength int32, maxLength int32, minValue int32, maxValue int32, choices interface{}, model string, ) *JobVariable`

NewJobVariable instantiates a new JobVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobVariableWithDefaults

`func NewJobVariableWithDefaults() *JobVariable`

NewJobVariableWithDefaults instantiates a new JobVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobVariable) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *JobVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobVariable) SetType(v string)`

SetType sets Type field to given value.


### GetLabel

`func (o *JobVariable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *JobVariable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *JobVariable) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetHelpText

`func (o *JobVariable) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *JobVariable) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *JobVariable) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.


### GetDefault

`func (o *JobVariable) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *JobVariable) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *JobVariable) SetDefault(v interface{})`

SetDefault sets Default field to given value.


### SetDefaultNil

`func (o *JobVariable) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *JobVariable) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetRequired

`func (o *JobVariable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *JobVariable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *JobVariable) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetMinLength

`func (o *JobVariable) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *JobVariable) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *JobVariable) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.


### GetMaxLength

`func (o *JobVariable) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *JobVariable) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *JobVariable) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.


### GetMinValue

`func (o *JobVariable) GetMinValue() int32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *JobVariable) GetMinValueOk() (*int32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *JobVariable) SetMinValue(v int32)`

SetMinValue sets MinValue field to given value.


### GetMaxValue

`func (o *JobVariable) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *JobVariable) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *JobVariable) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.


### GetChoices

`func (o *JobVariable) GetChoices() interface{}`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *JobVariable) GetChoicesOk() (*interface{}, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *JobVariable) SetChoices(v interface{})`

SetChoices sets Choices field to given value.


### SetChoicesNil

`func (o *JobVariable) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *JobVariable) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetModel

`func (o *JobVariable) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *JobVariable) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *JobVariable) SetModel(v string)`

SetModel sets Model field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


