# UserViewingVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTime** | **int32** | timestamp within the video, in seconds | 
**ViewEvent** | Pointer to **string** | Event since last viewing call:  * &#x60;seek&#x60; - If the user seeked the video  | [optional] 
**SessionId** | Pointer to **string** | Optional param to represent the current viewer session. Used by the backend to properly count one view per session per video. PeerTube admin can configure the server to not trust this &#x60;sessionId&#x60; parameter but use the request IP address instead to identify a viewer.  | [optional] 

## Methods

### NewUserViewingVideo

`func NewUserViewingVideo(currentTime int32, ) *UserViewingVideo`

NewUserViewingVideo instantiates a new UserViewingVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserViewingVideoWithDefaults

`func NewUserViewingVideoWithDefaults() *UserViewingVideo`

NewUserViewingVideoWithDefaults instantiates a new UserViewingVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentTime

`func (o *UserViewingVideo) GetCurrentTime() int32`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *UserViewingVideo) GetCurrentTimeOk() (*int32, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *UserViewingVideo) SetCurrentTime(v int32)`

SetCurrentTime sets CurrentTime field to given value.


### GetViewEvent

`func (o *UserViewingVideo) GetViewEvent() string`

GetViewEvent returns the ViewEvent field if non-nil, zero value otherwise.

### GetViewEventOk

`func (o *UserViewingVideo) GetViewEventOk() (*string, bool)`

GetViewEventOk returns a tuple with the ViewEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewEvent

`func (o *UserViewingVideo) SetViewEvent(v string)`

SetViewEvent sets ViewEvent field to given value.

### HasViewEvent

`func (o *UserViewingVideo) HasViewEvent() bool`

HasViewEvent returns a boolean if a field has been set.

### GetSessionId

`func (o *UserViewingVideo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UserViewingVideo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UserViewingVideo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *UserViewingVideo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


