# \SwaggerYamlAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwaggerYamlRetrieve**](SwaggerYamlAPI.md#SwaggerYamlRetrieve) | **Get** /swagger.yaml | 



## SwaggerYamlRetrieve

> map[string]interface{} SwaggerYamlRetrieve(ctx).Lang(lang).Depth(depth).Execute()





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
	lang := openapiclient.swagger_retrieve_lang_parameter("af") // SwaggerRetrieveLangParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwaggerYamlAPI.SwaggerYamlRetrieve(context.Background()).Lang(lang).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwaggerYamlAPI.SwaggerYamlRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwaggerYamlRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SwaggerYamlAPI.SwaggerYamlRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwaggerYamlRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lang** | [**SwaggerRetrieveLangParameter**](SwaggerRetrieveLangParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

