# ServerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUsers** | Pointer to **float32** |  | [optional] 
**TotalDailyActiveUsers** | Pointer to **float32** |  | [optional] 
**TotalWeeklyActiveUsers** | Pointer to **float32** |  | [optional] 
**TotalMonthlyActiveUsers** | Pointer to **float32** |  | [optional] 
**TotalModerators** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled total moderators stats | [optional] 
**TotalAdmins** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled total admins stats | [optional] 
**TotalLocalVideos** | Pointer to **float32** |  | [optional] 
**TotalLocalVideoViews** | Pointer to **float32** | Total video views made on the instance | [optional] 
**TotalLocalVideoComments** | Pointer to **float32** | Total comments made by local users | [optional] 
**TotalLocalVideoFilesSize** | Pointer to **float32** |  | [optional] 
**TotalVideos** | Pointer to **float32** |  | [optional] 
**TotalVideoComments** | Pointer to **float32** |  | [optional] 
**TotalLocalVideoChannels** | Pointer to **float32** |  | [optional] 
**TotalLocalDailyActiveVideoChannels** | Pointer to **float32** |  | [optional] 
**TotalLocalWeeklyActiveVideoChannels** | Pointer to **float32** |  | [optional] 
**TotalLocalMonthlyActiveVideoChannels** | Pointer to **float32** |  | [optional] 
**TotalLocalPlaylists** | Pointer to **float32** |  | [optional] 
**TotalInstanceFollowers** | Pointer to **float32** |  | [optional] 
**TotalInstanceFollowing** | Pointer to **float32** |  | [optional] 
**VideosRedundancy** | Pointer to [**[]ServerStatsVideosRedundancyInner**](ServerStatsVideosRedundancyInner.md) |  | [optional] 
**TotalActivityPubMessagesProcessed** | Pointer to **float32** |  | [optional] 
**TotalActivityPubMessagesSuccesses** | Pointer to **float32** |  | [optional] 
**TotalActivityPubMessagesErrors** | Pointer to **float32** |  | [optional] 
**ActivityPubMessagesProcessedPerSecond** | Pointer to **float32** |  | [optional] 
**TotalActivityPubMessagesWaiting** | Pointer to **float32** |  | [optional] 
**AverageRegistrationRequestResponseTimeMs** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled registration requests stats | [optional] 
**TotalRegistrationRequestsProcessed** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled registration requests stats | [optional] 
**TotalRegistrationRequests** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled registration requests stats | [optional] 
**AverageAbuseResponseTimeMs** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled abuses stats | [optional] 
**TotalAbusesProcessed** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled abuses stats | [optional] 
**TotalAbuses** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Value is null if the admin disabled abuses stats | [optional] 

## Methods

### NewServerStats

`func NewServerStats() *ServerStats`

NewServerStats instantiates a new ServerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatsWithDefaults

`func NewServerStatsWithDefaults() *ServerStats`

NewServerStatsWithDefaults instantiates a new ServerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUsers

`func (o *ServerStats) GetTotalUsers() float32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *ServerStats) GetTotalUsersOk() (*float32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *ServerStats) SetTotalUsers(v float32)`

SetTotalUsers sets TotalUsers field to given value.

### HasTotalUsers

`func (o *ServerStats) HasTotalUsers() bool`

HasTotalUsers returns a boolean if a field has been set.

### GetTotalDailyActiveUsers

`func (o *ServerStats) GetTotalDailyActiveUsers() float32`

GetTotalDailyActiveUsers returns the TotalDailyActiveUsers field if non-nil, zero value otherwise.

### GetTotalDailyActiveUsersOk

`func (o *ServerStats) GetTotalDailyActiveUsersOk() (*float32, bool)`

GetTotalDailyActiveUsersOk returns a tuple with the TotalDailyActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDailyActiveUsers

`func (o *ServerStats) SetTotalDailyActiveUsers(v float32)`

SetTotalDailyActiveUsers sets TotalDailyActiveUsers field to given value.

### HasTotalDailyActiveUsers

`func (o *ServerStats) HasTotalDailyActiveUsers() bool`

HasTotalDailyActiveUsers returns a boolean if a field has been set.

### GetTotalWeeklyActiveUsers

