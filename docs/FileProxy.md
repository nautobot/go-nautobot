# FileProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** |  | 
**UploadedAt** | **time.Time** |  | [readonly] 
**JobResult** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewFileProxy

`func NewFileProxy(id string, objectType string, display string, url string, naturalSlug string, name string, uploadedAt time.Time, ) *FileProxy`

NewFileProxy instantiates a new FileProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileProxyWithDefaults

`func NewFileProxyWithDefaults() *FileProxy`

NewFileProxyWithDefaults instantiates a new FileProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileProxy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileProxy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileProxy) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *FileProxy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FileProxy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FileProxy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *FileProxy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *FileProxy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *FileProxy) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *FileProxy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileProxy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileProxy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *FileProxy) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *FileProxy) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *FileProxy) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *FileProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileProxy) SetName(v string)`

SetName sets Name field to given value.


### GetUploadedAt

`func (o *FileProxy) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *FileProxy) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *FileProxy) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.


### GetJobResult

`func (o *FileProxy) GetJobResult() BulkWritableCircuitRequestTenant`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *FileProxy) GetJobResultOk() (*BulkWritableCircuitRequestTenant, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *FileProxy) SetJobResult(v BulkWritableCircuitRequestTenant)`

SetJobResult sets JobResult field to given value.

### HasJobResult

`func (o *FileProxy) HasJobResult() bool`

HasJobResult returns a boolean if a field has been set.

### SetJobResultNil

`func (o *FileProxy) SetJobResultNil(b bool)`

 SetJobResultNil sets the value for JobResult to be an explicit nil

### UnsetJobResult
`func (o *FileProxy) UnsetJobResult()`

UnsetJobResult ensures that no value is present for JobResult, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


