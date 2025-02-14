# AddVideoPlaylistVideoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | [**AddVideoPlaylistVideoRequestVideoId**](AddVideoPlaylistVideoRequestVideoId.md) |  | 
**StartTimestamp** | Pointer to **int32** | Start the video at this specific timestamp | [optional] 
**StopTimestamp** | Pointer to **int32** | Stop the video at this specific timestamp | [optional] 

## Methods

### NewAddVideoPlaylistVideoRequest

`func NewAddVideoPlaylistVideoRequest(videoId AddVideoPlaylistVideoRequestVideoId, ) *AddVideoPlaylistVideoRequest`

NewAddVideoPlaylistVideoRequest instantiates a new AddVideoPlaylistVideoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVideoPlaylistVideoRequestWithDefaults

`func NewAddVideoPlaylistVideoRequestWithDefaults() *AddVideoPlaylistVideoRequest`

NewAddVideoPlaylistVideoRequestWithDefaults instantiates a new AddVideoPlaylistVideoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *AddVideoPlaylistVideoRequest) GetVideoId() AddVideoPlaylistVideoRequestVideoId`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *AddVideoPlaylistVideoRequest) GetVideoIdOk() (*AddVideoPlaylistVideoRequestVideoId, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *AddVideoPlaylistVideoRequest) SetVideoId(v AddVideoPlaylistVideoRequestVideoId)`

SetVideoId sets VideoId field to given value.


### GetStartTimestamp

`func (o *AddVideoPlaylistVideoRequest) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *AddVideoPlaylistVideoRequest) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *AddVideoPlaylistVideoRequest) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *AddVideoPlaylistVideoRequest) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetStopTimestamp

`func (o *AddVideoPlaylistVideoRequest) GetStopTimestamp() int32`

GetStopTimestamp returns the StopTimestamp field if non-nil, zero value otherwise.

### GetStopTimestampOk

`func (o *AddVideoPlaylistVideoRequest) GetStopTimestampOk() (*int32, bool)`

GetStopTimestampOk returns a tuple with the StopTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimestamp

`func (o *AddVideoPlaylistVideoRequest) SetStopTimestamp(v int32)`

SetStopTimestamp sets StopTimestamp field to given value.

### HasStopTimestamp

`func (o *AddVideoPlaylistVideoRequest) HasStopTimestamp() bool`

HasStopTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


