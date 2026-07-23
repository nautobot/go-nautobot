# DataCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ComplianceClassName** | **string** |  | 
**LastValidationDate** | **time.Time** |  | [readonly] 
**ObjectId** | **string** |  | 
**ValidatedObjectStr** | Pointer to **string** |  | [optional] 
**ValidatedAttribute** | Pointer to **string** |  | [optional] [default to ""]
**ValidatedAttributeValue** | Pointer to **string** |  | [optional] 
**Valid** | **bool** |  | 
**Message** | Pointer to **string** |  | [optional] 
**ContentType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDataCompliance

`func NewDataCompliance(objectType string, display string, url string, naturalSlug string, complianceClassName string, lastValidationDate time.Time, objectId string, valid bool, contentType BulkWritableCableRequestStatus, notesUrl string, ) *DataCompliance`

NewDataCompliance instantiates a new DataCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataComplianceWithDefaults

`func NewDataComplianceWithDefaults() *DataCompliance`

NewDataComplianceWithDefaults instantiates a new DataCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataCompliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataCompliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataCompliance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataCompliance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *DataCompliance) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DataCompliance) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DataCompliance) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *DataCompliance) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DataCompliance) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DataCompliance) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *DataCompliance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DataCompliance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DataCompliance) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *DataCompliance) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *DataCompliance) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *DataCompliance) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetComplianceClassName

`func (o *DataCompliance) GetComplianceClassName() string`

GetComplianceClassName returns the ComplianceClassName field if non-nil, zero value otherwise.

### GetComplianceClassNameOk

`func (o *DataCompliance) GetComplianceClassNameOk() (*string, bool)`

GetComplianceClassNameOk returns a tuple with the ComplianceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceClassName

`func (o *DataCompliance) SetComplianceClassName(v string)`

SetComplianceClassName sets ComplianceClassName field to given value.


### GetLastValidationDate

`func (o *DataCompliance) GetLastValidationDate() time.Time`

GetLastValidationDate returns the LastValidationDate field if non-nil, zero value otherwise.

### GetLastValidationDateOk

`func (o *DataCompliance) GetLastValidationDateOk() (*time.Time, bool)`

GetLastValidationDateOk returns a tuple with the LastValidationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidationDate

`func (o *DataCompliance) SetLastValidationDate(v time.Time)`

SetLastValidationDate sets LastValidationDate field to given value.


### GetObjectId

`func (o *DataCompliance) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *DataCompliance) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *DataCompliance) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetValidatedObjectStr

`func (o *DataCompliance) GetValidatedObjectStr() string`

GetValidatedObjectStr returns the ValidatedObjectStr field if non-nil, zero value otherwise.

### GetValidatedObjectStrOk

`func (o *DataCompliance) GetValidatedObjectStrOk() (*string, bool)`

GetValidatedObjectStrOk returns a tuple with the ValidatedObjectStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedObjectStr

`func (o *DataCompliance) SetValidatedObjectStr(v string)`

SetValidatedObjectStr sets ValidatedObjectStr field to given value.

### HasValidatedObjectStr

`func (o *DataCompliance) HasValidatedObjectStr() bool`

HasValidatedObjectStr returns a boolean if a field has been set.

### GetValidatedAttribute

`func (o *DataCompliance) GetValidatedAttribute() string`

GetValidatedAttribute returns the ValidatedAttribute field if non-nil, zero value otherwise.

### GetValidatedAttributeOk

`func (o *DataCompliance) GetValidatedAttributeOk() (*string, bool)`

GetValidatedAttributeOk returns a tuple with the ValidatedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttribute

`func (o *DataCompliance) SetValidatedAttribute(v string)`

SetValidatedAttribute sets ValidatedAttribute field to given value.

### HasValidatedAttribute

`func (o *DataCompliance) HasValidatedAttribute() bool`

HasValidatedAttribute returns a boolean if a field has been set.

### GetValidatedAttributeValue

`func (o *DataCompliance) GetValidatedAttributeValue() string`

GetValidatedAttributeValue returns the ValidatedAttributeValue field if non-nil, zero value otherwise.

### GetValidatedAttributeValueOk

`func (o *DataCompliance) GetValidatedAttributeValueOk() (*string, bool)`

GetValidatedAttributeValueOk returns a tuple with the ValidatedAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttributeValue

`func (o *DataCompliance) SetValidatedAttributeValue(v string)`

SetValidatedAttributeValue sets ValidatedAttributeValue field to given value.

### HasValidatedAttributeValue

`func (o *DataCompliance) HasValidatedAttributeValue() bool`

HasValidatedAttributeValue returns a boolean if a field has been set.

### GetValid

`func (o *DataCompliance) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DataCompliance) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DataCompliance) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetMessage

`func (o *DataCompliance) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DataCompliance) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DataCompliance) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DataCompliance) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContentType

`func (o *DataCompliance) GetContentType() BulkWritableCableRequestStatus`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataCompliance) GetContentTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataCompliance) SetContentType(v BulkWritableCableRequestStatus)`

SetContentType sets ContentType field to given value.


### GetNotesUrl

`func (o *DataCompliance) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *DataCompliance) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *DataCompliance) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *DataCompliance) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DataCompliance) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DataCompliance) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DataCompliance) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


