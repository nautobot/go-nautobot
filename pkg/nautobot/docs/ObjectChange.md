# ObjectChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Action** | [**ObjectChangeAction**](ObjectChangeAction.md) |  | 
**ChangedObjectType** | **string** |  | [readonly] 
**RelatedObjectType** | **string** |  | [readonly] 
**ChangedObject** | [**NullableObjectChangeChangedObject**](ObjectChangeChangedObject.md) |  | [readonly] 
**Time** | **time.Time** |  | [readonly] 
**UserName** | **string** |  | [readonly] 
**RequestId** | **string** |  | [readonly] 
**ChangedObjectId** | **string** |  | 
**ChangeContext** | **string** |  | [readonly] 
**ChangeContextDetail** | **string** |  | [readonly] 
**RelatedObjectId** | Pointer to **NullableString** |  | [optional] 
**ObjectRepr** | **string** |  | [readonly] 
**ObjectData** | **map[string]interface{}** |  | [readonly] 
**ObjectDataV2** | **map[string]interface{}** |  | [readonly] 
**User** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewObjectChange

`func NewObjectChange(id string, objectType string, display string, url string, naturalSlug string, action ObjectChangeAction, changedObjectType string, relatedObjectType string, changedObject NullableObjectChangeChangedObject, time time.Time, userName string, requestId string, changedObjectId string, changeContext string, changeContextDetail string, objectRepr string, objectData map[string]interface{}, objectDataV2 map[string]interface{}, ) *ObjectChange`

NewObjectChange instantiates a new ObjectChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectChangeWithDefaults

`func NewObjectChangeWithDefaults() *ObjectChange`

NewObjectChangeWithDefaults instantiates a new ObjectChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectChange) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *ObjectChange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectChange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectChange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ObjectChange) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectChange) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectChange) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ObjectChange) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectChange) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectChange) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ObjectChange) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ObjectChange) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ObjectChange) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetAction

`func (o *ObjectChange) GetAction() ObjectChangeAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ObjectChange) GetActionOk() (*ObjectChangeAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ObjectChange) SetAction(v ObjectChangeAction)`

SetAction sets Action field to given value.


### GetChangedObjectType

`func (o *ObjectChange) GetChangedObjectType() string`

GetChangedObjectType returns the ChangedObjectType field if non-nil, zero value otherwise.

### GetChangedObjectTypeOk

`func (o *ObjectChange) GetChangedObjectTypeOk() (*string, bool)`

GetChangedObjectTypeOk returns a tuple with the ChangedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectType

`func (o *ObjectChange) SetChangedObjectType(v string)`

SetChangedObjectType sets ChangedObjectType field to given value.


### GetRelatedObjectType

`func (o *ObjectChange) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *ObjectChange) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *ObjectChange) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.


### GetChangedObject

`func (o *ObjectChange) GetChangedObject() ObjectChangeChangedObject`

GetChangedObject returns the ChangedObject field if non-nil, zero value otherwise.

### GetChangedObjectOk

`func (o *ObjectChange) GetChangedObjectOk() (*ObjectChangeChangedObject, bool)`

GetChangedObjectOk returns a tuple with the ChangedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObject

`func (o *ObjectChange) SetChangedObject(v ObjectChangeChangedObject)`

SetChangedObject sets ChangedObject field to given value.


### SetChangedObjectNil

`func (o *ObjectChange) SetChangedObjectNil(b bool)`

 SetChangedObjectNil sets the value for ChangedObject to be an explicit nil

### UnsetChangedObject
`func (o *ObjectChange) UnsetChangedObject()`

UnsetChangedObject ensures that no value is present for ChangedObject, not even an explicit nil
### GetTime

`func (o *ObjectChange) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ObjectChange) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ObjectChange) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetUserName

`func (o *ObjectChange) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ObjectChange) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ObjectChange) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetRequestId

`func (o *ObjectChange) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ObjectChange) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ObjectChange) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetChangedObjectId

`func (o *ObjectChange) GetChangedObjectId() string`

GetChangedObjectId returns the ChangedObjectId field if non-nil, zero value otherwise.

### GetChangedObjectIdOk

`func (o *ObjectChange) GetChangedObjectIdOk() (*string, bool)`

GetChangedObjectIdOk returns a tuple with the ChangedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectId

`func (o *ObjectChange) SetChangedObjectId(v string)`

