# PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LoadBalancerPoolMember** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CertificateProfile** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest

`func NewPatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest() *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewPatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest instantiates a new PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults

`func NewPatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults() *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewPatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults instantiates a new PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancerPoolMember

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMember() BulkWritableCableRequestStatus`

GetLoadBalancerPoolMember returns the LoadBalancerPoolMember field if non-nil, zero value otherwise.

### GetLoadBalancerPoolMemberOk

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMemberOk() (*BulkWritableCableRequestStatus, bool)`

GetLoadBalancerPoolMemberOk returns a tuple with the LoadBalancerPoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPoolMember

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetLoadBalancerPoolMember(v BulkWritableCableRequestStatus)`

SetLoadBalancerPoolMember sets LoadBalancerPoolMember field to given value.

### HasLoadBalancerPoolMember

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) HasLoadBalancerPoolMember() bool`

HasLoadBalancerPoolMember returns a boolean if a field has been set.

### GetCertificateProfile

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.

### HasCertificateProfile

`func (o *PatchedLoadBalancerPoolMemberCertificateProfileAssignmentRequest) HasCertificateProfile() bool`

HasCertificateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


