# PowerPortTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**Type** | Pointer to [**PowerPortType**](PowerPortType.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**DeviceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPowerPortTemplate

`func NewPowerPortTemplate(id string, objectType string, display string, url string, naturalSlug string, notesUrl string, name string, deviceType BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, ) *PowerPortTemplate`

NewPowerPortTemplate instantiates a new PowerPortTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPortTemplateWithDefaults

`func NewPowerPortTemplateWithDefaults() *PowerPortTemplate`

NewPowerPortTemplateWithDefaults instantiates a new PowerPortTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerPortTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerPortTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerPortTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PowerPortTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPortTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPortTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *PowerPortTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerPortTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerPortTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *PowerPortTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerPortTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerPortTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *PowerPortTemplate) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *PowerPortTemplate) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *PowerPortTemplate) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetNotesUrl

`func (o *PowerPortTemplate) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *PowerPortTemplate) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *PowerPortTemplate) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetType

`func (o *PowerPortTemplate) GetType() PowerPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerPortTemplate) GetTypeOk() (*PowerPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerPortTemplate) SetType(v PowerPortType)`

SetType sets Type field to given value.

### HasType

`func (o *PowerPortTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PowerPortTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPortTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPortTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PowerPortTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerPortTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerPortTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerPortTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PowerPortTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerPortTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerPortTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerPortTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *PowerPortTemplate) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *PowerPortTemplate) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *PowerPortTemplate) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *PowerPortTemplate) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *PowerPortTemplate) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *PowerPortTemplate) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *PowerPortTemplate) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *PowerPortTemplate) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *PowerPortTemplate) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *PowerPortTemplate) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *PowerPortTemplate) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *PowerPortTemplate) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDeviceType

`func (o *PowerPortTemplate) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PowerPortTemplate) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PowerPortTemplate) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.


### GetCreated

`func (o *PowerPortTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerPortTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerPortTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PowerPortTemplate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PowerPortTemplate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PowerPortTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerPortTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerPortTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PowerPortTemplate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PowerPortTemplate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetCustomFields

`func (o *PowerPortTemplate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerPortTemplate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerPortTemplate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerPortTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


