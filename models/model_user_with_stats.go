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

// checks if the UserWithStats type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserWithStats{}

// UserWithStats struct for UserWithStats
type UserWithStats struct {
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
	// Count of videos published
	VideosCount *int32 `json:"videosCount,omitempty"`
	// Count of reports/abuses of which the user is a target
	AbusesCount *int32 `json:"abusesCount,omitempty"`
	// Count of reports/abuses created by the user and accepted/acted upon by the moderation team
	AbusesAcceptedCount *int32 `json:"abusesAcceptedCount,omitempty"`
	// Count of reports/abuses created by the user
	AbusesCreatedCount *int32 `json:"abusesCreatedCount,omitempty"`
	// Count of comments published
	VideoCommentsCount *int32 `json:"videoCommentsCount,omitempty"`
}

// NewUserWithStats instantiates a new UserWithStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWithStats() *UserWithStats {
	this := UserWithStats{}
	return &this
}

// NewUserWithStatsWithDefaults instantiates a new UserWithStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithStatsWithDefaults() *UserWithStats {
	this := UserWithStats{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UserWithStats) GetAccount() Account {
	if o == nil || utils.IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAccountOk() (*Account, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *UserWithStats) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *UserWithStats) SetAccount(v Account) {
	o.Account = &v
}

// GetAutoPlayNextVideo returns the AutoPlayNextVideo field value if set, zero value otherwise.
func (o *UserWithStats) GetAutoPlayNextVideo() bool {
	if o == nil || utils.IsNil(o.AutoPlayNextVideo) {
		var ret bool
		return ret
	}
	return *o.AutoPlayNextVideo
}

// GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAutoPlayNextVideoOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AutoPlayNextVideo) {
		return nil, false
	}
	return o.AutoPlayNextVideo, true
}

// HasAutoPlayNextVideo returns a boolean if a field has been set.
func (o *UserWithStats) HasAutoPlayNextVideo() bool {
	if o != nil && !utils.IsNil(o.AutoPlayNextVideo) {
		return true
	}

	return false
}

// SetAutoPlayNextVideo gets a reference to the given bool and assigns it to the AutoPlayNextVideo field.
func (o *UserWithStats) SetAutoPlayNextVideo(v bool) {
	o.AutoPlayNextVideo = &v
}

// GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field value if set, zero value otherwise.
func (o *UserWithStats) GetAutoPlayNextVideoPlaylist() bool {
	if o == nil || utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		var ret bool
		return ret
	}
	return *o.AutoPlayNextVideoPlaylist
}

// GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAutoPlayNextVideoPlaylistOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		return nil, false
	}
	return o.AutoPlayNextVideoPlaylist, true
}

// HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.
func (o *UserWithStats) HasAutoPlayNextVideoPlaylist() bool {
	if o != nil && !utils.IsNil(o.AutoPlayNextVideoPlaylist) {
		return true
	}

	return false
}

// SetAutoPlayNextVideoPlaylist gets a reference to the given bool and assigns it to the AutoPlayNextVideoPlaylist field.
func (o *UserWithStats) SetAutoPlayNextVideoPlaylist(v bool) {
	o.AutoPlayNextVideoPlaylist = &v
}

// GetAutoPlayVideo returns the AutoPlayVideo field value if set, zero value otherwise.
func (o *UserWithStats) GetAutoPlayVideo() bool {
	if o == nil || utils.IsNil(o.AutoPlayVideo) {
		var ret bool
		return ret
	}
	return *o.AutoPlayVideo
}

// GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAutoPlayVideoOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AutoPlayVideo) {
		return nil, false
	}
	return o.AutoPlayVideo, true
}

// HasAutoPlayVideo returns a boolean if a field has been set.
func (o *UserWithStats) HasAutoPlayVideo() bool {
	if o != nil && !utils.IsNil(o.AutoPlayVideo) {
		return true
	}

	return false
}

// SetAutoPlayVideo gets a reference to the given bool and assigns it to the AutoPlayVideo field.
func (o *UserWithStats) SetAutoPlayVideo(v bool) {
	o.AutoPlayVideo = &v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *UserWithStats) GetBlocked() bool {
	if o == nil || utils.IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetBlockedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *UserWithStats) HasBlocked() bool {
	if o != nil && !utils.IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *UserWithStats) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetBlockedReason returns the BlockedReason field value if set, zero value otherwise.
