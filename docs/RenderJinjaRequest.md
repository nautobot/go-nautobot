# RenderJinjaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateCode** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRenderJinjaRequest

`func NewRenderJinjaRequest(templateCode string, ) *RenderJinjaRequest`

NewRenderJinjaRequest instantiates a new RenderJinjaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderJinjaRequestWithDefaults

`func NewRenderJinjaRequestWithDefaults() *RenderJinjaRequest`

NewRenderJinjaRequestWithDefaults instantiates a new RenderJinjaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateCode

`func (o *RenderJinjaRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *RenderJinjaRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *RenderJinjaRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetContext

`func (o *RenderJinjaRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RenderJinjaRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RenderJinjaRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *RenderJinjaRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


