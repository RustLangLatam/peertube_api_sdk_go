# OverviewVideosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]CategoryOverview**](CategoryOverview.md) |  | [optional] 
**Channels** | Pointer to [**[]ChannelOverview**](ChannelOverview.md) |  | [optional] 
**Tags** | Pointer to [**[]TagOverview**](TagOverview.md) |  | [optional] 

## Methods

### NewOverviewVideosResponse

`func NewOverviewVideosResponse() *OverviewVideosResponse`

NewOverviewVideosResponse instantiates a new OverviewVideosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewVideosResponseWithDefaults

`func NewOverviewVideosResponseWithDefaults() *OverviewVideosResponse`

NewOverviewVideosResponseWithDefaults instantiates a new OverviewVideosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *OverviewVideosResponse) GetCategories() []CategoryOverview`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *OverviewVideosResponse) GetCategoriesOk() (*[]CategoryOverview, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *OverviewVideosResponse) SetCategories(v []CategoryOverview)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *OverviewVideosResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetChannels

`func (o *OverviewVideosResponse) GetChannels() []ChannelOverview`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *OverviewVideosResponse) GetChannelsOk() (*[]ChannelOverview, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *OverviewVideosResponse) SetChannels(v []ChannelOverview)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *OverviewVideosResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetTags

`func (o *OverviewVideosResponse) GetTags() []TagOverview`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverviewVideosResponse) GetTagsOk() (*[]TagOverview, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverviewVideosResponse) SetTags(v []TagOverview)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverviewVideosResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


