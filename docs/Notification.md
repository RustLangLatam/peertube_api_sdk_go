# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** | Notification type, following the &#x60;UserNotificationType&#x60; enum: - &#x60;1&#x60; NEW_VIDEO_FROM_SUBSCRIPTION - &#x60;2&#x60; NEW_COMMENT_ON_MY_VIDEO - &#x60;3&#x60; NEW_ABUSE_FOR_MODERATORS - &#x60;4&#x60; BLACKLIST_ON_MY_VIDEO - &#x60;5&#x60; UNBLACKLIST_ON_MY_VIDEO - &#x60;6&#x60; MY_VIDEO_PUBLISHED - &#x60;7&#x60; MY_VIDEO_IMPORT_SUCCESS - &#x60;8&#x60; MY_VIDEO_IMPORT_ERROR - &#x60;9&#x60; NEW_USER_REGISTRATION - &#x60;10&#x60; NEW_FOLLOW - &#x60;11&#x60; COMMENT_MENTION - &#x60;12&#x60; VIDEO_AUTO_BLACKLIST_FOR_MODERATORS - &#x60;13&#x60; NEW_INSTANCE_FOLLOWER - &#x60;14&#x60; AUTO_INSTANCE_FOLLOWING - &#x60;15&#x60; ABUSE_STATE_CHANGE - &#x60;16&#x60; ABUSE_NEW_MESSAGE - &#x60;17&#x60; NEW_PLUGIN_VERSION - &#x60;18&#x60; NEW_PEERTUBE_VERSION - &#x60;19&#x60; MY_VIDEO_STUDIO_EDITION_FINISHED - &#x60;20&#x60; NEW_USER_REGISTRATION_REQUEST - &#x60;21&#x60; NEW_LIVE_FROM_SUBSCRIPTION  | [optional] 
**Read** | Pointer to **bool** |  | [optional] 
**Video** | Pointer to [**NullableNotificationVideo**](NotificationVideo.md) |  | [optional] 
**VideoImport** | Pointer to [**NullableNotificationVideoImport**](NotificationVideoImport.md) |  | [optional] 
**Comment** | Pointer to [**NullableNotificationComment**](NotificationComment.md) |  | [optional] 
**VideoAbuse** | Pointer to [**NullableNotificationVideoAbuse**](NotificationVideoAbuse.md) |  | [optional] 
**VideoBlacklist** | Pointer to [**NullableNotificationVideoAbuse**](NotificationVideoAbuse.md) |  | [optional] 
**Account** | Pointer to [**NullableActorInfo**](ActorInfo.md) |  | [optional] 
**ActorFollow** | Pointer to [**NullableNotificationActorFollow**](NotificationActorFollow.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Notification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Notification) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Notification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRead

`func (o *Notification) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *Notification) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *Notification) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *Notification) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetVideo

`func (o *Notification) GetVideo() NotificationVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *Notification) GetVideoOk() (*NotificationVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *Notification) SetVideo(v NotificationVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *Notification) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *Notification) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *Notification) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil
### GetVideoImport

`func (o *Notification) GetVideoImport() NotificationVideoImport`

GetVideoImport returns the VideoImport field if non-nil, zero value otherwise.

### GetVideoImportOk

`func (o *Notification) GetVideoImportOk() (*NotificationVideoImport, bool)`

GetVideoImportOk returns a tuple with the VideoImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoImport

`func (o *Notification) SetVideoImport(v NotificationVideoImport)`

SetVideoImport sets VideoImport field to given value.

### HasVideoImport

`func (o *Notification) HasVideoImport() bool`

HasVideoImport returns a boolean if a field has been set.

### SetVideoImportNil

`func (o *Notification) SetVideoImportNil(b bool)`

 SetVideoImportNil sets the value for VideoImport to be an explicit nil

### UnsetVideoImport
`func (o *Notification) UnsetVideoImport()`

UnsetVideoImport ensures that no value is present for VideoImport, not even an explicit nil
### GetComment

`func (o *Notification) GetComment() NotificationComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Notification) GetCommentOk() (*NotificationComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Notification) SetComment(v NotificationComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Notification) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Notification) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Notification) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetVideoAbuse

`func (o *Notification) GetVideoAbuse() NotificationVideoAbuse`

GetVideoAbuse returns the VideoAbuse field if non-nil, zero value otherwise.

### GetVideoAbuseOk

`func (o *Notification) GetVideoAbuseOk() (*NotificationVideoAbuse, bool)`

GetVideoAbuseOk returns a tuple with the VideoAbuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoAbuse

`func (o *Notification) SetVideoAbuse(v NotificationVideoAbuse)`

SetVideoAbuse sets VideoAbuse field to given value.

### HasVideoAbuse

`func (o *Notification) HasVideoAbuse() bool`

HasVideoAbuse returns a boolean if a field has been set.

### SetVideoAbuseNil

`func (o *Notification) SetVideoAbuseNil(b bool)`

 SetVideoAbuseNil sets the value for VideoAbuse to be an explicit nil

### UnsetVideoAbuse
`func (o *Notification) UnsetVideoAbuse()`

UnsetVideoAbuse ensures that no value is present for VideoAbuse, not even an explicit nil
### GetVideoBlacklist

`func (o *Notification) GetVideoBlacklist() NotificationVideoAbuse`

GetVideoBlacklist returns the VideoBlacklist field if non-nil, zero value otherwise.

### GetVideoBlacklistOk

`func (o *Notification) GetVideoBlacklistOk() (*NotificationVideoAbuse, bool)`

GetVideoBlacklistOk returns a tuple with the VideoBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBlacklist

`func (o *Notification) SetVideoBlacklist(v NotificationVideoAbuse)`

SetVideoBlacklist sets VideoBlacklist field to given value.

### HasVideoBlacklist

`func (o *Notification) HasVideoBlacklist() bool`

HasVideoBlacklist returns a boolean if a field has been set.

### SetVideoBlacklistNil

`func (o *Notification) SetVideoBlacklistNil(b bool)`

 SetVideoBlacklistNil sets the value for VideoBlacklist to be an explicit nil

### UnsetVideoBlacklist
`func (o *Notification) UnsetVideoBlacklist()`

UnsetVideoBlacklist ensures that no value is present for VideoBlacklist, not even an explicit nil
### GetAccount

`func (o *Notification) GetAccount() ActorInfo`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Notification) GetAccountOk() (*ActorInfo, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Notification) SetAccount(v ActorInfo)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Notification) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Notification) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Notification) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetActorFollow

`func (o *Notification) GetActorFollow() NotificationActorFollow`

GetActorFollow returns the ActorFollow field if non-nil, zero value otherwise.

### GetActorFollowOk

`func (o *Notification) GetActorFollowOk() (*NotificationActorFollow, bool)`

GetActorFollowOk returns a tuple with the ActorFollow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorFollow

`func (o *Notification) SetActorFollow(v NotificationActorFollow)`

SetActorFollow sets ActorFollow field to given value.

### HasActorFollow

`func (o *Notification) HasActorFollow() bool`

HasActorFollow returns a boolean if a field has been set.

### SetActorFollowNil

`func (o *Notification) SetActorFollowNil(b bool)`

 SetActorFollowNil sets the value for ActorFollow to be an explicit nil

### UnsetActorFollow
`func (o *Notification) UnsetActorFollow()`

UnsetActorFollow ensures that no value is present for ActorFollow, not even an explicit nil
### GetCreatedAt

`func (o *Notification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Notification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Notification) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Notification) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Notification) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Notification) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


