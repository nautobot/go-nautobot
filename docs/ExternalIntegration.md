# ExternalIntegration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**Name** | **string** |  | [default to null]
**RemoteUrl** | **string** |  | [default to null]
**VerifySsl** | **bool** | Verify SSL certificates when connecting to the remote system | [optional] [default to null]
**Timeout** | **int32** | Number of seconds to wait for a response | [optional] [default to null]
**ExtraConfig** | [**map[string]Object**](.md) | Optional user-defined JSON data for this integration | [optional] [default to null]
**HttpMethod** | [***OneOfExternalIntegrationHttpMethod**](OneOfExternalIntegrationHttpMethod.md) |  | [optional] [default to null]
**Headers** | [**map[string]Object**](.md) | Headers for the HTTP request | [optional] [default to null]
**CaFilePath** | **string** |  | [optional] [default to null]
**SecretsGroup** | [***BulkWritableExternalIntegrationRequestSecretsGroup**](BulkWritableExternalIntegrationRequest_secrets_group.md) |  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

