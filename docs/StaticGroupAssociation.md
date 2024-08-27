# StaticGroupAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**AssociatedObjectType** | **string** |  | 
**AssociatedObject** | [**DynamicGroupAssociatedObject**](DynamicGroupAssociatedObject.md) |  | [readonly] 
**AssociatedObjectId** | **string** |  | 
**DynamicGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStaticGroupAssociation

`func NewStaticGroupAssociation(id string, objectType string, display string, url string, naturalSlug string, associatedObjectType string, associatedObject DynamicGroupAssociatedObject, associatedObjectId string, dynamicGroup BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *StaticGroupAssociation`

NewStaticGroupAssociation instantiates a new StaticGroupAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticGroupAssociationWithDefaults

`func NewStaticGroupAssociationWithDefaults() *StaticGroupAssociation`

NewStaticGroupAssociationWithDefaults instantiates a new StaticGroupAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StaticGroupAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticGroupAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticGroupAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *StaticGroupAssociation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StaticGroupAssociation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StaticGroupAssociation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *StaticGroupAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *StaticGroupAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *StaticGroupAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *StaticGroupAssociation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StaticGroupAssociation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StaticGroupAssociation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *StaticGroupAssociation) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *StaticGroupAssociation) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *StaticGroupAssociation) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetAssociatedObjectType

`func (o *StaticGroupAssociation) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *StaticGroupAssociation) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *StaticGroupAssociation) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.


### GetAssociatedObject

`func (o *StaticGroupAssociation) GetAssociatedObject() DynamicGroupAssociatedObject`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *StaticGroupAssociation) GetAssociatedObjectOk() (*DynamicGroupAssociatedObject, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *StaticGroupAssociation) SetAssociatedObject(v DynamicGroupAssociatedObject)`

SetAssociatedObject sets AssociatedObject field to given value.


### GetAssociatedObjectId

`func (o *StaticGroupAssociation) GetAssociatedObjectId() string`

GetAssociatedObjectId returns the AssociatedObjectId field if non-nil, zero value otherwise.

### GetAssociatedObjectIdOk

`func (o *StaticGroupAssociation) GetAssociatedObjectIdOk() (*string, bool)`

GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectId

`func (o *StaticGroupAssociation) SetAssociatedObjectId(v string)`

SetAssociatedObjectId sets AssociatedObjectId field to given value.


### GetDynamicGroup

`func (o *StaticGroupAssociation) GetDynamicGroup() BulkWritableCableRequestStatus`

GetDynamicGroup returns the DynamicGroup field if non-nil, zero value otherwise.

### GetDynamicGroupOk

`func (o *StaticGroupAssociation) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetDynamicGroupOk returns a tuple with the DynamicGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicGroup

`func (o *StaticGroupAssociation) SetDynamicGroup(v BulkWritableCableRequestStatus)`

SetDynamicGroup sets DynamicGroup field to given value.


### GetCreated

`func (o *StaticGroupAssociation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StaticGroupAssociation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StaticGroupAssociation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *StaticGroupAssociation) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *StaticGroupAssociation) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *StaticGroupAssociation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *StaticGroupAssociation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *StaticGroupAssociation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *StaticGroupAssociation) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *StaticGroupAssociation) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *StaticGroupAssociation) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *StaticGroupAssociation) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *StaticGroupAssociation) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *StaticGroupAssociation) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StaticGroupAssociation) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StaticGroupAssociation) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StaticGroupAssociation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


