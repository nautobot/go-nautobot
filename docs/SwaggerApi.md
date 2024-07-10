# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwaggerRetrieve**](SwaggerApi.md#SwaggerRetrieve) | **Get** /swagger/ | 

# **SwaggerRetrieve**
> map[string]Object SwaggerRetrieve(ctx, optional)


OpenApi3 schema for this API. Format can be selected via content negotiation.  - YAML: application/vnd.oai.openapi - JSON: application/vnd.oai.openapi+json

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerApiSwaggerRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerApiSwaggerRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**|  | 
 **lang** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

**map[string]Object**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

