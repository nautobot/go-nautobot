# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Teams** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Name** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] [default to ""]
**Email** | Pointer to **string** |  | [optional] [default to ""]
**Address** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewContact

`func NewContact(id string, objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Contact) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Contact) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Contact) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Contact) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Contact) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Contact) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Contact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Contact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Contact) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Contact) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Contact) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Contact) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetTeams

`func (o *Contact) GetTeams() []BulkWritableCableRequestStatus`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Contact) GetTeamsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Contact) SetTeams(v []BulkWritableCableRequestStatus)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Contact) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetName

`func (o *Contact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contact) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *Contact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Contact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Contact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Contact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *Contact) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Contact) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Contact) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Contact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComments

`func (o *Contact) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Contact) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Contact) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Contact) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreated

`func (o *Contact) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Contact) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Contact) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Contact) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Contact) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Contact) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Contact) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Contact) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Contact) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Contact) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Contact) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Contact) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Contact) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Contact) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Contact) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Contact) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Contact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