func (o *UserWithStats) GetBlockedReason() string {
	if o == nil || utils.IsNil(o.BlockedReason) {
		var ret string
		return ret
	}
	return *o.BlockedReason
}

// GetBlockedReasonOk returns a tuple with the BlockedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetBlockedReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BlockedReason) {
		return nil, false
	}
	return o.BlockedReason, true
}

// HasBlockedReason returns a boolean if a field has been set.
func (o *UserWithStats) HasBlockedReason() bool {
	if o != nil && !utils.IsNil(o.BlockedReason) {
		return true
	}

	return false
}

// SetBlockedReason gets a reference to the given string and assigns it to the BlockedReason field.
func (o *UserWithStats) SetBlockedReason(v string) {
	o.BlockedReason = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserWithStats) GetCreatedAt() string {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetCreatedAtOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserWithStats) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UserWithStats) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserWithStats) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserWithStats) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserWithStats) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *UserWithStats) GetEmailVerified() bool {
	if o == nil || utils.IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *UserWithStats) HasEmailVerified() bool {
	if o != nil && !utils.IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *UserWithStats) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserWithStats) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserWithStats) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserWithStats) SetId(v int32) {
	o.Id = &v
}

// GetPluginAuth returns the PluginAuth field value if set, zero value otherwise.
func (o *UserWithStats) GetPluginAuth() string {
	if o == nil || utils.IsNil(o.PluginAuth) {
		var ret string
		return ret
	}
	return *o.PluginAuth
}

// GetPluginAuthOk returns a tuple with the PluginAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetPluginAuthOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PluginAuth) {
		return nil, false
	}
	return o.PluginAuth, true
}

// HasPluginAuth returns a boolean if a field has been set.
func (o *UserWithStats) HasPluginAuth() bool {
	if o != nil && !utils.IsNil(o.PluginAuth) {
		return true
	}

	return false
}

// SetPluginAuth gets a reference to the given string and assigns it to the PluginAuth field.
func (o *UserWithStats) SetPluginAuth(v string) {
	o.PluginAuth = &v
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise.
func (o *UserWithStats) GetLastLoginDate() time.Time {
	if o == nil || utils.IsNil(o.LastLoginDate) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LastLoginDate) {
		return nil, false
	}
	return o.LastLoginDate, true
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *UserWithStats) HasLastLoginDate() bool {
	if o != nil && !utils.IsNil(o.LastLoginDate) {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given time.Time and assigns it to the LastLoginDate field.
func (o *UserWithStats) SetLastLoginDate(v time.Time) {
	o.LastLoginDate = &v
}

// GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field value if set, zero value otherwise.
func (o *UserWithStats) GetNoInstanceConfigWarningModal() bool {
	if o == nil || utils.IsNil(o.NoInstanceConfigWarningModal) {
		var ret bool
		return ret
	}
	return *o.NoInstanceConfigWarningModal
}

// GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetNoInstanceConfigWarningModalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoInstanceConfigWarningModal) {
		return nil, false
	}
	return o.NoInstanceConfigWarningModal, true
}

// HasNoInstanceConfigWarningModal returns a boolean if a field has been set.
func (o *UserWithStats) HasNoInstanceConfigWarningModal() bool {
	if o != nil && !utils.IsNil(o.NoInstanceConfigWarningModal) {
		return true
	}

	return false
}

// SetNoInstanceConfigWarningModal gets a reference to the given bool and assigns it to the NoInstanceConfigWarningModal field.
func (o *UserWithStats) SetNoInstanceConfigWarningModal(v bool) {
	o.NoInstanceConfigWarningModal = &v
}

// GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field value if set, zero value otherwise.
func (o *UserWithStats) GetNoAccountSetupWarningModal() bool {
	if o == nil || utils.IsNil(o.NoAccountSetupWarningModal) {
		var ret bool
		return ret
	}
	return *o.NoAccountSetupWarningModal
}

// GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetNoAccountSetupWarningModalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoAccountSetupWarningModal) {
		return nil, false
	}
	return o.NoAccountSetupWarningModal, true
}

