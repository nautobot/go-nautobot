# DynamicGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ContentType** | **string** |  | 
**Name** | **string** | Dynamic Group name | 
**Description** | Pointer to **string** |  | [optional] 
**Filter** | **map[string]interface{}** | A JSON-encoded dictionary of filter parameters for group membership | 
**Children** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDynamicGroup

`func NewDynamicGroup(id string, objectType string, display string, url string, naturalSlug string, contentType string, name string, filter map[string]interface{}, children []BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *DynamicGroup`

NewDynamicGroup instantiates a new DynamicGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicGroupWithDefaults

`func NewDynamicGroupWithDefaults() *DynamicGroup`

NewDynamicGroupWithDefaults instantiates a new DynamicGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DynamicGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *DynamicGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DynamicGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DynamicGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *DynamicGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DynamicGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DynamicGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *DynamicGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DynamicGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DynamicGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *DynamicGroup) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *DynamicGroup) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *DynamicGroup) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetContentType

`func (o *DynamicGroup) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DynamicGroup) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DynamicGroup) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *DynamicGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DynamicGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilter

`func (o *DynamicGroup) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DynamicGroup) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DynamicGroup) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.


### GetChildren

`func (o *DynamicGroup) GetChildren() []BulkWritableCableRequestStatus`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DynamicGroup) GetChildrenOk() (*[]BulkWritableCableRequestStatus, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *DynamicGroup) SetChildren(v []BulkWritableCableRequestStatus)`

SetChildren sets Children field to given value.


### GetCreated

`func (o *DynamicGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DynamicGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DynamicGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *DynamicGroup) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *DynamicGroup) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *DynamicGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DynamicGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DynamicGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *DynamicGroup) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *DynamicGroup) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *DynamicGroup) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *DynamicGroup) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *DynamicGroup) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *DynamicGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DynamicGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DynamicGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DynamicGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


