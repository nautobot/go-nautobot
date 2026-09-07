# Cable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**TerminationAType** | Pointer to **NullableString** |  | [optional] 
**TerminationBType** | Pointer to **NullableString** |  | [optional] 
**TerminationA** | [**NullableCableTermination**](CableTermination.md) |  | [readonly] 
**TerminationB** | [**NullableCableTermination**](CableTermination.md) |  | [readonly] 
**TerminationAId** | Pointer to **NullableString** |  | [optional] 
**TerminationBId** | Pointer to **NullableString** |  | [optional] 
**LengthUnit** | Pointer to [**CableLengthUnit**](CableLengthUnit.md) |  | [optional] 
**Type** | Pointer to [**CableType**](CableType.md) |  | [optional] 
**TotalLanes** | **int32** |  | [readonly] 
**ConnectedLanes** | **int32** |  | [readonly] 
**Terminations** | [**map[string]CableTermination**](CableTermination.md) | Terminations on this cable, keyed by side (&#39;a&#39;/&#39;b&#39;) plus 1-indexed connector number. A standard cable has one connector per side (&#39;a1&#39;, &#39;b1&#39;); higher connector numbers (&#39;a2&#39;, &#39;b2&#39;, ...) appear only on breakout cables with multiple connectors per side. Each value is a brief representation (default depth) or the full nested serializer (&#x60;?depth&gt;&#x3D;1&#x60;) of the termination at that connector; polymorphic across Interface / CircuitTermination / ConsolePort / FrontPort / RearPort / PowerPort / PowerOutlet / PowerFeed. Uncabled connectors on breakout cables are represented as &#x60;null&#x60;. | [readonly] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** | RGB color in hexadecimal (e.g. 00ff00) | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**CableType** | Pointer to [**NullableBulkWritableCableRequestCableType**](BulkWritableCableRequestCableType.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewCable

`func NewCable(objectType string, display string, url string, naturalSlug string, terminationA NullableCableTermination, terminationB NullableCableTermination, totalLanes int32, connectedLanes int32, terminations map[string]*CableTermination, status BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Cable`

NewCable instantiates a new Cable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableWithDefaults

`func NewCableWithDefaults() *Cable`

NewCableWithDefaults instantiates a new Cable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *Cable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Cable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Cable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Cable) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Cable) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Cable) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Cable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Cable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Cable) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Cable) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Cable) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Cable) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetTerminationAType

`func (o *Cable) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *Cable) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *Cable) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.

### HasTerminationAType

`func (o *Cable) HasTerminationAType() bool`

HasTerminationAType returns a boolean if a field has been set.

### SetTerminationATypeNil

`func (o *Cable) SetTerminationATypeNil(b bool)`

 SetTerminationATypeNil sets the value for TerminationAType to be an explicit nil

### UnsetTerminationAType
`func (o *Cable) UnsetTerminationAType()`

UnsetTerminationAType ensures that no value is present for TerminationAType, not even an explicit nil
### GetTerminationBType

`func (o *Cable) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *Cable) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *Cable) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.

### HasTerminationBType

`func (o *Cable) HasTerminationBType() bool`

HasTerminationBType returns a boolean if a field has been set.

### SetTerminationBTypeNil

`func (o *Cable) SetTerminationBTypeNil(b bool)`

 SetTerminationBTypeNil sets the value for TerminationBType to be an explicit nil

### UnsetTerminationBType
`func (o *Cable) UnsetTerminationBType()`

UnsetTerminationBType ensures that no value is present for TerminationBType, not even an explicit nil
### GetTerminationA

`func (o *Cable) GetTerminationA() CableTermination`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *Cable) GetTerminationAOk() (*CableTermination, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *Cable) SetTerminationA(v CableTermination)`

SetTerminationA sets TerminationA field to given value.


### SetTerminationANil

`func (o *Cable) SetTerminationANil(b bool)`

 SetTerminationANil sets the value for TerminationA to be an explicit nil

### UnsetTerminationA
`func (o *Cable) UnsetTerminationA()`

UnsetTerminationA ensures that no value is present for TerminationA, not even an explicit nil
### GetTerminationB

`func (o *Cable) GetTerminationB() CableTermination`

GetTerminationB returns the TerminationB field if non-nil, zero value otherwise.

### GetTerminationBOk

`func (o *Cable) GetTerminationBOk() (*CableTermination, bool)`

GetTerminationBOk returns a tuple with the TerminationB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationB

`func (o *Cable) SetTerminationB(v CableTermination)`

SetTerminationB sets TerminationB field to given value.


### SetTerminationBNil

`func (o *Cable) SetTerminationBNil(b bool)`

 SetTerminationBNil sets the value for TerminationB to be an explicit nil

### UnsetTerminationB
`func (o *Cable) UnsetTerminationB()`

UnsetTerminationB ensures that no value is present for TerminationB, not even an explicit nil
### GetTerminationAId

`func (o *Cable) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *Cable) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *Cable) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.

