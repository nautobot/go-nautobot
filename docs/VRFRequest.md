# VRFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Rd** | **NullableString** | Unique route distinguisher (as defined in RFC 4364) | 
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

### NewVRFRequest

`func NewVRFRequest(name string, rd NullableString, ) *VRFRequest`

NewVRFRequest instantiates a new VRFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVRFRequestWithDefaults

`func NewVRFRequestWithDefaults() *VRFRequest`

NewVRFRequestWithDefaults instantiates a new VRFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VRFRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VRFRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VRFRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *VRFRequest) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *VRFRequest) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *VRFRequest) SetRd(v string)`

SetRd sets Rd field to given value.


### SetRdNil

`func (o *VRFRequest) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *VRFRequest) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *VRFRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VRFRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VRFRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VRFRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *VRFRequest) GetStatus() BulkWritableCircuitRequestTenant`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VRFRequest) GetStatusOk() (*BulkWritableCircuitRequestTenant, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VRFRequest) SetStatus(v BulkWritableCircuitRequestTenant)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VRFRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VRFRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VRFRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *VRFRequest) GetNamespace() BulkWritableCableRequestStatus`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VRFRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VRFRequest) SetNamespace(v BulkWritableCableRequestStatus)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *VRFRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *VRFRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VRFRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VRFRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VRFRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VRFRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VRFRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetImportTargets

`func (o *VRFRequest) GetImportTargets() []BulkWritableCableRequestStatus`

GetImportTargets returns the ImportTargets field if non-nil, zero value otherwise.

### GetImportTargetsOk

`func (o *VRFRequest) GetImportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetImportTargetsOk returns a tuple with the ImportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTargets

`func (o *VRFRequest) SetImportTargets(v []BulkWritableCableRequestStatus)`

SetImportTargets sets ImportTargets field to given value.

### HasImportTargets

`func (o *VRFRequest) HasImportTargets() bool`

HasImportTargets returns a boolean if a field has been set.

### GetExportTargets

`func (o *VRFRequest) GetExportTargets() []BulkWritableCableRequestStatus`

GetExportTargets returns the ExportTargets field if non-nil, zero value otherwise.

### GetExportTargetsOk

`func (o *VRFRequest) GetExportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetExportTargetsOk returns a tuple with the ExportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTargets

`func (o *VRFRequest) SetExportTargets(v []BulkWritableCableRequestStatus)`

SetExportTargets sets ExportTargets field to given value.

### HasExportTargets

`func (o *VRFRequest) HasExportTargets() bool`

HasExportTargets returns a boolean if a field has been set.

### GetTags

`func (o *VRFRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VRFRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VRFRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VRFRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VRFRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VRFRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VRFRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VRFRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VRFRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VRFRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VRFRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VRFRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


