# \TenancyAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenancyTenantGroupsBulkDestroy**](TenancyAPI.md#TenancyTenantGroupsBulkDestroy) | **Delete** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkPartialUpdate**](TenancyAPI.md#TenancyTenantGroupsBulkPartialUpdate) | **Patch** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkUpdate**](TenancyAPI.md#TenancyTenantGroupsBulkUpdate) | **Put** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsCreate**](TenancyAPI.md#TenancyTenantGroupsCreate) | **Post** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsDestroy**](TenancyAPI.md#TenancyTenantGroupsDestroy) | **Delete** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsList**](TenancyAPI.md#TenancyTenantGroupsList) | **Get** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsNotesCreate**](TenancyAPI.md#TenancyTenantGroupsNotesCreate) | **Post** /tenancy/tenant-groups/{id}/notes/ | 
[**TenancyTenantGroupsNotesList**](TenancyAPI.md#TenancyTenantGroupsNotesList) | **Get** /tenancy/tenant-groups/{id}/notes/ | 
[**TenancyTenantGroupsPartialUpdate**](TenancyAPI.md#TenancyTenantGroupsPartialUpdate) | **Patch** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsRetrieve**](TenancyAPI.md#TenancyTenantGroupsRetrieve) | **Get** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsUpdate**](TenancyAPI.md#TenancyTenantGroupsUpdate) | **Put** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantsBulkDestroy**](TenancyAPI.md#TenancyTenantsBulkDestroy) | **Delete** /tenancy/tenants/ | 
[**TenancyTenantsBulkPartialUpdate**](TenancyAPI.md#TenancyTenantsBulkPartialUpdate) | **Patch** /tenancy/tenants/ | 
[**TenancyTenantsBulkUpdate**](TenancyAPI.md#TenancyTenantsBulkUpdate) | **Put** /tenancy/tenants/ | 
[**TenancyTenantsCreate**](TenancyAPI.md#TenancyTenantsCreate) | **Post** /tenancy/tenants/ | 
[**TenancyTenantsDestroy**](TenancyAPI.md#TenancyTenantsDestroy) | **Delete** /tenancy/tenants/{id}/ | 
[**TenancyTenantsList**](TenancyAPI.md#TenancyTenantsList) | **Get** /tenancy/tenants/ | 
[**TenancyTenantsNotesCreate**](TenancyAPI.md#TenancyTenantsNotesCreate) | **Post** /tenancy/tenants/{id}/notes/ | 
[**TenancyTenantsNotesList**](TenancyAPI.md#TenancyTenantsNotesList) | **Get** /tenancy/tenants/{id}/notes/ | 
[**TenancyTenantsPartialUpdate**](TenancyAPI.md#TenancyTenantsPartialUpdate) | **Patch** /tenancy/tenants/{id}/ | 
[**TenancyTenantsRetrieve**](TenancyAPI.md#TenancyTenantsRetrieve) | **Get** /tenancy/tenants/{id}/ | 
[**TenancyTenantsUpdate**](TenancyAPI.md#TenancyTenantsUpdate) | **Put** /tenancy/tenants/{id}/ | 



## TenancyTenantGroupsBulkDestroy

> TenancyTenantGroupsBulkDestroy(ctx).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()





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
	r, err := apiClient.TenancyAPI.TenancyTenantGroupsBulkDestroy(context.Background()).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkDestroyRequest struct via the builder pattern


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


## TenancyTenantGroupsBulkPartialUpdate

> []TenantGroup TenancyTenantGroupsBulkPartialUpdate(ctx).PatchedBulkWritableTenantGroupRequest(patchedBulkWritableTenantGroupRequest).Format(format).Execute()





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
	patchedBulkWritableTenantGroupRequest := []openapiclient.PatchedBulkWritableTenantGroupRequest{*openapiclient.NewPatchedBulkWritableTenantGroupRequest("Id_example")} // []PatchedBulkWritableTenantGroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsBulkPartialUpdate(context.Background()).PatchedBulkWritableTenantGroupRequest(patchedBulkWritableTenantGroupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsBulkPartialUpdate`: []TenantGroup
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedBulkWritableTenantGroupRequest** | [**[]PatchedBulkWritableTenantGroupRequest**](PatchedBulkWritableTenantGroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsBulkUpdate

> []TenantGroup TenancyTenantGroupsBulkUpdate(ctx).BulkWritableTenantGroupRequest(bulkWritableTenantGroupRequest).Format(format).Execute()





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
	bulkWritableTenantGroupRequest := []openapiclient.BulkWritableTenantGroupRequest{*openapiclient.NewBulkWritableTenantGroupRequest("Id_example", "Name_example")} // []BulkWritableTenantGroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsBulkUpdate(context.Background()).BulkWritableTenantGroupRequest(bulkWritableTenantGroupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsBulkUpdate`: []TenantGroup
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWritableTenantGroupRequest** | [**[]BulkWritableTenantGroupRequest**](BulkWritableTenantGroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsCreate

> TenantGroup TenancyTenantGroupsCreate(ctx).TenantGroupRequest(tenantGroupRequest).Format(format).Execute()





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
	tenantGroupRequest := *openapiclient.NewTenantGroupRequest("Name_example") // TenantGroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsCreate(context.Background()).TenantGroupRequest(tenantGroupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsCreate`: TenantGroup
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantGroupRequest** | [**TenantGroupRequest**](TenantGroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsDestroy

> TenancyTenantGroupsDestroy(ctx, id).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenancyAPI.TenancyTenantGroupsDestroy(context.Background(), id).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsDestroyRequest struct via the builder pattern


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


## TenancyTenantGroupsList

> PaginatedTenantGroupList TenancyTenantGroupsList(ctx).Children(children).ChildrenIsnull(childrenIsnull).ChildrenN(childrenN).Contacts(contacts).ContactsIsnull(contactsIsnull).ContactsN(contactsN).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedIsnull(createdIsnull).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).DynamicGroups(dynamicGroups).DynamicGroupsN(dynamicGroupsN).Format(format).HasChildren(hasChildren).HasTenants(hasTenants).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).LastUpdated(lastUpdated).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedIsnull(lastUpdatedIsnull).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Parent(parent).ParentIsnull(parentIsnull).ParentN(parentN).Q(q).Sort(sort).Teams(teams).TeamsIsnull(teamsIsnull).TeamsN(teamsN).Tenants(tenants).TenantsIsnull(tenantsIsnull).TenantsN(tenantsN).Depth(depth).Execute()





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
	children := []string{"Inner_example"} // []string |  (optional)
	childrenIsnull := true // bool | Children (name or ID) is null (optional)
	childrenN := []string{"Inner_example"} // []string |  (optional)
	contacts := []string{"Inner_example"} // []string |  (optional)
	contactsIsnull := true // bool | Contacts (name or ID) is null (optional)
	contactsN := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Exclude Contacts (name or ID) (optional)
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdIsnull := true // bool |  (optional)
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
	dynamicGroups := []string{"Inner_example"} // []string |  (optional)
	dynamicGroupsN := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Exclude Dynamic groups (name or ID) (optional)
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	hasChildren := true // bool | Has children (optional)
	hasTenants := true // bool | Has tenants (optional)
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
	lastUpdated := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedIsnull := true // bool |  (optional)
	lastUpdatedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedN := []time.Time{time.Now()} // []time.Time |  (optional)
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
	parent := []string{"Inner_example"} // []string |  (optional)
	parentIsnull := true // bool | Parent tenant group (name or ID) is null (optional)
	parentN := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string | Search (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	teams := []string{"Inner_example"} // []string |  (optional)
	teamsIsnull := true // bool | Teams (name or ID) is null (optional)
	teamsN := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Exclude Teams (name or ID) (optional)
	tenants := []string{"Inner_example"} // []string |  (optional)
	tenantsIsnull := true // bool | Tenants (name or ID) is null (optional)
	tenantsN := []string{"Inner_example"} // []string |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsList(context.Background()).Children(children).ChildrenIsnull(childrenIsnull).ChildrenN(childrenN).Contacts(contacts).ContactsIsnull(contactsIsnull).ContactsN(contactsN).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedIsnull(createdIsnull).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).DynamicGroups(dynamicGroups).DynamicGroupsN(dynamicGroupsN).Format(format).HasChildren(hasChildren).HasTenants(hasTenants).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).LastUpdated(lastUpdated).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedIsnull(lastUpdatedIsnull).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Parent(parent).ParentIsnull(parentIsnull).ParentN(parentN).Q(q).Sort(sort).Teams(teams).TeamsIsnull(teamsIsnull).TeamsN(teamsN).Tenants(tenants).TenantsIsnull(tenantsIsnull).TenantsN(tenantsN).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsList`: PaginatedTenantGroupList
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **children** | **[]string** |  | 
 **childrenIsnull** | **bool** | Children (name or ID) is null | 
 **childrenN** | **[]string** |  | 
 **contacts** | **[]string** |  | 
 **contactsIsnull** | **bool** | Contacts (name or ID) is null | 
 **contactsN** | **string** | Exclude Contacts (name or ID) | 
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdIsnull** | **bool** |  | 
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
 **dynamicGroups** | **[]string** |  | 
 **dynamicGroupsN** | **string** | Exclude Dynamic groups (name or ID) | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **hasChildren** | **bool** | Has children | 
 **hasTenants** | **bool** | Has tenants | 
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
 **lastUpdated** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedIsnull** | **bool** |  | 
 **lastUpdatedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedN** | [**[]time.Time**](time.Time.md) |  | 
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
 **parent** | **[]string** |  | 
 **parentIsnull** | **bool** | Parent tenant group (name or ID) is null | 
 **parentN** | **[]string** |  | 
 **q** | **string** | Search | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **teams** | **[]string** |  | 
 **teamsIsnull** | **bool** | Teams (name or ID) is null | 
 **teamsN** | **string** | Exclude Teams (name or ID) | 
 **tenants** | **[]string** |  | 
 **tenantsIsnull** | **bool** | Tenants (name or ID) is null | 
 **tenantsN** | **[]string** |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedTenantGroupList**](PaginatedTenantGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsNotesCreate

> Note TenancyTenantGroupsNotesCreate(ctx, id).NoteInputRequest(noteInputRequest).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
	noteInputRequest := *openapiclient.NewNoteInputRequest("Note_example") // NoteInputRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsNotesCreate(context.Background(), id).NoteInputRequest(noteInputRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsNotesCreate`: Note
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteInputRequest** | [**NoteInputRequest**](NoteInputRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**Note**](Note.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsNotesList

> PaginatedNoteList TenancyTenantGroupsNotesList(ctx, id).Format(format).Limit(limit).Offset(offset).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsNotesList(context.Background(), id).Format(format).Limit(limit).Offset(offset).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsNotesList`: PaginatedNoteList
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedNoteList**](PaginatedNoteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsPartialUpdate

> TenantGroup TenancyTenantGroupsPartialUpdate(ctx, id).Format(format).PatchedTenantGroupRequest(patchedTenantGroupRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	patchedTenantGroupRequest := *openapiclient.NewPatchedTenantGroupRequest() // PatchedTenantGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsPartialUpdate(context.Background(), id).Format(format).PatchedTenantGroupRequest(patchedTenantGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsPartialUpdate`: TenantGroup
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **patchedTenantGroupRequest** | [**PatchedTenantGroupRequest**](PatchedTenantGroupRequest.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsRetrieve

> TenantGroup TenancyTenantGroupsRetrieve(ctx, id).Format(format).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsRetrieve(context.Background(), id).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsRetrieve`: TenantGroup
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsUpdate

> TenantGroup TenancyTenantGroupsUpdate(ctx, id).TenantGroupRequest(tenantGroupRequest).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
	tenantGroupRequest := *openapiclient.NewTenantGroupRequest("Name_example") // TenantGroupRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantGroupsUpdate(context.Background(), id).TenantGroupRequest(tenantGroupRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantGroupsUpdate`: TenantGroup
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantGroupRequest** | [**TenantGroupRequest**](TenantGroupRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsBulkDestroy

> TenancyTenantsBulkDestroy(ctx).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()





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
	r, err := apiClient.TenancyAPI.TenancyTenantsBulkDestroy(context.Background()).BulkOperationRequest(bulkOperationRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkDestroyRequest struct via the builder pattern


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


## TenancyTenantsBulkPartialUpdate

> []Tenant TenancyTenantsBulkPartialUpdate(ctx).PatchedBulkWritableTenantRequest(patchedBulkWritableTenantRequest).Format(format).Execute()





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
	patchedBulkWritableTenantRequest := []openapiclient.PatchedBulkWritableTenantRequest{*openapiclient.NewPatchedBulkWritableTenantRequest("Id_example")} // []PatchedBulkWritableTenantRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsBulkPartialUpdate(context.Background()).PatchedBulkWritableTenantRequest(patchedBulkWritableTenantRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsBulkPartialUpdate`: []Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedBulkWritableTenantRequest** | [**[]PatchedBulkWritableTenantRequest**](PatchedBulkWritableTenantRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsBulkUpdate

> []Tenant TenancyTenantsBulkUpdate(ctx).BulkWritableTenantRequest(bulkWritableTenantRequest).Format(format).Execute()





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
	bulkWritableTenantRequest := []openapiclient.BulkWritableTenantRequest{*openapiclient.NewBulkWritableTenantRequest("Id_example", "Name_example")} // []BulkWritableTenantRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsBulkUpdate(context.Background()).BulkWritableTenantRequest(bulkWritableTenantRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsBulkUpdate`: []Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWritableTenantRequest** | [**[]BulkWritableTenantRequest**](BulkWritableTenantRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsCreate

> Tenant TenancyTenantsCreate(ctx).TenantRequest(tenantRequest).Format(format).Execute()





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
	tenantRequest := *openapiclient.NewTenantRequest("Name_example") // TenantRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsCreate(context.Background()).TenantRequest(tenantRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsCreate`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantRequest** | [**TenantRequest**](TenantRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsDestroy

> TenancyTenantsDestroy(ctx, id).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenancyAPI.TenancyTenantsDestroy(context.Background(), id).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsDestroyRequest struct via the builder pattern


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


## TenancyTenantsList

> PaginatedTenantList TenancyTenantsList(ctx).Circuits(circuits).CircuitsIsnull(circuitsIsnull).CircuitsN(circuitsN).Clusters(clusters).ClustersIsnull(clustersIsnull).ClustersN(clustersN).Comments(comments).CommentsIc(commentsIc).CommentsIe(commentsIe).CommentsIew(commentsIew).CommentsIre(commentsIre).CommentsIsw(commentsIsw).CommentsN(commentsN).CommentsNic(commentsNic).CommentsNie(commentsNie).CommentsNiew(commentsNiew).CommentsNire(commentsNire).CommentsNisw(commentsNisw).CommentsNre(commentsNre).CommentsRe(commentsRe).Contacts(contacts).ContactsIsnull(contactsIsnull).ContactsN(contactsN).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedIsnull(createdIsnull).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Devices(devices).DevicesIsnull(devicesIsnull).DevicesN(devicesN).DynamicGroups(dynamicGroups).DynamicGroupsN(dynamicGroupsN).Format(format).HasCircuits(hasCircuits).HasClusters(hasClusters).HasDevices(hasDevices).HasIpAddresses(hasIpAddresses).HasLocations(hasLocations).HasPrefixes(hasPrefixes).HasRackReservations(hasRackReservations).HasRacks(hasRacks).HasRouteTargets(hasRouteTargets).HasVirtualMachines(hasVirtualMachines).HasVlans(hasVlans).HasVrfs(hasVrfs).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).IpAddresses(ipAddresses).IpAddressesIsnull(ipAddressesIsnull).IpAddressesN(ipAddressesN).LastUpdated(lastUpdated).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedIsnull(lastUpdatedIsnull).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).Locations(locations).LocationsIsnull(locationsIsnull).LocationsN(locationsN).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Prefixes(prefixes).PrefixesIsnull(prefixesIsnull).PrefixesN(prefixesN).Q(q).RackReservations(rackReservations).RackReservationsIsnull(rackReservationsIsnull).RackReservationsN(rackReservationsN).Racks(racks).RacksIsnull(racksIsnull).RacksN(racksN).RouteTargets(routeTargets).RouteTargetsIsnull(routeTargetsIsnull).RouteTargetsN(routeTargetsN).Sort(sort).Tags(tags).TagsIsnull(tagsIsnull).TagsN(tagsN).Teams(teams).TeamsIsnull(teamsIsnull).TeamsN(teamsN).TenantGroup(tenantGroup).TenantGroupIsnull(tenantGroupIsnull).TenantGroupN(tenantGroupN).VirtualMachines(virtualMachines).VirtualMachinesIsnull(virtualMachinesIsnull).VirtualMachinesN(virtualMachinesN).Vlans(vlans).VlansIsnull(vlansIsnull).VlansN(vlansN).Vrfs(vrfs).VrfsIsnull(vrfsIsnull).VrfsN(vrfsN).Depth(depth).Execute()





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
	circuits := []string{"Inner_example"} // []string | Circuits (ID) (optional)
	circuitsIsnull := true // bool | Circuits (ID) is null (optional)
	circuitsN := []string{"Inner_example"} // []string | Exclude Circuits (ID) (optional)
	clusters := []string{"Inner_example"} // []string |  (optional)
	clustersIsnull := true // bool | Clusters (name or ID) is null (optional)
	clustersN := []string{"Inner_example"} // []string |  (optional)
	comments := []string{"Inner_example"} // []string |  (optional)
	commentsIc := []string{"Inner_example"} // []string |  (optional)
	commentsIe := []string{"Inner_example"} // []string |  (optional)
	commentsIew := []string{"Inner_example"} // []string |  (optional)
	commentsIre := []string{"Inner_example"} // []string |  (optional)
	commentsIsw := []string{"Inner_example"} // []string |  (optional)
	commentsN := []string{"Inner_example"} // []string |  (optional)
	commentsNic := []string{"Inner_example"} // []string |  (optional)
	commentsNie := []string{"Inner_example"} // []string |  (optional)
	commentsNiew := []string{"Inner_example"} // []string |  (optional)
	commentsNire := []string{"Inner_example"} // []string |  (optional)
	commentsNisw := []string{"Inner_example"} // []string |  (optional)
	commentsNre := []string{"Inner_example"} // []string |  (optional)
	commentsRe := []string{"Inner_example"} // []string |  (optional)
	contacts := []string{"Inner_example"} // []string |  (optional)
	contactsIsnull := true // bool | Contacts (name or ID) is null (optional)
	contactsN := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Exclude Contacts (name or ID) (optional)
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdIsnull := true // bool |  (optional)
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
	devices := []string{"Inner_example"} // []string |  (optional)
	devicesIsnull := true // bool | Devices (name or ID) is null (optional)
	devicesN := []string{"Inner_example"} // []string |  (optional)
	dynamicGroups := []string{"Inner_example"} // []string |  (optional)
	dynamicGroupsN := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Exclude Dynamic groups (name or ID) (optional)
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	hasCircuits := true // bool | Has circuits (optional)
	hasClusters := true // bool | Has clusters (optional)
	hasDevices := true // bool | Has devices (optional)
	hasIpAddresses := true // bool | Has IP addresses (optional)
	hasLocations := true // bool | Has locations (optional)
	hasPrefixes := true // bool | Has prefixes (optional)
	hasRackReservations := true // bool | Has rack reservations (optional)
	hasRacks := true // bool | Has racks (optional)
	hasRouteTargets := true // bool | Has route targets (optional)
	hasVirtualMachines := true // bool | Has virtual machines (optional)
	hasVlans := true // bool | Has VLANs (optional)
	hasVrfs := true // bool | Has VRFs (optional)
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
	ipAddresses := []string{"Inner_example"} // []string | IP addresses (ID) (optional)
	ipAddressesIsnull := true // bool | IP addresses (ID) is null (optional)
	ipAddressesN := []string{"Inner_example"} // []string | Exclude IP addresses (ID) (optional)
	lastUpdated := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedIsnull := true // bool |  (optional)
	lastUpdatedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedN := []time.Time{time.Now()} // []time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	locations := []string{"Inner_example"} // []string |  (optional)
	locationsIsnull := true // bool | Locations (names and/or IDs) is null (optional)
	locationsN := []string{"Inner_example"} // []string |  (optional)
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
	prefixes := []string{"Inner_example"} // []string | Prefixes (ID) (optional)
	prefixesIsnull := true // bool | Prefixes (ID) is null (optional)
	prefixesN := []string{"Inner_example"} // []string | Exclude Prefixes (ID) (optional)
	q := "q_example" // string | Search (optional)
	rackReservations := []string{"Inner_example"} // []string | Rack reservations (ID) (optional)
	rackReservationsIsnull := true // bool | Rack reservations (ID) is null (optional)
	rackReservationsN := []string{"Inner_example"} // []string | Exclude Rack reservations (ID) (optional)
	racks := []string{"Inner_example"} // []string |  (optional)
	racksIsnull := true // bool | Racks (name or ID) is null (optional)
	racksN := []string{"Inner_example"} // []string |  (optional)
	routeTargets := []string{"Inner_example"} // []string |  (optional)
	routeTargetsIsnull := true // bool | Route targets (name or ID) is null (optional)
	routeTargetsN := []string{"Inner_example"} // []string |  (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	tagsIsnull := true // bool |  (optional)
	tagsN := []string{"Inner_example"} // []string |  (optional)
	teams := []string{"Inner_example"} // []string |  (optional)
	teamsIsnull := true // bool | Teams (name or ID) is null (optional)
	teamsN := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Exclude Teams (name or ID) (optional)
	tenantGroup := []string{"Inner_example"} // []string |  (optional)
	tenantGroupIsnull := true // bool | Tenant group (name or ID) is null (optional)
	tenantGroupN := []string{"Inner_example"} // []string |  (optional)
	virtualMachines := []string{"Inner_example"} // []string |  (optional)
	virtualMachinesIsnull := true // bool | Virtual machines (name or ID) is null (optional)
	virtualMachinesN := []string{"Inner_example"} // []string |  (optional)
	vlans := []string{"Inner_example"} // []string | VLANs (ID) (optional)
	vlansIsnull := true // bool | VLANs (ID) is null (optional)
	vlansN := []string{"Inner_example"} // []string | Exclude VLANs (ID) (optional)
	vrfs := []string{"Inner_example"} // []string |  (optional)
	vrfsIsnull := true // bool | VRFs (name or ID) is null (optional)
	vrfsN := []string{"Inner_example"} // []string |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsList(context.Background()).Circuits(circuits).CircuitsIsnull(circuitsIsnull).CircuitsN(circuitsN).Clusters(clusters).ClustersIsnull(clustersIsnull).ClustersN(clustersN).Comments(comments).CommentsIc(commentsIc).CommentsIe(commentsIe).CommentsIew(commentsIew).CommentsIre(commentsIre).CommentsIsw(commentsIsw).CommentsN(commentsN).CommentsNic(commentsNic).CommentsNie(commentsNie).CommentsNiew(commentsNiew).CommentsNire(commentsNire).CommentsNisw(commentsNisw).CommentsNre(commentsNre).CommentsRe(commentsRe).Contacts(contacts).ContactsIsnull(contactsIsnull).ContactsN(contactsN).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedIsnull(createdIsnull).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Devices(devices).DevicesIsnull(devicesIsnull).DevicesN(devicesN).DynamicGroups(dynamicGroups).DynamicGroupsN(dynamicGroupsN).Format(format).HasCircuits(hasCircuits).HasClusters(hasClusters).HasDevices(hasDevices).HasIpAddresses(hasIpAddresses).HasLocations(hasLocations).HasPrefixes(hasPrefixes).HasRackReservations(hasRackReservations).HasRacks(hasRacks).HasRouteTargets(hasRouteTargets).HasVirtualMachines(hasVirtualMachines).HasVlans(hasVlans).HasVrfs(hasVrfs).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).IpAddresses(ipAddresses).IpAddressesIsnull(ipAddressesIsnull).IpAddressesN(ipAddressesN).LastUpdated(lastUpdated).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedIsnull(lastUpdatedIsnull).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).Locations(locations).LocationsIsnull(locationsIsnull).LocationsN(locationsN).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Prefixes(prefixes).PrefixesIsnull(prefixesIsnull).PrefixesN(prefixesN).Q(q).RackReservations(rackReservations).RackReservationsIsnull(rackReservationsIsnull).RackReservationsN(rackReservationsN).Racks(racks).RacksIsnull(racksIsnull).RacksN(racksN).RouteTargets(routeTargets).RouteTargetsIsnull(routeTargetsIsnull).RouteTargetsN(routeTargetsN).Sort(sort).Tags(tags).TagsIsnull(tagsIsnull).TagsN(tagsN).Teams(teams).TeamsIsnull(teamsIsnull).TeamsN(teamsN).TenantGroup(tenantGroup).TenantGroupIsnull(tenantGroupIsnull).TenantGroupN(tenantGroupN).VirtualMachines(virtualMachines).VirtualMachinesIsnull(virtualMachinesIsnull).VirtualMachinesN(virtualMachinesN).Vlans(vlans).VlansIsnull(vlansIsnull).VlansN(vlansN).Vrfs(vrfs).VrfsIsnull(vrfsIsnull).VrfsN(vrfsN).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsList`: PaginatedTenantList
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **circuits** | **[]string** | Circuits (ID) | 
 **circuitsIsnull** | **bool** | Circuits (ID) is null | 
 **circuitsN** | **[]string** | Exclude Circuits (ID) | 
 **clusters** | **[]string** |  | 
 **clustersIsnull** | **bool** | Clusters (name or ID) is null | 
 **clustersN** | **[]string** |  | 
 **comments** | **[]string** |  | 
 **commentsIc** | **[]string** |  | 
 **commentsIe** | **[]string** |  | 
 **commentsIew** | **[]string** |  | 
 **commentsIre** | **[]string** |  | 
 **commentsIsw** | **[]string** |  | 
 **commentsN** | **[]string** |  | 
 **commentsNic** | **[]string** |  | 
 **commentsNie** | **[]string** |  | 
 **commentsNiew** | **[]string** |  | 
 **commentsNire** | **[]string** |  | 
 **commentsNisw** | **[]string** |  | 
 **commentsNre** | **[]string** |  | 
 **commentsRe** | **[]string** |  | 
 **contacts** | **[]string** |  | 
 **contactsIsnull** | **bool** | Contacts (name or ID) is null | 
 **contactsN** | **string** | Exclude Contacts (name or ID) | 
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdIsnull** | **bool** |  | 
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
 **devices** | **[]string** |  | 
 **devicesIsnull** | **bool** | Devices (name or ID) is null | 
 **devicesN** | **[]string** |  | 
 **dynamicGroups** | **[]string** |  | 
 **dynamicGroupsN** | **string** | Exclude Dynamic groups (name or ID) | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **hasCircuits** | **bool** | Has circuits | 
 **hasClusters** | **bool** | Has clusters | 
 **hasDevices** | **bool** | Has devices | 
 **hasIpAddresses** | **bool** | Has IP addresses | 
 **hasLocations** | **bool** | Has locations | 
 **hasPrefixes** | **bool** | Has prefixes | 
 **hasRackReservations** | **bool** | Has rack reservations | 
 **hasRacks** | **bool** | Has racks | 
 **hasRouteTargets** | **bool** | Has route targets | 
 **hasVirtualMachines** | **bool** | Has virtual machines | 
 **hasVlans** | **bool** | Has VLANs | 
 **hasVrfs** | **bool** | Has VRFs | 
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
 **ipAddresses** | **[]string** | IP addresses (ID) | 
 **ipAddressesIsnull** | **bool** | IP addresses (ID) is null | 
 **ipAddressesN** | **[]string** | Exclude IP addresses (ID) | 
 **lastUpdated** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedIsnull** | **bool** |  | 
 **lastUpdatedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedN** | [**[]time.Time**](time.Time.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **locations** | **[]string** |  | 
 **locationsIsnull** | **bool** | Locations (names and/or IDs) is null | 
 **locationsN** | **[]string** |  | 
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
 **prefixes** | **[]string** | Prefixes (ID) | 
 **prefixesIsnull** | **bool** | Prefixes (ID) is null | 
 **prefixesN** | **[]string** | Exclude Prefixes (ID) | 
 **q** | **string** | Search | 
 **rackReservations** | **[]string** | Rack reservations (ID) | 
 **rackReservationsIsnull** | **bool** | Rack reservations (ID) is null | 
 **rackReservationsN** | **[]string** | Exclude Rack reservations (ID) | 
 **racks** | **[]string** |  | 
 **racksIsnull** | **bool** | Racks (name or ID) is null | 
 **racksN** | **[]string** |  | 
 **routeTargets** | **[]string** |  | 
 **routeTargetsIsnull** | **bool** | Route targets (name or ID) is null | 
 **routeTargetsN** | **[]string** |  | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **tags** | **[]string** |  | 
 **tagsIsnull** | **bool** |  | 
 **tagsN** | **[]string** |  | 
 **teams** | **[]string** |  | 
 **teamsIsnull** | **bool** | Teams (name or ID) is null | 
 **teamsN** | **string** | Exclude Teams (name or ID) | 
 **tenantGroup** | **[]string** |  | 
 **tenantGroupIsnull** | **bool** | Tenant group (name or ID) is null | 
 **tenantGroupN** | **[]string** |  | 
 **virtualMachines** | **[]string** |  | 
 **virtualMachinesIsnull** | **bool** | Virtual machines (name or ID) is null | 
 **virtualMachinesN** | **[]string** |  | 
 **vlans** | **[]string** | VLANs (ID) | 
 **vlansIsnull** | **bool** | VLANs (ID) is null | 
 **vlansN** | **[]string** | Exclude VLANs (ID) | 
 **vrfs** | **[]string** |  | 
 **vrfsIsnull** | **bool** | VRFs (name or ID) is null | 
 **vrfsN** | **[]string** |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedTenantList**](PaginatedTenantList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsNotesCreate

> Note TenancyTenantsNotesCreate(ctx, id).NoteInputRequest(noteInputRequest).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
	noteInputRequest := *openapiclient.NewNoteInputRequest("Note_example") // NoteInputRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsNotesCreate(context.Background(), id).NoteInputRequest(noteInputRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsNotesCreate`: Note
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteInputRequest** | [**NoteInputRequest**](NoteInputRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**Note**](Note.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsNotesList

> PaginatedNoteList TenancyTenantsNotesList(ctx, id).Format(format).Limit(limit).Offset(offset).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsNotesList(context.Background(), id).Format(format).Limit(limit).Offset(offset).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsNotesList`: PaginatedNoteList
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**PaginatedNoteList**](PaginatedNoteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsPartialUpdate

> Tenant TenancyTenantsPartialUpdate(ctx, id).Format(format).PatchedTenantRequest(patchedTenantRequest).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	patchedTenantRequest := *openapiclient.NewPatchedTenantRequest() // PatchedTenantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsPartialUpdate(context.Background(), id).Format(format).PatchedTenantRequest(patchedTenantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsPartialUpdate`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **patchedTenantRequest** | [**PatchedTenantRequest**](PatchedTenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsRetrieve

> Tenant TenancyTenantsRetrieve(ctx, id).Format(format).Depth(depth).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)
	depth := int32(56) // int32 | Serializer Depth (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsRetrieve(context.Background(), id).Format(format).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsRetrieve`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 
 **depth** | **int32** | Serializer Depth | [default to 1]

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsUpdate

> Tenant TenancyTenantsUpdate(ctx, id).TenantRequest(tenantRequest).Format(format).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
	tenantRequest := *openapiclient.NewTenantRequest("Name_example") // TenantRequest | 
	format := openapiclient.circuits_circuit_terminations_list_format_parameter("csv") // CircuitsCircuitTerminationsListFormatParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenancyTenantsUpdate(context.Background(), id).TenantRequest(tenantRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenancyTenantsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenancyTenantsUpdate`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenancyTenantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantRequest** | [**TenantRequest**](TenantRequest.md) |  | 
 **format** | [**CircuitsCircuitTerminationsListFormatParameter**](CircuitsCircuitTerminationsListFormatParameter.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

