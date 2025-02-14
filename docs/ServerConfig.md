# ServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**ServerConfigInstance**](ServerConfigInstance.md) |  | [optional] 
**Search** | Pointer to [**ServerConfigSearch**](ServerConfigSearch.md) |  | [optional] 
**Plugin** | Pointer to [**ServerConfigPlugin**](ServerConfigPlugin.md) |  | [optional] 
**Theme** | Pointer to [**ServerConfigPlugin**](ServerConfigPlugin.md) |  | [optional] 
**Email** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**ContactForm** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**ServerVersion** | Pointer to **string** |  | [optional] 
**ServerCommit** | Pointer to **string** |  | [optional] 
**Signup** | Pointer to [**ServerConfigSignup**](ServerConfigSignup.md) |  | [optional] 
**Transcoding** | Pointer to [**ServerConfigTranscoding**](ServerConfigTranscoding.md) |  | [optional] 
**Import** | Pointer to [**ServerConfigImport**](ServerConfigImport.md) |  | [optional] 
**Export** | Pointer to [**ServerConfigExport**](ServerConfigExport.md) |  | [optional] 
**AutoBlacklist** | Pointer to [**ServerConfigAutoBlacklist**](ServerConfigAutoBlacklist.md) |  | [optional] 
**Avatar** | Pointer to [**ServerConfigAvatar**](ServerConfigAvatar.md) |  | [optional] 
**Video** | Pointer to [**ServerConfigVideo**](ServerConfigVideo.md) |  | [optional] 
**VideoCaption** | Pointer to [**ServerConfigVideoCaption**](ServerConfigVideoCaption.md) |  | [optional] 
**User** | Pointer to [**ServerConfigUser**](ServerConfigUser.md) |  | [optional] 
**Trending** | Pointer to [**ServerConfigTrending**](ServerConfigTrending.md) |  | [optional] 
**Tracker** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**Followings** | Pointer to [**ServerConfigFollowings**](ServerConfigFollowings.md) |  | [optional] 
**Federation** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**Homepage** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**OpenTelemetry** | Pointer to [**ServerConfigOpenTelemetry**](ServerConfigOpenTelemetry.md) |  | [optional] 
**Views** | Pointer to [**ServerConfigViews**](ServerConfigViews.md) |  | [optional] 

## Methods

### NewServerConfig

`func NewServerConfig() *ServerConfig`

NewServerConfig instantiates a new ServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigWithDefaults

`func NewServerConfigWithDefaults() *ServerConfig`

NewServerConfigWithDefaults instantiates a new ServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ServerConfig) GetInstance() ServerConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ServerConfig) GetInstanceOk() (*ServerConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ServerConfig) SetInstance(v ServerConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ServerConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetSearch

`func (o *ServerConfig) GetSearch() ServerConfigSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ServerConfig) GetSearchOk() (*ServerConfigSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ServerConfig) SetSearch(v ServerConfigSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ServerConfig) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetPlugin

`func (o *ServerConfig) GetPlugin() ServerConfigPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *ServerConfig) GetPluginOk() (*ServerConfigPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *ServerConfig) SetPlugin(v ServerConfigPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *ServerConfig) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetTheme

`func (o *ServerConfig) GetTheme() ServerConfigPlugin`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *ServerConfig) GetThemeOk() (*ServerConfigPlugin, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *ServerConfig) SetTheme(v ServerConfigPlugin)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *ServerConfig) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetEmail

`func (o *ServerConfig) GetEmail() ServerConfigEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServerConfig) GetEmailOk() (*ServerConfigEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServerConfig) SetEmail(v ServerConfigEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServerConfig) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactForm

`func (o *ServerConfig) GetContactForm() ServerConfigEmail`

GetContactForm returns the ContactForm field if non-nil, zero value otherwise.

### GetContactFormOk

`func (o *ServerConfig) GetContactFormOk() (*ServerConfigEmail, bool)`

GetContactFormOk returns a tuple with the ContactForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactForm

`func (o *ServerConfig) SetContactForm(v ServerConfigEmail)`

SetContactForm sets ContactForm field to given value.

### HasContactForm

`func (o *ServerConfig) HasContactForm() bool`

HasContactForm returns a boolean if a field has been set.

### GetServerVersion

`func (o *ServerConfig) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *ServerConfig) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *ServerConfig) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *ServerConfig) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerCommit

