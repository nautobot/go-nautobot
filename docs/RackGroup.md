# RackGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**TreeDepth** | **NullableInt32** |  | [readonly] 
**RackCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRackGroup

`func NewRackGroup(id string, objectType string, display string, url string, naturalSlug string, treeDepth NullableInt32, name string, location BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *RackGroup`

NewRackGroup instantiates a new RackGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackGroupWithDefaults

`func NewRackGroupWithDefaults() *RackGroup`

NewRackGroupWithDefaults instantiates a new RackGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RackGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RackGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RackGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *RackGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RackGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RackGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *RackGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RackGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RackGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *RackGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RackGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RackGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *RackGroup) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *RackGroup) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *RackGroup) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetTreeDepth

`func (o *RackGroup) GetTreeDepth() int32`

GetTreeDepth returns the TreeDepth field if non-nil, zero value otherwise.

### GetTreeDepthOk

`func (o *RackGroup) GetTreeDepthOk() (*int32, bool)`

GetTreeDepthOk returns a tuple with the TreeDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeDepth

`func (o *RackGroup) SetTreeDepth(v int32)`

SetTreeDepth sets TreeDepth field to given value.


### SetTreeDepthNil

`func (o *RackGroup) SetTreeDepthNil(b bool)`

 SetTreeDepthNil sets the value for TreeDepth to be an explicit nil

### UnsetTreeDepth
`func (o *RackGroup) UnsetTreeDepth()`

UnsetTreeDepth ensures that no value is present for TreeDepth, not even an explicit nil
### GetRackCount

`func (o *RackGroup) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *RackGroup) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *RackGroup) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.

### HasRackCount

`func (o *RackGroup) HasRackCount() bool`

HasRackCount returns a boolean if a field has been set.

### GetName

`func (o *RackGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RackGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RackGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RackGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RackGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RackGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RackGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *RackGroup) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RackGroup) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RackGroup) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RackGroup) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *RackGroup) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *RackGroup) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetLocation

`func (o *RackGroup) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RackGroup) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RackGroup) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetCreated

`func (o *RackGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RackGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RackGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *RackGroup) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *RackGroup) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *RackGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RackGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RackGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *RackGroup) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *RackGroup) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *RackGroup) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *RackGroup) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *RackGroup) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *RackGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RackGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RackGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RackGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


