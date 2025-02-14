# TagOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** |  | [optional] 
**Videos** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewTagOverview

`func NewTagOverview() *TagOverview`

NewTagOverview instantiates a new TagOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagOverviewWithDefaults

`func NewTagOverviewWithDefaults() *TagOverview`

NewTagOverviewWithDefaults instantiates a new TagOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *TagOverview) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagOverview) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagOverview) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TagOverview) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVideos

`func (o *TagOverview) GetVideos() []Video`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *TagOverview) GetVideosOk() (*[]Video, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *TagOverview) SetVideos(v []Video)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *TagOverview) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


