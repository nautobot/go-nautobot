# PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VirtualServer** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CertificateProfile** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest

`func NewPatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest(id string, ) *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest`

NewPatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest instantiates a new PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVirtualServerCertificateProfileAssignmentRequestWithDefaults

`func NewPatchedBulkWritableVirtualServerCertificateProfileAssignmentRequestWithDefaults() *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest`

NewPatchedBulkWritableVirtualServerCertificateProfileAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVirtualServer

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) GetVirtualServer() BulkWritableCableRequestStatus`

GetVirtualServer returns the VirtualServer field if non-nil, zero value otherwise.

### GetVirtualServerOk

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) GetVirtualServerOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualServerOk returns a tuple with the VirtualServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServer

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) SetVirtualServer(v BulkWritableCableRequestStatus)`

SetVirtualServer sets VirtualServer field to given value.

### HasVirtualServer

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) HasVirtualServer() bool`

HasVirtualServer returns a boolean if a field has been set.

### GetCertificateProfile

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.

### HasCertificateProfile

`func (o *PatchedBulkWritableVirtualServerCertificateProfileAssignmentRequest) HasCertificateProfile() bool`

HasCertificateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


