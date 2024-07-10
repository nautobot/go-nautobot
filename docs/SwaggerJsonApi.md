# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwaggerJsonRetrieve**](SwaggerJsonApi.md#SwaggerJsonRetrieve) | **Get** /swagger.json | 

# **SwaggerJsonRetrieve**
> map[string]Object SwaggerJsonRetrieve(ctx, optional)


OpenApi3 schema for this API. Format can be selected via content negotiation.  - YAML: application/vnd.oai.openapi - JSON: application/vnd.oai.openapi+json

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwaggerJsonApiSwaggerJsonRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwaggerJsonApiSwaggerJsonRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lang** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

**map[string]Object**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

