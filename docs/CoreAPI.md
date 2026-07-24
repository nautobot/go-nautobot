# \CoreAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreRenderJinjaTemplateCreate**](CoreAPI.md#CoreRenderJinjaTemplateCreate) | **Post** /core/render-jinja-template/ | 



## CoreRenderJinjaTemplateCreate

> RenderJinja CoreRenderJinjaTemplateCreate(ctx).RenderJinjaRequest(renderJinjaRequest).Format(format).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nautobot/go-nautobot/v3"
)

func main() {
	renderJinjaRequest := *openapiclient.NewRenderJinjaRequest("TemplateCode_example") // RenderJinjaRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreRenderJinjaTemplateCreate(context.Background()).RenderJinjaRequest(renderJinjaRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreRenderJinjaTemplateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreRenderJinjaTemplateCreate`: RenderJinja
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreRenderJinjaTemplateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreRenderJinjaTemplateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renderJinjaRequest** | [**RenderJinjaRequest**](RenderJinjaRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**RenderJinja**](RenderJinja.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

