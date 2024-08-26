# WritableCableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminationAType** | **string** |  | 
**TerminationBType** | **string** |  | 
**TerminationAId** | **string** |  | 
**TerminationBId** | **string** |  | 
**Type** | Pointer to [**PatchedWritableCableRequestType**](PatchedWritableCableRequestType.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** | RGB color in hexadecimal (e.g. 00ff00) | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**LengthUnit** | Pointer to [**PatchedWritableCableRequestLengthUnit**](PatchedWritableCableRequestLengthUnit.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewWritableCableRequest

`func NewWritableCableRequest(terminationAType string, terminationBType string, terminationAId string, terminationBId string, status BulkWritableCableRequestStatus, ) *WritableCableRequest`

NewWritableCableRequest instantiates a new WritableCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCableRequestWithDefaults

`func NewWritableCableRequestWithDefaults() *WritableCableRequest`

NewWritableCableRequestWithDefaults instantiates a new WritableCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminationAType

`func (o *WritableCableRequest) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *WritableCableRequest) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *WritableCableRequest) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.


### GetTerminationBType

`func (o *WritableCableRequest) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *WritableCableRequest) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *WritableCableRequest) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.


### GetTerminationAId

`func (o *WritableCableRequest) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *WritableCableRequest) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *WritableCableRequest) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.


### GetTerminationBId

`func (o *WritableCableRequest) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *WritableCableRequest) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *WritableCableRequest) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.


### GetType

`func (o *WritableCableRequest) GetType() PatchedWritableCableRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCableRequest) GetTypeOk() (*PatchedWritableCableRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCableRequest) SetType(v PatchedWritableCableRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabel

`func (o *WritableCableRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableCableRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableCableRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableCableRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *WritableCableRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WritableCableRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WritableCableRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WritableCableRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *WritableCableRequest) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WritableCableRequest) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WritableCableRequest) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *WritableCableRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *WritableCableRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *WritableCableRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *WritableCableRequest) GetLengthUnit() PatchedWritableCableRequestLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *WritableCableRequest) GetLengthUnitOk() (*PatchedWritableCableRequestLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *WritableCableRequest) SetLengthUnit(v PatchedWritableCableRequestLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *WritableCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetStatus

`func (o *WritableCableRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableCableRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableCableRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *WritableCableRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCableRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCableRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCableRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCableRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCableRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableCableRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableCableRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableCableRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableCableRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


