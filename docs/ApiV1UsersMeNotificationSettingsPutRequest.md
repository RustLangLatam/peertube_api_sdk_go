# ApiV1UsersMeNotificationSettingsPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewVideoFromSubscription** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**NewCommentOnMyVideo** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**AbuseAsModerator** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**VideoAutoBlacklistAsModerator** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**BlacklistOnMyVideo** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**MyVideoPublished** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**MyVideoImportFinished** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**NewFollow** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**NewUserRegistration** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**CommentMention** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**NewInstanceFollower** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 
**AutoInstanceFollowing** | Pointer to **int32** | Notification type. One of the following values, or a sum of multiple values: - &#x60;0&#x60; NONE - &#x60;1&#x60; WEB - &#x60;2&#x60; EMAIL  | [optional] 

## Methods

### NewApiV1UsersMeNotificationSettingsPutRequest

`func NewApiV1UsersMeNotificationSettingsPutRequest() *ApiV1UsersMeNotificationSettingsPutRequest`

NewApiV1UsersMeNotificationSettingsPutRequest instantiates a new ApiV1UsersMeNotificationSettingsPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1UsersMeNotificationSettingsPutRequestWithDefaults

`func NewApiV1UsersMeNotificationSettingsPutRequestWithDefaults() *ApiV1UsersMeNotificationSettingsPutRequest`

NewApiV1UsersMeNotificationSettingsPutRequestWithDefaults instantiates a new ApiV1UsersMeNotificationSettingsPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewVideoFromSubscription

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewVideoFromSubscription() int32`

GetNewVideoFromSubscription returns the NewVideoFromSubscription field if non-nil, zero value otherwise.

### GetNewVideoFromSubscriptionOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewVideoFromSubscriptionOk() (*int32, bool)`

GetNewVideoFromSubscriptionOk returns a tuple with the NewVideoFromSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVideoFromSubscription

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewVideoFromSubscription(v int32)`

SetNewVideoFromSubscription sets NewVideoFromSubscription field to given value.

### HasNewVideoFromSubscription

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewVideoFromSubscription() bool`

HasNewVideoFromSubscription returns a boolean if a field has been set.

### GetNewCommentOnMyVideo

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewCommentOnMyVideo() int32`

GetNewCommentOnMyVideo returns the NewCommentOnMyVideo field if non-nil, zero value otherwise.

### GetNewCommentOnMyVideoOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewCommentOnMyVideoOk() (*int32, bool)`

GetNewCommentOnMyVideoOk returns a tuple with the NewCommentOnMyVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCommentOnMyVideo

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewCommentOnMyVideo(v int32)`

SetNewCommentOnMyVideo sets NewCommentOnMyVideo field to given value.

### HasNewCommentOnMyVideo

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewCommentOnMyVideo() bool`

HasNewCommentOnMyVideo returns a boolean if a field has been set.

### GetAbuseAsModerator

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAbuseAsModerator() int32`

GetAbuseAsModerator returns the AbuseAsModerator field if non-nil, zero value otherwise.

### GetAbuseAsModeratorOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAbuseAsModeratorOk() (*int32, bool)`

GetAbuseAsModeratorOk returns a tuple with the AbuseAsModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseAsModerator

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetAbuseAsModerator(v int32)`

SetAbuseAsModerator sets AbuseAsModerator field to given value.

### HasAbuseAsModerator

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasAbuseAsModerator() bool`

HasAbuseAsModerator returns a boolean if a field has been set.

### GetVideoAutoBlacklistAsModerator

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetVideoAutoBlacklistAsModerator() int32`

GetVideoAutoBlacklistAsModerator returns the VideoAutoBlacklistAsModerator field if non-nil, zero value otherwise.

### GetVideoAutoBlacklistAsModeratorOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetVideoAutoBlacklistAsModeratorOk() (*int32, bool)`

GetVideoAutoBlacklistAsModeratorOk returns a tuple with the VideoAutoBlacklistAsModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoAutoBlacklistAsModerator

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetVideoAutoBlacklistAsModerator(v int32)`

SetVideoAutoBlacklistAsModerator sets VideoAutoBlacklistAsModerator field to given value.

### HasVideoAutoBlacklistAsModerator

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasVideoAutoBlacklistAsModerator() bool`

HasVideoAutoBlacklistAsModerator returns a boolean if a field has been set.

