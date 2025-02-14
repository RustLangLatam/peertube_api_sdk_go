# NotificationComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ThreadId** | Pointer to **int32** |  | [optional] 
**Video** | Pointer to [**VideoInfo**](VideoInfo.md) |  | [optional] 
**Account** | Pointer to [**ActorInfo**](ActorInfo.md) |  | [optional] 
**HeldForReview** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotificationComment

`func NewNotificationComment() *NotificationComment`

NewNotificationComment instantiates a new NotificationComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationCommentWithDefaults

`func NewNotificationCommentWithDefaults() *NotificationComment`

NewNotificationCommentWithDefaults instantiates a new NotificationComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationComment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetThreadId

`func (o *NotificationComment) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *NotificationComment) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *NotificationComment) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *NotificationComment) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetVideo

`func (o *NotificationComment) GetVideo() VideoInfo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *NotificationComment) GetVideoOk() (*VideoInfo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *NotificationComment) SetVideo(v VideoInfo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *NotificationComment) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAccount

`func (o *NotificationComment) GetAccount() ActorInfo`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NotificationComment) GetAccountOk() (*ActorInfo, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NotificationComment) SetAccount(v ActorInfo)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NotificationComment) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetHeldForReview

`func (o *NotificationComment) GetHeldForReview() bool`

GetHeldForReview returns the HeldForReview field if non-nil, zero value otherwise.

### GetHeldForReviewOk

`func (o *NotificationComment) GetHeldForReviewOk() (*bool, bool)`

GetHeldForReviewOk returns a tuple with the HeldForReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldForReview

`func (o *NotificationComment) SetHeldForReview(v bool)`

SetHeldForReview sets HeldForReview field to given value.

### HasHeldForReview

`func (o *NotificationComment) HasHeldForReview() bool`

HasHeldForReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


