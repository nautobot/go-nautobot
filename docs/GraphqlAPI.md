# \GraphqlAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphqlCreate**](GraphqlAPI.md#GraphqlCreate) | **Post** /graphql/ | 



## GraphqlCreate

> GraphqlCreate200Response GraphqlCreate(ctx).GraphQLAPIRequest(graphQLAPIRequest).Format(format).Execute()





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
	graphQLAPIRequest := *openapiclient.NewGraphQLAPIRequest("Query_example") // GraphQLAPIRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphqlAPI.GraphqlCreate(context.Background()).GraphQLAPIRequest(graphQLAPIRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphqlAPI.GraphqlCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphqlCreate`: GraphqlCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `GraphqlAPI.GraphqlCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphqlCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLAPIRequest** | [**GraphQLAPIRequest**](GraphQLAPIRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**GraphqlCreate200Response**](GraphqlCreate200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

