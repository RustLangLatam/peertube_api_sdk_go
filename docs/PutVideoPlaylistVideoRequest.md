# PutVideoPlaylistVideoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimestamp** | Pointer to **int32** | Start the video at this specific timestamp | [optional] 
**StopTimestamp** | Pointer to **int32** | Stop the video at this specific timestamp | [optional] 

## Methods

### NewPutVideoPlaylistVideoRequest

`func NewPutVideoPlaylistVideoRequest() *PutVideoPlaylistVideoRequest`

NewPutVideoPlaylistVideoRequest instantiates a new PutVideoPlaylistVideoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutVideoPlaylistVideoRequestWithDefaults

`func NewPutVideoPlaylistVideoRequestWithDefaults() *PutVideoPlaylistVideoRequest`

NewPutVideoPlaylistVideoRequestWithDefaults instantiates a new PutVideoPlaylistVideoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimestamp

`func (o *PutVideoPlaylistVideoRequest) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *PutVideoPlaylistVideoRequest) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *PutVideoPlaylistVideoRequest) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *PutVideoPlaylistVideoRequest) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetStopTimestamp

`func (o *PutVideoPlaylistVideoRequest) GetStopTimestamp() int32`

GetStopTimestamp returns the StopTimestamp field if non-nil, zero value otherwise.

### GetStopTimestampOk

`func (o *PutVideoPlaylistVideoRequest) GetStopTimestampOk() (*int32, bool)`

GetStopTimestampOk returns a tuple with the StopTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimestamp

`func (o *PutVideoPlaylistVideoRequest) SetStopTimestamp(v int32)`

SetStopTimestamp sets StopTimestamp field to given value.

### HasStopTimestamp

`func (o *PutVideoPlaylistVideoRequest) HasStopTimestamp() bool`

HasStopTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