`func (o *ServerStats) GetTotalWeeklyActiveUsers() float32`

GetTotalWeeklyActiveUsers returns the TotalWeeklyActiveUsers field if non-nil, zero value otherwise.

### GetTotalWeeklyActiveUsersOk

`func (o *ServerStats) GetTotalWeeklyActiveUsersOk() (*float32, bool)`

GetTotalWeeklyActiveUsersOk returns a tuple with the TotalWeeklyActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeeklyActiveUsers

`func (o *ServerStats) SetTotalWeeklyActiveUsers(v float32)`

SetTotalWeeklyActiveUsers sets TotalWeeklyActiveUsers field to given value.

### HasTotalWeeklyActiveUsers

`func (o *ServerStats) HasTotalWeeklyActiveUsers() bool`

HasTotalWeeklyActiveUsers returns a boolean if a field has been set.

### GetTotalMonthlyActiveUsers

`func (o *ServerStats) GetTotalMonthlyActiveUsers() float32`

GetTotalMonthlyActiveUsers returns the TotalMonthlyActiveUsers field if non-nil, zero value otherwise.

### GetTotalMonthlyActiveUsersOk

`func (o *ServerStats) GetTotalMonthlyActiveUsersOk() (*float32, bool)`

GetTotalMonthlyActiveUsersOk returns a tuple with the TotalMonthlyActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthlyActiveUsers

`func (o *ServerStats) SetTotalMonthlyActiveUsers(v float32)`

SetTotalMonthlyActiveUsers sets TotalMonthlyActiveUsers field to given value.

### HasTotalMonthlyActiveUsers

`func (o *ServerStats) HasTotalMonthlyActiveUsers() bool`

HasTotalMonthlyActiveUsers returns a boolean if a field has been set.

### GetTotalModerators

`func (o *ServerStats) GetTotalModerators() float32`

GetTotalModerators returns the TotalModerators field if non-nil, zero value otherwise.

### GetTotalModeratorsOk

`func (o *ServerStats) GetTotalModeratorsOk() (*float32, bool)`

GetTotalModeratorsOk returns a tuple with the TotalModerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalModerators

`func (o *ServerStats) SetTotalModerators(v float32)`

SetTotalModerators sets TotalModerators field to given value.

### HasTotalModerators

`func (o *ServerStats) HasTotalModerators() bool`

HasTotalModerators returns a boolean if a field has been set.

### GetTotalAdmins

`func (o *ServerStats) GetTotalAdmins() float32`

GetTotalAdmins returns the TotalAdmins field if non-nil, zero value otherwise.

### GetTotalAdminsOk

`func (o *ServerStats) GetTotalAdminsOk() (*float32, bool)`

GetTotalAdminsOk returns a tuple with the TotalAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAdmins

`func (o *ServerStats) SetTotalAdmins(v float32)`

SetTotalAdmins sets TotalAdmins field to given value.

### HasTotalAdmins

`func (o *ServerStats) HasTotalAdmins() bool`

HasTotalAdmins returns a boolean if a field has been set.

### GetTotalLocalVideos

`func (o *ServerStats) GetTotalLocalVideos() float32`

GetTotalLocalVideos returns the TotalLocalVideos field if non-nil, zero value otherwise.

### GetTotalLocalVideosOk

`func (o *ServerStats) GetTotalLocalVideosOk() (*float32, bool)`

GetTotalLocalVideosOk returns a tuple with the TotalLocalVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalVideos

`func (o *ServerStats) SetTotalLocalVideos(v float32)`

SetTotalLocalVideos sets TotalLocalVideos field to given value.

### HasTotalLocalVideos

`func (o *ServerStats) HasTotalLocalVideos() bool`

HasTotalLocalVideos returns a boolean if a field has been set.

### GetTotalLocalVideoViews

`func (o *ServerStats) GetTotalLocalVideoViews() float32`

GetTotalLocalVideoViews returns the TotalLocalVideoViews field if non-nil, zero value otherwise.

### GetTotalLocalVideoViewsOk

`func (o *ServerStats) GetTotalLocalVideoViewsOk() (*float32, bool)`

GetTotalLocalVideoViewsOk returns a tuple with the TotalLocalVideoViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalVideoViews

