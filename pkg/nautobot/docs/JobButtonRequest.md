# JobButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Text** | **string** | Jinja2 template code for button text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Buttons which render as empty text will not be displayed. | 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group. | [optional] 
**ButtonClass** | Pointer to [**ButtonClassEnum**](ButtonClassEnum.md) |  | [optional] 
**Confirmation** | Pointer to **bool** | Enable confirmation pop-up box. &lt;span class&#x3D;&#39;text-danger&#39;&gt;WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!&lt;/span&gt; | [optional] 
**Job** | [**BulkWritableJobButtonRequestJob**](BulkWritableJobButtonRequestJob.md) |  | 

## Methods

### NewJobButtonRequest

`func NewJobButtonRequest(contentTypes []string, name string, text string, job BulkWritableJobButtonRequestJob, ) *JobButtonRequest`

NewJobButtonRequest instantiates a new JobButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobButtonRequestWithDefaults

`func NewJobButtonRequestWithDefaults() *JobButtonRequest`

NewJobButtonRequestWithDefaults instantiates a new JobButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *JobButtonRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *JobButtonRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *JobButtonRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetName

`func (o *JobButtonRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobButtonRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobButtonRequest) SetName(v string)`

SetName sets Name field to given value.


### GetText

`func (o *JobButtonRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *JobButtonRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *JobButtonRequest) SetText(v string)`

SetText sets Text field to given value.


### GetWeight

`func (o *JobButtonRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *JobButtonRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *JobButtonRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *JobButtonRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *JobButtonRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *JobButtonRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *JobButtonRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *JobButtonRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *JobButtonRequest) GetButtonClass() ButtonClassEnum`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *JobButtonRequest) GetButtonClassOk() (*ButtonClassEnum, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *JobButtonRequest) SetButtonClass(v ButtonClassEnum)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *JobButtonRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetConfirmation

`func (o *JobButtonRequest) GetConfirmation() bool`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *JobButtonRequest) GetConfirmationOk() (*bool, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *JobButtonRequest) SetConfirmation(v bool)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *JobButtonRequest) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetJob

`func (o *JobButtonRequest) GetJob() BulkWritableJobButtonRequestJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobButtonRequest) GetJobOk() (*BulkWritableJobButtonRequestJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobButtonRequest) SetJob(v BulkWritableJobButtonRequestJob)`

SetJob sets Job field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


