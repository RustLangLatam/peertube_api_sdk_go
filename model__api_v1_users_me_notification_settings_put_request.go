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

// checks if the ApiV1UsersMeNotificationSettingsPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1UsersMeNotificationSettingsPutRequest{}

// ApiV1UsersMeNotificationSettingsPutRequest struct for ApiV1UsersMeNotificationSettingsPutRequest
type ApiV1UsersMeNotificationSettingsPutRequest struct {
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	NewVideoFromSubscription *int32 `json:"newVideoFromSubscription,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	NewCommentOnMyVideo *int32 `json:"newCommentOnMyVideo,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	AbuseAsModerator *int32 `json:"abuseAsModerator,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	VideoAutoBlacklistAsModerator *int32 `json:"videoAutoBlacklistAsModerator,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	BlacklistOnMyVideo *int32 `json:"blacklistOnMyVideo,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	MyVideoPublished *int32 `json:"myVideoPublished,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	MyVideoImportFinished *int32 `json:"myVideoImportFinished,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	NewFollow *int32 `json:"newFollow,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	NewUserRegistration *int32 `json:"newUserRegistration,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	CommentMention *int32 `json:"commentMention,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	NewInstanceFollower *int32 `json:"newInstanceFollower,omitempty"`
	// Notification type. One of the following values, or a sum of multiple values: - `0` NONE - `1` WEB - `2` EMAIL
	AutoInstanceFollowing *int32 `json:"autoInstanceFollowing,omitempty"`
}

// NewApiV1UsersMeNotificationSettingsPutRequest instantiates a new ApiV1UsersMeNotificationSettingsPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1UsersMeNotificationSettingsPutRequest() *ApiV1UsersMeNotificationSettingsPutRequest {
	this := ApiV1UsersMeNotificationSettingsPutRequest{}
	return &this
}

// NewApiV1UsersMeNotificationSettingsPutRequestWithDefaults instantiates a new ApiV1UsersMeNotificationSettingsPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1UsersMeNotificationSettingsPutRequestWithDefaults() *ApiV1UsersMeNotificationSettingsPutRequest {
	this := ApiV1UsersMeNotificationSettingsPutRequest{}
	return &this
}

// GetNewVideoFromSubscription returns the NewVideoFromSubscription field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewVideoFromSubscription() int32 {
	if o == nil || IsNil(o.NewVideoFromSubscription) {
		var ret int32
		return ret
	}
	return *o.NewVideoFromSubscription
}

// GetNewVideoFromSubscriptionOk returns a tuple with the NewVideoFromSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewVideoFromSubscriptionOk() (*int32, bool) {
	if o == nil || IsNil(o.NewVideoFromSubscription) {
		return nil, false
	}
	return o.NewVideoFromSubscription, true
}

// HasNewVideoFromSubscription returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewVideoFromSubscription() bool {
	if o != nil && !IsNil(o.NewVideoFromSubscription) {
		return true
	}

	return false
}

// SetNewVideoFromSubscription gets a reference to the given int32 and assigns it to the NewVideoFromSubscription field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewVideoFromSubscription(v int32) {
	o.NewVideoFromSubscription = &v
}

// GetNewCommentOnMyVideo returns the NewCommentOnMyVideo field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewCommentOnMyVideo() int32 {
	if o == nil || IsNil(o.NewCommentOnMyVideo) {
		var ret int32
		return ret
	}
	return *o.NewCommentOnMyVideo
}

// GetNewCommentOnMyVideoOk returns a tuple with the NewCommentOnMyVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewCommentOnMyVideoOk() (*int32, bool) {
	if o == nil || IsNil(o.NewCommentOnMyVideo) {
		return nil, false
	}
	return o.NewCommentOnMyVideo, true
}

// HasNewCommentOnMyVideo returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewCommentOnMyVideo() bool {
	if o != nil && !IsNil(o.NewCommentOnMyVideo) {
		return true
	}

	return false
}

// SetNewCommentOnMyVideo gets a reference to the given int32 and assigns it to the NewCommentOnMyVideo field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewCommentOnMyVideo(v int32) {
	o.NewCommentOnMyVideo = &v
}

// GetAbuseAsModerator returns the AbuseAsModerator field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAbuseAsModerator() int32 {
	if o == nil || IsNil(o.AbuseAsModerator) {
		var ret int32
		return ret
	}
	return *o.AbuseAsModerator
}

// GetAbuseAsModeratorOk returns a tuple with the AbuseAsModerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAbuseAsModeratorOk() (*int32, bool) {
	if o == nil || IsNil(o.AbuseAsModerator) {
		return nil, false
	}
	return o.AbuseAsModerator, true
}

// HasAbuseAsModerator returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasAbuseAsModerator() bool {
	if o != nil && !IsNil(o.AbuseAsModerator) {
		return true
	}

	return false
}

// SetAbuseAsModerator gets a reference to the given int32 and assigns it to the AbuseAsModerator field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetAbuseAsModerator(v int32) {
	o.AbuseAsModerator = &v
}

// GetVideoAutoBlacklistAsModerator returns the VideoAutoBlacklistAsModerator field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetVideoAutoBlacklistAsModerator() int32 {
	if o == nil || IsNil(o.VideoAutoBlacklistAsModerator) {
		var ret int32
		return ret
	}
	return *o.VideoAutoBlacklistAsModerator
}

// GetVideoAutoBlacklistAsModeratorOk returns a tuple with the VideoAutoBlacklistAsModerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetVideoAutoBlacklistAsModeratorOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoAutoBlacklistAsModerator) {
		return nil, false
	}
	return o.VideoAutoBlacklistAsModerator, true
}

// HasVideoAutoBlacklistAsModerator returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasVideoAutoBlacklistAsModerator() bool {
	if o != nil && !IsNil(o.VideoAutoBlacklistAsModerator) {
		return true
	}

	return false
}

// SetVideoAutoBlacklistAsModerator gets a reference to the given int32 and assigns it to the VideoAutoBlacklistAsModerator field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetVideoAutoBlacklistAsModerator(v int32) {
	o.VideoAutoBlacklistAsModerator = &v
}

// GetBlacklistOnMyVideo returns the BlacklistOnMyVideo field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetBlacklistOnMyVideo() int32 {
	if o == nil || IsNil(o.BlacklistOnMyVideo) {
		var ret int32
		return ret
	}
	return *o.BlacklistOnMyVideo
}

// GetBlacklistOnMyVideoOk returns a tuple with the BlacklistOnMyVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetBlacklistOnMyVideoOk() (*int32, bool) {
	if o == nil || IsNil(o.BlacklistOnMyVideo) {
		return nil, false
	}
	return o.BlacklistOnMyVideo, true
}

// HasBlacklistOnMyVideo returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasBlacklistOnMyVideo() bool {
	if o != nil && !IsNil(o.BlacklistOnMyVideo) {
		return true
	}

	return false
}

// SetBlacklistOnMyVideo gets a reference to the given int32 and assigns it to the BlacklistOnMyVideo field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetBlacklistOnMyVideo(v int32) {
	o.BlacklistOnMyVideo = &v
}

// GetMyVideoPublished returns the MyVideoPublished field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoPublished() int32 {
	if o == nil || IsNil(o.MyVideoPublished) {
		var ret int32
		return ret
	}
	return *o.MyVideoPublished
}

// GetMyVideoPublishedOk returns a tuple with the MyVideoPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoPublishedOk() (*int32, bool) {
	if o == nil || IsNil(o.MyVideoPublished) {
		return nil, false
	}
	return o.MyVideoPublished, true
}

// HasMyVideoPublished returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasMyVideoPublished() bool {
	if o != nil && !IsNil(o.MyVideoPublished) {
		return true
	}

	return false
}

// SetMyVideoPublished gets a reference to the given int32 and assigns it to the MyVideoPublished field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetMyVideoPublished(v int32) {
	o.MyVideoPublished = &v
}

// GetMyVideoImportFinished returns the MyVideoImportFinished field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoImportFinished() int32 {
	if o == nil || IsNil(o.MyVideoImportFinished) {
		var ret int32
		return ret
	}
	return *o.MyVideoImportFinished
}

// GetMyVideoImportFinishedOk returns a tuple with the MyVideoImportFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetMyVideoImportFinishedOk() (*int32, bool) {
	if o == nil || IsNil(o.MyVideoImportFinished) {
		return nil, false
	}
	return o.MyVideoImportFinished, true
}

// HasMyVideoImportFinished returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasMyVideoImportFinished() bool {
	if o != nil && !IsNil(o.MyVideoImportFinished) {
		return true
	}

	return false
}

// SetMyVideoImportFinished gets a reference to the given int32 and assigns it to the MyVideoImportFinished field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetMyVideoImportFinished(v int32) {
	o.MyVideoImportFinished = &v
}

// GetNewFollow returns the NewFollow field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewFollow() int32 {
	if o == nil || IsNil(o.NewFollow) {
		var ret int32
		return ret
	}
	return *o.NewFollow
}

// GetNewFollowOk returns a tuple with the NewFollow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewFollowOk() (*int32, bool) {
	if o == nil || IsNil(o.NewFollow) {
		return nil, false
	}
	return o.NewFollow, true
}

// HasNewFollow returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewFollow() bool {
	if o != nil && !IsNil(o.NewFollow) {
		return true
	}

	return false
}

// SetNewFollow gets a reference to the given int32 and assigns it to the NewFollow field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewFollow(v int32) {
	o.NewFollow = &v
}

// GetNewUserRegistration returns the NewUserRegistration field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewUserRegistration() int32 {
	if o == nil || IsNil(o.NewUserRegistration) {
		var ret int32
		return ret
	}
	return *o.NewUserRegistration
}

// GetNewUserRegistrationOk returns a tuple with the NewUserRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewUserRegistrationOk() (*int32, bool) {
	if o == nil || IsNil(o.NewUserRegistration) {
		return nil, false
	}
	return o.NewUserRegistration, true
}

// HasNewUserRegistration returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewUserRegistration() bool {
	if o != nil && !IsNil(o.NewUserRegistration) {
		return true
	}

	return false
}

// SetNewUserRegistration gets a reference to the given int32 and assigns it to the NewUserRegistration field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewUserRegistration(v int32) {
	o.NewUserRegistration = &v
}

// GetCommentMention returns the CommentMention field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetCommentMention() int32 {
	if o == nil || IsNil(o.CommentMention) {
		var ret int32
		return ret
	}
	return *o.CommentMention
}

// GetCommentMentionOk returns a tuple with the CommentMention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetCommentMentionOk() (*int32, bool) {
	if o == nil || IsNil(o.CommentMention) {
		return nil, false
	}
	return o.CommentMention, true
}

// HasCommentMention returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasCommentMention() bool {
	if o != nil && !IsNil(o.CommentMention) {
		return true
	}

	return false
}

// SetCommentMention gets a reference to the given int32 and assigns it to the CommentMention field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetCommentMention(v int32) {
	o.CommentMention = &v
}

// GetNewInstanceFollower returns the NewInstanceFollower field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewInstanceFollower() int32 {
	if o == nil || IsNil(o.NewInstanceFollower) {
		var ret int32
		return ret
	}
	return *o.NewInstanceFollower
}

// GetNewInstanceFollowerOk returns a tuple with the NewInstanceFollower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetNewInstanceFollowerOk() (*int32, bool) {
	if o == nil || IsNil(o.NewInstanceFollower) {
		return nil, false
	}
	return o.NewInstanceFollower, true
}

// HasNewInstanceFollower returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasNewInstanceFollower() bool {
	if o != nil && !IsNil(o.NewInstanceFollower) {
		return true
	}

	return false
}

// SetNewInstanceFollower gets a reference to the given int32 and assigns it to the NewInstanceFollower field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetNewInstanceFollower(v int32) {
	o.NewInstanceFollower = &v
}

// GetAutoInstanceFollowing returns the AutoInstanceFollowing field value if set, zero value otherwise.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAutoInstanceFollowing() int32 {
	if o == nil || IsNil(o.AutoInstanceFollowing) {
		var ret int32
		return ret
	}
	return *o.AutoInstanceFollowing
}

// GetAutoInstanceFollowingOk returns a tuple with the AutoInstanceFollowing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) GetAutoInstanceFollowingOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoInstanceFollowing) {
		return nil, false
	}
	return o.AutoInstanceFollowing, true
}

// HasAutoInstanceFollowing returns a boolean if a field has been set.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) HasAutoInstanceFollowing() bool {
	if o != nil && !IsNil(o.AutoInstanceFollowing) {
		return true
	}

	return false
}

// SetAutoInstanceFollowing gets a reference to the given int32 and assigns it to the AutoInstanceFollowing field.
func (o *ApiV1UsersMeNotificationSettingsPutRequest) SetAutoInstanceFollowing(v int32) {
	o.AutoInstanceFollowing = &v
}

func (o ApiV1UsersMeNotificationSettingsPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1UsersMeNotificationSettingsPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewVideoFromSubscription) {
		toSerialize["newVideoFromSubscription"] = o.NewVideoFromSubscription
	}
	if !IsNil(o.NewCommentOnMyVideo) {
		toSerialize["newCommentOnMyVideo"] = o.NewCommentOnMyVideo
	}
	if !IsNil(o.AbuseAsModerator) {
		toSerialize["abuseAsModerator"] = o.AbuseAsModerator
	}
	if !IsNil(o.VideoAutoBlacklistAsModerator) {
		toSerialize["videoAutoBlacklistAsModerator"] = o.VideoAutoBlacklistAsModerator
	}
	if !IsNil(o.BlacklistOnMyVideo) {
		toSerialize["blacklistOnMyVideo"] = o.BlacklistOnMyVideo
	}
	if !IsNil(o.MyVideoPublished) {
		toSerialize["myVideoPublished"] = o.MyVideoPublished
	}
	if !IsNil(o.MyVideoImportFinished) {
		toSerialize["myVideoImportFinished"] = o.MyVideoImportFinished
	}
	if !IsNil(o.NewFollow) {
		toSerialize["newFollow"] = o.NewFollow
	}
	if !IsNil(o.NewUserRegistration) {
		toSerialize["newUserRegistration"] = o.NewUserRegistration
	}
	if !IsNil(o.CommentMention) {
		toSerialize["commentMention"] = o.CommentMention
	}
	if !IsNil(o.NewInstanceFollower) {
		toSerialize["newInstanceFollower"] = o.NewInstanceFollower
	}
	if !IsNil(o.AutoInstanceFollowing) {
		toSerialize["autoInstanceFollowing"] = o.AutoInstanceFollowing
	}
	return toSerialize, nil
}

type NullableApiV1UsersMeNotificationSettingsPutRequest struct {
	value *ApiV1UsersMeNotificationSettingsPutRequest
	isSet bool
}

func (v NullableApiV1UsersMeNotificationSettingsPutRequest) Get() *ApiV1UsersMeNotificationSettingsPutRequest {
	return v.value
}

func (v *NullableApiV1UsersMeNotificationSettingsPutRequest) Set(val *ApiV1UsersMeNotificationSettingsPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1UsersMeNotificationSettingsPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1UsersMeNotificationSettingsPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1UsersMeNotificationSettingsPutRequest(val *ApiV1UsersMeNotificationSettingsPutRequest) *NullableApiV1UsersMeNotificationSettingsPutRequest {
	return &NullableApiV1UsersMeNotificationSettingsPutRequest{value: val, isSet: true}
}

func (v NullableApiV1UsersMeNotificationSettingsPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1UsersMeNotificationSettingsPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