`func (o *ServerConfig) GetServerCommit() string`

GetServerCommit returns the ServerCommit field if non-nil, zero value otherwise.

### GetServerCommitOk

`func (o *ServerConfig) GetServerCommitOk() (*string, bool)`

GetServerCommitOk returns a tuple with the ServerCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCommit

`func (o *ServerConfig) SetServerCommit(v string)`

SetServerCommit sets ServerCommit field to given value.

### HasServerCommit

`func (o *ServerConfig) HasServerCommit() bool`

HasServerCommit returns a boolean if a field has been set.

### GetSignup

`func (o *ServerConfig) GetSignup() ServerConfigSignup`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *ServerConfig) GetSignupOk() (*ServerConfigSignup, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *ServerConfig) SetSignup(v ServerConfigSignup)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *ServerConfig) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetTranscoding

`func (o *ServerConfig) GetTranscoding() ServerConfigTranscoding`

GetTranscoding returns the Transcoding field if non-nil, zero value otherwise.

### GetTranscodingOk

`func (o *ServerConfig) GetTranscodingOk() (*ServerConfigTranscoding, bool)`

GetTranscodingOk returns a tuple with the Transcoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscoding

`func (o *ServerConfig) SetTranscoding(v ServerConfigTranscoding)`

SetTranscoding sets Transcoding field to given value.

### HasTranscoding

`func (o *ServerConfig) HasTranscoding() bool`

HasTranscoding returns a boolean if a field has been set.

### GetImport

`func (o *ServerConfig) GetImport() ServerConfigImport`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *ServerConfig) GetImportOk() (*ServerConfigImport, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *ServerConfig) SetImport(v ServerConfigImport)`

SetImport sets Import field to given value.

### HasImport

`func (o *ServerConfig) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetExport

`func (o *ServerConfig) GetExport() ServerConfigExport`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *ServerConfig) GetExportOk() (*ServerConfigExport, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *ServerConfig) SetExport(v ServerConfigExport)`

SetExport sets Export field to given value.

### HasExport

`func (o *ServerConfig) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetAutoBlacklist

`func (o *ServerConfig) GetAutoBlacklist() ServerConfigAutoBlacklist`

GetAutoBlacklist returns the AutoBlacklist field if non-nil, zero value otherwise.

### GetAutoBlacklistOk

`func (o *ServerConfig) GetAutoBlacklistOk() (*ServerConfigAutoBlacklist, bool)`

GetAutoBlacklistOk returns a tuple with the AutoBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlacklist

`func (o *ServerConfig) SetAutoBlacklist(v ServerConfigAutoBlacklist)`

SetAutoBlacklist sets AutoBlacklist field to given value.

### HasAutoBlacklist

`func (o *ServerConfig) HasAutoBlacklist() bool`

HasAutoBlacklist returns a boolean if a field has been set.

### GetAvatar

`func (o *ServerConfig) GetAvatar() ServerConfigAvatar`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ServerConfig) GetAvatarOk() (*ServerConfigAvatar, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ServerConfig) SetAvatar(v ServerConfigAvatar)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ServerConfig) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetVideo

