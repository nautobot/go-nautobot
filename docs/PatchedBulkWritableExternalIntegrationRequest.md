# PatchedBulkWritableExternalIntegrationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Name** | **string** |  | [optional] [default to null]
**RemoteUrl** | **string** |  | [optional] [default to null]
**VerifySsl** | **bool** | Verify SSL certificates when connecting to the remote system | [optional] [default to null]
**Timeout** | **int32** | Number of seconds to wait for a response | [optional] [default to null]
**ExtraConfig** | [**map[string]Object**](.md) | Optional user-defined JSON data for this integration | [optional] [default to null]
**HttpMethod** | [***OneOfPatchedBulkWritableExternalIntegrationRequestHttpMethod**](OneOfPatchedBulkWritableExternalIntegrationRequestHttpMethod.md) |  | [optional] [default to null]
**Headers** | [**map[string]Object**](.md) | Headers for the HTTP request | [optional] [default to null]
**CaFilePath** | **string** |  | [optional] [default to null]
**SecretsGroup** | [***BulkWritableExternalIntegrationRequestSecretsGroup**](BulkWritableExternalIntegrationRequest_secrets_group.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

