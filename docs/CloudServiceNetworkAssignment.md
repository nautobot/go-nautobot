# CloudServiceNetworkAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**CloudNetwork** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CloudService** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewCloudServiceNetworkAssignment

`func NewCloudServiceNetworkAssignment(id string, objectType string, display string, url string, naturalSlug string, cloudNetwork BulkWritableCableRequestStatus, cloudService BulkWritableCableRequestStatus, ) *CloudServiceNetworkAssignment`

NewCloudServiceNetworkAssignment instantiates a new CloudServiceNetworkAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceNetworkAssignmentWithDefaults

`func NewCloudServiceNetworkAssignmentWithDefaults() *CloudServiceNetworkAssignment`

NewCloudServiceNetworkAssignmentWithDefaults instantiates a new CloudServiceNetworkAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudServiceNetworkAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudServiceNetworkAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudServiceNetworkAssignment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *CloudServiceNetworkAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudServiceNetworkAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudServiceNetworkAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *CloudServiceNetworkAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CloudServiceNetworkAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CloudServiceNetworkAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *CloudServiceNetworkAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudServiceNetworkAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudServiceNetworkAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *CloudServiceNetworkAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *CloudServiceNetworkAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *CloudServiceNetworkAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCloudNetwork

`func (o *CloudServiceNetworkAssignment) GetCloudNetwork() BulkWritableCableRequestStatus`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *CloudServiceNetworkAssignment) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *CloudServiceNetworkAssignment) SetCloudNetwork(v BulkWritableCableRequestStatus)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetCloudService

`func (o *CloudServiceNetworkAssignment) GetCloudService() BulkWritableCableRequestStatus`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *CloudServiceNetworkAssignment) GetCloudServiceOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *CloudServiceNetworkAssignment) SetCloudService(v BulkWritableCableRequestStatus)`

SetCloudService sets CloudService field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