`func (o *ServerConfig) GetVideo() ServerConfigVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *ServerConfig) GetVideoOk() (*ServerConfigVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *ServerConfig) SetVideo(v ServerConfigVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *ServerConfig) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVideoCaption

`func (o *ServerConfig) GetVideoCaption() ServerConfigVideoCaption`

GetVideoCaption returns the VideoCaption field if non-nil, zero value otherwise.

### GetVideoCaptionOk

`func (o *ServerConfig) GetVideoCaptionOk() (*ServerConfigVideoCaption, bool)`

GetVideoCaptionOk returns a tuple with the VideoCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCaption

`func (o *ServerConfig) SetVideoCaption(v ServerConfigVideoCaption)`

SetVideoCaption sets VideoCaption field to given value.

### HasVideoCaption

`func (o *ServerConfig) HasVideoCaption() bool`

HasVideoCaption returns a boolean if a field has been set.

### GetUser

`func (o *ServerConfig) GetUser() ServerConfigUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServerConfig) GetUserOk() (*ServerConfigUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServerConfig) SetUser(v ServerConfigUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ServerConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTrending

`func (o *ServerConfig) GetTrending() ServerConfigTrending`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *ServerConfig) GetTrendingOk() (*ServerConfigTrending, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *ServerConfig) SetTrending(v ServerConfigTrending)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *ServerConfig) HasTrending() bool`

HasTrending returns a boolean if a field has been set.

### GetTracker

`func (o *ServerConfig) GetTracker() ServerConfigEmail`

GetTracker returns the Tracker field if non-nil, zero value otherwise.

### GetTrackerOk

`func (o *ServerConfig) GetTrackerOk() (*ServerConfigEmail, bool)`

GetTrackerOk returns a tuple with the Tracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracker

`func (o *ServerConfig) SetTracker(v ServerConfigEmail)`

SetTracker sets Tracker field to given value.

### HasTracker

`func (o *ServerConfig) HasTracker() bool`

HasTracker returns a boolean if a field has been set.

### GetFollowings

`func (o *ServerConfig) GetFollowings() ServerConfigFollowings`

GetFollowings returns the Followings field if non-nil, zero value otherwise.

### GetFollowingsOk

`func (o *ServerConfig) GetFollowingsOk() (*ServerConfigFollowings, bool)`

GetFollowingsOk returns a tuple with the Followings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowings

`func (o *ServerConfig) SetFollowings(v ServerConfigFollowings)`

SetFollowings sets Followings field to given value.

### HasFollowings

`func (o *ServerConfig) HasFollowings() bool`

HasFollowings returns a boolean if a field has been set.

### GetFederation

`func (o *ServerConfig) GetFederation() ServerConfigEmail`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *ServerConfig) GetFederationOk() (*ServerConfigEmail, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *ServerConfig) SetFederation(v ServerConfigEmail)`

SetFederation sets Federation field to given value.

### HasFederation

`func (o *ServerConfig) HasFederation() bool`

HasFederation returns a boolean if a field has been set.

### GetHomepage

`func (o *ServerConfig) GetHomepage() ServerConfigEmail`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *ServerConfig) GetHomepageOk() (*ServerConfigEmail, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *ServerConfig) SetHomepage(v ServerConfigEmail)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *ServerConfig) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetOpenTelemetry

`func (o *ServerConfig) GetOpenTelemetry() ServerConfigOpenTelemetry`

GetOpenTelemetry returns the OpenTelemetry field if non-nil, zero value otherwise.

### GetOpenTelemetryOk

`func (o *ServerConfig) GetOpenTelemetryOk() (*ServerConfigOpenTelemetry, bool)`

GetOpenTelemetryOk returns a tuple with the OpenTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTelemetry

`func (o *ServerConfig) SetOpenTelemetry(v ServerConfigOpenTelemetry)`

SetOpenTelemetry sets OpenTelemetry field to given value.

### HasOpenTelemetry

`func (o *ServerConfig) HasOpenTelemetry() bool`

HasOpenTelemetry returns a boolean if a field has been set.

### GetViews

`func (o *ServerConfig) GetViews() ServerConfigViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ServerConfig) GetViewsOk() (*ServerConfigViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ServerConfig) SetViews(v ServerConfigViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *ServerConfig) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