`func (o *ServerStats) SetTotalLocalVideoViews(v float32)`

SetTotalLocalVideoViews sets TotalLocalVideoViews field to given value.

### HasTotalLocalVideoViews

`func (o *ServerStats) HasTotalLocalVideoViews() bool`

HasTotalLocalVideoViews returns a boolean if a field has been set.

### GetTotalLocalVideoComments

`func (o *ServerStats) GetTotalLocalVideoComments() float32`

GetTotalLocalVideoComments returns the TotalLocalVideoComments field if non-nil, zero value otherwise.

### GetTotalLocalVideoCommentsOk

`func (o *ServerStats) GetTotalLocalVideoCommentsOk() (*float32, bool)`

GetTotalLocalVideoCommentsOk returns a tuple with the TotalLocalVideoComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalVideoComments

`func (o *ServerStats) SetTotalLocalVideoComments(v float32)`

SetTotalLocalVideoComments sets TotalLocalVideoComments field to given value.

### HasTotalLocalVideoComments

`func (o *ServerStats) HasTotalLocalVideoComments() bool`

HasTotalLocalVideoComments returns a boolean if a field has been set.

### GetTotalLocalVideoFilesSize

`func (o *ServerStats) GetTotalLocalVideoFilesSize() float32`

GetTotalLocalVideoFilesSize returns the TotalLocalVideoFilesSize field if non-nil, zero value otherwise.

### GetTotalLocalVideoFilesSizeOk

`func (o *ServerStats) GetTotalLocalVideoFilesSizeOk() (*float32, bool)`

GetTotalLocalVideoFilesSizeOk returns a tuple with the TotalLocalVideoFilesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalVideoFilesSize

`func (o *ServerStats) SetTotalLocalVideoFilesSize(v float32)`

SetTotalLocalVideoFilesSize sets TotalLocalVideoFilesSize field to given value.

### HasTotalLocalVideoFilesSize

`func (o *ServerStats) HasTotalLocalVideoFilesSize() bool`

HasTotalLocalVideoFilesSize returns a boolean if a field has been set.

### GetTotalVideos

`func (o *ServerStats) GetTotalVideos() float32`

GetTotalVideos returns the TotalVideos field if non-nil, zero value otherwise.

### GetTotalVideosOk

`func (o *ServerStats) GetTotalVideosOk() (*float32, bool)`

GetTotalVideosOk returns a tuple with the TotalVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVideos

`func (o *ServerStats) SetTotalVideos(v float32)`

SetTotalVideos sets TotalVideos field to given value.

### HasTotalVideos

`func (o *ServerStats) HasTotalVideos() bool`

HasTotalVideos returns a boolean if a field has been set.

### GetTotalVideoComments

`func (o *ServerStats) GetTotalVideoComments() float32`

GetTotalVideoComments returns the TotalVideoComments field if non-nil, zero value otherwise.

### GetTotalVideoCommentsOk

`func (o *ServerStats) GetTotalVideoCommentsOk() (*float32, bool)`

GetTotalVideoCommentsOk returns a tuple with the TotalVideoComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVideoComments

`func (o *ServerStats) SetTotalVideoComments(v float32)`

SetTotalVideoComments sets TotalVideoComments field to given value.

### HasTotalVideoComments

`func (o *ServerStats) HasTotalVideoComments() bool`

HasTotalVideoComments returns a boolean if a field has been set.

### GetTotalLocalVideoChannels

`func (o *ServerStats) GetTotalLocalVideoChannels() float32`

GetTotalLocalVideoChannels returns the TotalLocalVideoChannels field if non-nil, zero value otherwise.

### GetTotalLocalVideoChannelsOk

`func (o *ServerStats) GetTotalLocalVideoChannelsOk() (*float32, bool)`

GetTotalLocalVideoChannelsOk returns a tuple with the TotalLocalVideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalVideoChannels

`func (o *ServerStats) SetTotalLocalVideoChannels(v float32)`

SetTotalLocalVideoChannels sets TotalLocalVideoChannels field to given value.

### HasTotalLocalVideoChannels

`func (o *ServerStats) HasTotalLocalVideoChannels() bool`

HasTotalLocalVideoChannels returns a boolean if a field has been set.