SetChangedObjectId sets ChangedObjectId field to given value.


### GetChangeContext

`func (o *ObjectChange) GetChangeContext() string`

GetChangeContext returns the ChangeContext field if non-nil, zero value otherwise.

### GetChangeContextOk

`func (o *ObjectChange) GetChangeContextOk() (*string, bool)`

GetChangeContextOk returns a tuple with the ChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeContext

`func (o *ObjectChange) SetChangeContext(v string)`

SetChangeContext sets ChangeContext field to given value.


### GetChangeContextDetail

`func (o *ObjectChange) GetChangeContextDetail() string`

GetChangeContextDetail returns the ChangeContextDetail field if non-nil, zero value otherwise.

### GetChangeContextDetailOk

`func (o *ObjectChange) GetChangeContextDetailOk() (*string, bool)`

GetChangeContextDetailOk returns a tuple with the ChangeContextDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeContextDetail

`func (o *ObjectChange) SetChangeContextDetail(v string)`

SetChangeContextDetail sets ChangeContextDetail field to given value.


### GetRelatedObjectId

`func (o *ObjectChange) GetRelatedObjectId() string`

GetRelatedObjectId returns the RelatedObjectId field if non-nil, zero value otherwise.

### GetRelatedObjectIdOk

`func (o *ObjectChange) GetRelatedObjectIdOk() (*string, bool)`

GetRelatedObjectIdOk returns a tuple with the RelatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectId

`func (o *ObjectChange) SetRelatedObjectId(v string)`

SetRelatedObjectId sets RelatedObjectId field to given value.

### HasRelatedObjectId

`func (o *ObjectChange) HasRelatedObjectId() bool`

HasRelatedObjectId returns a boolean if a field has been set.

### SetRelatedObjectIdNil

`func (o *ObjectChange) SetRelatedObjectIdNil(b bool)`

 SetRelatedObjectIdNil sets the value for RelatedObjectId to be an explicit nil

### UnsetRelatedObjectId
`func (o *ObjectChange) UnsetRelatedObjectId()`

UnsetRelatedObjectId ensures that no value is present for RelatedObjectId, not even an explicit nil
### GetObjectRepr

`func (o *ObjectChange) GetObjectRepr() string`

GetObjectRepr returns the ObjectRepr field if non-nil, zero value otherwise.

### GetObjectReprOk

`func (o *ObjectChange) GetObjectReprOk() (*string, bool)`

GetObjectReprOk returns a tuple with the ObjectRepr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRepr

`func (o *ObjectChange) SetObjectRepr(v string)`

SetObjectRepr sets ObjectRepr field to given value.


### GetObjectData

`func (o *ObjectChange) GetObjectData() map[string]interface{}`

GetObjectData returns the ObjectData field if non-nil, zero value otherwise.

### GetObjectDataOk

`func (o *ObjectChange) GetObjectDataOk() (*map[string]interface{}, bool)`

GetObjectDataOk returns a tuple with the ObjectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectData

`func (o *ObjectChange) SetObjectData(v map[string]interface{})`

SetObjectData sets ObjectData field to given value.


### GetObjectDataV2

`func (o *ObjectChange) GetObjectDataV2() map[string]interface{}`

GetObjectDataV2 returns the ObjectDataV2 field if non-nil, zero value otherwise.

### GetObjectDataV2Ok

`func (o *ObjectChange) GetObjectDataV2Ok() (*map[string]interface{}, bool)`

GetObjectDataV2Ok returns a tuple with the ObjectDataV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDataV2

`func (o *ObjectChange) SetObjectDataV2(v map[string]interface{})`

SetObjectDataV2 sets ObjectDataV2 field to given value.


### SetObjectDataV2Nil

`func (o *ObjectChange) SetObjectDataV2Nil(b bool)`

 SetObjectDataV2Nil sets the value for ObjectDataV2 to be an explicit nil

### UnsetObjectDataV2
`func (o *ObjectChange) UnsetObjectDataV2()`

UnsetObjectDataV2 ensures that no value is present for ObjectDataV2, not even an explicit nil
### GetUser

`func (o *ObjectChange) GetUser() BulkWritableCircuitRequestTenant`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectChange) GetUserOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectChange) SetUser(v BulkWritableCircuitRequestTenant)`

SetUser sets User field to given value.

### HasUser

`func (o *ObjectChange) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ObjectChange) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ObjectChange) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


