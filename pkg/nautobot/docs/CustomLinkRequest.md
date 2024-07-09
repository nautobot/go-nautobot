# CustomLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**Name** | **string** |  | 
**Text** | **string** | Jinja2 template code for link text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Links which render as empty text will not be displayed. | 
**TargetUrl** | **string** | Jinja2 template code for link URL. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. | 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to [**ButtonClassEnum**](ButtonClassEnum.md) | The class of the first link in a group will be used for the dropdown button | [optional] 
**NewWindow** | **bool** | Force link to open in a new window | 

## Methods

### NewCustomLinkRequest

`func NewCustomLinkRequest(contentType string, name string, text string, targetUrl string, newWindow bool, ) *CustomLinkRequest`

NewCustomLinkRequest instantiates a new CustomLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLinkRequestWithDefaults

`func NewCustomLinkRequestWithDefaults() *CustomLinkRequest`

NewCustomLinkRequestWithDefaults instantiates a new CustomLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *CustomLinkRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CustomLinkRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CustomLinkRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *CustomLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomLinkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetText

`func (o *CustomLinkRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CustomLinkRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CustomLinkRequest) SetText(v string)`

SetText sets Text field to given value.


### GetTargetUrl

`func (o *CustomLinkRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *CustomLinkRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *CustomLinkRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetWeight

`func (o *CustomLinkRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomLinkRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomLinkRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CustomLinkRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *CustomLinkRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CustomLinkRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CustomLinkRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CustomLinkRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *CustomLinkRequest) GetButtonClass() ButtonClassEnum`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *CustomLinkRequest) GetButtonClassOk() (*ButtonClassEnum, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *CustomLinkRequest) SetButtonClass(v ButtonClassEnum)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *CustomLinkRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetNewWindow

`func (o *CustomLinkRequest) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *CustomLinkRequest) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *CustomLinkRequest) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