// HasNoAccountSetupWarningModal returns a boolean if a field has been set.
func (o *UserWithStats) HasNoAccountSetupWarningModal() bool {
	if o != nil && !utils.IsNil(o.NoAccountSetupWarningModal) {
		return true
	}

	return false
}

// SetNoAccountSetupWarningModal gets a reference to the given bool and assigns it to the NoAccountSetupWarningModal field.
func (o *UserWithStats) SetNoAccountSetupWarningModal(v bool) {
	o.NoAccountSetupWarningModal = &v
}

// GetNoWelcomeModal returns the NoWelcomeModal field value if set, zero value otherwise.
func (o *UserWithStats) GetNoWelcomeModal() bool {
	if o == nil || utils.IsNil(o.NoWelcomeModal) {
		var ret bool
		return ret
	}
	return *o.NoWelcomeModal
}

// GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetNoWelcomeModalOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoWelcomeModal) {
		return nil, false
	}
	return o.NoWelcomeModal, true
}

// HasNoWelcomeModal returns a boolean if a field has been set.
func (o *UserWithStats) HasNoWelcomeModal() bool {
	if o != nil && !utils.IsNil(o.NoWelcomeModal) {
		return true
	}

	return false
}

// SetNoWelcomeModal gets a reference to the given bool and assigns it to the NoWelcomeModal field.
func (o *UserWithStats) SetNoWelcomeModal(v bool) {
	o.NoWelcomeModal = &v
}

// GetNsfwPolicy returns the NsfwPolicy field value if set, zero value otherwise.
func (o *UserWithStats) GetNsfwPolicy() NSFWPolicy {
	if o == nil || utils.IsNil(o.NsfwPolicy) {
		var ret NSFWPolicy
		return ret
	}
	return *o.NsfwPolicy
}

// GetNsfwPolicyOk returns a tuple with the NsfwPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetNsfwPolicyOk() (*NSFWPolicy, bool) {
	if o == nil || utils.IsNil(o.NsfwPolicy) {
		return nil, false
	}
	return o.NsfwPolicy, true
}

// HasNsfwPolicy returns a boolean if a field has been set.
func (o *UserWithStats) HasNsfwPolicy() bool {
	if o != nil && !utils.IsNil(o.NsfwPolicy) {
		return true
	}

	return false
}

// SetNsfwPolicy gets a reference to the given NSFWPolicy and assigns it to the NsfwPolicy field.
func (o *UserWithStats) SetNsfwPolicy(v NSFWPolicy) {
	o.NsfwPolicy = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserWithStats) GetRole() UserRole {
	if o == nil || utils.IsNil(o.Role) {
		var ret UserRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetRoleOk() (*UserRole, bool) {
	if o == nil || utils.IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserWithStats) HasRole() bool {
	if o != nil && !utils.IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRole and assigns it to the Role field.
func (o *UserWithStats) SetRole(v UserRole) {
	o.Role = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *UserWithStats) GetTheme() string {
	if o == nil || utils.IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetThemeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *UserWithStats) HasTheme() bool {
	if o != nil && !utils.IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *UserWithStats) SetTheme(v string) {
	o.Theme = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserWithStats) GetUsername() string {
	if o == nil || utils.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetUsernameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserWithStats) HasUsername() bool {
	if o != nil && !utils.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserWithStats) SetUsername(v string) {
	o.Username = &v
}

// GetVideoChannels returns the VideoChannels field value if set, zero value otherwise.
func (o *UserWithStats) GetVideoChannels() []VideoChannel {
	if o == nil || utils.IsNil(o.VideoChannels) {
		var ret []VideoChannel
		return ret
	}
	return o.VideoChannels
}

// GetVideoChannelsOk returns a tuple with the VideoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetVideoChannelsOk() ([]VideoChannel, bool) {
	if o == nil || utils.IsNil(o.VideoChannels) {
		return nil, false
	}
	return o.VideoChannels, true
}

// HasVideoChannels returns a boolean if a field has been set.
func (o *UserWithStats) HasVideoChannels() bool {
	if o != nil && !utils.IsNil(o.VideoChannels) {
		return true
	}

	return false
}

// SetVideoChannels gets a reference to the given []VideoChannel and assigns it to the VideoChannels field.
func (o *UserWithStats) SetVideoChannels(v []VideoChannel) {
	o.VideoChannels = v
}

// GetVideoQuota returns the VideoQuota field value if set, zero value otherwise.
func (o *UserWithStats) GetVideoQuota() int32 {
	if o == nil || utils.IsNil(o.VideoQuota) {
		var ret int32
		return ret
	}
	return *o.VideoQuota
}

// GetVideoQuotaOk returns a tuple with the VideoQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetVideoQuotaOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoQuota) {
		return nil, false
	}
	return o.VideoQuota, true
}

