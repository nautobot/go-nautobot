# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**LastLogin** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**IsSuperuser** | **bool** | Designates that this user has all permissions without explicitly assigning them. | [optional] [default to null]
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [default to null]
**FirstName** | **string** |  | [optional] [default to null]
**LastName** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [optional] [default to null]
**IsStaff** | **bool** | Designates whether the user can log into this admin site. | [optional] [default to null]
**IsActive** | **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] [default to null]
**DateJoined** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ConfigData** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Groups** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) | The groups this user belongs to. A user will get all permissions granted to each of their groups. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

