# GitRepositoryRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvidedContents** | [**[]OneOfGitRepositoryRequestProvidedContentsItems**](.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Slug** | **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] [default to null]
**RemoteUrl** | **string** | Only HTTP and HTTPS URLs are presently supported | [default to null]
**Branch** | **string** |  | [optional] [default to null]
**CurrentHead** | **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] [default to null]
**SecretsGroup** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

