# PatchedBulkWritableJobButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Jinja2 template code for button text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Buttons which render as empty text will not be displayed. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group. | [optional] 
**ButtonClass** | Pointer to [**ButtonClassEnum**](ButtonClassEnum.md) |  | [optional] 
**Confirmation** | Pointer to **bool** | Enable confirmation pop-up box. &lt;span class&#x3D;&#39;text-danger&#39;&gt;WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!&lt;/span&gt; | [optional] 
**Job** | Pointer to [**BulkWritableJobButtonRequestJob**](BulkWritableJobButtonRequestJob.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableJobButtonRequest

`func NewPatchedBulkWritableJobButtonRequest(id string, ) *PatchedBulkWritableJobButtonRequest`

NewPatchedBulkWritableJobButtonRequest instantiates a new PatchedBulkWritableJobButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableJobButtonRequestWithDefaults

`func NewPatchedBulkWritableJobButtonRequestWithDefaults() *PatchedBulkWritableJobButtonRequest`

NewPatchedBulkWritableJobButtonRequestWithDefaults instantiates a new PatchedBulkWritableJobButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableJobButtonRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableJobButtonRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableJobButtonRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentTypes

`func (o *PatchedBulkWritableJobButtonRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedBulkWritableJobButtonRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedBulkWritableJobButtonRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedBulkWritableJobButtonRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableJobButtonRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableJobButtonRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableJobButtonRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableJobButtonRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetText

`func (o *PatchedBulkWritableJobButtonRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PatchedBulkWritableJobButtonRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PatchedBulkWritableJobButtonRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PatchedBulkWritableJobButtonRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkWritableJobButtonRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableJobButtonRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableJobButtonRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableJobButtonRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedBulkWritableJobButtonRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedBulkWritableJobButtonRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedBulkWritableJobButtonRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedBulkWritableJobButtonRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *PatchedBulkWritableJobButtonRequest) GetButtonClass() ButtonClassEnum`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *PatchedBulkWritableJobButtonRequest) GetButtonClassOk() (*ButtonClassEnum, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *PatchedBulkWritableJobButtonRequest) SetButtonClass(v ButtonClassEnum)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *PatchedBulkWritableJobButtonRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetConfirmation

`func (o *PatchedBulkWritableJobButtonRequest) GetConfirmation() bool`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *PatchedBulkWritableJobButtonRequest) GetConfirmationOk() (*bool, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *PatchedBulkWritableJobButtonRequest) SetConfirmation(v bool)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *PatchedBulkWritableJobButtonRequest) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetJob

`func (o *PatchedBulkWritableJobButtonRequest) GetJob() BulkWritableJobButtonRequestJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *PatchedBulkWritableJobButtonRequest) GetJobOk() (*BulkWritableJobButtonRequestJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *PatchedBulkWritableJobButtonRequest) SetJob(v BulkWritableJobButtonRequestJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *PatchedBulkWritableJobButtonRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


