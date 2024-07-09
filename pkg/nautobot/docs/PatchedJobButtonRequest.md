# PatchedJobButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Jinja2 template code for button text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Buttons which render as empty text will not be displayed. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group. | [optional] 
**ButtonClass** | Pointer to [**ButtonClassEnum**](ButtonClassEnum.md) |  | [optional] 
**Confirmation** | Pointer to **bool** | Enable confirmation pop-up box. &lt;span class&#x3D;&#39;text-danger&#39;&gt;WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!&lt;/span&gt; | [optional] 
**Job** | Pointer to [**BulkWritableJobButtonRequestJob**](BulkWritableJobButtonRequestJob.md) |  | [optional] 

## Methods

### NewPatchedJobButtonRequest

`func NewPatchedJobButtonRequest() *PatchedJobButtonRequest`

NewPatchedJobButtonRequest instantiates a new PatchedJobButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJobButtonRequestWithDefaults

`func NewPatchedJobButtonRequestWithDefaults() *PatchedJobButtonRequest`

NewPatchedJobButtonRequestWithDefaults instantiates a new PatchedJobButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *PatchedJobButtonRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedJobButtonRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedJobButtonRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedJobButtonRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedJobButtonRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedJobButtonRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedJobButtonRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedJobButtonRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetText

`func (o *PatchedJobButtonRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PatchedJobButtonRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PatchedJobButtonRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PatchedJobButtonRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedJobButtonRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedJobButtonRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedJobButtonRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedJobButtonRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedJobButtonRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedJobButtonRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedJobButtonRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedJobButtonRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *PatchedJobButtonRequest) GetButtonClass() ButtonClassEnum`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *PatchedJobButtonRequest) GetButtonClassOk() (*ButtonClassEnum, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *PatchedJobButtonRequest) SetButtonClass(v ButtonClassEnum)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *PatchedJobButtonRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetConfirmation

`func (o *PatchedJobButtonRequest) GetConfirmation() bool`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *PatchedJobButtonRequest) GetConfirmationOk() (*bool, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *PatchedJobButtonRequest) SetConfirmation(v bool)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *PatchedJobButtonRequest) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetJob

`func (o *PatchedJobButtonRequest) GetJob() BulkWritableJobButtonRequestJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *PatchedJobButtonRequest) GetJobOk() (*BulkWritableJobButtonRequestJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *PatchedJobButtonRequest) SetJob(v BulkWritableJobButtonRequestJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *PatchedJobButtonRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


