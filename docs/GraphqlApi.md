# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphqlCreate**](GraphqlApi.md#GraphqlCreate) | **Post** /graphql/ | 

# **GraphqlCreate**
> InlineResponse2001 GraphqlCreate(ctx, body, optional)


Query the database using a GraphQL query

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GraphQlapiRequest**](GraphQlapiRequest.md)|  | 
 **optional** | ***GraphqlApiGraphqlCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GraphqlApiGraphqlCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

