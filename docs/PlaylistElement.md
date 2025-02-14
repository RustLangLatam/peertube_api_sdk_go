# PlaylistElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **int32** |  | [optional] 
**StartTimestamp** | Pointer to **int32** |  | [optional] 
**StopTimestamp** | Pointer to **int32** |  | [optional] 
**Video** | Pointer to [**NullableVideo**](Video.md) |  | [optional] 

## Methods

### NewPlaylistElement

`func NewPlaylistElement() *PlaylistElement`

NewPlaylistElement instantiates a new PlaylistElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistElementWithDefaults

`func NewPlaylistElementWithDefaults() *PlaylistElement`

NewPlaylistElementWithDefaults instantiates a new PlaylistElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *PlaylistElement) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PlaylistElement) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PlaylistElement) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PlaylistElement) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *PlaylistElement) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *PlaylistElement) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *PlaylistElement) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *PlaylistElement) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetStopTimestamp

`func (o *PlaylistElement) GetStopTimestamp() int32`

GetStopTimestamp returns the StopTimestamp field if non-nil, zero value otherwise.

### GetStopTimestampOk

`func (o *PlaylistElement) GetStopTimestampOk() (*int32, bool)`

GetStopTimestampOk returns a tuple with the StopTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimestamp

`func (o *PlaylistElement) SetStopTimestamp(v int32)`

SetStopTimestamp sets StopTimestamp field to given value.

### HasStopTimestamp

`func (o *PlaylistElement) HasStopTimestamp() bool`

HasStopTimestamp returns a boolean if a field has been set.

### GetVideo

`func (o *PlaylistElement) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PlaylistElement) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PlaylistElement) SetVideo(v Video)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *PlaylistElement) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *PlaylistElement) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *PlaylistElement) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