### GetTotalLocalDailyActiveVideoChannels

`func (o *ServerStats) GetTotalLocalDailyActiveVideoChannels() float32`

GetTotalLocalDailyActiveVideoChannels returns the TotalLocalDailyActiveVideoChannels field if non-nil, zero value otherwise.

### GetTotalLocalDailyActiveVideoChannelsOk

`func (o *ServerStats) GetTotalLocalDailyActiveVideoChannelsOk() (*float32, bool)`

GetTotalLocalDailyActiveVideoChannelsOk returns a tuple with the TotalLocalDailyActiveVideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalDailyActiveVideoChannels

`func (o *ServerStats) SetTotalLocalDailyActiveVideoChannels(v float32)`

SetTotalLocalDailyActiveVideoChannels sets TotalLocalDailyActiveVideoChannels field to given value.

### HasTotalLocalDailyActiveVideoChannels

`func (o *ServerStats) HasTotalLocalDailyActiveVideoChannels() bool`

HasTotalLocalDailyActiveVideoChannels returns a boolean if a field has been set.

### GetTotalLocalWeeklyActiveVideoChannels

`func (o *ServerStats) GetTotalLocalWeeklyActiveVideoChannels() float32`

GetTotalLocalWeeklyActiveVideoChannels returns the TotalLocalWeeklyActiveVideoChannels field if non-nil, zero value otherwise.

### GetTotalLocalWeeklyActiveVideoChannelsOk

`func (o *ServerStats) GetTotalLocalWeeklyActiveVideoChannelsOk() (*float32, bool)`

GetTotalLocalWeeklyActiveVideoChannelsOk returns a tuple with the TotalLocalWeeklyActiveVideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalWeeklyActiveVideoChannels

`func (o *ServerStats) SetTotalLocalWeeklyActiveVideoChannels(v float32)`

SetTotalLocalWeeklyActiveVideoChannels sets TotalLocalWeeklyActiveVideoChannels field to given value.

### HasTotalLocalWeeklyActiveVideoChannels

`func (o *ServerStats) HasTotalLocalWeeklyActiveVideoChannels() bool`

HasTotalLocalWeeklyActiveVideoChannels returns a boolean if a field has been set.

### GetTotalLocalMonthlyActiveVideoChannels

`func (o *ServerStats) GetTotalLocalMonthlyActiveVideoChannels() float32`

GetTotalLocalMonthlyActiveVideoChannels returns the TotalLocalMonthlyActiveVideoChannels field if non-nil, zero value otherwise.

### GetTotalLocalMonthlyActiveVideoChannelsOk

`func (o *ServerStats) GetTotalLocalMonthlyActiveVideoChannelsOk() (*float32, bool)`

GetTotalLocalMonthlyActiveVideoChannelsOk returns a tuple with the TotalLocalMonthlyActiveVideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalMonthlyActiveVideoChannels

`func (o *ServerStats) SetTotalLocalMonthlyActiveVideoChannels(v float32)`

SetTotalLocalMonthlyActiveVideoChannels sets TotalLocalMonthlyActiveVideoChannels field to given value.

### HasTotalLocalMonthlyActiveVideoChannels

`func (o *ServerStats) HasTotalLocalMonthlyActiveVideoChannels() bool`

HasTotalLocalMonthlyActiveVideoChannels returns a boolean if a field has been set.

### GetTotalLocalPlaylists

`func (o *ServerStats) GetTotalLocalPlaylists() float32`

GetTotalLocalPlaylists returns the TotalLocalPlaylists field if non-nil, zero value otherwise.

### GetTotalLocalPlaylistsOk

`func (o *ServerStats) GetTotalLocalPlaylistsOk() (*float32, bool)`

GetTotalLocalPlaylistsOk returns a tuple with the TotalLocalPlaylists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalPlaylists

`func (o *ServerStats) SetTotalLocalPlaylists(v float32)`

SetTotalLocalPlaylists sets TotalLocalPlaylists field to given value.

### HasTotalLocalPlaylists

`func (o *ServerStats) HasTotalLocalPlaylists() bool`

HasTotalLocalPlaylists returns a boolean if a field has been set.

### GetTotalInstanceFollowers

