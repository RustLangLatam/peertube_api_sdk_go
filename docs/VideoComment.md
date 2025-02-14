# VideoComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Text of the comment | [optional] 
**ThreadId** | Pointer to **int32** |  | [optional] 
**InReplyToCommentId** | Pointer to **NullableInt32** |  | [optional] 
**VideoId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] [default to false]
**HeldForReview** | Pointer to **bool** |  | [optional] 
**TotalRepliesFromVideoAuthor** | Pointer to **int32** |  | [optional] 
**TotalReplies** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewVideoComment

`func NewVideoComment() *VideoComment`

NewVideoComment instantiates a new VideoComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoCommentWithDefaults

`func NewVideoCommentWithDefaults() *VideoComment`

NewVideoCommentWithDefaults instantiates a new VideoComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoComment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *VideoComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoComment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoComment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetText

`func (o *VideoComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *VideoComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *VideoComment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *VideoComment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreadId

`func (o *VideoComment) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *VideoComment) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *VideoComment) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *VideoComment) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetInReplyToCommentId

`func (o *VideoComment) GetInReplyToCommentId() int32`

GetInReplyToCommentId returns the InReplyToCommentId field if non-nil, zero value otherwise.

### GetInReplyToCommentIdOk

`func (o *VideoComment) GetInReplyToCommentIdOk() (*int32, bool)`

GetInReplyToCommentIdOk returns a tuple with the InReplyToCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToCommentId

`func (o *VideoComment) SetInReplyToCommentId(v int32)`

SetInReplyToCommentId sets InReplyToCommentId field to given value.

### HasInReplyToCommentId

`func (o *VideoComment) HasInReplyToCommentId() bool`

HasInReplyToCommentId returns a boolean if a field has been set.

### SetInReplyToCommentIdNil

`func (o *VideoComment) SetInReplyToCommentIdNil(b bool)`

 SetInReplyToCommentIdNil sets the value for InReplyToCommentId to be an explicit nil

### UnsetInReplyToCommentId
`func (o *VideoComment) UnsetInReplyToCommentId()`

UnsetInReplyToCommentId ensures that no value is present for InReplyToCommentId, not even an explicit nil
### GetVideoId

`func (o *VideoComment) GetVideoId() int32`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *VideoComment) GetVideoIdOk() (*int32, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *VideoComment) SetVideoId(v int32)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *VideoComment) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoComment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *VideoComment) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *VideoComment) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *VideoComment) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *VideoComment) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *VideoComment) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *VideoComment) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetIsDeleted

`func (o *VideoComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *VideoComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *VideoComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *VideoComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetHeldForReview

`func (o *VideoComment) GetHeldForReview() bool`

GetHeldForReview returns the HeldForReview field if non-nil, zero value otherwise.

### GetHeldForReviewOk

`func (o *VideoComment) GetHeldForReviewOk() (*bool, bool)`

GetHeldForReviewOk returns a tuple with the HeldForReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldForReview

`func (o *VideoComment) SetHeldForReview(v bool)`

SetHeldForReview sets HeldForReview field to given value.

### HasHeldForReview

`func (o *VideoComment) HasHeldForReview() bool`

HasHeldForReview returns a boolean if a field has been set.

### GetTotalRepliesFromVideoAuthor

`func (o *VideoComment) GetTotalRepliesFromVideoAuthor() int32`

GetTotalRepliesFromVideoAuthor returns the TotalRepliesFromVideoAuthor field if non-nil, zero value otherwise.

### GetTotalRepliesFromVideoAuthorOk

`func (o *VideoComment) GetTotalRepliesFromVideoAuthorOk() (*int32, bool)`

GetTotalRepliesFromVideoAuthorOk returns a tuple with the TotalRepliesFromVideoAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepliesFromVideoAuthor

`func (o *VideoComment) SetTotalRepliesFromVideoAuthor(v int32)`

SetTotalRepliesFromVideoAuthor sets TotalRepliesFromVideoAuthor field to given value.

### HasTotalRepliesFromVideoAuthor

`func (o *VideoComment) HasTotalRepliesFromVideoAuthor() bool`

HasTotalRepliesFromVideoAuthor returns a boolean if a field has been set.

### GetTotalReplies

`func (o *VideoComment) GetTotalReplies() int32`

GetTotalReplies returns the TotalReplies field if non-nil, zero value otherwise.

### GetTotalRepliesOk

`func (o *VideoComment) GetTotalRepliesOk() (*int32, bool)`

GetTotalRepliesOk returns a tuple with the TotalReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReplies

`func (o *VideoComment) SetTotalReplies(v int32)`

SetTotalReplies sets TotalReplies field to given value.

### HasTotalReplies

`func (o *VideoComment) HasTotalReplies() bool`

HasTotalReplies returns a boolean if a field has been set.

### GetAccount

`func (o *VideoComment) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VideoComment) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VideoComment) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VideoComment) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