// HasVideoQuota returns a boolean if a field has been set.
func (o *UserWithStats) HasVideoQuota() bool {
	if o != nil && !utils.IsNil(o.VideoQuota) {
		return true
	}

	return false
}

// SetVideoQuota gets a reference to the given int32 and assigns it to the VideoQuota field.
func (o *UserWithStats) SetVideoQuota(v int32) {
	o.VideoQuota = &v
}

// GetVideoQuotaDaily returns the VideoQuotaDaily field value if set, zero value otherwise.
func (o *UserWithStats) GetVideoQuotaDaily() int32 {
	if o == nil || utils.IsNil(o.VideoQuotaDaily) {
		var ret int32
		return ret
	}
	return *o.VideoQuotaDaily
}

// GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetVideoQuotaDailyOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoQuotaDaily) {
		return nil, false
	}
	return o.VideoQuotaDaily, true
}

// HasVideoQuotaDaily returns a boolean if a field has been set.
func (o *UserWithStats) HasVideoQuotaDaily() bool {
	if o != nil && !utils.IsNil(o.VideoQuotaDaily) {
		return true
	}

	return false
}

// SetVideoQuotaDaily gets a reference to the given int32 and assigns it to the VideoQuotaDaily field.
func (o *UserWithStats) SetVideoQuotaDaily(v int32) {
	o.VideoQuotaDaily = &v
}

// GetP2pEnabled returns the P2pEnabled field value if set, zero value otherwise.
func (o *UserWithStats) GetP2pEnabled() bool {
	if o == nil || utils.IsNil(o.P2pEnabled) {
		var ret bool
		return ret
	}
	return *o.P2pEnabled
}

// GetP2pEnabledOk returns a tuple with the P2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetP2pEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.P2pEnabled) {
		return nil, false
	}
	return o.P2pEnabled, true
}

// HasP2pEnabled returns a boolean if a field has been set.
func (o *UserWithStats) HasP2pEnabled() bool {
	if o != nil && !utils.IsNil(o.P2pEnabled) {
		return true
	}

	return false
}

// SetP2pEnabled gets a reference to the given bool and assigns it to the P2pEnabled field.
func (o *UserWithStats) SetP2pEnabled(v bool) {
	o.P2pEnabled = &v
}

// GetVideosCount returns the VideosCount field value if set, zero value otherwise.
func (o *UserWithStats) GetVideosCount() int32 {
	if o == nil || utils.IsNil(o.VideosCount) {
		var ret int32
		return ret
	}
	return *o.VideosCount
}

// GetVideosCountOk returns a tuple with the VideosCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetVideosCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideosCount) {
		return nil, false
	}
	return o.VideosCount, true
}

// HasVideosCount returns a boolean if a field has been set.
func (o *UserWithStats) HasVideosCount() bool {
	if o != nil && !utils.IsNil(o.VideosCount) {
		return true
	}

	return false
}

// SetVideosCount gets a reference to the given int32 and assigns it to the VideosCount field.
func (o *UserWithStats) SetVideosCount(v int32) {
	o.VideosCount = &v
}

// GetAbusesCount returns the AbusesCount field value if set, zero value otherwise.
func (o *UserWithStats) GetAbusesCount() int32 {
	if o == nil || utils.IsNil(o.AbusesCount) {
		var ret int32
		return ret
	}
	return *o.AbusesCount
}

// GetAbusesCountOk returns a tuple with the AbusesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAbusesCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AbusesCount) {
		return nil, false
	}
	return o.AbusesCount, true
}

// HasAbusesCount returns a boolean if a field has been set.
func (o *UserWithStats) HasAbusesCount() bool {
	if o != nil && !utils.IsNil(o.AbusesCount) {
		return true
	}

	return false
}

// SetAbusesCount gets a reference to the given int32 and assigns it to the AbusesCount field.
func (o *UserWithStats) SetAbusesCount(v int32) {
	o.AbusesCount = &v
}

