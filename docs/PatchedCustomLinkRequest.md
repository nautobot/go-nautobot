# PatchedCustomLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Jinja2 template code for link text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Links which render as empty text will not be displayed. | [optional] 
**TargetUrl** | Pointer to **string** | Jinja2 template code for link URL. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to [**ButtonClassEnum**](ButtonClassEnum.md) | The class of the first link in a group will be used for the dropdown button | [optional] 
**NewWindow** | Pointer to **bool** | Force link to open in a new window | [optional] 

## Methods

### NewPatchedCustomLinkRequest

`func NewPatchedCustomLinkRequest() *PatchedCustomLinkRequest`

NewPatchedCustomLinkRequest instantiates a new PatchedCustomLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCustomLinkRequestWithDefaults

`func NewPatchedCustomLinkRequestWithDefaults() *PatchedCustomLinkRequest`

NewPatchedCustomLinkRequestWithDefaults instantiates a new PatchedCustomLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *PatchedCustomLinkRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedCustomLinkRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedCustomLinkRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedCustomLinkRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetName

`func (o *PatchedCustomLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCustomLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCustomLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCustomLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetText

`func (o *PatchedCustomLinkRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PatchedCustomLinkRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PatchedCustomLinkRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PatchedCustomLinkRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTargetUrl

`func (o *PatchedCustomLinkRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *PatchedCustomLinkRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *PatchedCustomLinkRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *PatchedCustomLinkRequest) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedCustomLinkRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedCustomLinkRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedCustomLinkRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedCustomLinkRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedCustomLinkRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedCustomLinkRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedCustomLinkRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedCustomLinkRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *PatchedCustomLinkRequest) GetButtonClass() ButtonClassEnum`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *PatchedCustomLinkRequest) GetButtonClassOk() (*ButtonClassEnum, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *PatchedCustomLinkRequest) SetButtonClass(v ButtonClassEnum)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *PatchedCustomLinkRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetNewWindow

`func (o *PatchedCustomLinkRequest) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *PatchedCustomLinkRequest) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *PatchedCustomLinkRequest) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.

### HasNewWindow

`func (o *PatchedCustomLinkRequest) HasNewWindow() bool`

HasNewWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


