# PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LoadBalancerPoolMember** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CertificateProfile** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest

`func NewPatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest(id string, ) *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewPatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest instantiates a new PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults

`func NewPatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults() *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewPatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLoadBalancerPoolMember

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMember() BulkWritableCableRequestStatus`

GetLoadBalancerPoolMember returns the LoadBalancerPoolMember field if non-nil, zero value otherwise.

### GetLoadBalancerPoolMemberOk

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMemberOk() (*BulkWritableCableRequestStatus, bool)`

GetLoadBalancerPoolMemberOk returns a tuple with the LoadBalancerPoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPoolMember

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetLoadBalancerPoolMember(v BulkWritableCableRequestStatus)`

SetLoadBalancerPoolMember sets LoadBalancerPoolMember field to given value.

### HasLoadBalancerPoolMember

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) HasLoadBalancerPoolMember() bool`

HasLoadBalancerPoolMember returns a boolean if a field has been set.

### GetCertificateProfile

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.

### HasCertificateProfile

`func (o *PatchedBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) HasCertificateProfile() bool`

HasCertificateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


