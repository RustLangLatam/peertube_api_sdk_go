/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"encoding/json"
)

// checks if the ServerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerConfig{}

// ServerConfig struct for ServerConfig
type ServerConfig struct {
	Instance      *ServerConfigInstance      `json:"instance,omitempty"`
	Search        *ServerConfigSearch        `json:"search,omitempty"`
	Plugin        *ServerConfigPlugin        `json:"plugin,omitempty"`
	Theme         *ServerConfigPlugin        `json:"theme,omitempty"`
	Email         *ServerConfigEmail         `json:"email,omitempty"`
	ContactForm   *ServerConfigEmail         `json:"contactForm,omitempty"`
	ServerVersion *string                    `json:"serverVersion,omitempty"`
	ServerCommit  *string                    `json:"serverCommit,omitempty"`
	Signup        *ServerConfigSignup        `json:"signup,omitempty"`
	Transcoding   *ServerConfigTranscoding   `json:"transcoding,omitempty"`
	Import        *ServerConfigImport        `json:"import,omitempty"`
	Export        *ServerConfigExport        `json:"export,omitempty"`
	AutoBlacklist *ServerConfigAutoBlacklist `json:"autoBlacklist,omitempty"`
	Avatar        *ServerConfigAvatar        `json:"avatar,omitempty"`
	Video         *ServerConfigVideo         `json:"video,omitempty"`
	VideoCaption  *ServerConfigVideoCaption  `json:"videoCaption,omitempty"`
	User          *ServerConfigUser          `json:"user,omitempty"`
	Trending      *ServerConfigTrending      `json:"trending,omitempty"`
	Tracker       *ServerConfigEmail         `json:"tracker,omitempty"`
	Followings    *ServerConfigFollowings    `json:"followings,omitempty"`
	Federation    *ServerConfigEmail         `json:"federation,omitempty"`
	Homepage      *ServerConfigEmail         `json:"homepage,omitempty"`
	OpenTelemetry *ServerConfigOpenTelemetry `json:"openTelemetry,omitempty"`
	Views         *ServerConfigViews         `json:"views,omitempty"`
}

// NewServerConfig instantiates a new ServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfig() *ServerConfig {
	this := ServerConfig{}
	return &this
}