// GetAbusesAcceptedCount returns the AbusesAcceptedCount field value if set, zero value otherwise.
func (o *UserWithStats) GetAbusesAcceptedCount() int32 {
	if o == nil || utils.IsNil(o.AbusesAcceptedCount) {
		var ret int32
		return ret
	}
	return *o.AbusesAcceptedCount
}

// GetAbusesAcceptedCountOk returns a tuple with the AbusesAcceptedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAbusesAcceptedCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AbusesAcceptedCount) {
		return nil, false
	}
	return o.AbusesAcceptedCount, true
}

// HasAbusesAcceptedCount returns a boolean if a field has been set.
func (o *UserWithStats) HasAbusesAcceptedCount() bool {
	if o != nil && !utils.IsNil(o.AbusesAcceptedCount) {
		return true
	}

	return false
}

// SetAbusesAcceptedCount gets a reference to the given int32 and assigns it to the AbusesAcceptedCount field.
func (o *UserWithStats) SetAbusesAcceptedCount(v int32) {
	o.AbusesAcceptedCount = &v
}

// GetAbusesCreatedCount returns the AbusesCreatedCount field value if set, zero value otherwise.
func (o *UserWithStats) GetAbusesCreatedCount() int32 {
	if o == nil || utils.IsNil(o.AbusesCreatedCount) {
		var ret int32
		return ret
	}
	return *o.AbusesCreatedCount
}

// GetAbusesCreatedCountOk returns a tuple with the AbusesCreatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetAbusesCreatedCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AbusesCreatedCount) {
		return nil, false
	}
	return o.AbusesCreatedCount, true
}

// HasAbusesCreatedCount returns a boolean if a field has been set.
func (o *UserWithStats) HasAbusesCreatedCount() bool {
	if o != nil && !utils.IsNil(o.AbusesCreatedCount) {
		return true
	}

	return false
}

// SetAbusesCreatedCount gets a reference to the given int32 and assigns it to the AbusesCreatedCount field.
func (o *UserWithStats) SetAbusesCreatedCount(v int32) {
	o.AbusesCreatedCount = &v
}

// GetVideoCommentsCount returns the VideoCommentsCount field value if set, zero value otherwise.
func (o *UserWithStats) GetVideoCommentsCount() int32 {
	if o == nil || utils.IsNil(o.VideoCommentsCount) {
		var ret int32
		return ret
	}
	return *o.VideoCommentsCount
}

// GetVideoCommentsCountOk returns a tuple with the VideoCommentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithStats) GetVideoCommentsCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoCommentsCount) {
		return nil, false
	}
	return o.VideoCommentsCount, true
}

// HasVideoCommentsCount returns a boolean if a field has been set.
func (o *UserWithStats) HasVideoCommentsCount() bool {
	if o != nil && !utils.IsNil(o.VideoCommentsCount) {
		return true
	}

	return false
}

// SetVideoCommentsCount gets a reference to the given int32 and assigns it to the VideoCommentsCount field.
func (o *UserWithStats) SetVideoCommentsCount(v int32) {
	o.VideoCommentsCount = &v
}

func (o UserWithStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserWithStats) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.VideosCount) {
		toSerialize["videosCount"] = o.VideosCount
	}
	if !utils.IsNil(o.AbusesCount) {
		toSerialize["abusesCount"] = o.AbusesCount
	}
	if !utils.IsNil(o.AbusesAcceptedCount) {
		toSerialize["abusesAcceptedCount"] = o.AbusesAcceptedCount
	}
	if !utils.IsNil(o.AbusesCreatedCount) {
		toSerialize["abusesCreatedCount"] = o.AbusesCreatedCount
	}
	if !utils.IsNil(o.VideoCommentsCount) {
		toSerialize["videoCommentsCount"] = o.VideoCommentsCount
	}
	return toSerialize, nil
}

type NullableUserWithStats struct {
	value *UserWithStats
	isSet bool
}

func (v NullableUserWithStats) Get() *UserWithStats {
	return v.value
}

func (v *NullableUserWithStats) Set(val *UserWithStats) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWithStats) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWithStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWithStats(val *UserWithStats) *NullableUserWithStats {
	return &NullableUserWithStats{value: val, isSet: true}
}

func (v NullableUserWithStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWithStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
