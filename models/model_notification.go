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

// checks if the Notification type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Notification{}

// Notification struct for Notification
type Notification struct {
	Id *int32 `json:"id,omitempty"`
	// Notification type, following the `UserNotificationType` enum: - `1` NEW_VIDEO_FROM_SUBSCRIPTION - `2` NEW_COMMENT_ON_MY_VIDEO - `3` NEW_ABUSE_FOR_MODERATORS - `4` BLACKLIST_ON_MY_VIDEO - `5` UNBLACKLIST_ON_MY_VIDEO - `6` MY_VIDEO_PUBLISHED - `7` MY_VIDEO_IMPORT_SUCCESS - `8` MY_VIDEO_IMPORT_ERROR - `9` NEW_USER_REGISTRATION - `10` NEW_FOLLOW - `11` COMMENT_MENTION - `12` VIDEO_AUTO_BLACKLIST_FOR_MODERATORS - `13` NEW_INSTANCE_FOLLOWER - `14` AUTO_INSTANCE_FOLLOWING - `15` ABUSE_STATE_CHANGE - `16` ABUSE_NEW_MESSAGE - `17` NEW_PLUGIN_VERSION - `18` NEW_PEERTUBE_VERSION - `19` MY_VIDEO_STUDIO_EDITION_FINISHED - `20` NEW_USER_REGISTRATION_REQUEST - `21` NEW_LIVE_FROM_SUBSCRIPTION
	Type           *int32                          `json:"type,omitempty"`
	Read           *bool                           `json:"read,omitempty"`
	Video          NullableNotificationVideo       `json:"video,omitempty"`
	VideoImport    NullableNotificationVideoImport `json:"videoImport,omitempty"`
	Comment        NullableNotificationComment     `json:"comment,omitempty"`
	VideoAbuse     NullableNotificationVideoAbuse  `json:"videoAbuse,omitempty"`
	VideoBlacklist NullableNotificationVideoAbuse  `json:"videoBlacklist,omitempty"`
	Account        NullableActorInfo               `json:"account,omitempty"`
	ActorFollow    NullableNotificationActorFollow `json:"actorFollow,omitempty"`
	CreatedAt      *time.Time                      `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time                      `json:"updatedAt,omitempty"`
}

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification() *Notification {
	this := Notification{}
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Notification) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Notification) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Notification) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Notification) GetType() int32 {
	if o == nil || utils.IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetTypeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Notification) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Notification) SetType(v int32) {
	o.Type = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *Notification) GetRead() bool {
	if o == nil || utils.IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetReadOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *Notification) HasRead() bool {
	if o != nil && !utils.IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *Notification) SetRead(v bool) {
	o.Read = &v
}

// GetVideo returns the Video field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetVideo() NotificationVideo {
	if o == nil || utils.IsNil(o.Video.Get()) {
		var ret NotificationVideo
		return ret
	}
	return *o.Video.Get()
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetVideoOk() (*NotificationVideo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Video.Get(), o.Video.IsSet()
}

// HasVideo returns a boolean if a field has been set.
func (o *Notification) HasVideo() bool {
	if o != nil && o.Video.IsSet() {
		return true
	}

	return false
}

// SetVideo gets a reference to the given NullableNotificationVideo and assigns it to the Video field.
func (o *Notification) SetVideo(v NotificationVideo) {
	o.Video.Set(&v)
}

// SetVideoNil sets the value for Video to be an explicit nil
func (o *Notification) SetVideoNil() {
	o.Video.Set(nil)
}

// UnsetVideo ensures that no value is present for Video, not even an explicit nil
func (o *Notification) UnsetVideo() {
	o.Video.Unset()
}

// GetVideoImport returns the VideoImport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetVideoImport() NotificationVideoImport {
	if o == nil || utils.IsNil(o.VideoImport.Get()) {
		var ret NotificationVideoImport
		return ret
	}
	return *o.VideoImport.Get()
}

// GetVideoImportOk returns a tuple with the VideoImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetVideoImportOk() (*NotificationVideoImport, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoImport.Get(), o.VideoImport.IsSet()
}

// HasVideoImport returns a boolean if a field has been set.
func (o *Notification) HasVideoImport() bool {
	if o != nil && o.VideoImport.IsSet() {
		return true
	}

	return false
}

// SetVideoImport gets a reference to the given NullableNotificationVideoImport and assigns it to the VideoImport field.
func (o *Notification) SetVideoImport(v NotificationVideoImport) {
	o.VideoImport.Set(&v)
}

// SetVideoImportNil sets the value for VideoImport to be an explicit nil
func (o *Notification) SetVideoImportNil() {
	o.VideoImport.Set(nil)
}

// UnsetVideoImport ensures that no value is present for VideoImport, not even an explicit nil
func (o *Notification) UnsetVideoImport() {
	o.VideoImport.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetComment() NotificationComment {
	if o == nil || utils.IsNil(o.Comment.Get()) {
		var ret NotificationComment
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetCommentOk() (*NotificationComment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Notification) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableNotificationComment and assigns it to the Comment field.
func (o *Notification) SetComment(v NotificationComment) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Notification) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Notification) UnsetComment() {
	o.Comment.Unset()
}

// GetVideoAbuse returns the VideoAbuse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetVideoAbuse() NotificationVideoAbuse {
	if o == nil || utils.IsNil(o.VideoAbuse.Get()) {
		var ret NotificationVideoAbuse
		return ret
	}
	return *o.VideoAbuse.Get()
}

// GetVideoAbuseOk returns a tuple with the VideoAbuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetVideoAbuseOk() (*NotificationVideoAbuse, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoAbuse.Get(), o.VideoAbuse.IsSet()
}

// HasVideoAbuse returns a boolean if a field has been set.
func (o *Notification) HasVideoAbuse() bool {
	if o != nil && o.VideoAbuse.IsSet() {
		return true
	}

	return false
}

// SetVideoAbuse gets a reference to the given NullableNotificationVideoAbuse and assigns it to the VideoAbuse field.
func (o *Notification) SetVideoAbuse(v NotificationVideoAbuse) {
	o.VideoAbuse.Set(&v)
}

// SetVideoAbuseNil sets the value for VideoAbuse to be an explicit nil
func (o *Notification) SetVideoAbuseNil() {
	o.VideoAbuse.Set(nil)
}

// UnsetVideoAbuse ensures that no value is present for VideoAbuse, not even an explicit nil
func (o *Notification) UnsetVideoAbuse() {
	o.VideoAbuse.Unset()
}

// GetVideoBlacklist returns the VideoBlacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetVideoBlacklist() NotificationVideoAbuse {
	if o == nil || utils.IsNil(o.VideoBlacklist.Get()) {
		var ret NotificationVideoAbuse
		return ret
	}
	return *o.VideoBlacklist.Get()
}

// GetVideoBlacklistOk returns a tuple with the VideoBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetVideoBlacklistOk() (*NotificationVideoAbuse, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoBlacklist.Get(), o.VideoBlacklist.IsSet()
}

// HasVideoBlacklist returns a boolean if a field has been set.
func (o *Notification) HasVideoBlacklist() bool {
	if o != nil && o.VideoBlacklist.IsSet() {
		return true
	}

	return false
}

// SetVideoBlacklist gets a reference to the given NullableNotificationVideoAbuse and assigns it to the VideoBlacklist field.
func (o *Notification) SetVideoBlacklist(v NotificationVideoAbuse) {
	o.VideoBlacklist.Set(&v)
}

// SetVideoBlacklistNil sets the value for VideoBlacklist to be an explicit nil
func (o *Notification) SetVideoBlacklistNil() {
	o.VideoBlacklist.Set(nil)
}

// UnsetVideoBlacklist ensures that no value is present for VideoBlacklist, not even an explicit nil
func (o *Notification) UnsetVideoBlacklist() {
	o.VideoBlacklist.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetAccount() ActorInfo {
	if o == nil || utils.IsNil(o.Account.Get()) {
		var ret ActorInfo
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetAccountOk() (*ActorInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *Notification) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableActorInfo and assigns it to the Account field.
func (o *Notification) SetAccount(v ActorInfo) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *Notification) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *Notification) UnsetAccount() {
	o.Account.Unset()
}

// GetActorFollow returns the ActorFollow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetActorFollow() NotificationActorFollow {
	if o == nil || utils.IsNil(o.ActorFollow.Get()) {
		var ret NotificationActorFollow
		return ret
	}
	return *o.ActorFollow.Get()
}

// GetActorFollowOk returns a tuple with the ActorFollow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetActorFollowOk() (*NotificationActorFollow, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorFollow.Get(), o.ActorFollow.IsSet()
}

// HasActorFollow returns a boolean if a field has been set.
func (o *Notification) HasActorFollow() bool {
	if o != nil && o.ActorFollow.IsSet() {
		return true
	}

	return false
}

// SetActorFollow gets a reference to the given NullableNotificationActorFollow and assigns it to the ActorFollow field.
func (o *Notification) SetActorFollow(v NotificationActorFollow) {
	o.ActorFollow.Set(&v)
}

// SetActorFollowNil sets the value for ActorFollow to be an explicit nil
func (o *Notification) SetActorFollowNil() {
	o.ActorFollow.Set(nil)
}

// UnsetActorFollow ensures that no value is present for ActorFollow, not even an explicit nil
func (o *Notification) UnsetActorFollow() {
	o.ActorFollow.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Notification) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Notification) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Notification) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Notification) GetUpdatedAt() time.Time {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Notification) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Notification) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if o.Video.IsSet() {
		toSerialize["video"] = o.Video.Get()
	}
	if o.VideoImport.IsSet() {
		toSerialize["videoImport"] = o.VideoImport.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.VideoAbuse.IsSet() {
		toSerialize["videoAbuse"] = o.VideoAbuse.Get()
	}
	if o.VideoBlacklist.IsSet() {
		toSerialize["videoBlacklist"] = o.VideoBlacklist.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.ActorFollow.IsSet() {
		toSerialize["actorFollow"] = o.ActorFollow.Get()
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
