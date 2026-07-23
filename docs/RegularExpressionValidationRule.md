# RegularExpressionValidationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ContentType** | **string** |  | 
**Name** | **string** |  | 
**Field** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** | Optional error message to display when validation fails. | [optional] 
**RegularExpression** | **string** |  | 
**ContextProcessing** | Pointer to **bool** | When enabled, the regular expression value is first processed as a Jinja2 template with access to the object being validated in a variable named &lt;code&gt;obj&lt;/code&gt;. | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRegularExpressionValidationRule

`func NewRegularExpressionValidationRule(objectType string, display string, url string, naturalSlug string, contentType string, name string, field string, regularExpression string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *RegularExpressionValidationRule`

NewRegularExpressionValidationRule instantiates a new RegularExpressionValidationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegularExpressionValidationRuleWithDefaults

`func NewRegularExpressionValidationRuleWithDefaults() *RegularExpressionValidationRule`

NewRegularExpressionValidationRuleWithDefaults instantiates a new RegularExpressionValidationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegularExpressionValidationRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegularExpressionValidationRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegularExpressionValidationRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegularExpressionValidationRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *RegularExpressionValidationRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RegularExpressionValidationRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RegularExpressionValidationRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *RegularExpressionValidationRule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RegularExpressionValidationRule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RegularExpressionValidationRule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *RegularExpressionValidationRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RegularExpressionValidationRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RegularExpressionValidationRule) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *RegularExpressionValidationRule) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *RegularExpressionValidationRule) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *RegularExpressionValidationRule) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetContentType

`func (o *RegularExpressionValidationRule) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *RegularExpressionValidationRule) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *RegularExpressionValidationRule) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *RegularExpressionValidationRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegularExpressionValidationRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegularExpressionValidationRule) SetName(v string)`

SetName sets Name field to given value.


### GetField

`func (o *RegularExpressionValidationRule) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RegularExpressionValidationRule) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RegularExpressionValidationRule) SetField(v string)`

SetField sets Field field to given value.


### GetEnabled

`func (o *RegularExpressionValidationRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RegularExpressionValidationRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RegularExpressionValidationRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RegularExpressionValidationRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *RegularExpressionValidationRule) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RegularExpressionValidationRule) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RegularExpressionValidationRule) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RegularExpressionValidationRule) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetRegularExpression

`func (o *RegularExpressionValidationRule) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *RegularExpressionValidationRule) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *RegularExpressionValidationRule) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.


### GetContextProcessing

`func (o *RegularExpressionValidationRule) GetContextProcessing() bool`

GetContextProcessing returns the ContextProcessing field if non-nil, zero value otherwise.

### GetContextProcessingOk

`func (o *RegularExpressionValidationRule) GetContextProcessingOk() (*bool, bool)`

GetContextProcessingOk returns a tuple with the ContextProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextProcessing

`func (o *RegularExpressionValidationRule) SetContextProcessing(v bool)`

SetContextProcessing sets ContextProcessing field to given value.

### HasContextProcessing

`func (o *RegularExpressionValidationRule) HasContextProcessing() bool`

HasContextProcessing returns a boolean if a field has been set.

### GetCreated

`func (o *RegularExpressionValidationRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RegularExpressionValidationRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RegularExpressionValidationRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *RegularExpressionValidationRule) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *RegularExpressionValidationRule) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *RegularExpressionValidationRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RegularExpressionValidationRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RegularExpressionValidationRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *RegularExpressionValidationRule) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *RegularExpressionValidationRule) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *RegularExpressionValidationRule) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RegularExpressionValidationRule) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RegularExpressionValidationRule) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RegularExpressionValidationRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *RegularExpressionValidationRule) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *RegularExpressionValidationRule) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *RegularExpressionValidationRule) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *RegularExpressionValidationRule) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RegularExpressionValidationRule) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RegularExpressionValidationRule) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RegularExpressionValidationRule) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


