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

// checks if the ServerStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStats{}

// ServerStats struct for ServerStats
type ServerStats struct {
	TotalUsers              *float32 `json:"totalUsers,omitempty"`
	TotalDailyActiveUsers   *float32 `json:"totalDailyActiveUsers,omitempty"`
	TotalWeeklyActiveUsers  *float32 `json:"totalWeeklyActiveUsers,omitempty"`
	TotalMonthlyActiveUsers *float32 `json:"totalMonthlyActiveUsers,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled total moderators stats
	TotalModerators *float32 `json:"totalModerators,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled total admins stats
	TotalAdmins      *float32 `json:"totalAdmins,omitempty"`
	TotalLocalVideos *float32 `json:"totalLocalVideos,omitempty"`
	// Total video views made on the instance
	TotalLocalVideoViews *float32 `json:"totalLocalVideoViews,omitempty"`
	// Total comments made by local users
	TotalLocalVideoComments               *float32                           `json:"totalLocalVideoComments,omitempty"`
	TotalLocalVideoFilesSize              *float32                           `json:"totalLocalVideoFilesSize,omitempty"`
	TotalVideos                           *float32                           `json:"totalVideos,omitempty"`
	TotalVideoComments                    *float32                           `json:"totalVideoComments,omitempty"`
	TotalLocalVideoChannels               *float32                           `json:"totalLocalVideoChannels,omitempty"`
	TotalLocalDailyActiveVideoChannels    *float32                           `json:"totalLocalDailyActiveVideoChannels,omitempty"`
	TotalLocalWeeklyActiveVideoChannels   *float32                           `json:"totalLocalWeeklyActiveVideoChannels,omitempty"`
	TotalLocalMonthlyActiveVideoChannels  *float32                           `json:"totalLocalMonthlyActiveVideoChannels,omitempty"`
	TotalLocalPlaylists                   *float32                           `json:"totalLocalPlaylists,omitempty"`
	TotalInstanceFollowers                *float32                           `json:"totalInstanceFollowers,omitempty"`
	TotalInstanceFollowing                *float32                           `json:"totalInstanceFollowing,omitempty"`
	VideosRedundancy                      []ServerStatsVideosRedundancyInner `json:"videosRedundancy,omitempty"`
	TotalActivityPubMessagesProcessed     *float32                           `json:"totalActivityPubMessagesProcessed,omitempty"`
	TotalActivityPubMessagesSuccesses     *float32                           `json:"totalActivityPubMessagesSuccesses,omitempty"`
	TotalActivityPubMessagesErrors        *float32                           `json:"totalActivityPubMessagesErrors,omitempty"`
	ActivityPubMessagesProcessedPerSecond *float32                           `json:"activityPubMessagesProcessedPerSecond,omitempty"`
	TotalActivityPubMessagesWaiting       *float32                           `json:"totalActivityPubMessagesWaiting,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled registration requests stats
	AverageRegistrationRequestResponseTimeMs *float32 `json:"averageRegistrationRequestResponseTimeMs,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled registration requests stats
	TotalRegistrationRequestsProcessed *float32 `json:"totalRegistrationRequestsProcessed,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled registration requests stats
	TotalRegistrationRequests *float32 `json:"totalRegistrationRequests,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled abuses stats
	AverageAbuseResponseTimeMs *float32 `json:"averageAbuseResponseTimeMs,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled abuses stats
	TotalAbusesProcessed *float32 `json:"totalAbusesProcessed,omitempty"`
	// **PeerTube >= 6.1** Value is null if the admin disabled abuses stats
	TotalAbuses *float32 `json:"totalAbuses,omitempty"`
}

// NewServerStats instantiates a new ServerStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStats() *ServerStats {
	this := ServerStats{}
	return &this
}

// NewServerStatsWithDefaults instantiates a new ServerStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatsWithDefaults() *ServerStats {
	this := ServerStats{}
	return &this
}

// GetTotalUsers returns the TotalUsers field value if set, zero value otherwise.
func (o *ServerStats) GetTotalUsers() float32 {
	if o == nil || IsNil(o.TotalUsers) {
		var ret float32
		return ret
	}
	return *o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalUsersOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalUsers) {
		return nil, false
	}
	return o.TotalUsers, true
}

// HasTotalUsers returns a boolean if a field has been set.
func (o *ServerStats) HasTotalUsers() bool {
	if o != nil && !IsNil(o.TotalUsers) {
		return true
	}

	return false
}

// SetTotalUsers gets a reference to the given float32 and assigns it to the TotalUsers field.
func (o *ServerStats) SetTotalUsers(v float32) {
	o.TotalUsers = &v
}

// GetTotalDailyActiveUsers returns the TotalDailyActiveUsers field value if set, zero value otherwise.
func (o *ServerStats) GetTotalDailyActiveUsers() float32 {
	if o == nil || IsNil(o.TotalDailyActiveUsers) {
		var ret float32
		return ret
	}
	return *o.TotalDailyActiveUsers
}

// GetTotalDailyActiveUsersOk returns a tuple with the TotalDailyActiveUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalDailyActiveUsersOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalDailyActiveUsers) {
		return nil, false
	}
	return o.TotalDailyActiveUsers, true
}

// HasTotalDailyActiveUsers returns a boolean if a field has been set.
func (o *ServerStats) HasTotalDailyActiveUsers() bool {
	if o != nil && !IsNil(o.TotalDailyActiveUsers) {
		return true
	}

	return false
}

// SetTotalDailyActiveUsers gets a reference to the given float32 and assigns it to the TotalDailyActiveUsers field.
func (o *ServerStats) SetTotalDailyActiveUsers(v float32) {
	o.TotalDailyActiveUsers = &v
}

// GetTotalWeeklyActiveUsers returns the TotalWeeklyActiveUsers field value if set, zero value otherwise.
func (o *ServerStats) GetTotalWeeklyActiveUsers() float32 {
	if o == nil || IsNil(o.TotalWeeklyActiveUsers) {
		var ret float32
		return ret
	}
	return *o.TotalWeeklyActiveUsers
}

// GetTotalWeeklyActiveUsersOk returns a tuple with the TotalWeeklyActiveUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalWeeklyActiveUsersOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalWeeklyActiveUsers) {
		return nil, false
	}
	return o.TotalWeeklyActiveUsers, true
}

// HasTotalWeeklyActiveUsers returns a boolean if a field has been set.
func (o *ServerStats) HasTotalWeeklyActiveUsers() bool {
	if o != nil && !IsNil(o.TotalWeeklyActiveUsers) {
		return true
	}

	return false
}

// SetTotalWeeklyActiveUsers gets a reference to the given float32 and assigns it to the TotalWeeklyActiveUsers field.
func (o *ServerStats) SetTotalWeeklyActiveUsers(v float32) {
	o.TotalWeeklyActiveUsers = &v
}

// GetTotalMonthlyActiveUsers returns the TotalMonthlyActiveUsers field value if set, zero value otherwise.
func (o *ServerStats) GetTotalMonthlyActiveUsers() float32 {
	if o == nil || IsNil(o.TotalMonthlyActiveUsers) {
		var ret float32
		return ret
	}
	return *o.TotalMonthlyActiveUsers
}

// GetTotalMonthlyActiveUsersOk returns a tuple with the TotalMonthlyActiveUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalMonthlyActiveUsersOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalMonthlyActiveUsers) {
		return nil, false
	}
	return o.TotalMonthlyActiveUsers, true
}

// HasTotalMonthlyActiveUsers returns a boolean if a field has been set.
func (o *ServerStats) HasTotalMonthlyActiveUsers() bool {
	if o != nil && !IsNil(o.TotalMonthlyActiveUsers) {
		return true
	}

	return false
}

// SetTotalMonthlyActiveUsers gets a reference to the given float32 and assigns it to the TotalMonthlyActiveUsers field.
func (o *ServerStats) SetTotalMonthlyActiveUsers(v float32) {
	o.TotalMonthlyActiveUsers = &v
}

// GetTotalModerators returns the TotalModerators field value if set, zero value otherwise.
func (o *ServerStats) GetTotalModerators() float32 {
	if o == nil || IsNil(o.TotalModerators) {
		var ret float32
		return ret
	}
	return *o.TotalModerators
}

// GetTotalModeratorsOk returns a tuple with the TotalModerators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalModeratorsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalModerators) {
		return nil, false
	}
	return o.TotalModerators, true
}

// HasTotalModerators returns a boolean if a field has been set.
func (o *ServerStats) HasTotalModerators() bool {
	if o != nil && !IsNil(o.TotalModerators) {
		return true
	}

	return false
}

// SetTotalModerators gets a reference to the given float32 and assigns it to the TotalModerators field.
func (o *ServerStats) SetTotalModerators(v float32) {
	o.TotalModerators = &v
}

// GetTotalAdmins returns the TotalAdmins field value if set, zero value otherwise.
func (o *ServerStats) GetTotalAdmins() float32 {
	if o == nil || IsNil(o.TotalAdmins) {
		var ret float32
		return ret
	}
	return *o.TotalAdmins
}

// GetTotalAdminsOk returns a tuple with the TotalAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalAdminsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalAdmins) {
		return nil, false
	}
	return o.TotalAdmins, true
}

// HasTotalAdmins returns a boolean if a field has been set.
func (o *ServerStats) HasTotalAdmins() bool {
	if o != nil && !IsNil(o.TotalAdmins) {
		return true
	}

	return false
}

// SetTotalAdmins gets a reference to the given float32 and assigns it to the TotalAdmins field.
func (o *ServerStats) SetTotalAdmins(v float32) {
	o.TotalAdmins = &v
}

// GetTotalLocalVideos returns the TotalLocalVideos field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalVideos() float32 {
	if o == nil || IsNil(o.TotalLocalVideos) {
		var ret float32
		return ret
	}
	return *o.TotalLocalVideos
}

// GetTotalLocalVideosOk returns a tuple with the TotalLocalVideos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalVideosOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalVideos) {
		return nil, false
	}
	return o.TotalLocalVideos, true
}

// HasTotalLocalVideos returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalVideos() bool {
	if o != nil && !IsNil(o.TotalLocalVideos) {
		return true
	}

	return false
}

// SetTotalLocalVideos gets a reference to the given float32 and assigns it to the TotalLocalVideos field.
func (o *ServerStats) SetTotalLocalVideos(v float32) {
	o.TotalLocalVideos = &v
}

// GetTotalLocalVideoViews returns the TotalLocalVideoViews field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalVideoViews() float32 {
	if o == nil || IsNil(o.TotalLocalVideoViews) {
		var ret float32
		return ret
	}
	return *o.TotalLocalVideoViews
}

// GetTotalLocalVideoViewsOk returns a tuple with the TotalLocalVideoViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalVideoViewsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalVideoViews) {
		return nil, false
	}
	return o.TotalLocalVideoViews, true
}

// HasTotalLocalVideoViews returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalVideoViews() bool {
	if o != nil && !IsNil(o.TotalLocalVideoViews) {
		return true
	}

	return false
}

// SetTotalLocalVideoViews gets a reference to the given float32 and assigns it to the TotalLocalVideoViews field.
func (o *ServerStats) SetTotalLocalVideoViews(v float32) {
	o.TotalLocalVideoViews = &v
}

// GetTotalLocalVideoComments returns the TotalLocalVideoComments field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalVideoComments() float32 {
	if o == nil || IsNil(o.TotalLocalVideoComments) {
		var ret float32
		return ret
	}
	return *o.TotalLocalVideoComments
}

// GetTotalLocalVideoCommentsOk returns a tuple with the TotalLocalVideoComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalVideoCommentsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalVideoComments) {
		return nil, false
	}
	return o.TotalLocalVideoComments, true
}

// HasTotalLocalVideoComments returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalVideoComments() bool {
	if o != nil && !IsNil(o.TotalLocalVideoComments) {
		return true
	}

	return false
}

// SetTotalLocalVideoComments gets a reference to the given float32 and assigns it to the TotalLocalVideoComments field.
func (o *ServerStats) SetTotalLocalVideoComments(v float32) {
	o.TotalLocalVideoComments = &v
}

// GetTotalLocalVideoFilesSize returns the TotalLocalVideoFilesSize field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalVideoFilesSize() float32 {
	if o == nil || IsNil(o.TotalLocalVideoFilesSize) {
		var ret float32
		return ret
	}
	return *o.TotalLocalVideoFilesSize
}

// GetTotalLocalVideoFilesSizeOk returns a tuple with the TotalLocalVideoFilesSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalVideoFilesSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalVideoFilesSize) {
		return nil, false
	}
	return o.TotalLocalVideoFilesSize, true
}

// HasTotalLocalVideoFilesSize returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalVideoFilesSize() bool {
	if o != nil && !IsNil(o.TotalLocalVideoFilesSize) {
		return true
	}

	return false
}

// SetTotalLocalVideoFilesSize gets a reference to the given float32 and assigns it to the TotalLocalVideoFilesSize field.
func (o *ServerStats) SetTotalLocalVideoFilesSize(v float32) {
	o.TotalLocalVideoFilesSize = &v
}

// GetTotalVideos returns the TotalVideos field value if set, zero value otherwise.
func (o *ServerStats) GetTotalVideos() float32 {
	if o == nil || IsNil(o.TotalVideos) {
		var ret float32
		return ret
	}
	return *o.TotalVideos
}

// GetTotalVideosOk returns a tuple with the TotalVideos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalVideosOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalVideos) {
		return nil, false
	}
	return o.TotalVideos, true
}

// HasTotalVideos returns a boolean if a field has been set.
func (o *ServerStats) HasTotalVideos() bool {
	if o != nil && !IsNil(o.TotalVideos) {
		return true
	}

	return false
}

// SetTotalVideos gets a reference to the given float32 and assigns it to the TotalVideos field.
func (o *ServerStats) SetTotalVideos(v float32) {
	o.TotalVideos = &v
}

// GetTotalVideoComments returns the TotalVideoComments field value if set, zero value otherwise.
func (o *ServerStats) GetTotalVideoComments() float32 {
	if o == nil || IsNil(o.TotalVideoComments) {
		var ret float32
		return ret
	}
	return *o.TotalVideoComments
}

// GetTotalVideoCommentsOk returns a tuple with the TotalVideoComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalVideoCommentsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalVideoComments) {
		return nil, false
	}
	return o.TotalVideoComments, true
}

// HasTotalVideoComments returns a boolean if a field has been set.
func (o *ServerStats) HasTotalVideoComments() bool {
	if o != nil && !IsNil(o.TotalVideoComments) {
		return true
	}

	return false
}

// SetTotalVideoComments gets a reference to the given float32 and assigns it to the TotalVideoComments field.
func (o *ServerStats) SetTotalVideoComments(v float32) {
	o.TotalVideoComments = &v
}

// GetTotalLocalVideoChannels returns the TotalLocalVideoChannels field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalVideoChannels() float32 {
	if o == nil || IsNil(o.TotalLocalVideoChannels) {
		var ret float32
		return ret
	}
	return *o.TotalLocalVideoChannels
}

// GetTotalLocalVideoChannelsOk returns a tuple with the TotalLocalVideoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalVideoChannelsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalVideoChannels) {
		return nil, false
	}
	return o.TotalLocalVideoChannels, true
}

// HasTotalLocalVideoChannels returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalVideoChannels() bool {
	if o != nil && !IsNil(o.TotalLocalVideoChannels) {
		return true
	}

	return false
}

// SetTotalLocalVideoChannels gets a reference to the given float32 and assigns it to the TotalLocalVideoChannels field.
func (o *ServerStats) SetTotalLocalVideoChannels(v float32) {
	o.TotalLocalVideoChannels = &v
}

// GetTotalLocalDailyActiveVideoChannels returns the TotalLocalDailyActiveVideoChannels field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalDailyActiveVideoChannels() float32 {
	if o == nil || IsNil(o.TotalLocalDailyActiveVideoChannels) {
		var ret float32
		return ret
	}
	return *o.TotalLocalDailyActiveVideoChannels
}

// GetTotalLocalDailyActiveVideoChannelsOk returns a tuple with the TotalLocalDailyActiveVideoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalDailyActiveVideoChannelsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalDailyActiveVideoChannels) {
		return nil, false
	}
	return o.TotalLocalDailyActiveVideoChannels, true
}

// HasTotalLocalDailyActiveVideoChannels returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalDailyActiveVideoChannels() bool {
	if o != nil && !IsNil(o.TotalLocalDailyActiveVideoChannels) {
		return true
	}

	return false
}

// SetTotalLocalDailyActiveVideoChannels gets a reference to the given float32 and assigns it to the TotalLocalDailyActiveVideoChannels field.
func (o *ServerStats) SetTotalLocalDailyActiveVideoChannels(v float32) {
	o.TotalLocalDailyActiveVideoChannels = &v
}

// GetTotalLocalWeeklyActiveVideoChannels returns the TotalLocalWeeklyActiveVideoChannels field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalWeeklyActiveVideoChannels() float32 {
	if o == nil || IsNil(o.TotalLocalWeeklyActiveVideoChannels) {
		var ret float32
		return ret
	}
	return *o.TotalLocalWeeklyActiveVideoChannels
}

// GetTotalLocalWeeklyActiveVideoChannelsOk returns a tuple with the TotalLocalWeeklyActiveVideoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalWeeklyActiveVideoChannelsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalWeeklyActiveVideoChannels) {
		return nil, false
	}
	return o.TotalLocalWeeklyActiveVideoChannels, true
}

// HasTotalLocalWeeklyActiveVideoChannels returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalWeeklyActiveVideoChannels() bool {
	if o != nil && !IsNil(o.TotalLocalWeeklyActiveVideoChannels) {
		return true
	}

	return false
}

// SetTotalLocalWeeklyActiveVideoChannels gets a reference to the given float32 and assigns it to the TotalLocalWeeklyActiveVideoChannels field.
func (o *ServerStats) SetTotalLocalWeeklyActiveVideoChannels(v float32) {
	o.TotalLocalWeeklyActiveVideoChannels = &v
}

// GetTotalLocalMonthlyActiveVideoChannels returns the TotalLocalMonthlyActiveVideoChannels field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalMonthlyActiveVideoChannels() float32 {
	if o == nil || IsNil(o.TotalLocalMonthlyActiveVideoChannels) {
		var ret float32
		return ret
	}
	return *o.TotalLocalMonthlyActiveVideoChannels
}

// GetTotalLocalMonthlyActiveVideoChannelsOk returns a tuple with the TotalLocalMonthlyActiveVideoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalMonthlyActiveVideoChannelsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalMonthlyActiveVideoChannels) {
		return nil, false
	}
	return o.TotalLocalMonthlyActiveVideoChannels, true
}

// HasTotalLocalMonthlyActiveVideoChannels returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalMonthlyActiveVideoChannels() bool {
	if o != nil && !IsNil(o.TotalLocalMonthlyActiveVideoChannels) {
		return true
	}

	return false
}

// SetTotalLocalMonthlyActiveVideoChannels gets a reference to the given float32 and assigns it to the TotalLocalMonthlyActiveVideoChannels field.
func (o *ServerStats) SetTotalLocalMonthlyActiveVideoChannels(v float32) {
	o.TotalLocalMonthlyActiveVideoChannels = &v
}

// GetTotalLocalPlaylists returns the TotalLocalPlaylists field value if set, zero value otherwise.
func (o *ServerStats) GetTotalLocalPlaylists() float32 {
	if o == nil || IsNil(o.TotalLocalPlaylists) {
		var ret float32
		return ret
	}
	return *o.TotalLocalPlaylists
}

// GetTotalLocalPlaylistsOk returns a tuple with the TotalLocalPlaylists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalLocalPlaylistsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalPlaylists) {
		return nil, false
	}
	return o.TotalLocalPlaylists, true
}

// HasTotalLocalPlaylists returns a boolean if a field has been set.
func (o *ServerStats) HasTotalLocalPlaylists() bool {
	if o != nil && !IsNil(o.TotalLocalPlaylists) {
		return true
	}

	return false
}

// SetTotalLocalPlaylists gets a reference to the given float32 and assigns it to the TotalLocalPlaylists field.
func (o *ServerStats) SetTotalLocalPlaylists(v float32) {
	o.TotalLocalPlaylists = &v
}

// GetTotalInstanceFollowers returns the TotalInstanceFollowers field value if set, zero value otherwise.
func (o *ServerStats) GetTotalInstanceFollowers() float32 {
	if o == nil || IsNil(o.TotalInstanceFollowers) {
		var ret float32
		return ret
	}
	return *o.TotalInstanceFollowers
}

// GetTotalInstanceFollowersOk returns a tuple with the TotalInstanceFollowers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalInstanceFollowersOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalInstanceFollowers) {
		return nil, false
	}
	return o.TotalInstanceFollowers, true
}

// HasTotalInstanceFollowers returns a boolean if a field has been set.
func (o *ServerStats) HasTotalInstanceFollowers() bool {
	if o != nil && !IsNil(o.TotalInstanceFollowers) {
		return true
	}

	return false
}

// SetTotalInstanceFollowers gets a reference to the given float32 and assigns it to the TotalInstanceFollowers field.
func (o *ServerStats) SetTotalInstanceFollowers(v float32) {
	o.TotalInstanceFollowers = &v
}

// GetTotalInstanceFollowing returns the TotalInstanceFollowing field value if set, zero value otherwise.
func (o *ServerStats) GetTotalInstanceFollowing() float32 {
	if o == nil || IsNil(o.TotalInstanceFollowing) {
		var ret float32
		return ret
	}
	return *o.TotalInstanceFollowing
}

// GetTotalInstanceFollowingOk returns a tuple with the TotalInstanceFollowing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalInstanceFollowingOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalInstanceFollowing) {
		return nil, false
	}
	return o.TotalInstanceFollowing, true
}

// HasTotalInstanceFollowing returns a boolean if a field has been set.
func (o *ServerStats) HasTotalInstanceFollowing() bool {
	if o != nil && !IsNil(o.TotalInstanceFollowing) {
		return true
	}

	return false
}

// SetTotalInstanceFollowing gets a reference to the given float32 and assigns it to the TotalInstanceFollowing field.
func (o *ServerStats) SetTotalInstanceFollowing(v float32) {
	o.TotalInstanceFollowing = &v
}

// GetVideosRedundancy returns the VideosRedundancy field value if set, zero value otherwise.
func (o *ServerStats) GetVideosRedundancy() []ServerStatsVideosRedundancyInner {
	if o == nil || IsNil(o.VideosRedundancy) {
		var ret []ServerStatsVideosRedundancyInner
		return ret
	}
	return o.VideosRedundancy
}

// GetVideosRedundancyOk returns a tuple with the VideosRedundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetVideosRedundancyOk() ([]ServerStatsVideosRedundancyInner, bool) {
	if o == nil || IsNil(o.VideosRedundancy) {
		return nil, false
	}
	return o.VideosRedundancy, true
}

// HasVideosRedundancy returns a boolean if a field has been set.
func (o *ServerStats) HasVideosRedundancy() bool {
	if o != nil && !IsNil(o.VideosRedundancy) {
		return true
	}

	return false
}

// SetVideosRedundancy gets a reference to the given []ServerStatsVideosRedundancyInner and assigns it to the VideosRedundancy field.
func (o *ServerStats) SetVideosRedundancy(v []ServerStatsVideosRedundancyInner) {
	o.VideosRedundancy = v
}

// GetTotalActivityPubMessagesProcessed returns the TotalActivityPubMessagesProcessed field value if set, zero value otherwise.
func (o *ServerStats) GetTotalActivityPubMessagesProcessed() float32 {
	if o == nil || IsNil(o.TotalActivityPubMessagesProcessed) {
		var ret float32
		return ret
	}
	return *o.TotalActivityPubMessagesProcessed
}

// GetTotalActivityPubMessagesProcessedOk returns a tuple with the TotalActivityPubMessagesProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalActivityPubMessagesProcessedOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalActivityPubMessagesProcessed) {
		return nil, false
	}
	return o.TotalActivityPubMessagesProcessed, true
}

// HasTotalActivityPubMessagesProcessed returns a boolean if a field has been set.
func (o *ServerStats) HasTotalActivityPubMessagesProcessed() bool {
	if o != nil && !IsNil(o.TotalActivityPubMessagesProcessed) {
		return true
	}

	return false
}

// SetTotalActivityPubMessagesProcessed gets a reference to the given float32 and assigns it to the TotalActivityPubMessagesProcessed field.
func (o *ServerStats) SetTotalActivityPubMessagesProcessed(v float32) {
	o.TotalActivityPubMessagesProcessed = &v
}

// GetTotalActivityPubMessagesSuccesses returns the TotalActivityPubMessagesSuccesses field value if set, zero value otherwise.
func (o *ServerStats) GetTotalActivityPubMessagesSuccesses() float32 {
	if o == nil || IsNil(o.TotalActivityPubMessagesSuccesses) {
		var ret float32
		return ret
	}
	return *o.TotalActivityPubMessagesSuccesses
}

// GetTotalActivityPubMessagesSuccessesOk returns a tuple with the TotalActivityPubMessagesSuccesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalActivityPubMessagesSuccessesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalActivityPubMessagesSuccesses) {
		return nil, false
	}
	return o.TotalActivityPubMessagesSuccesses, true
}

// HasTotalActivityPubMessagesSuccesses returns a boolean if a field has been set.
func (o *ServerStats) HasTotalActivityPubMessagesSuccesses() bool {
	if o != nil && !IsNil(o.TotalActivityPubMessagesSuccesses) {
		return true
	}

	return false
}

// SetTotalActivityPubMessagesSuccesses gets a reference to the given float32 and assigns it to the TotalActivityPubMessagesSuccesses field.
func (o *ServerStats) SetTotalActivityPubMessagesSuccesses(v float32) {
	o.TotalActivityPubMessagesSuccesses = &v
}

// GetTotalActivityPubMessagesErrors returns the TotalActivityPubMessagesErrors field value if set, zero value otherwise.
func (o *ServerStats) GetTotalActivityPubMessagesErrors() float32 {
	if o == nil || IsNil(o.TotalActivityPubMessagesErrors) {
		var ret float32
		return ret
	}
	return *o.TotalActivityPubMessagesErrors
}

// GetTotalActivityPubMessagesErrorsOk returns a tuple with the TotalActivityPubMessagesErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalActivityPubMessagesErrorsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalActivityPubMessagesErrors) {
		return nil, false
	}
	return o.TotalActivityPubMessagesErrors, true
}

// HasTotalActivityPubMessagesErrors returns a boolean if a field has been set.
func (o *ServerStats) HasTotalActivityPubMessagesErrors() bool {
	if o != nil && !IsNil(o.TotalActivityPubMessagesErrors) {
		return true
	}

	return false
}

// SetTotalActivityPubMessagesErrors gets a reference to the given float32 and assigns it to the TotalActivityPubMessagesErrors field.
func (o *ServerStats) SetTotalActivityPubMessagesErrors(v float32) {
	o.TotalActivityPubMessagesErrors = &v
}

// GetActivityPubMessagesProcessedPerSecond returns the ActivityPubMessagesProcessedPerSecond field value if set, zero value otherwise.
func (o *ServerStats) GetActivityPubMessagesProcessedPerSecond() float32 {
	if o == nil || IsNil(o.ActivityPubMessagesProcessedPerSecond) {
		var ret float32
		return ret
	}
	return *o.ActivityPubMessagesProcessedPerSecond
}

// GetActivityPubMessagesProcessedPerSecondOk returns a tuple with the ActivityPubMessagesProcessedPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetActivityPubMessagesProcessedPerSecondOk() (*float32, bool) {
	if o == nil || IsNil(o.ActivityPubMessagesProcessedPerSecond) {
		return nil, false
	}
	return o.ActivityPubMessagesProcessedPerSecond, true
}

// HasActivityPubMessagesProcessedPerSecond returns a boolean if a field has been set.
func (o *ServerStats) HasActivityPubMessagesProcessedPerSecond() bool {
	if o != nil && !IsNil(o.ActivityPubMessagesProcessedPerSecond) {
		return true
	}

	return false
}

// SetActivityPubMessagesProcessedPerSecond gets a reference to the given float32 and assigns it to the ActivityPubMessagesProcessedPerSecond field.
func (o *ServerStats) SetActivityPubMessagesProcessedPerSecond(v float32) {
	o.ActivityPubMessagesProcessedPerSecond = &v
}

// GetTotalActivityPubMessagesWaiting returns the TotalActivityPubMessagesWaiting field value if set, zero value otherwise.
func (o *ServerStats) GetTotalActivityPubMessagesWaiting() float32 {
	if o == nil || IsNil(o.TotalActivityPubMessagesWaiting) {
		var ret float32
		return ret
	}
	return *o.TotalActivityPubMessagesWaiting
}

// GetTotalActivityPubMessagesWaitingOk returns a tuple with the TotalActivityPubMessagesWaiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalActivityPubMessagesWaitingOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalActivityPubMessagesWaiting) {
		return nil, false
	}
	return o.TotalActivityPubMessagesWaiting, true
}

// HasTotalActivityPubMessagesWaiting returns a boolean if a field has been set.
func (o *ServerStats) HasTotalActivityPubMessagesWaiting() bool {
	if o != nil && !IsNil(o.TotalActivityPubMessagesWaiting) {
		return true
	}

	return false
}

// SetTotalActivityPubMessagesWaiting gets a reference to the given float32 and assigns it to the TotalActivityPubMessagesWaiting field.
func (o *ServerStats) SetTotalActivityPubMessagesWaiting(v float32) {
	o.TotalActivityPubMessagesWaiting = &v
}

// GetAverageRegistrationRequestResponseTimeMs returns the AverageRegistrationRequestResponseTimeMs field value if set, zero value otherwise.
func (o *ServerStats) GetAverageRegistrationRequestResponseTimeMs() float32 {
	if o == nil || IsNil(o.AverageRegistrationRequestResponseTimeMs) {
		var ret float32
		return ret
	}
	return *o.AverageRegistrationRequestResponseTimeMs
}

// GetAverageRegistrationRequestResponseTimeMsOk returns a tuple with the AverageRegistrationRequestResponseTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetAverageRegistrationRequestResponseTimeMsOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageRegistrationRequestResponseTimeMs) {
		return nil, false
	}
	return o.AverageRegistrationRequestResponseTimeMs, true
}

// HasAverageRegistrationRequestResponseTimeMs returns a boolean if a field has been set.
func (o *ServerStats) HasAverageRegistrationRequestResponseTimeMs() bool {
	if o != nil && !IsNil(o.AverageRegistrationRequestResponseTimeMs) {
		return true
	}

	return false
}

// SetAverageRegistrationRequestResponseTimeMs gets a reference to the given float32 and assigns it to the AverageRegistrationRequestResponseTimeMs field.
func (o *ServerStats) SetAverageRegistrationRequestResponseTimeMs(v float32) {
	o.AverageRegistrationRequestResponseTimeMs = &v
}

// GetTotalRegistrationRequestsProcessed returns the TotalRegistrationRequestsProcessed field value if set, zero value otherwise.
func (o *ServerStats) GetTotalRegistrationRequestsProcessed() float32 {
	if o == nil || IsNil(o.TotalRegistrationRequestsProcessed) {
		var ret float32
		return ret
	}
	return *o.TotalRegistrationRequestsProcessed
}

// GetTotalRegistrationRequestsProcessedOk returns a tuple with the TotalRegistrationRequestsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalRegistrationRequestsProcessedOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRegistrationRequestsProcessed) {
		return nil, false
	}
	return o.TotalRegistrationRequestsProcessed, true
}

// HasTotalRegistrationRequestsProcessed returns a boolean if a field has been set.
func (o *ServerStats) HasTotalRegistrationRequestsProcessed() bool {
	if o != nil && !IsNil(o.TotalRegistrationRequestsProcessed) {
		return true
	}

	return false
}

// SetTotalRegistrationRequestsProcessed gets a reference to the given float32 and assigns it to the TotalRegistrationRequestsProcessed field.
func (o *ServerStats) SetTotalRegistrationRequestsProcessed(v float32) {
	o.TotalRegistrationRequestsProcessed = &v
}

// GetTotalRegistrationRequests returns the TotalRegistrationRequests field value if set, zero value otherwise.
func (o *ServerStats) GetTotalRegistrationRequests() float32 {
	if o == nil || IsNil(o.TotalRegistrationRequests) {
		var ret float32
		return ret
	}
	return *o.TotalRegistrationRequests
}

// GetTotalRegistrationRequestsOk returns a tuple with the TotalRegistrationRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalRegistrationRequestsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRegistrationRequests) {
		return nil, false
	}
	return o.TotalRegistrationRequests, true
}

// HasTotalRegistrationRequests returns a boolean if a field has been set.
func (o *ServerStats) HasTotalRegistrationRequests() bool {
	if o != nil && !IsNil(o.TotalRegistrationRequests) {
		return true
	}

	return false
}

// SetTotalRegistrationRequests gets a reference to the given float32 and assigns it to the TotalRegistrationRequests field.
func (o *ServerStats) SetTotalRegistrationRequests(v float32) {
	o.TotalRegistrationRequests = &v
}

// GetAverageAbuseResponseTimeMs returns the AverageAbuseResponseTimeMs field value if set, zero value otherwise.
func (o *ServerStats) GetAverageAbuseResponseTimeMs() float32 {
	if o == nil || IsNil(o.AverageAbuseResponseTimeMs) {
		var ret float32
		return ret
	}
	return *o.AverageAbuseResponseTimeMs
}

// GetAverageAbuseResponseTimeMsOk returns a tuple with the AverageAbuseResponseTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetAverageAbuseResponseTimeMsOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageAbuseResponseTimeMs) {
		return nil, false
	}
	return o.AverageAbuseResponseTimeMs, true
}

// HasAverageAbuseResponseTimeMs returns a boolean if a field has been set.
func (o *ServerStats) HasAverageAbuseResponseTimeMs() bool {
	if o != nil && !IsNil(o.AverageAbuseResponseTimeMs) {
		return true
	}

	return false
}

// SetAverageAbuseResponseTimeMs gets a reference to the given float32 and assigns it to the AverageAbuseResponseTimeMs field.
func (o *ServerStats) SetAverageAbuseResponseTimeMs(v float32) {
	o.AverageAbuseResponseTimeMs = &v
}

// GetTotalAbusesProcessed returns the TotalAbusesProcessed field value if set, zero value otherwise.
func (o *ServerStats) GetTotalAbusesProcessed() float32 {
	if o == nil || IsNil(o.TotalAbusesProcessed) {
		var ret float32
		return ret
	}
	return *o.TotalAbusesProcessed
}

// GetTotalAbusesProcessedOk returns a tuple with the TotalAbusesProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalAbusesProcessedOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalAbusesProcessed) {
		return nil, false
	}
	return o.TotalAbusesProcessed, true
}

// HasTotalAbusesProcessed returns a boolean if a field has been set.
func (o *ServerStats) HasTotalAbusesProcessed() bool {
	if o != nil && !IsNil(o.TotalAbusesProcessed) {
		return true
	}

	return false
}

// SetTotalAbusesProcessed gets a reference to the given float32 and assigns it to the TotalAbusesProcessed field.
func (o *ServerStats) SetTotalAbusesProcessed(v float32) {
	o.TotalAbusesProcessed = &v
}

// GetTotalAbuses returns the TotalAbuses field value if set, zero value otherwise.
func (o *ServerStats) GetTotalAbuses() float32 {
	if o == nil || IsNil(o.TotalAbuses) {
		var ret float32
		return ret
	}
	return *o.TotalAbuses
}

// GetTotalAbusesOk returns a tuple with the TotalAbuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalAbusesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalAbuses) {
		return nil, false
	}
	return o.TotalAbuses, true
}

// HasTotalAbuses returns a boolean if a field has been set.
func (o *ServerStats) HasTotalAbuses() bool {
	if o != nil && !IsNil(o.TotalAbuses) {
		return true
	}

	return false
}

// SetTotalAbuses gets a reference to the given float32 and assigns it to the TotalAbuses field.
func (o *ServerStats) SetTotalAbuses(v float32) {
	o.TotalAbuses = &v
}

func (o ServerStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalUsers) {
		toSerialize["totalUsers"] = o.TotalUsers
	}
	if !IsNil(o.TotalDailyActiveUsers) {
		toSerialize["totalDailyActiveUsers"] = o.TotalDailyActiveUsers
	}
	if !IsNil(o.TotalWeeklyActiveUsers) {
		toSerialize["totalWeeklyActiveUsers"] = o.TotalWeeklyActiveUsers
	}
	if !IsNil(o.TotalMonthlyActiveUsers) {
		toSerialize["totalMonthlyActiveUsers"] = o.TotalMonthlyActiveUsers
	}
	if !IsNil(o.TotalModerators) {
		toSerialize["totalModerators"] = o.TotalModerators
	}
	if !IsNil(o.TotalAdmins) {
		toSerialize["totalAdmins"] = o.TotalAdmins
	}
	if !IsNil(o.TotalLocalVideos) {
		toSerialize["totalLocalVideos"] = o.TotalLocalVideos
	}
	if !IsNil(o.TotalLocalVideoViews) {
		toSerialize["totalLocalVideoViews"] = o.TotalLocalVideoViews
	}
	if !IsNil(o.TotalLocalVideoComments) {
		toSerialize["totalLocalVideoComments"] = o.TotalLocalVideoComments
	}
	if !IsNil(o.TotalLocalVideoFilesSize) {
		toSerialize["totalLocalVideoFilesSize"] = o.TotalLocalVideoFilesSize
	}
	if !IsNil(o.TotalVideos) {
		toSerialize["totalVideos"] = o.TotalVideos
	}
	if !IsNil(o.TotalVideoComments) {
		toSerialize["totalVideoComments"] = o.TotalVideoComments
	}
	if !IsNil(o.TotalLocalVideoChannels) {
		toSerialize["totalLocalVideoChannels"] = o.TotalLocalVideoChannels
	}
	if !IsNil(o.TotalLocalDailyActiveVideoChannels) {
		toSerialize["totalLocalDailyActiveVideoChannels"] = o.TotalLocalDailyActiveVideoChannels
	}
	if !IsNil(o.TotalLocalWeeklyActiveVideoChannels) {
		toSerialize["totalLocalWeeklyActiveVideoChannels"] = o.TotalLocalWeeklyActiveVideoChannels
	}
	if !IsNil(o.TotalLocalMonthlyActiveVideoChannels) {
		toSerialize["totalLocalMonthlyActiveVideoChannels"] = o.TotalLocalMonthlyActiveVideoChannels
	}
	if !IsNil(o.TotalLocalPlaylists) {
		toSerialize["totalLocalPlaylists"] = o.TotalLocalPlaylists
	}
	if !IsNil(o.TotalInstanceFollowers) {
		toSerialize["totalInstanceFollowers"] = o.TotalInstanceFollowers
	}
	if !IsNil(o.TotalInstanceFollowing) {
		toSerialize["totalInstanceFollowing"] = o.TotalInstanceFollowing
	}
	if !IsNil(o.VideosRedundancy) {
		toSerialize["videosRedundancy"] = o.VideosRedundancy
	}
	if !IsNil(o.TotalActivityPubMessagesProcessed) {
		toSerialize["totalActivityPubMessagesProcessed"] = o.TotalActivityPubMessagesProcessed
	}
	if !IsNil(o.TotalActivityPubMessagesSuccesses) {
		toSerialize["totalActivityPubMessagesSuccesses"] = o.TotalActivityPubMessagesSuccesses
	}
	if !IsNil(o.TotalActivityPubMessagesErrors) {
		toSerialize["totalActivityPubMessagesErrors"] = o.TotalActivityPubMessagesErrors
	}
	if !IsNil(o.ActivityPubMessagesProcessedPerSecond) {
		toSerialize["activityPubMessagesProcessedPerSecond"] = o.ActivityPubMessagesProcessedPerSecond
	}
	if !IsNil(o.TotalActivityPubMessagesWaiting) {
		toSerialize["totalActivityPubMessagesWaiting"] = o.TotalActivityPubMessagesWaiting
	}
	if !IsNil(o.AverageRegistrationRequestResponseTimeMs) {
		toSerialize["averageRegistrationRequestResponseTimeMs"] = o.AverageRegistrationRequestResponseTimeMs
	}
	if !IsNil(o.TotalRegistrationRequestsProcessed) {
		toSerialize["totalRegistrationRequestsProcessed"] = o.TotalRegistrationRequestsProcessed
	}
	if !IsNil(o.TotalRegistrationRequests) {
		toSerialize["totalRegistrationRequests"] = o.TotalRegistrationRequests
	}
	if !IsNil(o.AverageAbuseResponseTimeMs) {
		toSerialize["averageAbuseResponseTimeMs"] = o.AverageAbuseResponseTimeMs
	}
	if !IsNil(o.TotalAbusesProcessed) {
		toSerialize["totalAbusesProcessed"] = o.TotalAbusesProcessed
	}
	if !IsNil(o.TotalAbuses) {
		toSerialize["totalAbuses"] = o.TotalAbuses
	}
	return toSerialize, nil
}

type NullableServerStats struct {
	value *ServerStats
	isSet bool
}

func (v NullableServerStats) Get() *ServerStats {
	return v.value
}

func (v *NullableServerStats) Set(val *ServerStats) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStats) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStats(val *ServerStats) *NullableServerStats {
	return &NullableServerStats{value: val, isSet: true}
}

func (v NullableServerStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
