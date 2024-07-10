# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenancyTenantGroupsBulkDestroy**](TenancyApi.md#TenancyTenantGroupsBulkDestroy) | **Delete** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkPartialUpdate**](TenancyApi.md#TenancyTenantGroupsBulkPartialUpdate) | **Patch** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkUpdate**](TenancyApi.md#TenancyTenantGroupsBulkUpdate) | **Put** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsCreate**](TenancyApi.md#TenancyTenantGroupsCreate) | **Post** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsDestroy**](TenancyApi.md#TenancyTenantGroupsDestroy) | **Delete** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsList**](TenancyApi.md#TenancyTenantGroupsList) | **Get** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsNotesCreate**](TenancyApi.md#TenancyTenantGroupsNotesCreate) | **Post** /tenancy/tenant-groups/{id}/notes/ | 
[**TenancyTenantGroupsNotesList**](TenancyApi.md#TenancyTenantGroupsNotesList) | **Get** /tenancy/tenant-groups/{id}/notes/ | 
[**TenancyTenantGroupsPartialUpdate**](TenancyApi.md#TenancyTenantGroupsPartialUpdate) | **Patch** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsRetrieve**](TenancyApi.md#TenancyTenantGroupsRetrieve) | **Get** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsUpdate**](TenancyApi.md#TenancyTenantGroupsUpdate) | **Put** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantsBulkDestroy**](TenancyApi.md#TenancyTenantsBulkDestroy) | **Delete** /tenancy/tenants/ | 
[**TenancyTenantsBulkPartialUpdate**](TenancyApi.md#TenancyTenantsBulkPartialUpdate) | **Patch** /tenancy/tenants/ | 
[**TenancyTenantsBulkUpdate**](TenancyApi.md#TenancyTenantsBulkUpdate) | **Put** /tenancy/tenants/ | 
[**TenancyTenantsCreate**](TenancyApi.md#TenancyTenantsCreate) | **Post** /tenancy/tenants/ | 
[**TenancyTenantsDestroy**](TenancyApi.md#TenancyTenantsDestroy) | **Delete** /tenancy/tenants/{id}/ | 
[**TenancyTenantsList**](TenancyApi.md#TenancyTenantsList) | **Get** /tenancy/tenants/ | 
[**TenancyTenantsNotesCreate**](TenancyApi.md#TenancyTenantsNotesCreate) | **Post** /tenancy/tenants/{id}/notes/ | 
[**TenancyTenantsNotesList**](TenancyApi.md#TenancyTenantsNotesList) | **Get** /tenancy/tenants/{id}/notes/ | 
[**TenancyTenantsPartialUpdate**](TenancyApi.md#TenancyTenantsPartialUpdate) | **Patch** /tenancy/tenants/{id}/ | 
[**TenancyTenantsRetrieve**](TenancyApi.md#TenancyTenantsRetrieve) | **Get** /tenancy/tenants/{id}/ | 
[**TenancyTenantsUpdate**](TenancyApi.md#TenancyTenantsUpdate) | **Put** /tenancy/tenants/{id}/ | 

# **TenancyTenantGroupsBulkDestroy**
> TenancyTenantGroupsBulkDestroy(ctx, body, optional)


Destroy a list of tenant group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkOperationRequest**](BulkOperationRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantGroupsBulkDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsBulkDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsBulkPartialUpdate**
> []TenantGroup TenancyTenantGroupsBulkPartialUpdate(ctx, body, optional)


Partial update a list of tenant group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchedBulkWritableTenantGroupRequest**](PatchedBulkWritableTenantGroupRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantGroupsBulkPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsBulkPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsBulkUpdate**
> []TenantGroup TenancyTenantGroupsBulkUpdate(ctx, body, optional)


Update a list of tenant group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkWritableTenantGroupRequest**](BulkWritableTenantGroupRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantGroupsBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsBulkUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsCreate**
> TenantGroup TenancyTenantGroupsCreate(ctx, body, optional)


Create one or more tenant group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TenantGroupRequest**](TenantGroupRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsDestroy**
> TenancyTenantGroupsDestroy(ctx, id, optional)


Destroy a tenant group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant group. | 
 **optional** | ***TenancyApiTenancyTenantGroupsDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsList**
> PaginatedTenantGroupList TenancyTenantGroupsList(ctx, optional)


Retrieve a list of tenant group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenancyApiTenancyTenantGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **children** | [**optional.Interface of []string**](string.md)|  | 
 **childrenIsnull** | **optional.Bool**|  | 
 **childrenN** | [**optional.Interface of []string**](string.md)|  | 
 **created** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdGt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdGte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdIsnull** | **optional.Bool**|  | 
 **createdLt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdLte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdN** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **description** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIc** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIe** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIew** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIre** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIsw** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionN** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNic** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNie** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNiew** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNire** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNisw** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNre** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionRe** | [**optional.Interface of []string**](string.md)|  | 
 **format** | **optional.String**|  | 
 **hasChildren** | **optional.Bool**| Has children | 
 **hasTenants** | **optional.Bool**| Has tenants | 
 **id** | [**optional.Interface of []string**](string.md)| Unique object identifier, either a UUID primary key or a composite key. | 
 **idIc** | [**optional.Interface of []string**](string.md)|  | 
 **idIe** | [**optional.Interface of []string**](string.md)|  | 
 **idIew** | [**optional.Interface of []string**](string.md)|  | 
 **idIre** | [**optional.Interface of []string**](string.md)|  | 
 **idIsw** | [**optional.Interface of []string**](string.md)|  | 
 **idN** | [**optional.Interface of []string**](string.md)|  | 
 **idNic** | [**optional.Interface of []string**](string.md)|  | 
 **idNie** | [**optional.Interface of []string**](string.md)|  | 
 **idNiew** | [**optional.Interface of []string**](string.md)|  | 
 **idNire** | [**optional.Interface of []string**](string.md)|  | 
 **idNisw** | [**optional.Interface of []string**](string.md)|  | 
 **idNre** | [**optional.Interface of []string**](string.md)|  | 
 **idRe** | [**optional.Interface of []string**](string.md)|  | 
 **lastUpdated** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedGt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedGte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedIsnull** | **optional.Bool**|  | 
 **lastUpdatedLt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedLte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedN** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **name** | [**optional.Interface of []string**](string.md)|  | 
 **nameIc** | [**optional.Interface of []string**](string.md)|  | 
 **nameIe** | [**optional.Interface of []string**](string.md)|  | 
 **nameIew** | [**optional.Interface of []string**](string.md)|  | 
 **nameIre** | [**optional.Interface of []string**](string.md)|  | 
 **nameIsw** | [**optional.Interface of []string**](string.md)|  | 
 **nameN** | [**optional.Interface of []string**](string.md)|  | 
 **nameNic** | [**optional.Interface of []string**](string.md)|  | 
 **nameNie** | [**optional.Interface of []string**](string.md)|  | 
 **nameNiew** | [**optional.Interface of []string**](string.md)|  | 
 **nameNire** | [**optional.Interface of []string**](string.md)|  | 
 **nameNisw** | [**optional.Interface of []string**](string.md)|  | 
 **nameNre** | [**optional.Interface of []string**](string.md)|  | 
 **nameRe** | [**optional.Interface of []string**](string.md)|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **parent** | [**optional.Interface of []string**](string.md)|  | 
 **parentIsnull** | **optional.Bool**|  | 
 **parentN** | [**optional.Interface of []string**](string.md)|  | 
 **q** | **optional.String**| Search | 
 **sort** | **optional.String**| Which field to use when ordering the results. | 
 **tenants** | [**optional.Interface of []string**](string.md)|  | 
 **tenantsIsnull** | **optional.Bool**|  | 
 **tenantsN** | [**optional.Interface of []string**](string.md)|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedTenantGroupList**](PaginatedTenantGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsNotesCreate**
> Note TenancyTenantGroupsNotesCreate(ctx, body, id, optional)


API methods for returning or creating notes on an object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NoteInputRequest**](NoteInputRequest.md)|  | 
  **id** | [**string**](.md)| A UUID string identifying this tenant group. | 
 **optional** | ***TenancyApiTenancyTenantGroupsNotesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsNotesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**Note**](Note.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsNotesList**
> PaginatedNoteList TenancyTenantGroupsNotesList(ctx, id, optional)


API methods for returning or creating notes on an object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant group. | 
 **optional** | ***TenancyApiTenancyTenantGroupsNotesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsNotesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedNoteList**](PaginatedNoteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsPartialUpdate**
> TenantGroup TenancyTenantGroupsPartialUpdate(ctx, id, optional)


Partial update a tenant group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant group. | 
 **optional** | ***TenancyApiTenancyTenantGroupsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedTenantGroupRequest**](PatchedTenantGroupRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsRetrieve**
> TenantGroup TenancyTenantGroupsRetrieve(ctx, id, optional)


Retrieve a tenant group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant group. | 
 **optional** | ***TenancyApiTenancyTenantGroupsRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsUpdate**
> TenantGroup TenancyTenantGroupsUpdate(ctx, body, id, optional)


Update a tenant group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TenantGroupRequest**](TenantGroupRequest.md)|  | 
  **id** | [**string**](.md)| A UUID string identifying this tenant group. | 
 **optional** | ***TenancyApiTenancyTenantGroupsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsBulkDestroy**
> TenancyTenantsBulkDestroy(ctx, body, optional)


Destroy a list of tenant objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkOperationRequest**](BulkOperationRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantsBulkDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsBulkDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsBulkPartialUpdate**
> []Tenant TenancyTenantsBulkPartialUpdate(ctx, body, optional)


Partial update a list of tenant objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchedBulkWritableTenantRequest**](PatchedBulkWritableTenantRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantsBulkPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsBulkPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsBulkUpdate**
> []Tenant TenancyTenantsBulkUpdate(ctx, body, optional)


Update a list of tenant objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkWritableTenantRequest**](BulkWritableTenantRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantsBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsBulkUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsCreate**
> Tenant TenancyTenantsCreate(ctx, body, optional)


Create one or more tenant objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TenantRequest**](TenantRequest.md)|  | 
 **optional** | ***TenancyApiTenancyTenantsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsDestroy**
> TenancyTenantsDestroy(ctx, id, optional)


Destroy a tenant object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant. | 
 **optional** | ***TenancyApiTenancyTenantsDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsList**
> PaginatedTenantList TenancyTenantsList(ctx, optional)


Retrieve a list of tenant objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenancyApiTenancyTenantsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **circuits** | [**optional.Interface of []string**](string.md)| Circuits (ID) | 
 **circuitsIsnull** | **optional.Bool**|  | 
 **circuitsN** | [**optional.Interface of []string**](string.md)| Circuits (ID) | 
 **clusters** | [**optional.Interface of []string**](string.md)|  | 
 **clustersIsnull** | **optional.Bool**|  | 
 **clustersN** | [**optional.Interface of []string**](string.md)|  | 
 **comments** | [**optional.Interface of []string**](string.md)|  | 
 **commentsIc** | [**optional.Interface of []string**](string.md)|  | 
 **commentsIe** | [**optional.Interface of []string**](string.md)|  | 
 **commentsIew** | [**optional.Interface of []string**](string.md)|  | 
 **commentsIre** | [**optional.Interface of []string**](string.md)|  | 
 **commentsIsw** | [**optional.Interface of []string**](string.md)|  | 
 **commentsN** | [**optional.Interface of []string**](string.md)|  | 
 **commentsNic** | [**optional.Interface of []string**](string.md)|  | 
 **commentsNie** | [**optional.Interface of []string**](string.md)|  | 
 **commentsNiew** | [**optional.Interface of []string**](string.md)|  | 
 **commentsNire** | [**optional.Interface of []string**](string.md)|  | 
 **commentsNisw** | [**optional.Interface of []string**](string.md)|  | 
 **commentsNre** | [**optional.Interface of []string**](string.md)|  | 
 **commentsRe** | [**optional.Interface of []string**](string.md)|  | 
 **created** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdGt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdGte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdIsnull** | **optional.Bool**|  | 
 **createdLt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdLte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdN** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **description** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIc** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIe** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIew** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIre** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionIsw** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionN** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNic** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNie** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNiew** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNire** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNisw** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionNre** | [**optional.Interface of []string**](string.md)|  | 
 **descriptionRe** | [**optional.Interface of []string**](string.md)|  | 
 **devices** | [**optional.Interface of []string**](string.md)|  | 
 **devicesIsnull** | **optional.Bool**|  | 
 **devicesN** | [**optional.Interface of []string**](string.md)|  | 
 **format** | **optional.String**|  | 
 **hasCircuits** | **optional.Bool**| Has circuits | 
 **hasClusters** | **optional.Bool**| Has clusters | 
 **hasDevices** | **optional.Bool**| Has devices | 
 **hasIpAddresses** | **optional.Bool**| Has IP addresses | 
 **hasLocations** | **optional.Bool**| Has locations | 
 **hasPrefixes** | **optional.Bool**| Has prefixes | 
 **hasRackReservations** | **optional.Bool**| Has rack reservations | 
 **hasRacks** | **optional.Bool**| Has racks | 
 **hasRouteTargets** | **optional.Bool**| Has route targets | 
 **hasVirtualMachines** | **optional.Bool**| Has virtual machines | 
 **hasVlans** | **optional.Bool**| Has VLANs | 
 **hasVrfs** | **optional.Bool**| Has VRFs | 
 **id** | [**optional.Interface of []string**](string.md)| Unique object identifier, either a UUID primary key or a composite key. | 
 **idIc** | [**optional.Interface of []string**](string.md)|  | 
 **idIe** | [**optional.Interface of []string**](string.md)|  | 
 **idIew** | [**optional.Interface of []string**](string.md)|  | 
 **idIre** | [**optional.Interface of []string**](string.md)|  | 
 **idIsw** | [**optional.Interface of []string**](string.md)|  | 
 **idN** | [**optional.Interface of []string**](string.md)|  | 
 **idNic** | [**optional.Interface of []string**](string.md)|  | 
 **idNie** | [**optional.Interface of []string**](string.md)|  | 
 **idNiew** | [**optional.Interface of []string**](string.md)|  | 
 **idNire** | [**optional.Interface of []string**](string.md)|  | 
 **idNisw** | [**optional.Interface of []string**](string.md)|  | 
 **idNre** | [**optional.Interface of []string**](string.md)|  | 
 **idRe** | [**optional.Interface of []string**](string.md)|  | 
 **ipAddresses** | [**optional.Interface of []string**](string.md)| IP addresses (ID) | 
 **ipAddressesIsnull** | **optional.Bool**|  | 
 **ipAddressesN** | [**optional.Interface of []string**](string.md)| IP addresses (ID) | 
 **lastUpdated** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedGt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedGte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedIsnull** | **optional.Bool**|  | 
 **lastUpdatedLt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedLte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **lastUpdatedN** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **locations** | [**optional.Interface of []string**](string.md)|  | 
 **locationsIsnull** | **optional.Bool**|  | 
 **locationsN** | [**optional.Interface of []string**](string.md)|  | 
 **name** | [**optional.Interface of []string**](string.md)|  | 
 **nameIc** | [**optional.Interface of []string**](string.md)|  | 
 **nameIe** | [**optional.Interface of []string**](string.md)|  | 
 **nameIew** | [**optional.Interface of []string**](string.md)|  | 
 **nameIre** | [**optional.Interface of []string**](string.md)|  | 
 **nameIsw** | [**optional.Interface of []string**](string.md)|  | 
 **nameN** | [**optional.Interface of []string**](string.md)|  | 
 **nameNic** | [**optional.Interface of []string**](string.md)|  | 
 **nameNie** | [**optional.Interface of []string**](string.md)|  | 
 **nameNiew** | [**optional.Interface of []string**](string.md)|  | 
 **nameNire** | [**optional.Interface of []string**](string.md)|  | 
 **nameNisw** | [**optional.Interface of []string**](string.md)|  | 
 **nameNre** | [**optional.Interface of []string**](string.md)|  | 
 **nameRe** | [**optional.Interface of []string**](string.md)|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **prefixes** | [**optional.Interface of []string**](string.md)| Prefixes (ID) | 
 **prefixesIsnull** | **optional.Bool**|  | 
 **prefixesN** | [**optional.Interface of []string**](string.md)| Prefixes (ID) | 
 **q** | **optional.String**| Search | 
 **rackReservations** | [**optional.Interface of []string**](string.md)| Rack reservations (ID) | 
 **rackReservationsIsnull** | **optional.Bool**|  | 
 **rackReservationsN** | [**optional.Interface of []string**](string.md)| Rack reservations (ID) | 
 **racks** | [**optional.Interface of []string**](string.md)|  | 
 **racksIsnull** | **optional.Bool**|  | 
 **racksN** | [**optional.Interface of []string**](string.md)|  | 
 **routeTargets** | [**optional.Interface of []string**](string.md)|  | 
 **routeTargetsIsnull** | **optional.Bool**|  | 
 **routeTargetsN** | [**optional.Interface of []string**](string.md)|  | 
 **sort** | **optional.String**| Which field to use when ordering the results. | 
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **tagsIsnull** | **optional.Bool**|  | 
 **tagsN** | [**optional.Interface of []string**](string.md)|  | 
 **tenantGroup** | [**optional.Interface of []string**](string.md)|  | 
 **tenantGroupIsnull** | **optional.Bool**|  | 
 **tenantGroupN** | [**optional.Interface of []string**](string.md)|  | 
 **virtualMachines** | [**optional.Interface of []string**](string.md)|  | 
 **virtualMachinesIsnull** | **optional.Bool**|  | 
 **virtualMachinesN** | [**optional.Interface of []string**](string.md)|  | 
 **vlans** | [**optional.Interface of []string**](string.md)| VLANs (ID) | 
 **vlansIsnull** | **optional.Bool**|  | 
 **vlansN** | [**optional.Interface of []string**](string.md)| VLANs (ID) | 
 **vrfs** | [**optional.Interface of []string**](string.md)|  | 
 **vrfsIsnull** | **optional.Bool**|  | 
 **vrfsN** | [**optional.Interface of []string**](string.md)|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedTenantList**](PaginatedTenantList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsNotesCreate**
> Note TenancyTenantsNotesCreate(ctx, body, id, optional)


API methods for returning or creating notes on an object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NoteInputRequest**](NoteInputRequest.md)|  | 
  **id** | [**string**](.md)| A UUID string identifying this tenant. | 
 **optional** | ***TenancyApiTenancyTenantsNotesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsNotesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**Note**](Note.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsNotesList**
> PaginatedNoteList TenancyTenantsNotesList(ctx, id, optional)


API methods for returning or creating notes on an object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant. | 
 **optional** | ***TenancyApiTenancyTenantsNotesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsNotesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedNoteList**](PaginatedNoteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsPartialUpdate**
> Tenant TenancyTenantsPartialUpdate(ctx, id, optional)


Partial update a tenant object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant. | 
 **optional** | ***TenancyApiTenancyTenantsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedTenantRequest**](PatchedTenantRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsRetrieve**
> Tenant TenancyTenantsRetrieve(ctx, id, optional)


Retrieve a tenant object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this tenant. | 
 **optional** | ***TenancyApiTenancyTenantsRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsUpdate**
> Tenant TenancyTenantsUpdate(ctx, body, id, optional)


Update a tenant object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TenantRequest**](TenantRequest.md)|  | 
  **id** | [**string**](.md)| A UUID string identifying this tenant. | 
 **optional** | ***TenancyApiTenancyTenantsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

