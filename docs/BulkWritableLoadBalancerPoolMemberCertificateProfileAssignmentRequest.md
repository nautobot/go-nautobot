# BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LoadBalancerPoolMember** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CertificateProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest

`func NewBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest(id string, loadBalancerPoolMember BulkWritableCableRequestStatus, certificateProfile BulkWritableCableRequestStatus, ) *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest instantiates a new BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults

`func NewBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults() *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest`

NewBulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequestWithDefaults instantiates a new BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLoadBalancerPoolMember

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMember() BulkWritableCableRequestStatus`

GetLoadBalancerPoolMember returns the LoadBalancerPoolMember field if non-nil, zero value otherwise.

### GetLoadBalancerPoolMemberOk

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetLoadBalancerPoolMemberOk() (*BulkWritableCableRequestStatus, bool)`

GetLoadBalancerPoolMemberOk returns a tuple with the LoadBalancerPoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPoolMember

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetLoadBalancerPoolMember(v BulkWritableCableRequestStatus)`

SetLoadBalancerPoolMember sets LoadBalancerPoolMember field to given value.


### GetCertificateProfile

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *BulkWritableLoadBalancerPoolMemberCertificateProfileAssignmentRequest) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