// NewServerConfigWithDefaults instantiates a new ServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigWithDefaults() *ServerConfig {
	this := ServerConfig{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ServerConfig) GetInstance() ServerConfigInstance {
	if o == nil || IsNil(o.Instance) {
		var ret ServerConfigInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetInstanceOk() (*ServerConfigInstance, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ServerConfig) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given ServerConfigInstance and assigns it to the Instance field.
func (o *ServerConfig) SetInstance(v ServerConfigInstance) {
	o.Instance = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *ServerConfig) GetSearch() ServerConfigSearch {
	if o == nil || IsNil(o.Search) {
		var ret ServerConfigSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetSearchOk() (*ServerConfigSearch, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *ServerConfig) HasSearch() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given ServerConfigSearch and assigns it to the Search field.
func (o *ServerConfig) SetSearch(v ServerConfigSearch) {
	o.Search = &v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *ServerConfig) GetPlugin() ServerConfigPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret ServerConfigPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetPluginOk() (*ServerConfigPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *ServerConfig) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given ServerConfigPlugin and assigns it to the Plugin field.
func (o *ServerConfig) SetPlugin(v ServerConfigPlugin) {
	o.Plugin = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ServerConfig) GetTheme() ServerConfigPlugin {
	if o == nil || IsNil(o.Theme) {
		var ret ServerConfigPlugin
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetThemeOk() (*ServerConfigPlugin, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ServerConfig) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given ServerConfigPlugin and assigns it to the Theme field.
func (o *ServerConfig) SetTheme(v ServerConfigPlugin) {
	o.Theme = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ServerConfig) GetEmail() ServerConfigEmail {
	if o == nil || IsNil(o.Email) {
		var ret ServerConfigEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetEmailOk() (*ServerConfigEmail, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ServerConfig) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given ServerConfigEmail and assigns it to the Email field.
func (o *ServerConfig) SetEmail(v ServerConfigEmail) {
	o.Email = &v
}

// GetContactForm returns the ContactForm field value if set, zero value otherwise.
func (o *ServerConfig) GetContactForm() ServerConfigEmail {
	if o == nil || IsNil(o.ContactForm) {
		var ret ServerConfigEmail
		return ret
	}
	return *o.ContactForm
}

// GetContactFormOk returns a tuple with the ContactForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetContactFormOk() (*ServerConfigEmail, bool) {
	if o == nil || IsNil(o.ContactForm) {
		return nil, false
	}
	return o.ContactForm, true
}

// HasContactForm returns a boolean if a field has been set.
func (o *ServerConfig) HasContactForm() bool {
	if o != nil && !IsNil(o.ContactForm) {
		return true
	}

	return false
}

// SetContactForm gets a reference to the given ServerConfigEmail and assigns it to the ContactForm field.
func (o *ServerConfig) SetContactForm(v ServerConfigEmail) {
	o.ContactForm = &v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *ServerConfig) GetServerVersion() string {
	if o == nil || IsNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetServerVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerVersion) {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *ServerConfig) HasServerVersion() bool {
	if o != nil && !IsNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *ServerConfig) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetServerCommit returns the ServerCommit field value if set, zero value otherwise.
func (o *ServerConfig) GetServerCommit() string {
	if o == nil || IsNil(o.ServerCommit) {
		var ret string
		return ret
	}
	return *o.ServerCommit
}

// GetServerCommitOk returns a tuple with the ServerCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetServerCommitOk() (*string, bool) {
	if o == nil || IsNil(o.ServerCommit) {
		return nil, false
	}
	return o.ServerCommit, true
}

// HasServerCommit returns a boolean if a field has been set.
func (o *ServerConfig) HasServerCommit() bool {
	if o != nil && !IsNil(o.ServerCommit) {
		return true
	}

	return false
}

// SetServerCommit gets a reference to the given string and assigns it to the ServerCommit field.
func (o *ServerConfig) SetServerCommit(v string) {
	o.ServerCommit = &v
}

// GetSignup returns the Signup field value if set, zero value otherwise.
func (o *ServerConfig) GetSignup() ServerConfigSignup {
	if o == nil || IsNil(o.Signup) {
		var ret ServerConfigSignup
		return ret
	}
	return *o.Signup
}

// GetSignupOk returns a tuple with the Signup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetSignupOk() (*ServerConfigSignup, bool) {
	if o == nil || IsNil(o.Signup) {
		return nil, false
	}
	return o.Signup, true
}

// HasSignup returns a boolean if a field has been set.
func (o *ServerConfig) HasSignup() bool {
	if o != nil && !IsNil(o.Signup) {
		return true
	}

	return false
}

// SetSignup gets a reference to the given ServerConfigSignup and assigns it to the Signup field.
func (o *ServerConfig) SetSignup(v ServerConfigSignup) {
	o.Signup = &v
}

// GetTranscoding returns the Transcoding field value if set, zero value otherwise.
func (o *ServerConfig) GetTranscoding() ServerConfigTranscoding {
	if o == nil || IsNil(o.Transcoding) {
		var ret ServerConfigTranscoding
		return ret
	}
	return *o.Transcoding
}

// GetTranscodingOk returns a tuple with the Transcoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetTranscodingOk() (*ServerConfigTranscoding, bool) {
	if o == nil || IsNil(o.Transcoding) {
		return nil, false
	}
	return o.Transcoding, true
}

// HasTranscoding returns a boolean if a field has been set.
func (o *ServerConfig) HasTranscoding() bool {
	if o != nil && !IsNil(o.Transcoding) {
		return true
	}

	return false
}

// SetTranscoding gets a reference to the given ServerConfigTranscoding and assigns it to the Transcoding field.
func (o *ServerConfig) SetTranscoding(v ServerConfigTranscoding) {
	o.Transcoding = &v
}

// GetImport returns the Import field value if set, zero value otherwise.
func (o *ServerConfig) GetImport() ServerConfigImport {
	if o == nil || IsNil(o.Import) {
		var ret ServerConfigImport
		return ret
	}
	return *o.Import
}

// GetImportOk returns a tuple with the Import field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetImportOk() (*ServerConfigImport, bool) {
	if o == nil || IsNil(o.Import) {
		return nil, false
	}
	return o.Import, true
}

// HasImport returns a boolean if a field has been set.
func (o *ServerConfig) HasImport() bool {
	if o != nil && !IsNil(o.Import) {
		return true
	}

	return false
}

// SetImport gets a reference to the given ServerConfigImport and assigns it to the Import field.
func (o *ServerConfig) SetImport(v ServerConfigImport) {
	o.Import = &v
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *ServerConfig) GetExport() ServerConfigExport {
	if o == nil || IsNil(o.Export) {
		var ret ServerConfigExport
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetExportOk() (*ServerConfigExport, bool) {
	if o == nil || IsNil(o.Export) {
		return nil, false
	}
	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *ServerConfig) HasExport() bool {
	if o != nil && !IsNil(o.Export) {
		return true
	}

	return false
}

// SetExport gets a reference to the given ServerConfigExport and assigns it to the Export field.
func (o *ServerConfig) SetExport(v ServerConfigExport) {
	o.Export = &v
}

// GetAutoBlacklist returns the AutoBlacklist field value if set, zero value otherwise.
func (o *ServerConfig) GetAutoBlacklist() ServerConfigAutoBlacklist {
	if o == nil || IsNil(o.AutoBlacklist) {
		var ret ServerConfigAutoBlacklist
		return ret
	}
	return *o.AutoBlacklist
}

// GetAutoBlacklistOk returns a tuple with the AutoBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetAutoBlacklistOk() (*ServerConfigAutoBlacklist, bool) {
	if o == nil || IsNil(o.AutoBlacklist) {
		return nil, false
	}
	return o.AutoBlacklist, true
}

// HasAutoBlacklist returns a boolean if a field has been set.
func (o *ServerConfig) HasAutoBlacklist() bool {
	if o != nil && !IsNil(o.AutoBlacklist) {
		return true
	}

	return false
}

// SetAutoBlacklist gets a reference to the given ServerConfigAutoBlacklist and assigns it to the AutoBlacklist field.
func (o *ServerConfig) SetAutoBlacklist(v ServerConfigAutoBlacklist) {
	o.AutoBlacklist = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ServerConfig) GetAvatar() ServerConfigAvatar {
	if o == nil || IsNil(o.Avatar) {
		var ret ServerConfigAvatar
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetAvatarOk() (*ServerConfigAvatar, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ServerConfig) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given ServerConfigAvatar and assigns it to the Avatar field.
func (o *ServerConfig) SetAvatar(v ServerConfigAvatar) {
	o.Avatar = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *ServerConfig) GetVideo() ServerConfigVideo {
	if o == nil || IsNil(o.Video) {
		var ret ServerConfigVideo
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetVideoOk() (*ServerConfigVideo, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *ServerConfig) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given ServerConfigVideo and assigns it to the Video field.
func (o *ServerConfig) SetVideo(v ServerConfigVideo) {
	o.Video = &v
}

// GetVideoCaption returns the VideoCaption field value if set, zero value otherwise.
func (o *ServerConfig) GetVideoCaption() ServerConfigVideoCaption {
	if o == nil || IsNil(o.VideoCaption) {
		var ret ServerConfigVideoCaption
		return ret
	}
	return *o.VideoCaption
}

// GetVideoCaptionOk returns a tuple with the VideoCaption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetVideoCaptionOk() (*ServerConfigVideoCaption, bool) {
	if o == nil || IsNil(o.VideoCaption) {
		return nil, false
	}
	return o.VideoCaption, true
}

// HasVideoCaption returns a boolean if a field has been set.
func (o *ServerConfig) HasVideoCaption() bool {
	if o != nil && !IsNil(o.VideoCaption) {
		return true
	}

	return false
}

// SetVideoCaption gets a reference to the given ServerConfigVideoCaption and assigns it to the VideoCaption field.
func (o *ServerConfig) SetVideoCaption(v ServerConfigVideoCaption) {
	o.VideoCaption = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ServerConfig) GetUser() ServerConfigUser {
	if o == nil || IsNil(o.User) {
		var ret ServerConfigUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetUserOk() (*ServerConfigUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ServerConfig) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given ServerConfigUser and assigns it to the User field.
func (o *ServerConfig) SetUser(v ServerConfigUser) {
	o.User = &v
}

// GetTrending returns the Trending field value if set, zero value otherwise.
func (o *ServerConfig) GetTrending() ServerConfigTrending {
	if o == nil || IsNil(o.Trending) {
		var ret ServerConfigTrending
		return ret
	}
	return *o.Trending
}

// GetTrendingOk returns a tuple with the Trending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetTrendingOk() (*ServerConfigTrending, bool) {
	if o == nil || IsNil(o.Trending) {
		return nil, false
	}
	return o.Trending, true
}

// HasTrending returns a boolean if a field has been set.
func (o *ServerConfig) HasTrending() bool {
	if o != nil && !IsNil(o.Trending) {
		return true
	}

	return false
}

// SetTrending gets a reference to the given ServerConfigTrending and assigns it to the Trending field.
func (o *ServerConfig) SetTrending(v ServerConfigTrending) {
	o.Trending = &v
}

// GetTracker returns the Tracker field value if set, zero value otherwise.
func (o *ServerConfig) GetTracker() ServerConfigEmail {
	if o == nil || IsNil(o.Tracker) {
		var ret ServerConfigEmail
		return ret
	}
	return *o.Tracker
}

// GetTrackerOk returns a tuple with the Tracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetTrackerOk() (*ServerConfigEmail, bool) {
	if o == nil || IsNil(o.Tracker) {
		return nil, false
	}
	return o.Tracker, true
}

// HasTracker returns a boolean if a field has been set.
func (o *ServerConfig) HasTracker() bool {
	if o != nil && !IsNil(o.Tracker) {
		return true
	}

	return false
}

// SetTracker gets a reference to the given ServerConfigEmail and assigns it to the Tracker field.
func (o *ServerConfig) SetTracker(v ServerConfigEmail) {
	o.Tracker = &v
}

// GetFollowings returns the Followings field value if set, zero value otherwise.
func (o *ServerConfig) GetFollowings() ServerConfigFollowings {
	if o == nil || IsNil(o.Followings) {
		var ret ServerConfigFollowings
		return ret
	}
	return *o.Followings
}

// GetFollowingsOk returns a tuple with the Followings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetFollowingsOk() (*ServerConfigFollowings, bool) {
	if o == nil || IsNil(o.Followings) {
		return nil, false
	}
	return o.Followings, true
}

// HasFollowings returns a boolean if a field has been set.
func (o *ServerConfig) HasFollowings() bool {
	if o != nil && !IsNil(o.Followings) {
		return true
	}

	return false
}

// SetFollowings gets a reference to the given ServerConfigFollowings and assigns it to the Followings field.
func (o *ServerConfig) SetFollowings(v ServerConfigFollowings) {
	o.Followings = &v
}

// GetFederation returns the Federation field value if set, zero value otherwise.
func (o *ServerConfig) GetFederation() ServerConfigEmail {
	if o == nil || IsNil(o.Federation) {
		var ret ServerConfigEmail
		return ret
	}
	return *o.Federation
}

// GetFederationOk returns a tuple with the Federation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetFederationOk() (*ServerConfigEmail, bool) {
	if o == nil || IsNil(o.Federation) {
		return nil, false
	}
	return o.Federation, true
}

// HasFederation returns a boolean if a field has been set.
func (o *ServerConfig) HasFederation() bool {
	if o != nil && !IsNil(o.Federation) {
		return true
	}

	return false
}

// SetFederation gets a reference to the given ServerConfigEmail and assigns it to the Federation field.
func (o *ServerConfig) SetFederation(v ServerConfigEmail) {
	o.Federation = &v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *ServerConfig) GetHomepage() ServerConfigEmail {
	if o == nil || IsNil(o.Homepage) {
		var ret ServerConfigEmail
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetHomepageOk() (*ServerConfigEmail, bool) {
	if o == nil || IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *ServerConfig) HasHomepage() bool {
	if o != nil && !IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given ServerConfigEmail and assigns it to the Homepage field.
func (o *ServerConfig) SetHomepage(v ServerConfigEmail) {
	o.Homepage = &v
}

// GetOpenTelemetry returns the OpenTelemetry field value if set, zero value otherwise.
func (o *ServerConfig) GetOpenTelemetry() ServerConfigOpenTelemetry {
	if o == nil || IsNil(o.OpenTelemetry) {
		var ret ServerConfigOpenTelemetry
		return ret
	}
	return *o.OpenTelemetry
}

// GetOpenTelemetryOk returns a tuple with the OpenTelemetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetOpenTelemetryOk() (*ServerConfigOpenTelemetry, bool) {
	if o == nil || IsNil(o.OpenTelemetry) {
		return nil, false
	}
	return o.OpenTelemetry, true
}

// HasOpenTelemetry returns a boolean if a field has been set.
func (o *ServerConfig) HasOpenTelemetry() bool {
	if o != nil && !IsNil(o.OpenTelemetry) {
		return true
	}

	return false
}

// SetOpenTelemetry gets a reference to the given ServerConfigOpenTelemetry and assigns it to the OpenTelemetry field.
func (o *ServerConfig) SetOpenTelemetry(v ServerConfigOpenTelemetry) {
	o.OpenTelemetry = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *ServerConfig) GetViews() ServerConfigViews {
	if o == nil || IsNil(o.Views) {
		var ret ServerConfigViews
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetViewsOk() (*ServerConfigViews, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *ServerConfig) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given ServerConfigViews and assigns it to the Views field.
func (o *ServerConfig) SetViews(v ServerConfigViews) {
	o.Views = &v
}

func (o ServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ContactForm) {
		toSerialize["contactForm"] = o.ContactForm
	}
	if !IsNil(o.ServerVersion) {
		toSerialize["serverVersion"] = o.ServerVersion
	}
	if !IsNil(o.ServerCommit) {
		toSerialize["serverCommit"] = o.ServerCommit
	}
	if !IsNil(o.Signup) {
		toSerialize["signup"] = o.Signup
	}
	if !IsNil(o.Transcoding) {
		toSerialize["transcoding"] = o.Transcoding
	}
	if !IsNil(o.Import) {
		toSerialize["import"] = o.Import
	}
	if !IsNil(o.Export) {
		toSerialize["export"] = o.Export
	}
	if !IsNil(o.AutoBlacklist) {
		toSerialize["autoBlacklist"] = o.AutoBlacklist
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.VideoCaption) {
		toSerialize["videoCaption"] = o.VideoCaption
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Trending) {
		toSerialize["trending"] = o.Trending
	}
	if !IsNil(o.Tracker) {
		toSerialize["tracker"] = o.Tracker
	}
	if !IsNil(o.Followings) {
		toSerialize["followings"] = o.Followings
	}
	if !IsNil(o.Federation) {
		toSerialize["federation"] = o.Federation
	}
	if !IsNil(o.Homepage) {
		toSerialize["homepage"] = o.Homepage
	}
	if !IsNil(o.OpenTelemetry) {
		toSerialize["openTelemetry"] = o.OpenTelemetry
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	return toSerialize, nil
}

type NullableServerConfig struct {
	value *ServerConfig
	isSet bool
}

func (v NullableServerConfig) Get() *ServerConfig {
	return v.value
}

func (v *NullableServerConfig) Set(val *ServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfig(val *ServerConfig) *NullableServerConfig {
	return &NullableServerConfig{value: val, isSet: true}
}

func (v NullableServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