### GetBlacklistOnMyVideo

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetBlacklistOnMyVideo() int32`

GetBlacklistOnMyVideo returns the BlacklistOnMyVideo field if non-nil, zero value otherwise.

### GetBlacklistOnMyVideoOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetBlacklistOnMyVideoOk() (*int32, bool)`

GetBlacklistOnMyVideoOk returns a tuple with the BlacklistOnMyVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistOnMyVideo

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetBlacklistOnMyVideo(v int32)`

SetBlacklistOnMyVideo sets BlacklistOnMyVideo field to given value.

### HasBlacklistOnMyVideo

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasBlacklistOnMyVideo() bool`

HasBlacklistOnMyVideo returns a boolean if a field has been set.

### GetMyVideoPublished

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoPublished() int32`

GetMyVideoPublished returns the MyVideoPublished field if non-nil, zero value otherwise.

### GetMyVideoPublishedOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoPublishedOk() (*int32, bool)`

GetMyVideoPublishedOk returns a tuple with the MyVideoPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyVideoPublished

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetMyVideoPublished(v int32)`

SetMyVideoPublished sets MyVideoPublished field to given value.

### HasMyVideoPublished

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasMyVideoPublished() bool`

HasMyVideoPublished returns a boolean if a field has been set.

### GetMyVideoImportFinished

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoImportFinished() int32`

GetMyVideoImportFinished returns the MyVideoImportFinished field if non-nil, zero value otherwise.

### GetMyVideoImportFinishedOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoImportFinishedOk() (*int32, bool)`

GetMyVideoImportFinishedOk returns a tuple with the MyVideoImportFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyVideoImportFinished

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetMyVideoImportFinished(v int32)`

SetMyVideoImportFinished sets MyVideoImportFinished field to given value.

### HasMyVideoImportFinished

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasMyVideoImportFinished() bool`

HasMyVideoImportFinished returns a boolean if a field has been set.

### GetNewFollow

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewFollow() int32`

GetNewFollow returns the NewFollow field if non-nil, zero value otherwise.

### GetNewFollowOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewFollowOk() (*int32, bool)`

GetNewFollowOk returns a tuple with the NewFollow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFollow

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewFollow(v int32)`

SetNewFollow sets NewFollow field to given value.

### HasNewFollow

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewFollow() bool`

HasNewFollow returns a boolean if a field has been set.

### GetNewUserRegistration

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewUserRegistration() int32`

GetNewUserRegistration returns the NewUserRegistration field if non-nil, zero value otherwise.

### GetNewUserRegistrationOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewUserRegistrationOk() (*int32, bool)`

GetNewUserRegistrationOk returns a tuple with the NewUserRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUserRegistration

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewUserRegistration(v int32)`

SetNewUserRegistration sets NewUserRegistration field to given value.

### HasNewUserRegistration

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewUserRegistration() bool`

HasNewUserRegistration returns a boolean if a field has been set.

### GetCommentMention

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetCommentMention() int32`

GetCommentMention returns the CommentMention field if non-nil, zero value otherwise.

### GetCommentMentionOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetCommentMentionOk() (*int32, bool)`

GetCommentMentionOk returns a tuple with the CommentMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentMention

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetCommentMention(v int32)`

SetCommentMention sets CommentMention field to given value.

### HasCommentMention

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasCommentMention() bool`

HasCommentMention returns a boolean if a field has been set.

### GetNewInstanceFollower

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewInstanceFollower() int32`

GetNewInstanceFollower returns the NewInstanceFollower field if non-nil, zero value otherwise.

### GetNewInstanceFollowerOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewInstanceFollowerOk() (*int32, bool)`

GetNewInstanceFollowerOk returns a tuple with the NewInstanceFollower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewInstanceFollower

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewInstanceFollower(v int32)`

SetNewInstanceFollower sets NewInstanceFollower field to given value.

### HasNewInstanceFollower

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewInstanceFollower() bool`

HasNewInstanceFollower returns a boolean if a field has been set.

### GetAutoInstanceFollowing

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAutoInstanceFollowing() int32`

GetAutoInstanceFollowing returns the AutoInstanceFollowing field if non-nil, zero value otherwise.

### GetAutoInstanceFollowingOk

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAutoInstanceFollowingOk() (*int32, bool)`

GetAutoInstanceFollowingOk returns a tuple with the AutoInstanceFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInstanceFollowing

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetAutoInstanceFollowing(v int32)`

SetAutoInstanceFollowing sets AutoInstanceFollowing field to given value.

### HasAutoInstanceFollowing

`func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasAutoInstanceFollowing() bool`

HasAutoInstanceFollowing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


