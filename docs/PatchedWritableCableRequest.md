# PatchedWritableCableRequest

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
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedWritableCableRequest

`func NewPatchedWritableCableRequest() *PatchedWritableCableRequest`

NewPatchedWritableCableRequest instantiates a new PatchedWritableCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCableRequestWithDefaults

`func NewPatchedWritableCableRequestWithDefaults() *PatchedWritableCableRequest`

NewPatchedWritableCableRequestWithDefaults instantiates a new PatchedWritableCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableCableRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableCableRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableCableRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableCableRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTerminationAType

`func (o *PatchedWritableCableRequest) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *PatchedWritableCableRequest) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *PatchedWritableCableRequest) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.

### HasTerminationAType

`func (o *PatchedWritableCableRequest) HasTerminationAType() bool`

HasTerminationAType returns a boolean if a field has been set.

### SetTerminationATypeNil

`func (o *PatchedWritableCableRequest) SetTerminationATypeNil(b bool)`

 SetTerminationATypeNil sets the value for TerminationAType to be an explicit nil

### UnsetTerminationAType
`func (o *PatchedWritableCableRequest) UnsetTerminationAType()`

UnsetTerminationAType ensures that no value is present for TerminationAType, not even an explicit nil
### GetTerminationBType

`func (o *PatchedWritableCableRequest) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *PatchedWritableCableRequest) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *PatchedWritableCableRequest) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.

### HasTerminationBType

`func (o *PatchedWritableCableRequest) HasTerminationBType() bool`

HasTerminationBType returns a boolean if a field has been set.

### SetTerminationBTypeNil

`func (o *PatchedWritableCableRequest) SetTerminationBTypeNil(b bool)`

 SetTerminationBTypeNil sets the value for TerminationBType to be an explicit nil

### UnsetTerminationBType
`func (o *PatchedWritableCableRequest) UnsetTerminationBType()`

UnsetTerminationBType ensures that no value is present for TerminationBType, not even an explicit nil
### GetTerminationAId

`func (o *PatchedWritableCableRequest) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *PatchedWritableCableRequest) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *PatchedWritableCableRequest) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.

### HasTerminationAId

`func (o *PatchedWritableCableRequest) HasTerminationAId() bool`

HasTerminationAId returns a boolean if a field has been set.

### SetTerminationAIdNil

`func (o *PatchedWritableCableRequest) SetTerminationAIdNil(b bool)`

 SetTerminationAIdNil sets the value for TerminationAId to be an explicit nil

### UnsetTerminationAId
`func (o *PatchedWritableCableRequest) UnsetTerminationAId()`

UnsetTerminationAId ensures that no value is present for TerminationAId, not even an explicit nil
### GetTerminationBId

`func (o *PatchedWritableCableRequest) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *PatchedWritableCableRequest) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *PatchedWritableCableRequest) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.

### HasTerminationBId

`func (o *PatchedWritableCableRequest) HasTerminationBId() bool`

HasTerminationBId returns a boolean if a field has been set.

### SetTerminationBIdNil

`func (o *PatchedWritableCableRequest) SetTerminationBIdNil(b bool)`

 SetTerminationBIdNil sets the value for TerminationBId to be an explicit nil

### UnsetTerminationBId
`func (o *PatchedWritableCableRequest) UnsetTerminationBId()`

UnsetTerminationBId ensures that no value is present for TerminationBId, not even an explicit nil
### GetLengthUnit

`func (o *PatchedWritableCableRequest) GetLengthUnit() LengthUnitEnum`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *PatchedWritableCableRequest) GetLengthUnitOk() (*LengthUnitEnum, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *PatchedWritableCableRequest) SetLengthUnit(v LengthUnitEnum)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *PatchedWritableCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableCableRequest) GetType() CableTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableCableRequest) GetTypeOk() (*CableTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableCableRequest) SetType(v CableTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTerminations

`func (o *PatchedWritableCableRequest) GetTerminations() map[string]PatchedWritableCableRequestTerminationsValue`

GetTerminations returns the Terminations field if non-nil, zero value otherwise.

### GetTerminationsOk

`func (o *PatchedWritableCableRequest) GetTerminationsOk() (*map[string]PatchedWritableCableRequestTerminationsValue, bool)`

GetTerminationsOk returns a tuple with the Terminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminations

`func (o *PatchedWritableCableRequest) SetTerminations(v map[string]PatchedWritableCableRequestTerminationsValue)`

SetTerminations sets Terminations field to given value.

### HasTerminations

`func (o *PatchedWritableCableRequest) HasTerminations() bool`

HasTerminations returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableCableRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableCableRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableCableRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableCableRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *PatchedWritableCableRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedWritableCableRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedWritableCableRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedWritableCableRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *PatchedWritableCableRequest) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PatchedWritableCableRequest) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PatchedWritableCableRequest) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PatchedWritableCableRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *PatchedWritableCableRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *PatchedWritableCableRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetCableType

`func (o *PatchedWritableCableRequest) GetCableType() BulkWritableCableRequestCableType`

GetCableType returns the CableType field if non-nil, zero value otherwise.

### GetCableTypeOk

`func (o *PatchedWritableCableRequest) GetCableTypeOk() (*BulkWritableCableRequestCableType, bool)`

GetCableTypeOk returns a tuple with the CableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableType

`func (o *PatchedWritableCableRequest) SetCableType(v BulkWritableCableRequestCableType)`

SetCableType sets CableType field to given value.

### HasCableType

`func (o *PatchedWritableCableRequest) HasCableType() bool`

HasCableType returns a boolean if a field has been set.

### SetCableTypeNil

`func (o *PatchedWritableCableRequest) SetCableTypeNil(b bool)`

 SetCableTypeNil sets the value for CableType to be an explicit nil

### UnsetCableType
`func (o *PatchedWritableCableRequest) UnsetCableType()`

UnsetCableType ensures that no value is present for CableType, not even an explicit nil
### GetStatus

`func (o *PatchedWritableCableRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableCableRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableCableRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableCableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableCableRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableCableRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableCableRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableCableRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableCableRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableCableRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableCableRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableCableRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableCableRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableCableRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableCableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


