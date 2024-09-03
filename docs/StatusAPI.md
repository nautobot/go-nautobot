# \StatusAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusRetrieve**](StatusAPI.md#StatusRetrieve) | **Get** /status/ | 



## StatusRetrieve

> StatusRetrieve200Response StatusRetrieve(ctx).Format(format).Depth(depth).Execute()





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
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.StatusRetrieve(context.Background()).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.StatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusRetrieve`: StatusRetrieve200Response
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.StatusRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**StatusRetrieve200Response**](StatusRetrieve200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

