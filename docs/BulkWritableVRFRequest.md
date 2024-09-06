# BulkWritableVRFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableVRFRequest

`func NewBulkWritableVRFRequest(id string, name string, rd NullableString, ) *BulkWritableVRFRequest`

NewBulkWritableVRFRequest instantiates a new BulkWritableVRFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVRFRequestWithDefaults

`func NewBulkWritableVRFRequestWithDefaults() *BulkWritableVRFRequest`

NewBulkWritableVRFRequestWithDefaults instantiates a new BulkWritableVRFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVRFRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVRFRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVRFRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableVRFRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVRFRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVRFRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *BulkWritableVRFRequest) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *BulkWritableVRFRequest) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *BulkWritableVRFRequest) SetRd(v string)`

SetRd sets Rd field to given value.


### SetRdNil

`func (o *BulkWritableVRFRequest) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *BulkWritableVRFRequest) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *BulkWritableVRFRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableVRFRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableVRFRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableVRFRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *BulkWritableVRFRequest) GetStatus() BulkWritableCircuitRequestTenant`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableVRFRequest) GetStatusOk() (*BulkWritableCircuitRequestTenant, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableVRFRequest) SetStatus(v BulkWritableCircuitRequestTenant)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkWritableVRFRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BulkWritableVRFRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BulkWritableVRFRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *BulkWritableVRFRequest) GetNamespace() BulkWritableCableRequestStatus`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BulkWritableVRFRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BulkWritableVRFRequest) SetNamespace(v BulkWritableCableRequestStatus)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BulkWritableVRFRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *BulkWritableVRFRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVRFRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVRFRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVRFRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVRFRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVRFRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetImportTargets

`func (o *BulkWritableVRFRequest) GetImportTargets() []BulkWritableCableRequestStatus`

GetImportTargets returns the ImportTargets field if non-nil, zero value otherwise.

### GetImportTargetsOk

`func (o *BulkWritableVRFRequest) GetImportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetImportTargetsOk returns a tuple with the ImportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTargets

`func (o *BulkWritableVRFRequest) SetImportTargets(v []BulkWritableCableRequestStatus)`

SetImportTargets sets ImportTargets field to given value.

### HasImportTargets

`func (o *BulkWritableVRFRequest) HasImportTargets() bool`

HasImportTargets returns a boolean if a field has been set.

### GetExportTargets

`func (o *BulkWritableVRFRequest) GetExportTargets() []BulkWritableCableRequestStatus`

GetExportTargets returns the ExportTargets field if non-nil, zero value otherwise.

### GetExportTargetsOk

`func (o *BulkWritableVRFRequest) GetExportTargetsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetExportTargetsOk returns a tuple with the ExportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTargets

`func (o *BulkWritableVRFRequest) SetExportTargets(v []BulkWritableCableRequestStatus)`

SetExportTargets sets ExportTargets field to given value.

### HasExportTargets

`func (o *BulkWritableVRFRequest) HasExportTargets() bool`

HasExportTargets returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVRFRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVRFRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVRFRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVRFRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableVRFRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVRFRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVRFRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVRFRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVRFRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVRFRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVRFRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVRFRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