`func (o *ServerStats) GetTotalInstanceFollowers() float32`

GetTotalInstanceFollowers returns the TotalInstanceFollowers field if non-nil, zero value otherwise.

### GetTotalInstanceFollowersOk

`func (o *ServerStats) GetTotalInstanceFollowersOk() (*float32, bool)`

GetTotalInstanceFollowersOk returns a tuple with the TotalInstanceFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstanceFollowers

`func (o *ServerStats) SetTotalInstanceFollowers(v float32)`

SetTotalInstanceFollowers sets TotalInstanceFollowers field to given value.

### HasTotalInstanceFollowers

`func (o *ServerStats) HasTotalInstanceFollowers() bool`

HasTotalInstanceFollowers returns a boolean if a field has been set.

### GetTotalInstanceFollowing

`func (o *ServerStats) GetTotalInstanceFollowing() float32`

GetTotalInstanceFollowing returns the TotalInstanceFollowing field if non-nil, zero value otherwise.

### GetTotalInstanceFollowingOk

`func (o *ServerStats) GetTotalInstanceFollowingOk() (*float32, bool)`

GetTotalInstanceFollowingOk returns a tuple with the TotalInstanceFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstanceFollowing

`func (o *ServerStats) SetTotalInstanceFollowing(v float32)`

SetTotalInstanceFollowing sets TotalInstanceFollowing field to given value.

### HasTotalInstanceFollowing

`func (o *ServerStats) HasTotalInstanceFollowing() bool`

HasTotalInstanceFollowing returns a boolean if a field has been set.

### GetVideosRedundancy

`func (o *ServerStats) GetVideosRedundancy() []ServerStatsVideosRedundancyInner`

GetVideosRedundancy returns the VideosRedundancy field if non-nil, zero value otherwise.

### GetVideosRedundancyOk

`func (o *ServerStats) GetVideosRedundancyOk() (*[]ServerStatsVideosRedundancyInner, bool)`

GetVideosRedundancyOk returns a tuple with the VideosRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideosRedundancy

`func (o *ServerStats) SetVideosRedundancy(v []ServerStatsVideosRedundancyInner)`

SetVideosRedundancy sets VideosRedundancy field to given value.

### HasVideosRedundancy

`func (o *ServerStats) HasVideosRedundancy() bool`

HasVideosRedundancy returns a boolean if a field has been set.

### GetTotalActivityPubMessagesProcessed

`func (o *ServerStats) GetTotalActivityPubMessagesProcessed() float32`

GetTotalActivityPubMessagesProcessed returns the TotalActivityPubMessagesProcessed field if non-nil, zero value otherwise.

### GetTotalActivityPubMessagesProcessedOk

`func (o *ServerStats) GetTotalActivityPubMessagesProcessedOk() (*float32, bool)`

GetTotalActivityPubMessagesProcessedOk returns a tuple with the TotalActivityPubMessagesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivityPubMessagesProcessed

`func (o *ServerStats) SetTotalActivityPubMessagesProcessed(v float32)`

SetTotalActivityPubMessagesProcessed sets TotalActivityPubMessagesProcessed field to given value.

### HasTotalActivityPubMessagesProcessed

`func (o *ServerStats) HasTotalActivityPubMessagesProcessed() bool`

HasTotalActivityPubMessagesProcessed returns a boolean if a field has been set.

### GetTotalActivityPubMessagesSuccesses

`func (o *ServerStats) GetTotalActivityPubMessagesSuccesses() float32`

GetTotalActivityPubMessagesSuccesses returns the TotalActivityPubMessagesSuccesses field if non-nil, zero value otherwise.

### GetTotalActivityPubMessagesSuccessesOk

`func (o *ServerStats) GetTotalActivityPubMessagesSuccessesOk() (*float32, bool)`

GetTotalActivityPubMessagesSuccessesOk returns a tuple with the TotalActivityPubMessagesSuccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivityPubMessagesSuccesses

`func (o *ServerStats) SetTotalActivityPubMessagesSuccesses(v float32)`

SetTotalActivityPubMessagesSuccesses sets TotalActivityPubMessagesSuccesses field to given value.

### HasTotalActivityPubMessagesSuccesses

