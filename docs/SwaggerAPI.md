# \SwaggerAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwaggerRetrieve**](SwaggerAPI.md#SwaggerRetrieve) | **Get** /swagger/ | 



## SwaggerRetrieve

> map[string]interface{} SwaggerRetrieve(ctx).Format(format).Lang(lang).Depth(depth).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TobiPeterG/go-nautobot"
)

func main() {
	format := openapiclient.swagger_retrieve_format_parameter("json") // SwaggerRetrieveFormatParameter |  (optional)
	lang := openapiclient.swagger_retrieve_lang_parameter("af") // SwaggerRetrieveLangParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwaggerAPI.SwaggerRetrieve(context.Background()).Format(format).Lang(lang).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwaggerAPI.SwaggerRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwaggerRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SwaggerAPI.SwaggerRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwaggerRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**SwaggerRetrieveFormatParameter**](SwaggerRetrieveFormatParameter.md) |  | 
 **lang** | [**SwaggerRetrieveLangParameter**](SwaggerRetrieveLangParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

