# VideoUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Video** | Pointer to [**VideoUploadResponseVideo**](VideoUploadResponseVideo.md) |  | [optional] 

## Methods

### NewVideoUploadResponse

`func NewVideoUploadResponse() *VideoUploadResponse`

NewVideoUploadResponse instantiates a new VideoUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoUploadResponseWithDefaults

`func NewVideoUploadResponseWithDefaults() *VideoUploadResponse`

NewVideoUploadResponseWithDefaults instantiates a new VideoUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideo

`func (o *VideoUploadResponse) GetVideo() VideoUploadResponseVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *VideoUploadResponse) GetVideoOk() (*VideoUploadResponseVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *VideoUploadResponse) SetVideo(v VideoUploadResponseVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *VideoUploadResponse) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


