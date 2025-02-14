# VideoRating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Video** | [**Video**](Video.md) |  | 
**Rating** | **string** | Rating of the video | 

## Methods

### NewVideoRating

`func NewVideoRating(video Video, rating string, ) *VideoRating`

NewVideoRating instantiates a new VideoRating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoRatingWithDefaults

`func NewVideoRatingWithDefaults() *VideoRating`

NewVideoRatingWithDefaults instantiates a new VideoRating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideo

`func (o *VideoRating) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *VideoRating) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *VideoRating) SetVideo(v Video)`

SetVideo sets Video field to given value.


### GetRating

`func (o *VideoRating) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *VideoRating) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *VideoRating) SetRating(v string)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


