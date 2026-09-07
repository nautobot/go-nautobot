# WritableCableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TerminationAType** | Pointer to **NullableString** |  | [optional] 
**TerminationBType** | Pointer to **NullableString** |  | [optional] 
**TerminationAId** | Pointer to **NullableString** |  | [optional] 
**TerminationBId** | Pointer to **NullableString** |  | [optional] 
**LengthUnit** | Pointer to [**LengthUnitEnum**](LengthUnitEnum.md) |  | [optional] 
**Type** | Pointer to [**CableTypeChoices**](CableTypeChoices.md) |  | [optional] 
**Terminations** | Pointer to [**map[string]PatchedWritableCableRequestTerminationsValue**](PatchedWritableCableRequestTerminationsValue.md) | Terminations to apply, keyed by side (&#39;a&#39;/&#39;b&#39;) plus 1-indexed connector number. A standard cable has one connector per side (&#39;a1&#39;, &#39;b1&#39;); higher connector numbers (&#39;a2&#39;, &#39;b2&#39;, ...) apply only to breakout cables with multiple connectors per side. Each value is either &#x60;null&#x60; (delete the existing termination at this connector) or an &#x60;{\&quot;object_type\&quot;: \&quot;&lt;app.model&gt;\&quot;, \&quot;id\&quot;: \&quot;&lt;uuid&gt;\&quot;}&#x60; reference to the termination to plug in. Connectors omitted from the payload are left untouched (PATCH-style merge semantics). | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** | RGB color in hexadecimal (e.g. 00ff00) | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**CableType** | Pointer to [**NullableBulkWritableCableRequestCableType**](BulkWritableCableRequestCableType.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewWritableCableRequest

`func NewWritableCableRequest(status BulkWritableCableRequestStatus, ) *WritableCableRequest`

NewWritableCableRequest instantiates a new WritableCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCableRequestWithDefaults

`func NewWritableCableRequestWithDefaults() *WritableCableRequest`

NewWritableCableRequestWithDefaults instantiates a new WritableCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCableRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCableRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCableRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WritableCableRequest) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasTerminationAType

`func (o *WritableCableRequest) HasTerminationAType() bool`

HasTerminationAType returns a boolean if a field has been set.

### SetTerminationATypeNil

`func (o *WritableCableRequest) SetTerminationATypeNil(b bool)`

 SetTerminationATypeNil sets the value for TerminationAType to be an explicit nil

### UnsetTerminationAType
`func (o *WritableCableRequest) UnsetTerminationAType()`

UnsetTerminationAType ensures that no value is present for TerminationAType, not even an explicit nil
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

### HasTerminationBType

`func (o *WritableCableRequest) HasTerminationBType() bool`

HasTerminationBType returns a boolean if a field has been set.

### SetTerminationBTypeNil

`func (o *WritableCableRequest) SetTerminationBTypeNil(b bool)`

 SetTerminationBTypeNil sets the value for TerminationBType to be an explicit nil

### UnsetTerminationBType
`func (o *WritableCableRequest) UnsetTerminationBType()`

UnsetTerminationBType ensures that no value is present for TerminationBType, not even an explicit nil
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

### HasTerminationAId

`func (o *WritableCableRequest) HasTerminationAId() bool`

HasTerminationAId returns a boolean if a field has been set.

### SetTerminationAIdNil

`func (o *WritableCableRequest) SetTerminationAIdNil(b bool)`

 SetTerminationAIdNil sets the value for TerminationAId to be an explicit nil

### UnsetTerminationAId
`func (o *WritableCableRequest) UnsetTerminationAId()`

UnsetTerminationAId ensures that no value is present for TerminationAId, not even an explicit nil
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

### HasTerminationBId

`func (o *WritableCableRequest) HasTerminationBId() bool`

HasTerminationBId returns a boolean if a field has been set.

### SetTerminationBIdNil

`func (o *WritableCableRequest) SetTerminationBIdNil(b bool)`

 SetTerminationBIdNil sets the value for TerminationBId to be an explicit nil

### UnsetTerminationBId
`func (o *WritableCableRequest) UnsetTerminationBId()`

UnsetTerminationBId ensures that no value is present for TerminationBId, not even an explicit nil
### GetLengthUnit

`func (o *WritableCableRequest) GetLengthUnit() LengthUnitEnum`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *WritableCableRequest) GetLengthUnitOk() (*LengthUnitEnum, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *WritableCableRequest) SetLengthUnit(v LengthUnitEnum)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *WritableCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetType

`func (o *WritableCableRequest) GetType() CableTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCableRequest) GetTypeOk() (*CableTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCableRequest) SetType(v CableTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTerminations

`func (o *WritableCableRequest) GetTerminations() map[string]PatchedWritableCableRequestTerminationsValue`

GetTerminations returns the Terminations field if non-nil, zero value otherwise.

### GetTerminationsOk

`func (o *WritableCableRequest) GetTerminationsOk() (*map[string]PatchedWritableCableRequestTerminationsValue, bool)`

GetTerminationsOk returns a tuple with the Terminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminations

`func (o *WritableCableRequest) SetTerminations(v map[string]PatchedWritableCableRequestTerminationsValue)`

SetTerminations sets Terminations field to given value.

### HasTerminations

`func (o *WritableCableRequest) HasTerminations() bool`

HasTerminations returns a boolean if a field has been set.

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
### GetCableType

`func (o *WritableCableRequest) GetCableType() BulkWritableCableRequestCableType`

GetCableType returns the CableType field if non-nil, zero value otherwise.

### GetCableTypeOk

`func (o *WritableCableRequest) GetCableTypeOk() (*BulkWritableCableRequestCableType, bool)`

GetCableTypeOk returns a tuple with the CableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableType

`func (o *WritableCableRequest) SetCableType(v BulkWritableCableRequestCableType)`

SetCableType sets CableType field to given value.

### HasCableType

`func (o *WritableCableRequest) HasCableType() bool`

HasCableType returns a boolean if a field has been set.

### SetCableTypeNil

`func (o *WritableCableRequest) SetCableTypeNil(b bool)`

 SetCableTypeNil sets the value for CableType to be an explicit nil

### UnsetCableType
`func (o *WritableCableRequest) UnsetCableType()`

UnsetCableType ensures that no value is present for CableType, not even an explicit nil
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


### GetCustomFields

`func (o *WritableCableRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCableRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCableRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableCableRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableCableRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableCableRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableCableRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


