# VirtualServerCertificateProfileAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**VirtualServer** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CertificateProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewVirtualServerCertificateProfileAssignment

`func NewVirtualServerCertificateProfileAssignment(objectType string, display string, url string, naturalSlug string, virtualServer BulkWritableCableRequestStatus, certificateProfile BulkWritableCableRequestStatus, ) *VirtualServerCertificateProfileAssignment`

NewVirtualServerCertificateProfileAssignment instantiates a new VirtualServerCertificateProfileAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualServerCertificateProfileAssignmentWithDefaults

`func NewVirtualServerCertificateProfileAssignmentWithDefaults() *VirtualServerCertificateProfileAssignment`

NewVirtualServerCertificateProfileAssignmentWithDefaults instantiates a new VirtualServerCertificateProfileAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualServerCertificateProfileAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualServerCertificateProfileAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualServerCertificateProfileAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualServerCertificateProfileAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VirtualServerCertificateProfileAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualServerCertificateProfileAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualServerCertificateProfileAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VirtualServerCertificateProfileAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualServerCertificateProfileAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualServerCertificateProfileAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VirtualServerCertificateProfileAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualServerCertificateProfileAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualServerCertificateProfileAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VirtualServerCertificateProfileAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VirtualServerCertificateProfileAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VirtualServerCertificateProfileAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetVirtualServer

`func (o *VirtualServerCertificateProfileAssignment) GetVirtualServer() BulkWritableCableRequestStatus`

GetVirtualServer returns the VirtualServer field if non-nil, zero value otherwise.

### GetVirtualServerOk

`func (o *VirtualServerCertificateProfileAssignment) GetVirtualServerOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualServerOk returns a tuple with the VirtualServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServer

`func (o *VirtualServerCertificateProfileAssignment) SetVirtualServer(v BulkWritableCableRequestStatus)`

SetVirtualServer sets VirtualServer field to given value.


### GetCertificateProfile

`func (o *VirtualServerCertificateProfileAssignment) GetCertificateProfile() BulkWritableCableRequestStatus`

GetCertificateProfile returns the CertificateProfile field if non-nil, zero value otherwise.

### GetCertificateProfileOk

`func (o *VirtualServerCertificateProfileAssignment) GetCertificateProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetCertificateProfileOk returns a tuple with the CertificateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfile

`func (o *VirtualServerCertificateProfileAssignment) SetCertificateProfile(v BulkWritableCableRequestStatus)`

SetCertificateProfile sets CertificateProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


