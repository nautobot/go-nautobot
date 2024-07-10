# Cable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**TerminationAType** | **string** |  | [default to null]
**TerminationBType** | **string** |  | [default to null]
**TerminationA** | [***AllOfCableTerminationA**](AllOfCableTerminationA.md) |  | [default to null]
**TerminationB** | [***AllOfCableTerminationB**](AllOfCableTerminationB.md) |  | [default to null]
**LengthUnit** | [***CableLengthUnit**](Cable_length_unit.md) |  | [optional] [default to null]
**Type_** | [***CableType**](Cable_type.md) |  | [optional] [default to null]
**TerminationAId** | **string** |  | [default to null]
**TerminationBId** | **string** |  | [default to null]
**Label** | **string** |  | [optional] [default to null]
**Color** | **string** | RGB color in hexadecimal (e.g. 00ff00) | [optional] [default to null]
**Length** | **int32** |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