### HasTerminationAId

`func (o *Cable) HasTerminationAId() bool`

HasTerminationAId returns a boolean if a field has been set.

### SetTerminationAIdNil

`func (o *Cable) SetTerminationAIdNil(b bool)`

 SetTerminationAIdNil sets the value for TerminationAId to be an explicit nil

### UnsetTerminationAId
`func (o *Cable) UnsetTerminationAId()`

UnsetTerminationAId ensures that no value is present for TerminationAId, not even an explicit nil
### GetTerminationBId

`func (o *Cable) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *Cable) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *Cable) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.

### HasTerminationBId

`func (o *Cable) HasTerminationBId() bool`

HasTerminationBId returns a boolean if a field has been set.

### SetTerminationBIdNil

`func (o *Cable) SetTerminationBIdNil(b bool)`

 SetTerminationBIdNil sets the value for TerminationBId to be an explicit nil

### UnsetTerminationBId
`func (o *Cable) UnsetTerminationBId()`

UnsetTerminationBId ensures that no value is present for TerminationBId, not even an explicit nil
### GetLengthUnit

`func (o *Cable) GetLengthUnit() CableLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *Cable) GetLengthUnitOk() (*CableLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *Cable) SetLengthUnit(v CableLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *Cable) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetType

`func (o *Cable) GetType() CableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cable) GetTypeOk() (*CableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cable) SetType(v CableType)`

SetType sets Type field to given value.

### HasType

`func (o *Cable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTotalLanes

`func (o *Cable) GetTotalLanes() int32`

GetTotalLanes returns the TotalLanes field if non-nil, zero value otherwise.

### GetTotalLanesOk

`func (o *Cable) GetTotalLanesOk() (*int32, bool)`

GetTotalLanesOk returns a tuple with the TotalLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLanes

`func (o *Cable) SetTotalLanes(v int32)`

SetTotalLanes sets TotalLanes field to given value.


### GetConnectedLanes

`func (o *Cable) GetConnectedLanes() int32`

GetConnectedLanes returns the ConnectedLanes field if non-nil, zero value otherwise.

### GetConnectedLanesOk

`func (o *Cable) GetConnectedLanesOk() (*int32, bool)`

GetConnectedLanesOk returns a tuple with the ConnectedLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedLanes

`func (o *Cable) SetConnectedLanes(v int32)`

SetConnectedLanes sets ConnectedLanes field to given value.


### GetTerminations

`func (o *Cable) GetTerminations() map[string]*CableTermination`

GetTerminations returns the Terminations field if non-nil, zero value otherwise.

### GetTerminationsOk

`func (o *Cable) GetTerminationsOk() (*map[string]*CableTermination, bool)`

GetTerminationsOk returns a tuple with the Terminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminations

`func (o *Cable) SetTerminations(v map[string]*CableTermination)`

SetTerminations sets Terminations field to given value.


### GetLabel

`func (o *Cable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Cable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Cable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Cable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *Cable) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Cable) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Cable) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Cable) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *Cable) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Cable) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Cable) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Cable) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *Cable) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *Cable) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetCableType

`func (o *Cable) GetCableType() BulkWritableCableRequestCableType`

GetCableType returns the CableType field if non-nil, zero value otherwise.

### GetCableTypeOk

`func (o *Cable) GetCableTypeOk() (*BulkWritableCableRequestCableType, bool)`

GetCableTypeOk returns a tuple with the CableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableType

`func (o *Cable) SetCableType(v BulkWritableCableRequestCableType)`

SetCableType sets CableType field to given value.

### HasCableType

`func (o *Cable) HasCableType() bool`

HasCableType returns a boolean if a field has been set.

### SetCableTypeNil

`func (o *Cable) SetCableTypeNil(b bool)`

 SetCableTypeNil sets the value for CableType to be an explicit nil

### UnsetCableType
`func (o *Cable) UnsetCableType()`

UnsetCableType ensures that no value is present for CableType, not even an explicit nil
### GetStatus

`func (o *Cable) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cable) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cable) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *Cable) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Cable) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Cable) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Cable) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Cable) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Cable) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Cable) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Cable) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Cable) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Cable) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Cable) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Cable) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Cable) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Cable) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Cable) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Cable) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Cable) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *Cable) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Cable) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Cable) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Cable) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


