# ChannelOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**VideoChannelSummary**](VideoChannelSummary.md) |  | [optional] 
**Videos** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewChannelOverview

`func NewChannelOverview() *ChannelOverview`

NewChannelOverview instantiates a new ChannelOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOverviewWithDefaults

`func NewChannelOverviewWithDefaults() *ChannelOverview`

NewChannelOverviewWithDefaults instantiates a new ChannelOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelOverview) GetChannel() VideoChannelSummary`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelOverview) GetChannelOk() (*VideoChannelSummary, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelOverview) SetChannel(v VideoChannelSummary)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelOverview) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetVideos

`func (o *ChannelOverview) GetVideos() []Video`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *ChannelOverview) GetVideosOk() (*[]Video, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *ChannelOverview) SetVideos(v []Video)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *ChannelOverview) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


