# ModuleFamily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ModuleTypeCount** | Pointer to **int32** |  | [optional] [readonly] 
**ModuleBayCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModuleFamily

`func NewModuleFamily(objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ModuleFamily`

NewModuleFamily instantiates a new ModuleFamily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleFamilyWithDefaults

`func NewModuleFamilyWithDefaults() *ModuleFamily`

NewModuleFamilyWithDefaults instantiates a new ModuleFamily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModuleFamily) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModuleFamily) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModuleFamily) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModuleFamily) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ModuleFamily) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ModuleFamily) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ModuleFamily) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ModuleFamily) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ModuleFamily) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ModuleFamily) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ModuleFamily) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModuleFamily) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModuleFamily) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ModuleFamily) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ModuleFamily) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ModuleFamily) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetModuleTypeCount

`func (o *ModuleFamily) GetModuleTypeCount() int32`

GetModuleTypeCount returns the ModuleTypeCount field if non-nil, zero value otherwise.

### GetModuleTypeCountOk

`func (o *ModuleFamily) GetModuleTypeCountOk() (*int32, bool)`

GetModuleTypeCountOk returns a tuple with the ModuleTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleTypeCount

`func (o *ModuleFamily) SetModuleTypeCount(v int32)`

SetModuleTypeCount sets ModuleTypeCount field to given value.

### HasModuleTypeCount

`func (o *ModuleFamily) HasModuleTypeCount() bool`

HasModuleTypeCount returns a boolean if a field has been set.

### GetModuleBayCount

`func (o *ModuleFamily) GetModuleBayCount() int32`

GetModuleBayCount returns the ModuleBayCount field if non-nil, zero value otherwise.

### GetModuleBayCountOk

`func (o *ModuleFamily) GetModuleBayCountOk() (*int32, bool)`

GetModuleBayCountOk returns a tuple with the ModuleBayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayCount

`func (o *ModuleFamily) SetModuleBayCount(v int32)`

SetModuleBayCount sets ModuleBayCount field to given value.

### HasModuleBayCount

`func (o *ModuleFamily) HasModuleBayCount() bool`

HasModuleBayCount returns a boolean if a field has been set.

### GetName

`func (o *ModuleFamily) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleFamily) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleFamily) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ModuleFamily) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleFamily) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleFamily) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleFamily) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *ModuleFamily) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModuleFamily) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModuleFamily) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ModuleFamily) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ModuleFamily) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ModuleFamily) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ModuleFamily) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ModuleFamily) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ModuleFamily) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ModuleFamily) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *ModuleFamily) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ModuleFamily) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ModuleFamily) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *ModuleFamily) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleFamily) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleFamily) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleFamily) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


