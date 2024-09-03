# \UsersAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigRetrieve**](UsersAPI.md#UsersConfigRetrieve) | **Get** /users/config/ | 
[**UsersGroupsBulkDestroy**](UsersAPI.md#UsersGroupsBulkDestroy) | **Delete** /users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersAPI.md#UsersGroupsBulkPartialUpdate) | **Patch** /users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersAPI.md#UsersGroupsBulkUpdate) | **Put** /users/groups/ | 
[**UsersGroupsCreate**](UsersAPI.md#UsersGroupsCreate) | **Post** /users/groups/ | 
[**UsersGroupsDestroy**](UsersAPI.md#UsersGroupsDestroy) | **Delete** /users/groups/{id}/ | 
[**UsersGroupsList**](UsersAPI.md#UsersGroupsList) | **Get** /users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersAPI.md#UsersGroupsPartialUpdate) | **Patch** /users/groups/{id}/ | 
[**UsersGroupsRetrieve**](UsersAPI.md#UsersGroupsRetrieve) | **Get** /users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersAPI.md#UsersGroupsUpdate) | **Put** /users/groups/{id}/ | 
[**UsersPermissionsBulkDestroy**](UsersAPI.md#UsersPermissionsBulkDestroy) | **Delete** /users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersAPI.md#UsersPermissionsBulkPartialUpdate) | **Patch** /users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersAPI.md#UsersPermissionsBulkUpdate) | **Put** /users/permissions/ | 
[**UsersPermissionsCreate**](UsersAPI.md#UsersPermissionsCreate) | **Post** /users/permissions/ | 
[**UsersPermissionsDestroy**](UsersAPI.md#UsersPermissionsDestroy) | **Delete** /users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersAPI.md#UsersPermissionsList) | **Get** /users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersAPI.md#UsersPermissionsPartialUpdate) | **Patch** /users/permissions/{id}/ | 
[**UsersPermissionsRetrieve**](UsersAPI.md#UsersPermissionsRetrieve) | **Get** /users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersAPI.md#UsersPermissionsUpdate) | **Put** /users/permissions/{id}/ | 
[**UsersTokensBulkDestroy**](UsersAPI.md#UsersTokensBulkDestroy) | **Delete** /users/tokens/ | 
[**UsersTokensBulkPartialUpdate**](UsersAPI.md#UsersTokensBulkPartialUpdate) | **Patch** /users/tokens/ | 
[**UsersTokensBulkUpdate**](UsersAPI.md#UsersTokensBulkUpdate) | **Put** /users/tokens/ | 
[**UsersTokensCreate**](UsersAPI.md#UsersTokensCreate) | **Post** /users/tokens/ | 
[**UsersTokensDestroy**](UsersAPI.md#UsersTokensDestroy) | **Delete** /users/tokens/{id}/ | 
[**UsersTokensList**](UsersAPI.md#UsersTokensList) | **Get** /users/tokens/ | 
[**UsersTokensPartialUpdate**](UsersAPI.md#UsersTokensPartialUpdate) | **Patch** /users/tokens/{id}/ | 
[**UsersTokensRetrieve**](UsersAPI.md#UsersTokensRetrieve) | **Get** /users/tokens/{id}/ | 
[**UsersTokensUpdate**](UsersAPI.md#UsersTokensUpdate) | **Put** /users/tokens/{id}/ | 
[**UsersUsersBulkDestroy**](UsersAPI.md#UsersUsersBulkDestroy) | **Delete** /users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersAPI.md#UsersUsersBulkPartialUpdate) | **Patch** /users/users/ | 
[**UsersUsersBulkUpdate**](UsersAPI.md#UsersUsersBulkUpdate) | **Put** /users/users/ | 
[**UsersUsersCreate**](UsersAPI.md#UsersUsersCreate) | **Post** /users/users/ | 
[**UsersUsersDestroy**](UsersAPI.md#UsersUsersDestroy) | **Delete** /users/users/{id}/ | 
[**UsersUsersList**](UsersAPI.md#UsersUsersList) | **Get** /users/users/ | 
[**UsersUsersPartialUpdate**](UsersAPI.md#UsersUsersPartialUpdate) | **Patch** /users/users/{id}/ | 
[**UsersUsersRetrieve**](UsersAPI.md#UsersUsersRetrieve) | **Get** /users/users/{id}/ | 
[**UsersUsersUpdate**](UsersAPI.md#UsersUsersUpdate) | **Put** /users/users/{id}/ | 



## UsersConfigRetrieve

> map[string]interface{} UsersConfigRetrieve(ctx).Format(format).Depth(depth).Execute()





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
	resp, r, err := apiClient.UsersAPI.UsersConfigRetrieve(context.Background()).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersConfigRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersConfigRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersConfigRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersConfigRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkDestroy

> UsersGroupsBulkDestroy(ctx).BulkOperationIntegerIDRequest(bulkOperationIntegerIDRequest).Format(format).Execute()





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
	bulkOperationIntegerIDRequest := []openapiclient.BulkOperationIntegerIDRequest{*openapiclient.NewBulkOperationIntegerIDRequest(int32(123))} // []BulkOperationIntegerIDRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersGroupsBulkDestroy(context.Background()).BulkOperationIntegerIDRequest(bulkOperationIntegerIDRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkOperationIntegerIDRequest** | [**[]BulkOperationIntegerIDRequest**](BulkOperationIntegerIDRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkPartialUpdate

> []Group UsersGroupsBulkPartialUpdate(ctx).PatchedBulkWritableGroupRequest(patchedBulkWritableGroupRequest).Format(format).Execute()





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
	patchedBulkWritableGroupRequest := []openapiclient.PatchedBulkWritableGroupRequest{*openapiclient.NewPatchedBulkWritableGroupRequest(int32(123))} // []PatchedBulkWritableGroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsBulkPartialUpdate(context.Background()).PatchedBulkWritableGroupRequest(patchedBulkWritableGroupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsBulkPartialUpdate`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedBulkWritableGroupRequest** | [**[]PatchedBulkWritableGroupRequest**](PatchedBulkWritableGroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkUpdate

> []Group UsersGroupsBulkUpdate(ctx).BulkWritableGroupRequest(bulkWritableGroupRequest).Format(format).Execute()





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
	bulkWritableGroupRequest := []openapiclient.BulkWritableGroupRequest{*openapiclient.NewBulkWritableGroupRequest(int32(123), "Name_example")} // []BulkWritableGroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsBulkUpdate(context.Background()).BulkWritableGroupRequest(bulkWritableGroupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsBulkUpdate`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWritableGroupRequest** | [**[]BulkWritableGroupRequest**](BulkWritableGroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsCreate

> Group UsersGroupsCreate(ctx).GroupRequest(groupRequest).Format(format).Execute()





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
	groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsCreate(context.Background()).GroupRequest(groupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsCreate`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequest** | [**GroupRequest**](GroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsDestroy

> UsersGroupsDestroy(ctx, id).Format(format).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersGroupsDestroy(context.Background(), id).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsList

> PaginatedGroupList UsersGroupsList(ctx).Format(format).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Q(q).Sort(sort).Depth(depth).Execute()





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
	id := []int32{int32(123)} // []int32 | Unique object identifier, either a UUID primary key or a composite key. (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := []string{"Inner_example"} // []string |  (optional)
	nameIc := []string{"Inner_example"} // []string |  (optional)
	nameIe := []string{"Inner_example"} // []string |  (optional)
	nameIew := []string{"Inner_example"} // []string |  (optional)
	nameIre := []string{"Inner_example"} // []string |  (optional)
	nameIsw := []string{"Inner_example"} // []string |  (optional)
	nameN := []string{"Inner_example"} // []string |  (optional)
	nameNic := []string{"Inner_example"} // []string |  (optional)
	nameNie := []string{"Inner_example"} // []string |  (optional)
	nameNiew := []string{"Inner_example"} // []string |  (optional)
	nameNire := []string{"Inner_example"} // []string |  (optional)
	nameNisw := []string{"Inner_example"} // []string |  (optional)
	nameNre := []string{"Inner_example"} // []string |  (optional)
	nameRe := []string{"Inner_example"} // []string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	q := "q_example" // string | Search (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsList(context.Background()).Format(format).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Q(q).Sort(sort).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsList`: PaginatedGroupList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **id** | **[]int32** | Unique object identifier, either a UUID primary key or a composite key. | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **[]string** |  | 
 **nameIc** | **[]string** |  | 
 **nameIe** | **[]string** |  | 
 **nameIew** | **[]string** |  | 
 **nameIre** | **[]string** |  | 
 **nameIsw** | **[]string** |  | 
 **nameN** | **[]string** |  | 
 **nameNic** | **[]string** |  | 
 **nameNie** | **[]string** |  | 
 **nameNiew** | **[]string** |  | 
 **nameNire** | **[]string** |  | 
 **nameNisw** | **[]string** |  | 
 **nameNre** | **[]string** |  | 
 **nameRe** | **[]string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedGroupList**](PaginatedGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsPartialUpdate

> Group UsersGroupsPartialUpdate(ctx, id).Format(format).PatchedGroupRequest(patchedGroupRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	patchedGroupRequest := *openapiclient.NewPatchedGroupRequest() // PatchedGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsPartialUpdate(context.Background(), id).Format(format).PatchedGroupRequest(patchedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsPartialUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **patchedGroupRequest** | [**PatchedGroupRequest**](PatchedGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsRetrieve

> Group UsersGroupsRetrieve(ctx, id).Format(format).Depth(depth).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsRetrieve(context.Background(), id).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsRetrieve`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsUpdate

> Group UsersGroupsUpdate(ctx, id).GroupRequest(groupRequest).Format(format).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group.
	groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsUpdate(context.Background(), id).GroupRequest(groupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRequest** | [**GroupRequest**](GroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkDestroy

> UsersPermissionsBulkDestroy(ctx).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()





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
	bulkOperationRequest := []openapiclient.BulkOperationRequest{*openapiclient.NewBulkOperationRequest("Id_example")} // []BulkOperationRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersPermissionsBulkDestroy(context.Background()).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkOperationRequest** | [**[]BulkOperationRequest**](BulkOperationRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkPartialUpdate

> []ObjectPermission UsersPermissionsBulkPartialUpdate(ctx).PatchedBulkWritableObjectPermissionRequest(patchedBulkWritableObjectPermissionRequest).Format(format).Execute()





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
	patchedBulkWritableObjectPermissionRequest := []openapiclient.PatchedBulkWritableObjectPermissionRequest{*openapiclient.NewPatchedBulkWritableObjectPermissionRequest("Id_example")} // []PatchedBulkWritableObjectPermissionRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsBulkPartialUpdate(context.Background()).PatchedBulkWritableObjectPermissionRequest(patchedBulkWritableObjectPermissionRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsBulkPartialUpdate`: []ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedBulkWritableObjectPermissionRequest** | [**[]PatchedBulkWritableObjectPermissionRequest**](PatchedBulkWritableObjectPermissionRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkUpdate

> []ObjectPermission UsersPermissionsBulkUpdate(ctx).BulkWritableObjectPermissionRequest(bulkWritableObjectPermissionRequest).Format(format).Execute()





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
	bulkWritableObjectPermissionRequest := []openapiclient.BulkWritableObjectPermissionRequest{*openapiclient.NewBulkWritableObjectPermissionRequest("Id_example", []string{"ObjectTypes_example"}, "Name_example", interface{}(123))} // []BulkWritableObjectPermissionRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsBulkUpdate(context.Background()).BulkWritableObjectPermissionRequest(bulkWritableObjectPermissionRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsBulkUpdate`: []ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWritableObjectPermissionRequest** | [**[]BulkWritableObjectPermissionRequest**](BulkWritableObjectPermissionRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsCreate

> ObjectPermission UsersPermissionsCreate(ctx).ObjectPermissionRequest(objectPermissionRequest).Format(format).Execute()





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
	objectPermissionRequest := *openapiclient.NewObjectPermissionRequest([]string{"ObjectTypes_example"}, "Name_example", interface{}(123)) // ObjectPermissionRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsCreate(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsCreate`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectPermissionRequest** | [**ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsDestroy

> UsersPermissionsDestroy(ctx, id).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersPermissionsDestroy(context.Background(), id).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsList

> PaginatedObjectPermissionList UsersPermissionsList(ctx).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Enabled(enabled).Format(format).Groups(groups).GroupsN(groupsN).GroupsId(groupsId).GroupsIdN(groupsIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).Q(q).Sort(sort).Users(users).UsersN(usersN).Depth(depth).Execute()





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
	description := []string{"Inner_example"} // []string |  (optional)
	descriptionIc := []string{"Inner_example"} // []string |  (optional)
	descriptionIe := []string{"Inner_example"} // []string |  (optional)
	descriptionIew := []string{"Inner_example"} // []string |  (optional)
	descriptionIre := []string{"Inner_example"} // []string |  (optional)
	descriptionIsw := []string{"Inner_example"} // []string |  (optional)
	descriptionN := []string{"Inner_example"} // []string |  (optional)
	descriptionNic := []string{"Inner_example"} // []string |  (optional)
	descriptionNie := []string{"Inner_example"} // []string |  (optional)
	descriptionNiew := []string{"Inner_example"} // []string |  (optional)
	descriptionNire := []string{"Inner_example"} // []string |  (optional)
	descriptionNisw := []string{"Inner_example"} // []string |  (optional)
	descriptionNre := []string{"Inner_example"} // []string |  (optional)
	descriptionRe := []string{"Inner_example"} // []string |  (optional)
	enabled := true // bool |  (optional)
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	groups := []string{"Inner_example"} // []string | Group (name) (optional)
	groupsN := []string{"Inner_example"} // []string | Exclude Group (name) (optional)
	groupsId := []int32{int32(123)} // []int32 | Group (ID) (optional)
	groupsIdN := []int32{int32(123)} // []int32 | Exclude Group (ID) (optional)
	id := []string{"Inner_example"} // []string | Unique object identifier, either a UUID primary key or a composite key. (optional)
	idIc := []string{"Inner_example"} // []string |  (optional)
	idIe := []string{"Inner_example"} // []string |  (optional)
	idIew := []string{"Inner_example"} // []string |  (optional)
	idIre := []string{"Inner_example"} // []string |  (optional)
	idIsw := []string{"Inner_example"} // []string |  (optional)
	idN := []string{"Inner_example"} // []string |  (optional)
	idNic := []string{"Inner_example"} // []string |  (optional)
	idNie := []string{"Inner_example"} // []string |  (optional)
	idNiew := []string{"Inner_example"} // []string |  (optional)
	idNire := []string{"Inner_example"} // []string |  (optional)
	idNisw := []string{"Inner_example"} // []string |  (optional)
	idNre := []string{"Inner_example"} // []string |  (optional)
	idRe := []string{"Inner_example"} // []string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := []string{"Inner_example"} // []string |  (optional)
	nameIc := []string{"Inner_example"} // []string |  (optional)
	nameIe := []string{"Inner_example"} // []string |  (optional)
	nameIew := []string{"Inner_example"} // []string |  (optional)
	nameIre := []string{"Inner_example"} // []string |  (optional)
	nameIsw := []string{"Inner_example"} // []string |  (optional)
	nameN := []string{"Inner_example"} // []string |  (optional)
	nameNic := []string{"Inner_example"} // []string |  (optional)
	nameNie := []string{"Inner_example"} // []string |  (optional)
	nameNiew := []string{"Inner_example"} // []string |  (optional)
	nameNire := []string{"Inner_example"} // []string |  (optional)
	nameNisw := []string{"Inner_example"} // []string |  (optional)
	nameNre := []string{"Inner_example"} // []string |  (optional)
	nameRe := []string{"Inner_example"} // []string |  (optional)
	objectTypes := []int32{int32(123)} // []int32 |  (optional)
	objectTypesN := []int32{int32(123)} // []int32 |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	q := "q_example" // string | Search (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	users := []string{"Inner_example"} // []string |  (optional)
	usersN := []string{"Inner_example"} // []string |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsList(context.Background()).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Enabled(enabled).Format(format).Groups(groups).GroupsN(groupsN).GroupsId(groupsId).GroupsIdN(groupsIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).Q(q).Sort(sort).Users(users).UsersN(usersN).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsList`: PaginatedObjectPermissionList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **[]string** |  | 
 **descriptionIc** | **[]string** |  | 
 **descriptionIe** | **[]string** |  | 
 **descriptionIew** | **[]string** |  | 
 **descriptionIre** | **[]string** |  | 
 **descriptionIsw** | **[]string** |  | 
 **descriptionN** | **[]string** |  | 
 **descriptionNic** | **[]string** |  | 
 **descriptionNie** | **[]string** |  | 
 **descriptionNiew** | **[]string** |  | 
 **descriptionNire** | **[]string** |  | 
 **descriptionNisw** | **[]string** |  | 
 **descriptionNre** | **[]string** |  | 
 **descriptionRe** | **[]string** |  | 
 **enabled** | **bool** |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **groups** | **[]string** | Group (name) | 
 **groupsN** | **[]string** | Exclude Group (name) | 
 **groupsId** | **[]int32** | Group (ID) | 
 **groupsIdN** | **[]int32** | Exclude Group (ID) | 
 **id** | **[]string** | Unique object identifier, either a UUID primary key or a composite key. | 
 **idIc** | **[]string** |  | 
 **idIe** | **[]string** |  | 
 **idIew** | **[]string** |  | 
 **idIre** | **[]string** |  | 
 **idIsw** | **[]string** |  | 
 **idN** | **[]string** |  | 
 **idNic** | **[]string** |  | 
 **idNie** | **[]string** |  | 
 **idNiew** | **[]string** |  | 
 **idNire** | **[]string** |  | 
 **idNisw** | **[]string** |  | 
 **idNre** | **[]string** |  | 
 **idRe** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **[]string** |  | 
 **nameIc** | **[]string** |  | 
 **nameIe** | **[]string** |  | 
 **nameIew** | **[]string** |  | 
 **nameIre** | **[]string** |  | 
 **nameIsw** | **[]string** |  | 
 **nameN** | **[]string** |  | 
 **nameNic** | **[]string** |  | 
 **nameNie** | **[]string** |  | 
 **nameNiew** | **[]string** |  | 
 **nameNire** | **[]string** |  | 
 **nameNisw** | **[]string** |  | 
 **nameNre** | **[]string** |  | 
 **nameRe** | **[]string** |  | 
 **objectTypes** | **[]int32** |  | 
 **objectTypesN** | **[]int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **users** | **[]string** |  | 
 **usersN** | **[]string** |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedObjectPermissionList**](PaginatedObjectPermissionList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsPartialUpdate

> ObjectPermission UsersPermissionsPartialUpdate(ctx, id).Format(format).PatchedObjectPermissionRequest(patchedObjectPermissionRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	patchedObjectPermissionRequest := *openapiclient.NewPatchedObjectPermissionRequest() // PatchedObjectPermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsPartialUpdate(context.Background(), id).Format(format).PatchedObjectPermissionRequest(patchedObjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsPartialUpdate`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **patchedObjectPermissionRequest** | [**PatchedObjectPermissionRequest**](PatchedObjectPermissionRequest.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsRetrieve

> ObjectPermission UsersPermissionsRetrieve(ctx, id).Format(format).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsRetrieve(context.Background(), id).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsRetrieve`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsUpdate

> ObjectPermission UsersPermissionsUpdate(ctx, id).ObjectPermissionRequest(objectPermissionRequest).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.
	objectPermissionRequest := *openapiclient.NewObjectPermissionRequest([]string{"ObjectTypes_example"}, "Name_example", interface{}(123)) // ObjectPermissionRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsUpdate(context.Background(), id).ObjectPermissionRequest(objectPermissionRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsUpdate`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectPermissionRequest** | [**ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkDestroy

> UsersTokensBulkDestroy(ctx).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()





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
	bulkOperationRequest := []openapiclient.BulkOperationRequest{*openapiclient.NewBulkOperationRequest("Id_example")} // []BulkOperationRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTokensBulkDestroy(context.Background()).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkOperationRequest** | [**[]BulkOperationRequest**](BulkOperationRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkPartialUpdate

> []Token UsersTokensBulkPartialUpdate(ctx).PatchedBulkWritableTokenRequest(patchedBulkWritableTokenRequest).Format(format).Execute()





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
	patchedBulkWritableTokenRequest := []openapiclient.PatchedBulkWritableTokenRequest{*openapiclient.NewPatchedBulkWritableTokenRequest("Id_example")} // []PatchedBulkWritableTokenRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensBulkPartialUpdate(context.Background()).PatchedBulkWritableTokenRequest(patchedBulkWritableTokenRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensBulkPartialUpdate`: []Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedBulkWritableTokenRequest** | [**[]PatchedBulkWritableTokenRequest**](PatchedBulkWritableTokenRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkUpdate

> []Token UsersTokensBulkUpdate(ctx).BulkWritableTokenRequest(bulkWritableTokenRequest).Format(format).Execute()





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
	bulkWritableTokenRequest := []openapiclient.BulkWritableTokenRequest{*openapiclient.NewBulkWritableTokenRequest("Id_example")} // []BulkWritableTokenRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensBulkUpdate(context.Background()).BulkWritableTokenRequest(bulkWritableTokenRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensBulkUpdate`: []Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWritableTokenRequest** | [**[]BulkWritableTokenRequest**](BulkWritableTokenRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensCreate

> Token UsersTokensCreate(ctx).Format(format).TokenRequest(tokenRequest).Execute()





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
	tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensCreate(context.Background()).Format(format).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensCreate`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensDestroy

> UsersTokensDestroy(ctx, id).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTokensDestroy(context.Background(), id).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensList

> PaginatedTokenList UsersTokensList(ctx).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Expires(expires).ExpiresGt(expiresGt).ExpiresGte(expiresGte).ExpiresIsnull(expiresIsnull).ExpiresLt(expiresLt).ExpiresLte(expiresLte).ExpiresN(expiresN).Format(format).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Key(key).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIre(keyIre).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNire(keyNire).KeyNisw(keyNisw).KeyNre(keyNre).KeyRe(keyRe).Limit(limit).Offset(offset).Q(q).Sort(sort).WriteEnabled(writeEnabled).Depth(depth).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/TobiPeterG/go-nautobot"
)

func main() {
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdN := []time.Time{time.Now()} // []time.Time |  (optional)
	description := []string{"Inner_example"} // []string |  (optional)
	descriptionIc := []string{"Inner_example"} // []string |  (optional)
	descriptionIe := []string{"Inner_example"} // []string |  (optional)
	descriptionIew := []string{"Inner_example"} // []string |  (optional)
	descriptionIre := []string{"Inner_example"} // []string |  (optional)
	descriptionIsw := []string{"Inner_example"} // []string |  (optional)
	descriptionN := []string{"Inner_example"} // []string |  (optional)
	descriptionNic := []string{"Inner_example"} // []string |  (optional)
	descriptionNie := []string{"Inner_example"} // []string |  (optional)
	descriptionNiew := []string{"Inner_example"} // []string |  (optional)
	descriptionNire := []string{"Inner_example"} // []string |  (optional)
	descriptionNisw := []string{"Inner_example"} // []string |  (optional)
	descriptionNre := []string{"Inner_example"} // []string |  (optional)
	descriptionRe := []string{"Inner_example"} // []string |  (optional)
	expires := []time.Time{time.Now()} // []time.Time |  (optional)
	expiresGt := []time.Time{time.Now()} // []time.Time |  (optional)
	expiresGte := []time.Time{time.Now()} // []time.Time |  (optional)
	expiresIsnull := true // bool |  (optional)
	expiresLt := []time.Time{time.Now()} // []time.Time |  (optional)
	expiresLte := []time.Time{time.Now()} // []time.Time |  (optional)
	expiresN := []time.Time{time.Now()} // []time.Time |  (optional)
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	id := []string{"Inner_example"} // []string | Unique object identifier, either a UUID primary key or a composite key. (optional)
	idIc := []string{"Inner_example"} // []string |  (optional)
	idIe := []string{"Inner_example"} // []string |  (optional)
	idIew := []string{"Inner_example"} // []string |  (optional)
	idIre := []string{"Inner_example"} // []string |  (optional)
	idIsw := []string{"Inner_example"} // []string |  (optional)
	idN := []string{"Inner_example"} // []string |  (optional)
	idNic := []string{"Inner_example"} // []string |  (optional)
	idNie := []string{"Inner_example"} // []string |  (optional)
	idNiew := []string{"Inner_example"} // []string |  (optional)
	idNire := []string{"Inner_example"} // []string |  (optional)
	idNisw := []string{"Inner_example"} // []string |  (optional)
	idNre := []string{"Inner_example"} // []string |  (optional)
	idRe := []string{"Inner_example"} // []string |  (optional)
	key := []string{"Inner_example"} // []string |  (optional)
	keyIc := []string{"Inner_example"} // []string |  (optional)
	keyIe := []string{"Inner_example"} // []string |  (optional)
	keyIew := []string{"Inner_example"} // []string |  (optional)
	keyIre := []string{"Inner_example"} // []string |  (optional)
	keyIsw := []string{"Inner_example"} // []string |  (optional)
	keyN := []string{"Inner_example"} // []string |  (optional)
	keyNic := []string{"Inner_example"} // []string |  (optional)
	keyNie := []string{"Inner_example"} // []string |  (optional)
	keyNiew := []string{"Inner_example"} // []string |  (optional)
	keyNire := []string{"Inner_example"} // []string |  (optional)
	keyNisw := []string{"Inner_example"} // []string |  (optional)
	keyNre := []string{"Inner_example"} // []string |  (optional)
	keyRe := []string{"Inner_example"} // []string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	q := "q_example" // string | Search (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	writeEnabled := true // bool |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensList(context.Background()).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Expires(expires).ExpiresGt(expiresGt).ExpiresGte(expiresGte).ExpiresIsnull(expiresIsnull).ExpiresLt(expiresLt).ExpiresLte(expiresLte).ExpiresN(expiresN).Format(format).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Key(key).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIre(keyIre).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNire(keyNire).KeyNisw(keyNisw).KeyNre(keyNre).KeyRe(keyRe).Limit(limit).Offset(offset).Q(q).Sort(sort).WriteEnabled(writeEnabled).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensList`: PaginatedTokenList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdLt** | [**[]time.Time**](time.Time.md) |  | 
 **createdLte** | [**[]time.Time**](time.Time.md) |  | 
 **createdN** | [**[]time.Time**](time.Time.md) |  | 
 **description** | **[]string** |  | 
 **descriptionIc** | **[]string** |  | 
 **descriptionIe** | **[]string** |  | 
 **descriptionIew** | **[]string** |  | 
 **descriptionIre** | **[]string** |  | 
 **descriptionIsw** | **[]string** |  | 
 **descriptionN** | **[]string** |  | 
 **descriptionNic** | **[]string** |  | 
 **descriptionNie** | **[]string** |  | 
 **descriptionNiew** | **[]string** |  | 
 **descriptionNire** | **[]string** |  | 
 **descriptionNisw** | **[]string** |  | 
 **descriptionNre** | **[]string** |  | 
 **descriptionRe** | **[]string** |  | 
 **expires** | [**[]time.Time**](time.Time.md) |  | 
 **expiresGt** | [**[]time.Time**](time.Time.md) |  | 
 **expiresGte** | [**[]time.Time**](time.Time.md) |  | 
 **expiresIsnull** | **bool** |  | 
 **expiresLt** | [**[]time.Time**](time.Time.md) |  | 
 **expiresLte** | [**[]time.Time**](time.Time.md) |  | 
 **expiresN** | [**[]time.Time**](time.Time.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **id** | **[]string** | Unique object identifier, either a UUID primary key or a composite key. | 
 **idIc** | **[]string** |  | 
 **idIe** | **[]string** |  | 
 **idIew** | **[]string** |  | 
 **idIre** | **[]string** |  | 
 **idIsw** | **[]string** |  | 
 **idN** | **[]string** |  | 
 **idNic** | **[]string** |  | 
 **idNie** | **[]string** |  | 
 **idNiew** | **[]string** |  | 
 **idNire** | **[]string** |  | 
 **idNisw** | **[]string** |  | 
 **idNre** | **[]string** |  | 
 **idRe** | **[]string** |  | 
 **key** | **[]string** |  | 
 **keyIc** | **[]string** |  | 
 **keyIe** | **[]string** |  | 
 **keyIew** | **[]string** |  | 
 **keyIre** | **[]string** |  | 
 **keyIsw** | **[]string** |  | 
 **keyN** | **[]string** |  | 
 **keyNic** | **[]string** |  | 
 **keyNie** | **[]string** |  | 
 **keyNiew** | **[]string** |  | 
 **keyNire** | **[]string** |  | 
 **keyNisw** | **[]string** |  | 
 **keyNre** | **[]string** |  | 
 **keyRe** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **writeEnabled** | **bool** |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedTokenList**](PaginatedTokenList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensPartialUpdate

> Token UsersTokensPartialUpdate(ctx, id).Format(format).PatchedTokenRequest(patchedTokenRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	patchedTokenRequest := *openapiclient.NewPatchedTokenRequest() // PatchedTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensPartialUpdate(context.Background(), id).Format(format).PatchedTokenRequest(patchedTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensPartialUpdate`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **patchedTokenRequest** | [**PatchedTokenRequest**](PatchedTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensRetrieve

> Token UsersTokensRetrieve(ctx, id).Format(format).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensRetrieve(context.Background(), id).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensRetrieve`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensUpdate

> Token UsersTokensUpdate(ctx, id).Format(format).TokenRequest(tokenRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensUpdate(context.Background(), id).Format(format).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensUpdate`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkDestroy

> UsersUsersBulkDestroy(ctx).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()





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
	bulkOperationRequest := []openapiclient.BulkOperationRequest{*openapiclient.NewBulkOperationRequest("Id_example")} // []BulkOperationRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUsersBulkDestroy(context.Background()).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkOperationRequest** | [**[]BulkOperationRequest**](BulkOperationRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkPartialUpdate

> []User UsersUsersBulkPartialUpdate(ctx).PatchedBulkWritableUserRequest(patchedBulkWritableUserRequest).Format(format).Execute()





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
	patchedBulkWritableUserRequest := []openapiclient.PatchedBulkWritableUserRequest{*openapiclient.NewPatchedBulkWritableUserRequest("Id_example")} // []PatchedBulkWritableUserRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersBulkPartialUpdate(context.Background()).PatchedBulkWritableUserRequest(patchedBulkWritableUserRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersBulkPartialUpdate`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedBulkWritableUserRequest** | [**[]PatchedBulkWritableUserRequest**](PatchedBulkWritableUserRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkUpdate

> []User UsersUsersBulkUpdate(ctx).BulkWritableUserRequest(bulkWritableUserRequest).Format(format).Execute()





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
	bulkWritableUserRequest := []openapiclient.BulkWritableUserRequest{*openapiclient.NewBulkWritableUserRequest("Id_example", "Username_example")} // []BulkWritableUserRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersBulkUpdate(context.Background()).BulkWritableUserRequest(bulkWritableUserRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersBulkUpdate`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWritableUserRequest** | [**[]BulkWritableUserRequest**](BulkWritableUserRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersCreate

> User UsersUsersCreate(ctx).UserRequest(userRequest).Format(format).Execute()





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
	userRequest := *openapiclient.NewUserRequest("Username_example") // UserRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersCreate(context.Background()).UserRequest(userRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersCreate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**UserRequest**](UserRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersDestroy

> UsersUsersDestroy(ctx, id).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUsersDestroy(context.Background(), id).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersList

> PaginatedUserList UsersUsersList(ctx).Email(email).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIre(emailIre).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNire(emailNire).EmailNisw(emailNisw).EmailNre(emailNre).EmailRe(emailRe).FirstName(firstName).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIre(firstNameIre).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNire(firstNameNire).FirstNameNisw(firstNameNisw).FirstNameNre(firstNameNre).FirstNameRe(firstNameRe).Format(format).Groups(groups).GroupsN(groupsN).GroupsId(groupsId).GroupsIdN(groupsIdN).HasObjectChanges(hasObjectChanges).HasObjectPermissions(hasObjectPermissions).HasRackReservations(hasRackReservations).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIre(lastNameIre).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNire(lastNameNire).LastNameNisw(lastNameNisw).LastNameNre(lastNameNre).LastNameRe(lastNameRe).Limit(limit).ObjectChanges(objectChanges).ObjectChangesIsnull(objectChangesIsnull).ObjectChangesN(objectChangesN).ObjectPermissions(objectPermissions).ObjectPermissionsIsnull(objectPermissionsIsnull).ObjectPermissionsN(objectPermissionsN).Offset(offset).Q(q).RackReservationsId(rackReservationsId).RackReservationsIdIsnull(rackReservationsIdIsnull).RackReservationsIdN(rackReservationsIdN).Sort(sort).Username(username).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIre(usernameIre).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNire(usernameNire).UsernameNisw(usernameNisw).UsernameNre(usernameNre).UsernameRe(usernameRe).Depth(depth).Execute()





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
	email := []string{"Inner_example"} // []string |  (optional)
	emailIc := []string{"Inner_example"} // []string |  (optional)
	emailIe := []string{"Inner_example"} // []string |  (optional)
	emailIew := []string{"Inner_example"} // []string |  (optional)
	emailIre := []string{"Inner_example"} // []string |  (optional)
	emailIsw := []string{"Inner_example"} // []string |  (optional)
	emailN := []string{"Inner_example"} // []string |  (optional)
	emailNic := []string{"Inner_example"} // []string |  (optional)
	emailNie := []string{"Inner_example"} // []string |  (optional)
	emailNiew := []string{"Inner_example"} // []string |  (optional)
	emailNire := []string{"Inner_example"} // []string |  (optional)
	emailNisw := []string{"Inner_example"} // []string |  (optional)
	emailNre := []string{"Inner_example"} // []string |  (optional)
	emailRe := []string{"Inner_example"} // []string |  (optional)
	firstName := []string{"Inner_example"} // []string |  (optional)
	firstNameIc := []string{"Inner_example"} // []string |  (optional)
	firstNameIe := []string{"Inner_example"} // []string |  (optional)
	firstNameIew := []string{"Inner_example"} // []string |  (optional)
	firstNameIre := []string{"Inner_example"} // []string |  (optional)
	firstNameIsw := []string{"Inner_example"} // []string |  (optional)
	firstNameN := []string{"Inner_example"} // []string |  (optional)
	firstNameNic := []string{"Inner_example"} // []string |  (optional)
	firstNameNie := []string{"Inner_example"} // []string |  (optional)
	firstNameNiew := []string{"Inner_example"} // []string |  (optional)
	firstNameNire := []string{"Inner_example"} // []string |  (optional)
	firstNameNisw := []string{"Inner_example"} // []string |  (optional)
	firstNameNre := []string{"Inner_example"} // []string |  (optional)
	firstNameRe := []string{"Inner_example"} // []string |  (optional)
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	groups := []string{"Inner_example"} // []string | Group (name) (optional)
	groupsN := []string{"Inner_example"} // []string | Exclude Group (name) (optional)
	groupsId := []int32{int32(123)} // []int32 | Group (ID) (optional)
	groupsIdN := []int32{int32(123)} // []int32 | Exclude Group (ID) (optional)
	hasObjectChanges := true // bool | Has Changes (optional)
	hasObjectPermissions := true // bool | Has object permissions (optional)
	hasRackReservations := true // bool | Has Rack Reservations (optional)
	id := []string{"Inner_example"} // []string | Unique object identifier, either a UUID primary key or a composite key. (optional)
	idIc := []string{"Inner_example"} // []string |  (optional)
	idIe := []string{"Inner_example"} // []string |  (optional)
	idIew := []string{"Inner_example"} // []string |  (optional)
	idIre := []string{"Inner_example"} // []string |  (optional)
	idIsw := []string{"Inner_example"} // []string |  (optional)
	idN := []string{"Inner_example"} // []string |  (optional)
	idNic := []string{"Inner_example"} // []string |  (optional)
	idNie := []string{"Inner_example"} // []string |  (optional)
	idNiew := []string{"Inner_example"} // []string |  (optional)
	idNire := []string{"Inner_example"} // []string |  (optional)
	idNisw := []string{"Inner_example"} // []string |  (optional)
	idNre := []string{"Inner_example"} // []string |  (optional)
	idRe := []string{"Inner_example"} // []string |  (optional)
	isActive := true // bool |  (optional)
	isStaff := true // bool |  (optional)
	lastName := []string{"Inner_example"} // []string |  (optional)
	lastNameIc := []string{"Inner_example"} // []string |  (optional)
	lastNameIe := []string{"Inner_example"} // []string |  (optional)
	lastNameIew := []string{"Inner_example"} // []string |  (optional)
	lastNameIre := []string{"Inner_example"} // []string |  (optional)
	lastNameIsw := []string{"Inner_example"} // []string |  (optional)
	lastNameN := []string{"Inner_example"} // []string |  (optional)
	lastNameNic := []string{"Inner_example"} // []string |  (optional)
	lastNameNie := []string{"Inner_example"} // []string |  (optional)
	lastNameNiew := []string{"Inner_example"} // []string |  (optional)
	lastNameNire := []string{"Inner_example"} // []string |  (optional)
	lastNameNisw := []string{"Inner_example"} // []string |  (optional)
	lastNameNre := []string{"Inner_example"} // []string |  (optional)
	lastNameRe := []string{"Inner_example"} // []string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	objectChanges := []string{"Inner_example"} // []string | Object Changes (ID) (optional)
	objectChangesIsnull := true // bool | Object Changes (ID) is null (optional)
	objectChangesN := []string{"Inner_example"} // []string | Exclude Object Changes (ID) (optional)
	objectPermissions := []string{"Inner_example"} // []string |  (optional)
	objectPermissionsIsnull := true // bool | Object Permission (ID or name) is null (optional)
	objectPermissionsN := []string{"Inner_example"} // []string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	q := "q_example" // string | Search (optional)
	rackReservationsId := []string{"Inner_example"} // []string | Rack Reservation (ID) (optional)
	rackReservationsIdIsnull := true // bool | Rack Reservation (ID) is null (optional)
	rackReservationsIdN := []string{"Inner_example"} // []string | Exclude Rack Reservation (ID) (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	username := []string{"Inner_example"} // []string |  (optional)
	usernameIc := []string{"Inner_example"} // []string |  (optional)
	usernameIe := []string{"Inner_example"} // []string |  (optional)
	usernameIew := []string{"Inner_example"} // []string |  (optional)
	usernameIre := []string{"Inner_example"} // []string |  (optional)
	usernameIsw := []string{"Inner_example"} // []string |  (optional)
	usernameN := []string{"Inner_example"} // []string |  (optional)
	usernameNic := []string{"Inner_example"} // []string |  (optional)
	usernameNie := []string{"Inner_example"} // []string |  (optional)
	usernameNiew := []string{"Inner_example"} // []string |  (optional)
	usernameNire := []string{"Inner_example"} // []string |  (optional)
	usernameNisw := []string{"Inner_example"} // []string |  (optional)
	usernameNre := []string{"Inner_example"} // []string |  (optional)
	usernameRe := []string{"Inner_example"} // []string |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersList(context.Background()).Email(email).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIre(emailIre).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNire(emailNire).EmailNisw(emailNisw).EmailNre(emailNre).EmailRe(emailRe).FirstName(firstName).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIre(firstNameIre).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNire(firstNameNire).FirstNameNisw(firstNameNisw).FirstNameNre(firstNameNre).FirstNameRe(firstNameRe).Format(format).Groups(groups).GroupsN(groupsN).GroupsId(groupsId).GroupsIdN(groupsIdN).HasObjectChanges(hasObjectChanges).HasObjectPermissions(hasObjectPermissions).HasRackReservations(hasRackReservations).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIre(lastNameIre).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNire(lastNameNire).LastNameNisw(lastNameNisw).LastNameNre(lastNameNre).LastNameRe(lastNameRe).Limit(limit).ObjectChanges(objectChanges).ObjectChangesIsnull(objectChangesIsnull).ObjectChangesN(objectChangesN).ObjectPermissions(objectPermissions).ObjectPermissionsIsnull(objectPermissionsIsnull).ObjectPermissionsN(objectPermissionsN).Offset(offset).Q(q).RackReservationsId(rackReservationsId).RackReservationsIdIsnull(rackReservationsIdIsnull).RackReservationsIdN(rackReservationsIdN).Sort(sort).Username(username).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIre(usernameIre).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNire(usernameNire).UsernameNisw(usernameNisw).UsernameNre(usernameNre).UsernameRe(usernameRe).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersList`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **[]string** |  | 
 **emailIc** | **[]string** |  | 
 **emailIe** | **[]string** |  | 
 **emailIew** | **[]string** |  | 
 **emailIre** | **[]string** |  | 
 **emailIsw** | **[]string** |  | 
 **emailN** | **[]string** |  | 
 **emailNic** | **[]string** |  | 
 **emailNie** | **[]string** |  | 
 **emailNiew** | **[]string** |  | 
 **emailNire** | **[]string** |  | 
 **emailNisw** | **[]string** |  | 
 **emailNre** | **[]string** |  | 
 **emailRe** | **[]string** |  | 
 **firstName** | **[]string** |  | 
 **firstNameIc** | **[]string** |  | 
 **firstNameIe** | **[]string** |  | 
 **firstNameIew** | **[]string** |  | 
 **firstNameIre** | **[]string** |  | 
 **firstNameIsw** | **[]string** |  | 
 **firstNameN** | **[]string** |  | 
 **firstNameNic** | **[]string** |  | 
 **firstNameNie** | **[]string** |  | 
 **firstNameNiew** | **[]string** |  | 
 **firstNameNire** | **[]string** |  | 
 **firstNameNisw** | **[]string** |  | 
 **firstNameNre** | **[]string** |  | 
 **firstNameRe** | **[]string** |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **groups** | **[]string** | Group (name) | 
 **groupsN** | **[]string** | Exclude Group (name) | 
 **groupsId** | **[]int32** | Group (ID) | 
 **groupsIdN** | **[]int32** | Exclude Group (ID) | 
 **hasObjectChanges** | **bool** | Has Changes | 
 **hasObjectPermissions** | **bool** | Has object permissions | 
 **hasRackReservations** | **bool** | Has Rack Reservations | 
 **id** | **[]string** | Unique object identifier, either a UUID primary key or a composite key. | 
 **idIc** | **[]string** |  | 
 **idIe** | **[]string** |  | 
 **idIew** | **[]string** |  | 
 **idIre** | **[]string** |  | 
 **idIsw** | **[]string** |  | 
 **idN** | **[]string** |  | 
 **idNic** | **[]string** |  | 
 **idNie** | **[]string** |  | 
 **idNiew** | **[]string** |  | 
 **idNire** | **[]string** |  | 
 **idNisw** | **[]string** |  | 
 **idNre** | **[]string** |  | 
 **idRe** | **[]string** |  | 
 **isActive** | **bool** |  | 
 **isStaff** | **bool** |  | 
 **lastName** | **[]string** |  | 
 **lastNameIc** | **[]string** |  | 
 **lastNameIe** | **[]string** |  | 
 **lastNameIew** | **[]string** |  | 
 **lastNameIre** | **[]string** |  | 
 **lastNameIsw** | **[]string** |  | 
 **lastNameN** | **[]string** |  | 
 **lastNameNic** | **[]string** |  | 
 **lastNameNie** | **[]string** |  | 
 **lastNameNiew** | **[]string** |  | 
 **lastNameNire** | **[]string** |  | 
 **lastNameNisw** | **[]string** |  | 
 **lastNameNre** | **[]string** |  | 
 **lastNameRe** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **objectChanges** | **[]string** | Object Changes (ID) | 
 **objectChangesIsnull** | **bool** | Object Changes (ID) is null | 
 **objectChangesN** | **[]string** | Exclude Object Changes (ID) | 
 **objectPermissions** | **[]string** |  | 
 **objectPermissionsIsnull** | **bool** | Object Permission (ID or name) is null | 
 **objectPermissionsN** | **[]string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 
 **rackReservationsId** | **[]string** | Rack Reservation (ID) | 
 **rackReservationsIdIsnull** | **bool** | Rack Reservation (ID) is null | 
 **rackReservationsIdN** | **[]string** | Exclude Rack Reservation (ID) | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **username** | **[]string** |  | 
 **usernameIc** | **[]string** |  | 
 **usernameIe** | **[]string** |  | 
 **usernameIew** | **[]string** |  | 
 **usernameIre** | **[]string** |  | 
 **usernameIsw** | **[]string** |  | 
 **usernameN** | **[]string** |  | 
 **usernameNic** | **[]string** |  | 
 **usernameNie** | **[]string** |  | 
 **usernameNiew** | **[]string** |  | 
 **usernameNire** | **[]string** |  | 
 **usernameNisw** | **[]string** |  | 
 **usernameNre** | **[]string** |  | 
 **usernameRe** | **[]string** |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersPartialUpdate

> User UsersUsersPartialUpdate(ctx, id).Format(format).PatchedUserRequest(patchedUserRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	patchedUserRequest := *openapiclient.NewPatchedUserRequest() // PatchedUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersPartialUpdate(context.Background(), id).Format(format).PatchedUserRequest(patchedUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersPartialUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **patchedUserRequest** | [**PatchedUserRequest**](PatchedUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersRetrieve

> User UsersUsersRetrieve(ctx, id).Format(format).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersRetrieve(context.Background(), id).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersRetrieve`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersUpdate

> User UsersUsersUpdate(ctx, id).UserRequest(userRequest).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.
	userRequest := *openapiclient.NewUserRequest("Username_example") // UserRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersUpdate(context.Background(), id).UserRequest(userRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRequest** | [**UserRequest**](UserRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

