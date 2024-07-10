# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigRetrieve**](UsersApi.md#UsersConfigRetrieve) | **Get** /users/config/ | 
[**UsersGroupsBulkDestroy**](UsersApi.md#UsersGroupsBulkDestroy) | **Delete** /users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersApi.md#UsersGroupsBulkPartialUpdate) | **Patch** /users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersApi.md#UsersGroupsBulkUpdate) | **Put** /users/groups/ | 
[**UsersGroupsCreate**](UsersApi.md#UsersGroupsCreate) | **Post** /users/groups/ | 
[**UsersGroupsDestroy**](UsersApi.md#UsersGroupsDestroy) | **Delete** /users/groups/{id}/ | 
[**UsersGroupsList**](UsersApi.md#UsersGroupsList) | **Get** /users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersApi.md#UsersGroupsPartialUpdate) | **Patch** /users/groups/{id}/ | 
[**UsersGroupsRetrieve**](UsersApi.md#UsersGroupsRetrieve) | **Get** /users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersApi.md#UsersGroupsUpdate) | **Put** /users/groups/{id}/ | 
[**UsersPermissionsBulkDestroy**](UsersApi.md#UsersPermissionsBulkDestroy) | **Delete** /users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersApi.md#UsersPermissionsBulkPartialUpdate) | **Patch** /users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersApi.md#UsersPermissionsBulkUpdate) | **Put** /users/permissions/ | 
[**UsersPermissionsCreate**](UsersApi.md#UsersPermissionsCreate) | **Post** /users/permissions/ | 
[**UsersPermissionsDestroy**](UsersApi.md#UsersPermissionsDestroy) | **Delete** /users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersApi.md#UsersPermissionsList) | **Get** /users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersApi.md#UsersPermissionsPartialUpdate) | **Patch** /users/permissions/{id}/ | 
[**UsersPermissionsRetrieve**](UsersApi.md#UsersPermissionsRetrieve) | **Get** /users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersApi.md#UsersPermissionsUpdate) | **Put** /users/permissions/{id}/ | 
[**UsersTokensBulkDestroy**](UsersApi.md#UsersTokensBulkDestroy) | **Delete** /users/tokens/ | 
[**UsersTokensBulkPartialUpdate**](UsersApi.md#UsersTokensBulkPartialUpdate) | **Patch** /users/tokens/ | 
[**UsersTokensBulkUpdate**](UsersApi.md#UsersTokensBulkUpdate) | **Put** /users/tokens/ | 
[**UsersTokensCreate**](UsersApi.md#UsersTokensCreate) | **Post** /users/tokens/ | 
[**UsersTokensDestroy**](UsersApi.md#UsersTokensDestroy) | **Delete** /users/tokens/{id}/ | 
[**UsersTokensList**](UsersApi.md#UsersTokensList) | **Get** /users/tokens/ | 
[**UsersTokensPartialUpdate**](UsersApi.md#UsersTokensPartialUpdate) | **Patch** /users/tokens/{id}/ | 
[**UsersTokensRetrieve**](UsersApi.md#UsersTokensRetrieve) | **Get** /users/tokens/{id}/ | 
[**UsersTokensUpdate**](UsersApi.md#UsersTokensUpdate) | **Put** /users/tokens/{id}/ | 
[**UsersUsersBulkDestroy**](UsersApi.md#UsersUsersBulkDestroy) | **Delete** /users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersApi.md#UsersUsersBulkPartialUpdate) | **Patch** /users/users/ | 
[**UsersUsersBulkUpdate**](UsersApi.md#UsersUsersBulkUpdate) | **Put** /users/users/ | 
[**UsersUsersCreate**](UsersApi.md#UsersUsersCreate) | **Post** /users/users/ | 
[**UsersUsersDestroy**](UsersApi.md#UsersUsersDestroy) | **Delete** /users/users/{id}/ | 
[**UsersUsersList**](UsersApi.md#UsersUsersList) | **Get** /users/users/ | 
[**UsersUsersPartialUpdate**](UsersApi.md#UsersUsersPartialUpdate) | **Patch** /users/users/{id}/ | 
[**UsersUsersRetrieve**](UsersApi.md#UsersUsersRetrieve) | **Get** /users/users/{id}/ | 
[**UsersUsersUpdate**](UsersApi.md#UsersUsersUpdate) | **Put** /users/users/{id}/ | 

# **UsersConfigRetrieve**
> map[string]Object UsersConfigRetrieve(ctx, optional)


Return the config_data for the currently authenticated User.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersConfigRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersConfigRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

**map[string]Object**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsBulkDestroy**
> UsersGroupsBulkDestroy(ctx, body, optional)


Destroy a list of group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkOperationIntegerIdRequest**](BulkOperationIntegerIDRequest.md)|  | 
 **optional** | ***UsersApiUsersGroupsBulkDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsBulkDestroyOpts struct
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

# **UsersGroupsBulkPartialUpdate**
> []Group UsersGroupsBulkPartialUpdate(ctx, body, optional)


Partial update a list of group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchedBulkWritableGroupRequest**](PatchedBulkWritableGroupRequest.md)|  | 
 **optional** | ***UsersApiUsersGroupsBulkPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsBulkPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsBulkUpdate**
> []Group UsersGroupsBulkUpdate(ctx, body, optional)


Update a list of group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkWritableGroupRequest**](BulkWritableGroupRequest.md)|  | 
 **optional** | ***UsersApiUsersGroupsBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsBulkUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsCreate**
> Group UsersGroupsCreate(ctx, body, optional)


Create one or more group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupRequest**](GroupRequest.md)|  | 
 **optional** | ***UsersApiUsersGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsDestroy**
> UsersGroupsDestroy(ctx, id, optional)


Destroy a group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 
 **optional** | ***UsersApiUsersGroupsDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsDestroyOpts struct
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

# **UsersGroupsList**
> PaginatedGroupList UsersGroupsList(ctx, optional)


Retrieve a list of group objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**|  | 
 **id** | [**optional.Interface of []int32**](int32.md)| Unique object identifier, either a UUID primary key or a composite key. | 
 **idGt** | [**optional.Interface of []int32**](int32.md)|  | 
 **idGte** | [**optional.Interface of []int32**](int32.md)|  | 
 **idLt** | [**optional.Interface of []int32**](int32.md)|  | 
 **idLte** | [**optional.Interface of []int32**](int32.md)|  | 
 **idN** | [**optional.Interface of []int32**](int32.md)|  | 
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
 **q** | **optional.String**| Search | 
 **sort** | **optional.String**| Which field to use when ordering the results. | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedGroupList**](PaginatedGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsPartialUpdate**
> Group UsersGroupsPartialUpdate(ctx, id, optional)


Partial update a group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 
 **optional** | ***UsersApiUsersGroupsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedGroupRequest**](PatchedGroupRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsRetrieve**
> Group UsersGroupsRetrieve(ctx, id, optional)


Retrieve a group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 
 **optional** | ***UsersApiUsersGroupsRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsUpdate**
> Group UsersGroupsUpdate(ctx, body, id, optional)


Update a group object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupRequest**](GroupRequest.md)|  | 
  **id** | **int32**| A unique integer value identifying this group. | 
 **optional** | ***UsersApiUsersGroupsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsBulkDestroy**
> UsersPermissionsBulkDestroy(ctx, body, optional)


Destroy a list of permission objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkOperationRequest**](BulkOperationRequest.md)|  | 
 **optional** | ***UsersApiUsersPermissionsBulkDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsBulkDestroyOpts struct
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

# **UsersPermissionsBulkPartialUpdate**
> []ObjectPermission UsersPermissionsBulkPartialUpdate(ctx, body, optional)


Partial update a list of permission objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchedBulkWritableObjectPermissionRequest**](PatchedBulkWritableObjectPermissionRequest.md)|  | 
 **optional** | ***UsersApiUsersPermissionsBulkPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsBulkPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsBulkUpdate**
> []ObjectPermission UsersPermissionsBulkUpdate(ctx, body, optional)


Update a list of permission objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkWritableObjectPermissionRequest**](BulkWritableObjectPermissionRequest.md)|  | 
 **optional** | ***UsersApiUsersPermissionsBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsBulkUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsCreate**
> ObjectPermission UsersPermissionsCreate(ctx, body, optional)


Create one or more permission objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObjectPermissionRequest**](ObjectPermissionRequest.md)|  | 
 **optional** | ***UsersApiUsersPermissionsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsDestroy**
> UsersPermissionsDestroy(ctx, id, optional)


Destroy a permission object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this permission. | 
 **optional** | ***UsersApiUsersPermissionsDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsDestroyOpts struct
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

# **UsersPermissionsList**
> PaginatedObjectPermissionList UsersPermissionsList(ctx, optional)


Retrieve a list of permission objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersPermissionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
 **enabled** | **optional.Bool**|  | 
 **format** | **optional.String**|  | 
 **groups** | [**optional.Interface of []string**](string.md)| Group (name) | 
 **groupsN** | [**optional.Interface of []string**](string.md)| Group (name) | 
 **groupsId** | [**optional.Interface of []int32**](int32.md)| Group (ID) | 
 **groupsIdN** | [**optional.Interface of []int32**](int32.md)| Group (ID) | 
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
 **objectTypes** | [**optional.Interface of []int32**](int32.md)|  | 
 **objectTypesN** | [**optional.Interface of []int32**](int32.md)|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **q** | **optional.String**| Search | 
 **sort** | **optional.String**| Which field to use when ordering the results. | 
 **users** | [**optional.Interface of []string**](string.md)|  | 
 **usersN** | [**optional.Interface of []string**](string.md)|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedObjectPermissionList**](PaginatedObjectPermissionList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsPartialUpdate**
> ObjectPermission UsersPermissionsPartialUpdate(ctx, id, optional)


Partial update a permission object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this permission. | 
 **optional** | ***UsersApiUsersPermissionsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedObjectPermissionRequest**](PatchedObjectPermissionRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsRetrieve**
> ObjectPermission UsersPermissionsRetrieve(ctx, id, optional)


Retrieve a permission object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this permission. | 
 **optional** | ***UsersApiUsersPermissionsRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsUpdate**
> ObjectPermission UsersPermissionsUpdate(ctx, body, id, optional)


Update a permission object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObjectPermissionRequest**](ObjectPermissionRequest.md)|  | 
  **id** | [**string**](.md)| A UUID string identifying this permission. | 
 **optional** | ***UsersApiUsersPermissionsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensBulkDestroy**
> UsersTokensBulkDestroy(ctx, body, optional)


Destroy a list of token objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkOperationRequest**](BulkOperationRequest.md)|  | 
 **optional** | ***UsersApiUsersTokensBulkDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensBulkDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensBulkPartialUpdate**
> []Token UsersTokensBulkPartialUpdate(ctx, body, optional)


Partial update a list of token objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchedBulkWritableTokenRequest**](PatchedBulkWritableTokenRequest.md)|  | 
 **optional** | ***UsersApiUsersTokensBulkPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensBulkPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensBulkUpdate**
> []Token UsersTokensBulkUpdate(ctx, body, optional)


Update a list of token objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkWritableTokenRequest**](BulkWritableTokenRequest.md)|  | 
 **optional** | ***UsersApiUsersTokensBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensBulkUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensCreate**
> Token UsersTokensCreate(ctx, optional)


Create one or more token objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersTokensCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TokenRequest**](TokenRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensDestroy**
> UsersTokensDestroy(ctx, id, optional)


Destroy a token object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this token. | 
 **optional** | ***UsersApiUsersTokensDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensList**
> PaginatedTokenList UsersTokensList(ctx, optional)


Retrieve a list of token objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersTokensListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdGt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **createdGte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
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
 **expires** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **expiresGt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **expiresGte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **expiresIsnull** | **optional.Bool**|  | 
 **expiresLt** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **expiresLte** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **expiresN** | [**optional.Interface of []time.Time**](time.Time.md)|  | 
 **format** | **optional.String**|  | 
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
 **key** | [**optional.Interface of []string**](string.md)|  | 
 **keyIc** | [**optional.Interface of []string**](string.md)|  | 
 **keyIe** | [**optional.Interface of []string**](string.md)|  | 
 **keyIew** | [**optional.Interface of []string**](string.md)|  | 
 **keyIre** | [**optional.Interface of []string**](string.md)|  | 
 **keyIsw** | [**optional.Interface of []string**](string.md)|  | 
 **keyN** | [**optional.Interface of []string**](string.md)|  | 
 **keyNic** | [**optional.Interface of []string**](string.md)|  | 
 **keyNie** | [**optional.Interface of []string**](string.md)|  | 
 **keyNiew** | [**optional.Interface of []string**](string.md)|  | 
 **keyNire** | [**optional.Interface of []string**](string.md)|  | 
 **keyNisw** | [**optional.Interface of []string**](string.md)|  | 
 **keyNre** | [**optional.Interface of []string**](string.md)|  | 
 **keyRe** | [**optional.Interface of []string**](string.md)|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **q** | **optional.String**| Search | 
 **sort** | **optional.String**| Which field to use when ordering the results. | 
 **writeEnabled** | **optional.Bool**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedTokenList**](PaginatedTokenList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensPartialUpdate**
> Token UsersTokensPartialUpdate(ctx, id, optional)


Partial update a token object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this token. | 
 **optional** | ***UsersApiUsersTokensPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedTokenRequest**](PatchedTokenRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensRetrieve**
> Token UsersTokensRetrieve(ctx, id, optional)


Retrieve a token object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this token. | 
 **optional** | ***UsersApiUsersTokensRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersTokensUpdate**
> Token UsersTokensUpdate(ctx, id, optional)


Update a token object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this token. | 
 **optional** | ***UsersApiUsersTokensUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersTokensUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TokenRequest**](TokenRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersBulkDestroy**
> UsersUsersBulkDestroy(ctx, body, optional)


Destroy a list of user objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkOperationRequest**](BulkOperationRequest.md)|  | 
 **optional** | ***UsersApiUsersUsersBulkDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersBulkDestroyOpts struct
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

# **UsersUsersBulkPartialUpdate**
> []User UsersUsersBulkPartialUpdate(ctx, body, optional)


Partial update a list of user objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PatchedBulkWritableUserRequest**](PatchedBulkWritableUserRequest.md)|  | 
 **optional** | ***UsersApiUsersUsersBulkPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersBulkPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersBulkUpdate**
> []User UsersUsersBulkUpdate(ctx, body, optional)


Update a list of user objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]BulkWritableUserRequest**](BulkWritableUserRequest.md)|  | 
 **optional** | ***UsersApiUsersUsersBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersBulkUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**[]User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersCreate**
> User UsersUsersCreate(ctx, body, optional)


Create one or more user objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserRequest**](UserRequest.md)|  | 
 **optional** | ***UsersApiUsersUsersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.**|  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersDestroy**
> UsersUsersDestroy(ctx, id, optional)


Destroy a user object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this user. | 
 **optional** | ***UsersApiUsersUsersDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersDestroyOpts struct
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

# **UsersUsersList**
> PaginatedUserList UsersUsersList(ctx, optional)


Retrieve a list of user objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | [**optional.Interface of []string**](string.md)|  | 
 **emailIc** | [**optional.Interface of []string**](string.md)|  | 
 **emailIe** | [**optional.Interface of []string**](string.md)|  | 
 **emailIew** | [**optional.Interface of []string**](string.md)|  | 
 **emailIre** | [**optional.Interface of []string**](string.md)|  | 
 **emailIsw** | [**optional.Interface of []string**](string.md)|  | 
 **emailN** | [**optional.Interface of []string**](string.md)|  | 
 **emailNic** | [**optional.Interface of []string**](string.md)|  | 
 **emailNie** | [**optional.Interface of []string**](string.md)|  | 
 **emailNiew** | [**optional.Interface of []string**](string.md)|  | 
 **emailNire** | [**optional.Interface of []string**](string.md)|  | 
 **emailNisw** | [**optional.Interface of []string**](string.md)|  | 
 **emailNre** | [**optional.Interface of []string**](string.md)|  | 
 **emailRe** | [**optional.Interface of []string**](string.md)|  | 
 **firstName** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameIc** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameIe** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameIew** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameIre** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameIsw** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameN** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameNic** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameNie** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameNiew** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameNire** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameNisw** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameNre** | [**optional.Interface of []string**](string.md)|  | 
 **firstNameRe** | [**optional.Interface of []string**](string.md)|  | 
 **format** | **optional.String**|  | 
 **groups** | [**optional.Interface of []string**](string.md)| Group (name) | 
 **groupsN** | [**optional.Interface of []string**](string.md)| Group (name) | 
 **groupsId** | [**optional.Interface of []int32**](int32.md)| Group (ID) | 
 **groupsIdN** | [**optional.Interface of []int32**](int32.md)| Group (ID) | 
 **hasObjectChanges** | **optional.Bool**| Has Changes | 
 **hasObjectPermissions** | **optional.Bool**| Has object permissions | 
 **hasRackReservations** | **optional.Bool**| Has Rack Reservations | 
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
 **isActive** | **optional.Bool**|  | 
 **isStaff** | **optional.Bool**|  | 
 **lastName** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameIc** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameIe** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameIew** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameIre** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameIsw** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameN** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameNic** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameNie** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameNiew** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameNire** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameNisw** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameNre** | [**optional.Interface of []string**](string.md)|  | 
 **lastNameRe** | [**optional.Interface of []string**](string.md)|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **objectChanges** | [**optional.Interface of []string**](string.md)| Object Changes (ID) | 
 **objectChangesIsnull** | **optional.Bool**|  | 
 **objectChangesN** | [**optional.Interface of []string**](string.md)| Object Changes (ID) | 
 **objectPermissions** | [**optional.Interface of []string**](string.md)|  | 
 **objectPermissionsIsnull** | **optional.Bool**|  | 
 **objectPermissionsN** | [**optional.Interface of []string**](string.md)|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **q** | **optional.String**| Search | 
 **rackReservationsId** | [**optional.Interface of []string**](string.md)| Rack Reservation (ID) | 
 **rackReservationsIdIsnull** | **optional.Bool**|  | 
 **rackReservationsIdN** | [**optional.Interface of []string**](string.md)| Rack Reservation (ID) | 
 **sort** | **optional.String**| Which field to use when ordering the results. | 
 **username** | [**optional.Interface of []string**](string.md)|  | 
 **usernameIc** | [**optional.Interface of []string**](string.md)|  | 
 **usernameIe** | [**optional.Interface of []string**](string.md)|  | 
 **usernameIew** | [**optional.Interface of []string**](string.md)|  | 
 **usernameIre** | [**optional.Interface of []string**](string.md)|  | 
 **usernameIsw** | [**optional.Interface of []string**](string.md)|  | 
 **usernameN** | [**optional.Interface of []string**](string.md)|  | 
 **usernameNic** | [**optional.Interface of []string**](string.md)|  | 
 **usernameNie** | [**optional.Interface of []string**](string.md)|  | 
 **usernameNiew** | [**optional.Interface of []string**](string.md)|  | 
 **usernameNire** | [**optional.Interface of []string**](string.md)|  | 
 **usernameNisw** | [**optional.Interface of []string**](string.md)|  | 
 **usernameNre** | [**optional.Interface of []string**](string.md)|  | 
 **usernameRe** | [**optional.Interface of []string**](string.md)|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPartialUpdate**
> User UsersUsersPartialUpdate(ctx, id, optional)


Partial update a user object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this user. | 
 **optional** | ***UsersApiUsersUsersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchedUserRequest**](PatchedUserRequest.md)|  | 
 **format** | **optional.**|  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersRetrieve**
> User UsersUsersRetrieve(ctx, id, optional)


Retrieve a user object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A UUID string identifying this user. | 
 **optional** | ***UsersApiUsersUsersRetrieveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersRetrieveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | 
 **depth** | **optional.Int32**| Serializer Depth | [default to 1]

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersUpdate**
> User UsersUsersUpdate(ctx, body, id, optional)


Update a user object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserRequest**](UserRequest.md)|  | 
  **id** | [**string**](.md)| A UUID string identifying this user. | 
 **optional** | ***UsersApiUsersUsersUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.**|  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

