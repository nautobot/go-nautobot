# GitRepository

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**ProvidedContents** | [**[]OneOfGitRepositoryProvidedContentsItems**](.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Slug** | **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] [default to null]
**RemoteUrl** | **string** | Only HTTP and HTTPS URLs are presently supported | [default to null]
**Branch** | **string** |  | [optional] [default to null]
**CurrentHead** | **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] [default to null]
**SecretsGroup** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