`func (o *ServerStats) HasTotalActivityPubMessagesSuccesses() bool`

HasTotalActivityPubMessagesSuccesses returns a boolean if a field has been set.

### GetTotalActivityPubMessagesErrors

`func (o *ServerStats) GetTotalActivityPubMessagesErrors() float32`

GetTotalActivityPubMessagesErrors returns the TotalActivityPubMessagesErrors field if non-nil, zero value otherwise.

### GetTotalActivityPubMessagesErrorsOk

`func (o *ServerStats) GetTotalActivityPubMessagesErrorsOk() (*float32, bool)`

GetTotalActivityPubMessagesErrorsOk returns a tuple with the TotalActivityPubMessagesErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivityPubMessagesErrors

`func (o *ServerStats) SetTotalActivityPubMessagesErrors(v float32)`

SetTotalActivityPubMessagesErrors sets TotalActivityPubMessagesErrors field to given value.

### HasTotalActivityPubMessagesErrors

`func (o *ServerStats) HasTotalActivityPubMessagesErrors() bool`

HasTotalActivityPubMessagesErrors returns a boolean if a field has been set.

### GetActivityPubMessagesProcessedPerSecond

`func (o *ServerStats) GetActivityPubMessagesProcessedPerSecond() float32`

GetActivityPubMessagesProcessedPerSecond returns the ActivityPubMessagesProcessedPerSecond field if non-nil, zero value otherwise.

### GetActivityPubMessagesProcessedPerSecondOk

`func (o *ServerStats) GetActivityPubMessagesProcessedPerSecondOk() (*float32, bool)`

GetActivityPubMessagesProcessedPerSecondOk returns a tuple with the ActivityPubMessagesProcessedPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityPubMessagesProcessedPerSecond

`func (o *ServerStats) SetActivityPubMessagesProcessedPerSecond(v float32)`

SetActivityPubMessagesProcessedPerSecond sets ActivityPubMessagesProcessedPerSecond field to given value.

### HasActivityPubMessagesProcessedPerSecond

`func (o *ServerStats) HasActivityPubMessagesProcessedPerSecond() bool`

HasActivityPubMessagesProcessedPerSecond returns a boolean if a field has been set.

### GetTotalActivityPubMessagesWaiting

`func (o *ServerStats) GetTotalActivityPubMessagesWaiting() float32`

GetTotalActivityPubMessagesWaiting returns the TotalActivityPubMessagesWaiting field if non-nil, zero value otherwise.

### GetTotalActivityPubMessagesWaitingOk

`func (o *ServerStats) GetTotalActivityPubMessagesWaitingOk() (*float32, bool)`

GetTotalActivityPubMessagesWaitingOk returns a tuple with the TotalActivityPubMessagesWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivityPubMessagesWaiting

`func (o *ServerStats) SetTotalActivityPubMessagesWaiting(v float32)`

SetTotalActivityPubMessagesWaiting sets TotalActivityPubMessagesWaiting field to given value.

### HasTotalActivityPubMessagesWaiting

`func (o *ServerStats) HasTotalActivityPubMessagesWaiting() bool`

HasTotalActivityPubMessagesWaiting returns a boolean if a field has been set.

### GetAverageRegistrationRequestResponseTimeMs

`func (o *ServerStats) GetAverageRegistrationRequestResponseTimeMs() float32`

GetAverageRegistrationRequestResponseTimeMs returns the AverageRegistrationRequestResponseTimeMs field if non-nil, zero value otherwise.

### GetAverageRegistrationRequestResponseTimeMsOk

`func (o *ServerStats) GetAverageRegistrationRequestResponseTimeMsOk() (*float32, bool)`

GetAverageRegistrationRequestResponseTimeMsOk returns a tuple with the AverageRegistrationRequestResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRegistrationRequestResponseTimeMs

`func (o *ServerStats) SetAverageRegistrationRequestResponseTimeMs(v float32)`

SetAverageRegistrationRequestResponseTimeMs sets AverageRegistrationRequestResponseTimeMs field to given value.

### HasAverageRegistrationRequestResponseTimeMs

`func (o *ServerStats) HasAverageRegistrationRequestResponseTimeMs() bool`

