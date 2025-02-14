# VideoCommentForOwnerOrAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Text of the comment | [optional] 
**HeldForReview** | Pointer to **bool** |  | [optional] 
**ThreadId** | Pointer to **int32** |  | [optional] 
**InReplyToCommentId** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Video** | Pointer to [**VideoInfo**](VideoInfo.md) |  | [optional] 
**AutomaticTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVideoCommentForOwnerOrAdmin

`func NewVideoCommentForOwnerOrAdmin() *VideoCommentForOwnerOrAdmin`

NewVideoCommentForOwnerOrAdmin instantiates a new VideoCommentForOwnerOrAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoCommentForOwnerOrAdminWithDefaults

`func NewVideoCommentForOwnerOrAdminWithDefaults() *VideoCommentForOwnerOrAdmin`

NewVideoCommentForOwnerOrAdminWithDefaults instantiates a new VideoCommentForOwnerOrAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoCommentForOwnerOrAdmin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoCommentForOwnerOrAdmin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoCommentForOwnerOrAdmin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoCommentForOwnerOrAdmin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *VideoCommentForOwnerOrAdmin) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoCommentForOwnerOrAdmin) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoCommentForOwnerOrAdmin) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoCommentForOwnerOrAdmin) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetText

`func (o *VideoCommentForOwnerOrAdmin) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *VideoCommentForOwnerOrAdmin) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *VideoCommentForOwnerOrAdmin) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *VideoCommentForOwnerOrAdmin) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHeldForReview

`func (o *VideoCommentForOwnerOrAdmin) GetHeldForReview() bool`

GetHeldForReview returns the HeldForReview field if non-nil, zero value otherwise.

### GetHeldForReviewOk

`func (o *VideoCommentForOwnerOrAdmin) GetHeldForReviewOk() (*bool, bool)`

GetHeldForReviewOk returns a tuple with the HeldForReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldForReview

`func (o *VideoCommentForOwnerOrAdmin) SetHeldForReview(v bool)`

SetHeldForReview sets HeldForReview field to given value.

### HasHeldForReview

`func (o *VideoCommentForOwnerOrAdmin) HasHeldForReview() bool`

HasHeldForReview returns a boolean if a field has been set.

### GetThreadId

`func (o *VideoCommentForOwnerOrAdmin) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *VideoCommentForOwnerOrAdmin) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *VideoCommentForOwnerOrAdmin) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *VideoCommentForOwnerOrAdmin) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetInReplyToCommentId

`func (o *VideoCommentForOwnerOrAdmin) GetInReplyToCommentId() int32`

GetInReplyToCommentId returns the InReplyToCommentId field if non-nil, zero value otherwise.

### GetInReplyToCommentIdOk

`func (o *VideoCommentForOwnerOrAdmin) GetInReplyToCommentIdOk() (*int32, bool)`

GetInReplyToCommentIdOk returns a tuple with the InReplyToCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToCommentId

`func (o *VideoCommentForOwnerOrAdmin) SetInReplyToCommentId(v int32)`

SetInReplyToCommentId sets InReplyToCommentId field to given value.

### HasInReplyToCommentId

`func (o *VideoCommentForOwnerOrAdmin) HasInReplyToCommentId() bool`

HasInReplyToCommentId returns a boolean if a field has been set.

### SetInReplyToCommentIdNil

`func (o *VideoCommentForOwnerOrAdmin) SetInReplyToCommentIdNil(b bool)`

 SetInReplyToCommentIdNil sets the value for InReplyToCommentId to be an explicit nil

### UnsetInReplyToCommentId
`func (o *VideoCommentForOwnerOrAdmin) UnsetInReplyToCommentId()`

UnsetInReplyToCommentId ensures that no value is present for InReplyToCommentId, not even an explicit nil
### GetCreatedAt

`func (o *VideoCommentForOwnerOrAdmin) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoCommentForOwnerOrAdmin) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoCommentForOwnerOrAdmin) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoCommentForOwnerOrAdmin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoCommentForOwnerOrAdmin) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoCommentForOwnerOrAdmin) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoCommentForOwnerOrAdmin) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoCommentForOwnerOrAdmin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAccount

`func (o *VideoCommentForOwnerOrAdmin) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VideoCommentForOwnerOrAdmin) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VideoCommentForOwnerOrAdmin) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VideoCommentForOwnerOrAdmin) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVideo

`func (o *VideoCommentForOwnerOrAdmin) GetVideo() VideoInfo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *VideoCommentForOwnerOrAdmin) GetVideoOk() (*VideoInfo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *VideoCommentForOwnerOrAdmin) SetVideo(v VideoInfo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *VideoCommentForOwnerOrAdmin) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAutomaticTags

`func (o *VideoCommentForOwnerOrAdmin) GetAutomaticTags() []string`

GetAutomaticTags returns the AutomaticTags field if non-nil, zero value otherwise.

### GetAutomaticTagsOk

`func (o *VideoCommentForOwnerOrAdmin) GetAutomaticTagsOk() (*[]string, bool)`

GetAutomaticTagsOk returns a tuple with the AutomaticTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTags

`func (o *VideoCommentForOwnerOrAdmin) SetAutomaticTags(v []string)`

SetAutomaticTags sets AutomaticTags field to given value.

### HasAutomaticTags

`func (o *VideoCommentForOwnerOrAdmin) HasAutomaticTags() bool`

HasAutomaticTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


