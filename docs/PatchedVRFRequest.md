# PatchedVRFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPatchedVRFRequest

`func NewPatchedVRFRequest() *PatchedVRFRequest`

NewPatchedVRFRequest instantiates a new PatchedVRFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVRFRequestWithDefaults

`func NewPatchedVRFRequestWithDefaults() *PatchedVRFRequest`

NewPatchedVRFRequestWithDefaults instantiates a new PatchedVRFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedVRFRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVRFRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVRFRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVRFRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRd

`func (o *PatchedVRFRequest) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *PatchedVRFRequest) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *PatchedVRFRequest) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *PatchedVRFRequest) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *PatchedVRFRequest) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *PatchedVRFRequest) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *PatchedVRFRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVRFRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVRFRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVRFRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedVRFRequest) GetStatus() BulkWritableCircuitRequestTenant`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedVRFRequest) GetStatusOk() (*BulkWritableCircuitRequestTenant, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedVRFRequest) SetStatus(v BulkWritableCircuitRequestTenant)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedVRFRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedVRFRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedVRFRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *PatchedVRFRequest) GetNamespace() BulkWritableCableRequestStatus`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedVRFRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedVRFRequest) SetNamespace(v BulkWritableCableRequestStatus)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedVRFRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedVRFRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedVRFRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedVRFRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedVRFRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedVRFRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedVRFRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetImportTargets

`func (o *PatchedVRFRequest) GetImportTargets() []BulkWritableCableRequestStatus`

GetImportTargets returns the ImportTargets field if non-nil, zero value otherwise.

### GetImportTargetsOk

`func (o *PatchedVRFRequest) GetImportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetImportTargetsOk returns a tuple with the ImportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTargets

`func (o *PatchedVRFRequest) SetImportTargets(v []BulkWritableCableRequestStatus)`

SetImportTargets sets ImportTargets field to given value.

### HasImportTargets

`func (o *PatchedVRFRequest) HasImportTargets() bool`

HasImportTargets returns a boolean if a field has been set.

### GetExportTargets

`func (o *PatchedVRFRequest) GetExportTargets() []BulkWritableCableRequestStatus`

GetExportTargets returns the ExportTargets field if non-nil, zero value otherwise.

### GetExportTargetsOk

`func (o *PatchedVRFRequest) GetExportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetExportTargetsOk returns a tuple with the ExportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTargets

`func (o *PatchedVRFRequest) SetExportTargets(v []BulkWritableCableRequestStatus)`

SetExportTargets sets ExportTargets field to given value.

### HasExportTargets

`func (o *PatchedVRFRequest) HasExportTargets() bool`

HasExportTargets returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVRFRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVRFRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVRFRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVRFRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVRFRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVRFRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVRFRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVRFRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVRFRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVRFRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVRFRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVRFRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


