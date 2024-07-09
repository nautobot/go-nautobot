# PatchedJobHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeCreate** | Pointer to **bool** | Call this job hook when a matching object is created. | [optional] 
**TypeDelete** | Pointer to **bool** | Call this job hook when a matching object is deleted. | [optional] 
**TypeUpdate** | Pointer to **bool** | Call this job hook when a matching object is updated. | [optional] 
**Job** | Pointer to [**BulkWritableJobHookRequestJob**](BulkWritableJobHookRequestJob.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedJobHookRequest

`func NewPatchedJobHookRequest() *PatchedJobHookRequest`

NewPatchedJobHookRequest instantiates a new PatchedJobHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJobHookRequestWithDefaults

`func NewPatchedJobHookRequestWithDefaults() *PatchedJobHookRequest`

NewPatchedJobHookRequestWithDefaults instantiates a new PatchedJobHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *PatchedJobHookRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedJobHookRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedJobHookRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedJobHookRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedJobHookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedJobHookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedJobHookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedJobHookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *PatchedJobHookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedJobHookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedJobHookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedJobHookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeCreate

`func (o *PatchedJobHookRequest) GetTypeCreate() bool`

GetTypeCreate returns the TypeCreate field if non-nil, zero value otherwise.

### GetTypeCreateOk

`func (o *PatchedJobHookRequest) GetTypeCreateOk() (*bool, bool)`

GetTypeCreateOk returns a tuple with the TypeCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCreate

`func (o *PatchedJobHookRequest) SetTypeCreate(v bool)`

SetTypeCreate sets TypeCreate field to given value.

### HasTypeCreate

`func (o *PatchedJobHookRequest) HasTypeCreate() bool`

HasTypeCreate returns a boolean if a field has been set.

### GetTypeDelete

`func (o *PatchedJobHookRequest) GetTypeDelete() bool`

GetTypeDelete returns the TypeDelete field if non-nil, zero value otherwise.

### GetTypeDeleteOk

`func (o *PatchedJobHookRequest) GetTypeDeleteOk() (*bool, bool)`

GetTypeDeleteOk returns a tuple with the TypeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDelete

`func (o *PatchedJobHookRequest) SetTypeDelete(v bool)`

SetTypeDelete sets TypeDelete field to given value.

### HasTypeDelete

`func (o *PatchedJobHookRequest) HasTypeDelete() bool`

HasTypeDelete returns a boolean if a field has been set.

### GetTypeUpdate

`func (o *PatchedJobHookRequest) GetTypeUpdate() bool`

GetTypeUpdate returns the TypeUpdate field if non-nil, zero value otherwise.

### GetTypeUpdateOk

`func (o *PatchedJobHookRequest) GetTypeUpdateOk() (*bool, bool)`

GetTypeUpdateOk returns a tuple with the TypeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUpdate

`func (o *PatchedJobHookRequest) SetTypeUpdate(v bool)`

SetTypeUpdate sets TypeUpdate field to given value.

### HasTypeUpdate

`func (o *PatchedJobHookRequest) HasTypeUpdate() bool`

HasTypeUpdate returns a boolean if a field has been set.

### GetJob

`func (o *PatchedJobHookRequest) GetJob() BulkWritableJobHookRequestJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *PatchedJobHookRequest) GetJobOk() (*BulkWritableJobHookRequestJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *PatchedJobHookRequest) SetJob(v BulkWritableJobHookRequestJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *PatchedJobHookRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedJobHookRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedJobHookRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedJobHookRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedJobHookRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedJobHookRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedJobHookRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedJobHookRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedJobHookRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


