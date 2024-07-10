# PowerFeed

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**CablePeerType** | **string** |  | [default to null]
**CablePeer** | [***AllOfPowerFeedCablePeer**](AllOfPowerFeedCablePeer.md) |  | [default to null]
**ConnectedEndpointType** | **string** |  | [default to null]
**ConnectedEndpoint** | [***AllOfPowerFeedConnectedEndpoint**](AllOfPowerFeedConnectedEndpoint.md) |  | [default to null]
**ConnectedEndpointReachable** | **bool** |  | [default to null]
**Type_** | [***PowerFeedType**](PowerFeed_type.md) |  | [optional] [default to null]
**Supply** | [***PowerFeedSupply**](PowerFeed_supply.md) |  | [optional] [default to null]
**Phase** | [***PowerFeedPhase**](PowerFeed_phase.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Voltage** | **int32** |  | [optional] [default to null]
**Amperage** | **int32** |  | [optional] [default to null]
**MaxUtilization** | **int32** | Maximum permissible draw (percentage) | [optional] [default to null]
**AvailablePower** | **int32** |  | [default to null]
**Comments** | **string** |  | [optional] [default to null]
**Cable** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**PowerPanel** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Rack** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

