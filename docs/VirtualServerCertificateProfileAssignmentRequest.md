# VirtualServerCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VirtualServer** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CertificateProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewVirtualServerCertificateProfileAssignmentRequest

`func NewVirtualServerCertificateProfileAssignmentRequest(virtualServer BulkWritableCableRequestStatus, certificateProfile BulkWritableCableRequestStatus, ) *VirtualServerCertificateProfileAssignmentRequest`

NewVirtualServerCertificateProfileAssignmentRequest instantiates a new VirtualServerCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualServerCertificateProfileAssignmentRequestWithDefaults

`func NewVirtualServerCertificateProfileAssignmentRequestWithDefaults() *VirtualServerCertificateProfileAssignmentRequest`

NewVirtualServerCertificateProfileAssignmentRequestWithDefaults instantiates a new VirtualServerCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualServerCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualServerCertificateProfileAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVirtualServer

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetVirtualServer() BulkWritableCableRequestStatus`

GetVirtualServer returns the VirtualServer field if non-nil, zero value otherwise.

### GetVirtualServerOk

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetVirtualServerOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualServerOk returns a tuple with the VirtualServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServer

`func (o *VirtualServerCertificateProfileAssignmentRequest) SetVirtualServer(v BulkWritableCableRequestStatus)`

SetVirtualServer sets VirtualServer field to given value.


### GetCertificateProfile

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *VirtualServerCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *VirtualServerCertificateProfileAssignmentRequest) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


