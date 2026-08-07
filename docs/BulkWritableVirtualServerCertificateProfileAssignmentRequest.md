# BulkWritableVirtualServerCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VirtualServer** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CertificateProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableVirtualServerCertificateProfileAssignmentRequest

`func NewBulkWritableVirtualServerCertificateProfileAssignmentRequest(id string, virtualServer BulkWritableCableRequestStatus, certificateProfile BulkWritableCableRequestStatus, ) *BulkWritableVirtualServerCertificateProfileAssignmentRequest`

NewBulkWritableVirtualServerCertificateProfileAssignmentRequest instantiates a new BulkWritableVirtualServerCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVirtualServerCertificateProfileAssignmentRequestWithDefaults

`func NewBulkWritableVirtualServerCertificateProfileAssignmentRequestWithDefaults() *BulkWritableVirtualServerCertificateProfileAssignmentRequest`

NewBulkWritableVirtualServerCertificateProfileAssignmentRequestWithDefaults instantiates a new BulkWritableVirtualServerCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVirtualServer

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) GetVirtualServer() BulkWritableCableRequestStatus`

GetVirtualServer returns the VirtualServer field if non-nil, zero value otherwise.

### GetVirtualServerOk

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) GetVirtualServerOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualServerOk returns a tuple with the VirtualServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServer

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) SetVirtualServer(v BulkWritableCableRequestStatus)`

SetVirtualServer sets VirtualServer field to given value.


### GetCertificateProfile

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *BulkWritableVirtualServerCertificateProfileAssignmentRequest) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


