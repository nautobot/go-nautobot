# RenderJinja

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateCode** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**RenderedTemplate** | **string** |  | [readonly] 
**RenderedTemplateLines** | **[]string** |  | [readonly] 

## Methods

### NewRenderJinja

`func NewRenderJinja(templateCode string, renderedTemplate string, renderedTemplateLines []string, ) *RenderJinja`

NewRenderJinja instantiates a new RenderJinja object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderJinjaWithDefaults

`func NewRenderJinjaWithDefaults() *RenderJinja`

NewRenderJinjaWithDefaults instantiates a new RenderJinja object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateCode

`func (o *RenderJinja) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *RenderJinja) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *RenderJinja) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetContext

`func (o *RenderJinja) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RenderJinja) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RenderJinja) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *RenderJinja) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRenderedTemplate

`func (o *RenderJinja) GetRenderedTemplate() string`

GetRenderedTemplate returns the RenderedTemplate field if non-nil, zero value otherwise.

### GetRenderedTemplateOk

`func (o *RenderJinja) GetRenderedTemplateOk() (*string, bool)`

GetRenderedTemplateOk returns a tuple with the RenderedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedTemplate

`func (o *RenderJinja) SetRenderedTemplate(v string)`

SetRenderedTemplate sets RenderedTemplate field to given value.


### GetRenderedTemplateLines

`func (o *RenderJinja) GetRenderedTemplateLines() []string`

GetRenderedTemplateLines returns the RenderedTemplateLines field if non-nil, zero value otherwise.

### GetRenderedTemplateLinesOk

`func (o *RenderJinja) GetRenderedTemplateLinesOk() (*[]string, bool)`

GetRenderedTemplateLinesOk returns a tuple with the RenderedTemplateLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedTemplateLines

`func (o *RenderJinja) SetRenderedTemplateLines(v []string)`

SetRenderedTemplateLines sets RenderedTemplateLines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


