# LiveVideoSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** | Start date of the live session | [optional] 
**EndDate** | Pointer to **NullableTime** | End date of the live session | [optional] 
**Error** | Pointer to **NullableInt32** | Error type if an error occurred during the live session:   - &#x60;1&#x60;: Bad socket health (transcoding is too slow)   - &#x60;2&#x60;: Max duration exceeded   - &#x60;3&#x60;: Quota exceeded   - &#x60;4&#x60;: Quota FFmpeg error   - &#x60;5&#x60;: Video has been blacklisted during the live  | [optional] 
**ReplayVideo** | Pointer to [**LiveVideoSessionResponseReplayVideo**](LiveVideoSessionResponseReplayVideo.md) |  | [optional] 

## Methods

### NewLiveVideoSessionResponse

`func NewLiveVideoSessionResponse() *LiveVideoSessionResponse`

NewLiveVideoSessionResponse instantiates a new LiveVideoSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveVideoSessionResponseWithDefaults

`func NewLiveVideoSessionResponseWithDefaults() *LiveVideoSessionResponse`

NewLiveVideoSessionResponseWithDefaults instantiates a new LiveVideoSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveVideoSessionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveVideoSessionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveVideoSessionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LiveVideoSessionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartDate

`func (o *LiveVideoSessionResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LiveVideoSessionResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LiveVideoSessionResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LiveVideoSessionResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LiveVideoSessionResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LiveVideoSessionResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LiveVideoSessionResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LiveVideoSessionResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *LiveVideoSessionResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *LiveVideoSessionResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetError

`func (o *LiveVideoSessionResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LiveVideoSessionResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LiveVideoSessionResponse) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *LiveVideoSessionResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *LiveVideoSessionResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *LiveVideoSessionResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetReplayVideo

`func (o *LiveVideoSessionResponse) GetReplayVideo() LiveVideoSessionResponseReplayVideo`

GetReplayVideo returns the ReplayVideo field if non-nil, zero value otherwise.

### GetReplayVideoOk

`func (o *LiveVideoSessionResponse) GetReplayVideoOk() (*LiveVideoSessionResponseReplayVideo, bool)`

GetReplayVideoOk returns a tuple with the ReplayVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayVideo

`func (o *LiveVideoSessionResponse) SetReplayVideo(v LiveVideoSessionResponseReplayVideo)`

SetReplayVideo sets ReplayVideo field to given value.

### HasReplayVideo

`func (o *LiveVideoSessionResponse) HasReplayVideo() bool`

HasReplayVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


