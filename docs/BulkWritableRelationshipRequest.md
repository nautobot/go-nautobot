# BulkWritableRelationshipRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**SourceType** | **string** |  | [default to null]
**DestinationType** | **string** |  | [default to null]
**Label** | **string** | Label of the relationship as displayed to users | [default to null]
**Key** | **string** | Internal relationship key. Please use underscores rather than dashes in this key. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Type_** | [***AllOfBulkWritableRelationshipRequestType_**](AllOfBulkWritableRelationshipRequestType_.md) | Cardinality of this relationship | [optional] [default to null]
**RequiredOn** | [***OneOfBulkWritableRelationshipRequestRequiredOn**](OneOfBulkWritableRelationshipRequestRequiredOn.md) | Objects on the specified side MUST implement this relationship. Not permitted for symmetric relationships. | [optional] [default to null]
**SourceLabel** | **string** | Label for related destination objects, as displayed on the source object. | [optional] [default to null]
**SourceHidden** | **bool** | Hide this relationship on the source object. | [optional] [default to null]
**SourceFilter** | [**map[string]Object**](.md) | Filterset filter matching the applicable source objects of the selected type | [optional] [default to null]
**DestinationLabel** | **string** | Label for related source objects, as displayed on the destination object. | [optional] [default to null]
**DestinationHidden** | **bool** | Hide this relationship on the destination object. | [optional] [default to null]
**DestinationFilter** | [**map[string]Object**](.md) | Filterset filter matching the applicable destination objects of the selected type | [optional] [default to null]
**AdvancedUi** | **bool** | Hide this field from the object&#x27;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