HasAverageRegistrationRequestResponseTimeMs returns a boolean if a field has been set.

### GetTotalRegistrationRequestsProcessed

`func (o *ServerStats) GetTotalRegistrationRequestsProcessed() float32`

GetTotalRegistrationRequestsProcessed returns the TotalRegistrationRequestsProcessed field if non-nil, zero value otherwise.

### GetTotalRegistrationRequestsProcessedOk

`func (o *ServerStats) GetTotalRegistrationRequestsProcessedOk() (*float32, bool)`

GetTotalRegistrationRequestsProcessedOk returns a tuple with the TotalRegistrationRequestsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRegistrationRequestsProcessed

`func (o *ServerStats) SetTotalRegistrationRequestsProcessed(v float32)`

SetTotalRegistrationRequestsProcessed sets TotalRegistrationRequestsProcessed field to given value.

### HasTotalRegistrationRequestsProcessed

`func (o *ServerStats) HasTotalRegistrationRequestsProcessed() bool`

HasTotalRegistrationRequestsProcessed returns a boolean if a field has been set.

### GetTotalRegistrationRequests

`func (o *ServerStats) GetTotalRegistrationRequests() float32`

GetTotalRegistrationRequests returns the TotalRegistrationRequests field if non-nil, zero value otherwise.

### GetTotalRegistrationRequestsOk

`func (o *ServerStats) GetTotalRegistrationRequestsOk() (*float32, bool)`

GetTotalRegistrationRequestsOk returns a tuple with the TotalRegistrationRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRegistrationRequests

`func (o *ServerStats) SetTotalRegistrationRequests(v float32)`

SetTotalRegistrationRequests sets TotalRegistrationRequests field to given value.

### HasTotalRegistrationRequests

`func (o *ServerStats) HasTotalRegistrationRequests() bool`

HasTotalRegistrationRequests returns a boolean if a field has been set.

### GetAverageAbuseResponseTimeMs

`func (o *ServerStats) GetAverageAbuseResponseTimeMs() float32`

GetAverageAbuseResponseTimeMs returns the AverageAbuseResponseTimeMs field if non-nil, zero value otherwise.

### GetAverageAbuseResponseTimeMsOk

`func (o *ServerStats) GetAverageAbuseResponseTimeMsOk() (*float32, bool)`

GetAverageAbuseResponseTimeMsOk returns a tuple with the AverageAbuseResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageAbuseResponseTimeMs

`func (o *ServerStats) SetAverageAbuseResponseTimeMs(v float32)`

SetAverageAbuseResponseTimeMs sets AverageAbuseResponseTimeMs field to given value.

### HasAverageAbuseResponseTimeMs

`func (o *ServerStats) HasAverageAbuseResponseTimeMs() bool`

HasAverageAbuseResponseTimeMs returns a boolean if a field has been set.

### GetTotalAbusesProcessed

`func (o *ServerStats) GetTotalAbusesProcessed() float32`

GetTotalAbusesProcessed returns the TotalAbusesProcessed field if non-nil, zero value otherwise.

### GetTotalAbusesProcessedOk

`func (o *ServerStats) GetTotalAbusesProcessedOk() (*float32, bool)`

GetTotalAbusesProcessedOk returns a tuple with the TotalAbusesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAbusesProcessed

`func (o *ServerStats) SetTotalAbusesProcessed(v float32)`

SetTotalAbusesProcessed sets TotalAbusesProcessed field to given value.

### HasTotalAbusesProcessed

`func (o *ServerStats) HasTotalAbusesProcessed() bool`

HasTotalAbusesProcessed returns a boolean if a field has been set.

### GetTotalAbuses

`func (o *ServerStats) GetTotalAbuses() float32`

GetTotalAbuses returns the TotalAbuses field if non-nil, zero value otherwise.

### GetTotalAbusesOk

`func (o *ServerStats) GetTotalAbusesOk() (*float32, bool)`

GetTotalAbusesOk returns a tuple with the TotalAbuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAbuses

`func (o *ServerStats) SetTotalAbuses(v float32)`

SetTotalAbuses sets TotalAbuses field to given value.

### HasTotalAbuses

`func (o *ServerStats) HasTotalAbuses() bool`

HasTotalAbuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


