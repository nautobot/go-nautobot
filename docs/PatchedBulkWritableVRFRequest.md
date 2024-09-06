# PatchedBulkWritableVRFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Namespace** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ImportTargets** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ExportTargets** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVRFRequest

`func NewPatchedBulkWritableVRFRequest(id string, ) *PatchedBulkWritableVRFRequest`

NewPatchedBulkWritableVRFRequest instantiates a new PatchedBulkWritableVRFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVRFRequestWithDefaults

`func NewPatchedBulkWritableVRFRequestWithDefaults() *PatchedBulkWritableVRFRequest`

NewPatchedBulkWritableVRFRequestWithDefaults instantiates a new PatchedBulkWritableVRFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVRFRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVRFRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVRFRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableVRFRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVRFRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVRFRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVRFRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRd

`func (o *PatchedBulkWritableVRFRequest) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *PatchedBulkWritableVRFRequest) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *PatchedBulkWritableVRFRequest) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *PatchedBulkWritableVRFRequest) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *PatchedBulkWritableVRFRequest) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *PatchedBulkWritableVRFRequest) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *PatchedBulkWritableVRFRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVRFRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVRFRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVRFRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableVRFRequest) GetStatus() BulkWritableCircuitRequestTenant`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableVRFRequest) GetStatusOk() (*BulkWritableCircuitRequestTenant, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableVRFRequest) SetStatus(v BulkWritableCircuitRequestTenant)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableVRFRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedBulkWritableVRFRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedBulkWritableVRFRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *PatchedBulkWritableVRFRequest) GetNamespace() BulkWritableCableRequestStatus`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedBulkWritableVRFRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedBulkWritableVRFRequest) SetNamespace(v BulkWritableCableRequestStatus)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedBulkWritableVRFRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWritableVRFRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVRFRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVRFRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVRFRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVRFRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVRFRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetImportTargets

`func (o *PatchedBulkWritableVRFRequest) GetImportTargets() []BulkWritableCableRequestStatus`

GetImportTargets returns the ImportTargets field if non-nil, zero value otherwise.

### GetImportTargetsOk

`func (o *PatchedBulkWritableVRFRequest) GetImportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetImportTargetsOk returns a tuple with the ImportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTargets

`func (o *PatchedBulkWritableVRFRequest) SetImportTargets(v []BulkWritableCableRequestStatus)`

SetImportTargets sets ImportTargets field to given value.

### HasImportTargets

`func (o *PatchedBulkWritableVRFRequest) HasImportTargets() bool`

HasImportTargets returns a boolean if a field has been set.

### GetExportTargets

`func (o *PatchedBulkWritableVRFRequest) GetExportTargets() []BulkWritableCableRequestStatus`

GetExportTargets returns the ExportTargets field if non-nil, zero value otherwise.

### GetExportTargetsOk

`func (o *PatchedBulkWritableVRFRequest) GetExportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetExportTargetsOk returns a tuple with the ExportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTargets

`func (o *PatchedBulkWritableVRFRequest) SetExportTargets(v []BulkWritableCableRequestStatus)`

SetExportTargets sets ExportTargets field to given value.

### HasExportTargets

`func (o *PatchedBulkWritableVRFRequest) HasExportTargets() bool`

HasExportTargets returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVRFRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVRFRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVRFRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVRFRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableVRFRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVRFRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVRFRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVRFRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVRFRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVRFRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVRFRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVRFRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


