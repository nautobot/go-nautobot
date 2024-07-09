# PatchedBulkWritableCableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TerminationAType** | Pointer to **string** |  | [optional] 
**TerminationBType** | Pointer to **string** |  | [optional] 
**LengthUnit** | Pointer to [**LengthUnitEnum**](LengthUnitEnum.md) |  | [optional] 
**Type** | Pointer to [**CableTypeChoices**](CableTypeChoices.md) |  | [optional] 
**TerminationAId** | Pointer to **string** |  | [optional] 
**TerminationBId** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** | RGB color in hexadecimal (e.g. 00ff00) | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableCableRequest

`func NewPatchedBulkWritableCableRequest(id string, ) *PatchedBulkWritableCableRequest`

NewPatchedBulkWritableCableRequest instantiates a new PatchedBulkWritableCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCableRequestWithDefaults

`func NewPatchedBulkWritableCableRequestWithDefaults() *PatchedBulkWritableCableRequest`

NewPatchedBulkWritableCableRequestWithDefaults instantiates a new PatchedBulkWritableCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCableRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCableRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCableRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTerminationAType

`func (o *PatchedBulkWritableCableRequest) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *PatchedBulkWritableCableRequest) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *PatchedBulkWritableCableRequest) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.

### HasTerminationAType

`func (o *PatchedBulkWritableCableRequest) HasTerminationAType() bool`

HasTerminationAType returns a boolean if a field has been set.

### GetTerminationBType

`func (o *PatchedBulkWritableCableRequest) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *PatchedBulkWritableCableRequest) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *PatchedBulkWritableCableRequest) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.

### HasTerminationBType

`func (o *PatchedBulkWritableCableRequest) HasTerminationBType() bool`

HasTerminationBType returns a boolean if a field has been set.

### GetLengthUnit

`func (o *PatchedBulkWritableCableRequest) GetLengthUnit() LengthUnitEnum`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *PatchedBulkWritableCableRequest) GetLengthUnitOk() (*LengthUnitEnum, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *PatchedBulkWritableCableRequest) SetLengthUnit(v LengthUnitEnum)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *PatchedBulkWritableCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkWritableCableRequest) GetType() CableTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkWritableCableRequest) GetTypeOk() (*CableTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkWritableCableRequest) SetType(v CableTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkWritableCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTerminationAId

`func (o *PatchedBulkWritableCableRequest) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *PatchedBulkWritableCableRequest) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *PatchedBulkWritableCableRequest) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.

### HasTerminationAId

`func (o *PatchedBulkWritableCableRequest) HasTerminationAId() bool`

HasTerminationAId returns a boolean if a field has been set.

### GetTerminationBId

`func (o *PatchedBulkWritableCableRequest) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *PatchedBulkWritableCableRequest) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *PatchedBulkWritableCableRequest) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.

### HasTerminationBId

`func (o *PatchedBulkWritableCableRequest) HasTerminationBId() bool`

HasTerminationBId returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkWritableCableRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkWritableCableRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkWritableCableRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkWritableCableRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *PatchedBulkWritableCableRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedBulkWritableCableRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedBulkWritableCableRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedBulkWritableCableRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *PatchedBulkWritableCableRequest) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PatchedBulkWritableCableRequest) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PatchedBulkWritableCableRequest) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PatchedBulkWritableCableRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *PatchedBulkWritableCableRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *PatchedBulkWritableCableRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetStatus

`func (o *PatchedBulkWritableCableRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableCableRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableCableRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableCableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableCableRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableCableRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableCableRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableCableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableCableRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableCableRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableCableRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableCableRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableCableRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableCableRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableCableRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


