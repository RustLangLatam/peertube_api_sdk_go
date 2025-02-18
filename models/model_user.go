/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/RustLangLatam/peertube_api_sdk_go/utils"
	"time"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &User{}

// User struct for User
type User struct {
	Account *Account `json:"account,omitempty"`
	// Automatically start playing the upcoming video after the currently playing video
	AutoPlayNextVideo *bool `json:"autoPlayNextVideo,omitempty"`
	// Automatically start playing the video on the playlist after the currently playing video
	AutoPlayNextVideoPlaylist *bool `json:"autoPlayNextVideoPlaylist,omitempty"`
	// Automatically start playing the video on the watch page
	AutoPlayVideo *bool   `json:"autoPlayVideo,omitempty"`
	Blocked       *bool   `json:"blocked,omitempty"`
	BlockedReason *string `json:"blockedReason,omitempty"`
	CreatedAt     *string `json:"createdAt,omitempty"`
	// The user email
	Email *string `json:"email,omitempty"`
	// Has the user confirmed their email address?
	EmailVerified *bool  `json:"emailVerified,omitempty"`
	Id            *int32 `json:"id,omitempty"`
	// Auth plugin to use to authenticate the user
	PluginAuth                   *string     `json:"pluginAuth,omitempty"`
	LastLoginDate                *time.Time  `json:"lastLoginDate,omitempty"`
	NoInstanceConfigWarningModal *bool       `json:"noInstanceConfigWarningModal,omitempty"`
	NoAccountSetupWarningModal   *bool       `json:"noAccountSetupWarningModal,omitempty"`
	NoWelcomeModal               *bool       `json:"noWelcomeModal,omitempty"`
	NsfwPolicy                   *NSFWPolicy `json:"nsfwPolicy,omitempty"`
	Role                         *UserRole   `json:"role,omitempty"`
	// Theme enabled by this user
	Theme *string `json:"theme,omitempty"`
	// immutable name of the user, used to find or mention its actor
	Username      *string        `json:"username,omitempty" validate:"regexp=^[a-z0-9._]+$"`
	VideoChannels []VideoChannel `json:"videoChannels,omitempty"`
	// The user video quota in bytes
	VideoQuota *int32 `json:"videoQuota,omitempty"`
	// The user daily video quota in bytes
	VideoQuotaDaily *int32 `json:"videoQuotaDaily,omitempty"`
	// Enable P2P in the player
	P2pEnabled *bool `json:"p2pEnabled,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *User) GetAccount() Account {
	if o == nil || utils.IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountOk() (*Account, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *User) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *User) SetAccount(v Account) {
	o.Account = &v
}

// GetAutoPlayNextVideo returns the AutoPlayNextVideo field value if set, zero value otherwise.
func (o *User) GetAutoPlayNextVideo() bool {
	if o == nil || utils.IsNil(o.AutoPlayNextVideo) {
		var ret bool
		return ret
	}
	return *o.AutoPlayNextVideo
}

// GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAutoPlayNextVideoOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AutoPlayNextVideo) {
		return nil, false
	}
	return o.AutoPlayNextVideo, true
}

// HasAutoPlayNextVideo returns a boolean if a field has been set.
func (o *User) HasAutoPlayNextVideo() bool {
	if o != nil && !utils.IsNil(o.AutoPlayNextVideo) {
		return true
	}

	return false
}

// SetAutoPlayNextVideo gets a reference to the given bool and assigns it to the AutoPlayNextVideo field.
func (o *User) SetAutoPlayNextVideo(v bool) {
	o.AutoPlayNextVideo = &v
}

// GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field value if set, zero value otherwise.
func (o *User) GetAutoPlayNextVideoPlaylist() bool {
	if o == nil || utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		var ret bool
		return ret
	}
	return *o.AutoPlayNextVideoPlaylist
}

// GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAutoPlayNextVideoPlaylistOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		return nil, false
	}
	return o.AutoPlayNextVideoPlaylist, true
}

// HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.
func (o *User) HasAutoPlayNextVideoPlaylist() bool {
	if o != nil && !utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		return true
	}

	return false
}

// SetAutoPlayNextVideoPlaylist gets a reference to the given bool and assigns it to the AutoPlayNextVideoPlaylist field.
func (o *User) SetAutoPlayNextVideoPlaylist(v bool) {
	o.AutoPlayNextVideoPlaylist = &v
}

// GetAutoPlayVideo returns the AutoPlayVideo field value if set, zero value otherwise.
func (o *User) GetAutoPlayVideo() bool {
	if o == nil || utils.IsNil(o.AutoPlayVideo) {
		var ret bool
		return ret
	}
	return *o.AutoPlayVideo
}

// GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAutoPlayVideoOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AutoPlayVideo) {
		return nil, false
	}
	return o.AutoPlayVideo, true
}

// HasAutoPlayVideo returns a boolean if a field has been set.
func (o *User) HasAutoPlayVideo() bool {
	if o != nil && !utils.IsNil(o.AutoPlayVideo) {
		return true
	}

	return false
}

// SetAutoPlayVideo gets a reference to the given bool and assigns it to the AutoPlayVideo field.
func (o *User) SetAutoPlayVideo(v bool) {
	o.AutoPlayVideo = &v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *User) GetBlocked() bool {
	if o == nil || utils.IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBlockedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *User) HasBlocked() bool {
	if o != nil && !utils.IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *User) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetBlockedReason returns the BlockedReason field value if set, zero value otherwise.
func (o *User) GetBlockedReason() string {
	if o == nil || utils.IsNil(o.BlockedReason) {
		var ret string
		return ret
	}
	return *o.BlockedReason
}

// GetBlockedReasonOk returns a tuple with the BlockedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBlockedReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BlockedReason) {
		return nil, false
	}
	return o.BlockedReason, true
}

// HasBlockedReason returns a boolean if a field has been set.
func (o *User) HasBlockedReason() bool {
	if o != nil && !utils.IsNil(o.BlockedReason) {
		return true
	}

	return false
}

// SetBlockedReason gets a reference to the given string and assigns it to the BlockedReason field.
func (o *User) SetBlockedReason(v string) {
	o.BlockedReason = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() string {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *User) GetEmailVerified() bool {
	if o == nil || utils.IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *User) HasEmailVerified() bool {
	if o != nil && !utils.IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *User) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *User) SetId(v int32) {
	o.Id = &v
}

// GetPluginAuth returns the PluginAuth field value if set, zero value otherwise.
func (o *User) GetPluginAuth() string {
	if o == nil || utils.IsNil(o.PluginAuth) {
		var ret string
		return ret
	}
	return *o.PluginAuth
}

// GetPluginAuthOk returns a tuple with the PluginAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPluginAuthOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PluginAuth) {
		return nil, false
	}
	return o.PluginAuth, true
}

// HasPluginAuth returns a boolean if a field has been set.
func (o *User) HasPluginAuth() bool {
	if o != nil && !utils.IsNil(o.PluginAuth) {
		return true
	}

	return false
}

// SetPluginAuth gets a reference to the given string and assigns it to the PluginAuth field.
func (o *User) SetPluginAuth(v string) {
	o.PluginAuth = &v
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise.
func (o *User) GetLastLoginDate() time.Time {
	if o == nil || utils.IsNil(o.LastLoginDate) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LastLoginDate) {
		return nil, false
	}
	return o.LastLoginDate, true
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *User) HasLastLoginDate() bool {
	if o != nil && !utils.IsNil(o.LastLoginDate) {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given time.Time and assigns it to the LastLoginDate field.
func (o *User) SetLastLoginDate(v time.Time) {
	o.LastLoginDate = &v
}

// GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field value if set, zero value otherwise.
func (o *User) GetNoInstanceConfigWarningModal() bool {
	if o == nil || utils.IsNil(o.NoInstanceConfigWarningModal) {
		var ret bool
		return ret
	}
	return *o.NoInstanceConfigWarningModal
}

// GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNoInstanceConfigWarningModalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoInstanceConfigWarningModal) {
		return nil, false
	}
	return o.NoInstanceConfigWarningModal, true
}

// HasNoInstanceConfigWarningModal returns a boolean if a field has been set.
func (o *User) HasNoInstanceConfigWarningModal() bool {
	if o != nil && !utils.IsNil(o.NoInstanceConfigWarningModal) {
		return true
	}

	return false
}

// SetNoInstanceConfigWarningModal gets a reference to the given bool and assigns it to the NoInstanceConfigWarningModal field.
func (o *User) SetNoInstanceConfigWarningModal(v bool) {
	o.NoInstanceConfigWarningModal = &v
}

// GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field value if set, zero value otherwise.
func (o *User) GetNoAccountSetupWarningModal() bool {
	if o == nil || utils.IsNil(o.NoAccountSetupWarningModal) {
		var ret bool
		return ret
	}
	return *o.NoAccountSetupWarningModal
}

// GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNoAccountSetupWarningModalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoAccountSetupWarningModal) {
		return nil, false
	}
	return o.NoAccountSetupWarningModal, true
}

// HasNoAccountSetupWarningModal returns a boolean if a field has been set.
func (o *User) HasNoAccountSetupWarningModal() bool {
	if o != nil && !utils.IsNil(o.NoAccountSetupWarningModal) {
		return true
	}

	return false
}

// SetNoAccountSetupWarningModal gets a reference to the given bool and assigns it to the NoAccountSetupWarningModal field.
func (o *User) SetNoAccountSetupWarningModal(v bool) {
	o.NoAccountSetupWarningModal = &v
}

// GetNoWelcomeModal returns the NoWelcomeModal field value if set, zero value otherwise.
func (o *User) GetNoWelcomeModal() bool {
	if o == nil || utils.IsNil(o.NoWelcomeModal) {
		var ret bool
		return ret
	}
	return *o.NoWelcomeModal
}

// GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNoWelcomeModalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoWelcomeModal) {
		return nil, false
	}
	return o.NoWelcomeModal, true
}

// HasNoWelcomeModal returns a boolean if a field has been set.
func (o *User) HasNoWelcomeModal() bool {
	if o != nil && !utils.IsNil(o.NoWelcomeModal) {
		return true
	}

	return false
}

// SetNoWelcomeModal gets a reference to the given bool and assigns it to the NoWelcomeModal field.
func (o *User) SetNoWelcomeModal(v bool) {
	o.NoWelcomeModal = &v
}

// GetNsfwPolicy returns the NsfwPolicy field value if set, zero value otherwise.
func (o *User) GetNsfwPolicy() NSFWPolicy {
	if o == nil || utils.IsNil(o.NsfwPolicy) {
		var ret NSFWPolicy
		return ret
	}
	return *o.NsfwPolicy
}

// GetNsfwPolicyOk returns a tuple with the NsfwPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNsfwPolicyOk() (*NSFWPolicy, bool) {
	if o == nil || utils.IsNil(o.NsfwPolicy) {
		return nil, false
	}
	return o.NsfwPolicy, true
}

// HasNsfwPolicy returns a boolean if a field has been set.
func (o *User) HasNsfwPolicy() bool {
	if o != nil && !utils.IsNil(o.NsfwPolicy) {
		return true
	}

	return false
}

// SetNsfwPolicy gets a reference to the given NSFWPolicy and assigns it to the NsfwPolicy field.
func (o *User) SetNsfwPolicy(v NSFWPolicy) {
	o.NsfwPolicy = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *User) GetRole() UserRole {
	if o == nil || utils.IsNil(o.Role) {
		var ret UserRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRoleOk() (*UserRole, bool) {
	if o == nil || utils.IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *User) HasRole() bool {
	if o != nil && !utils.IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRole and assigns it to the Role field.
func (o *User) SetRole(v UserRole) {
	o.Role = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *User) GetTheme() string {
	if o == nil || utils.IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetThemeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *User) HasTheme() bool {
	if o != nil && !utils.IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *User) SetTheme(v string) {
	o.Theme = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || utils.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && !utils.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetVideoChannels returns the VideoChannels field value if set, zero value otherwise.
func (o *User) GetVideoChannels() []VideoChannel {
	if o == nil || utils.IsNil(o.VideoChannels) {
		var ret []VideoChannel
		return ret
	}
	return o.VideoChannels
}

// GetVideoChannelsOk returns a tuple with the VideoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVideoChannelsOk() ([]VideoChannel, bool) {
	if o == nil || utils.IsNil(o.VideoChannels) {
		return nil, false
	}
	return o.VideoChannels, true
}

// HasVideoChannels returns a boolean if a field has been set.
func (o *User) HasVideoChannels() bool {
	if o != nil && !utils.IsNil(o.VideoChannels) {
		return true
	}

	return false
}

// SetVideoChannels gets a reference to the given []VideoChannel and assigns it to the VideoChannels field.
func (o *User) SetVideoChannels(v []VideoChannel) {
	o.VideoChannels = v
}

// GetVideoQuota returns the VideoQuota field value if set, zero value otherwise.
func (o *User) GetVideoQuota() int32 {
	if o == nil || utils.IsNil(o.VideoQuota) {
		var ret int32
		return ret
	}
	return *o.VideoQuota
}

// GetVideoQuotaOk returns a tuple with the VideoQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVideoQuotaOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoQuota) {
		return nil, false
	}
	return o.VideoQuota, true
}

// HasVideoQuota returns a boolean if a field has been set.
func (o *User) HasVideoQuota() bool {
	if o != nil && !utils.IsNil(o.VideoQuota) {
		return true
	}

	return false
}

// SetVideoQuota gets a reference to the given int32 and assigns it to the VideoQuota field.
func (o *User) SetVideoQuota(v int32) {
	o.VideoQuota = &v
}

// GetVideoQuotaDaily returns the VideoQuotaDaily field value if set, zero value otherwise.
func (o *User) GetVideoQuotaDaily() int32 {
	if o == nil || utils.IsNil(o.VideoQuotaDaily) {
		var ret int32
		return ret
	}
	return *o.VideoQuotaDaily
}

// GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVideoQuotaDailyOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoQuotaDaily) {
		return nil, false
	}
	return o.VideoQuotaDaily, true
}

// HasVideoQuotaDaily returns a boolean if a field has been set.
func (o *User) HasVideoQuotaDaily() bool {
	if o != nil && !utils.IsNil(o.VideoQuotaDaily) {
		return true
	}

	return false
}

// SetVideoQuotaDaily gets a reference to the given int32 and assigns it to the VideoQuotaDaily field.
func (o *User) SetVideoQuotaDaily(v int32) {
	o.VideoQuotaDaily = &v
}

// GetP2pEnabled returns the P2pEnabled field value if set, zero value otherwise.
func (o *User) GetP2pEnabled() bool {
	if o == nil || utils.IsNil(o.P2pEnabled) {
		var ret bool
		return ret
	}
	return *o.P2pEnabled
}

// GetP2pEnabledOk returns a tuple with the P2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetP2pEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.P2pEnabled) {
		return nil, false
	}
	return o.P2pEnabled, true
}

// HasP2pEnabled returns a boolean if a field has been set.
func (o *User) HasP2pEnabled() bool {
	if o != nil && !utils.IsNil(o.P2pEnabled) {
		return true
	}

	return false
}

// SetP2pEnabled gets a reference to the given bool and assigns it to the P2pEnabled field.
func (o *User) SetP2pEnabled(v bool) {
	o.P2pEnabled = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !utils.IsNil(o.AutoPlayNextVideo) {
		toSerialize["autoPlayNextVideo"] = o.AutoPlayNextVideo
	}
	if !utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		toSerialize["autoPlayNextVideoPlaylist"] = o.AutoPlayNextVideoPlaylist
	}
	if !utils.IsNil(o.AutoPlayVideo) {
		toSerialize["autoPlayVideo"] = o.AutoPlayVideo
	}
	if !utils.IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !utils.IsNil(o.BlockedReason) {
		toSerialize["blockedReason"] = o.BlockedReason
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.EmailVerified) {
		toSerialize["emailVerified"] = o.EmailVerified
	}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.PluginAuth) {
		toSerialize["pluginAuth"] = o.PluginAuth
	}
	if !utils.IsNil(o.LastLoginDate) {
		toSerialize["lastLoginDate"] = o.LastLoginDate
	}
	if !utils.IsNil(o.NoInstanceConfigWarningModal) {
		toSerialize["noInstanceConfigWarningModal"] = o.NoInstanceConfigWarningModal
	}
	if !utils.IsNil(o.NoAccountSetupWarningModal) {
		toSerialize["noAccountSetupWarningModal"] = o.NoAccountSetupWarningModal
	}
	if !utils.IsNil(o.NoWelcomeModal) {
		toSerialize["noWelcomeModal"] = o.NoWelcomeModal
	}
	if !utils.IsNil(o.NsfwPolicy) {
		toSerialize["nsfwPolicy"] = o.NsfwPolicy
	}
	if !utils.IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !utils.IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !utils.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !utils.IsNil(o.VideoChannels) {
		toSerialize["videoChannels"] = o.VideoChannels
	}
	if !utils.IsNil(o.VideoQuota) {
		toSerialize["videoQuota"] = o.VideoQuota
	}
	if !utils.IsNil(o.VideoQuotaDaily) {
		toSerialize["videoQuotaDaily"] = o.VideoQuotaDaily
	}
	if !utils.IsNil(o.P2pEnabled) {
		toSerialize["p2pEnabled"] = o.P2pEnabled
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
