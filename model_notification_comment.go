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

// checks if the NotificationComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationComment{}

// NotificationComment struct for NotificationComment
type NotificationComment struct {
	Id            *int32     `json:"id,omitempty"`
	ThreadId      *int32     `json:"threadId,omitempty"`
	Video         *VideoInfo `json:"video,omitempty"`
	Account       *ActorInfo `json:"account,omitempty"`
	HeldForReview *bool      `json:"heldForReview,omitempty"`
}

// NewNotificationComment instantiates a new NotificationComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationComment() *NotificationComment {
	this := NotificationComment{}
	return &this
}

// NewNotificationCommentWithDefaults instantiates a new NotificationComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationCommentWithDefaults() *NotificationComment {
	this := NotificationComment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationComment) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationComment) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationComment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NotificationComment) SetId(v int32) {
	o.Id = &v
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *NotificationComment) GetThreadId() int32 {
	if o == nil || IsNil(o.ThreadId) {
		var ret int32
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationComment) GetThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *NotificationComment) HasThreadId() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given int32 and assigns it to the ThreadId field.
func (o *NotificationComment) SetThreadId(v int32) {
	o.ThreadId = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *NotificationComment) GetVideo() VideoInfo {
	if o == nil || IsNil(o.Video) {
		var ret VideoInfo
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationComment) GetVideoOk() (*VideoInfo, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *NotificationComment) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given VideoInfo and assigns it to the Video field.
func (o *NotificationComment) SetVideo(v VideoInfo) {
	o.Video = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *NotificationComment) GetAccount() ActorInfo {
	if o == nil || IsNil(o.Account) {
		var ret ActorInfo
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationComment) GetAccountOk() (*ActorInfo, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *NotificationComment) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ActorInfo and assigns it to the Account field.
func (o *NotificationComment) SetAccount(v ActorInfo) {
	o.Account = &v
}

// GetHeldForReview returns the HeldForReview field value if set, zero value otherwise.
func (o *NotificationComment) GetHeldForReview() bool {
	if o == nil || IsNil(o.HeldForReview) {
		var ret bool
		return ret
	}
	return *o.HeldForReview
}

// GetHeldForReviewOk returns a tuple with the HeldForReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationComment) GetHeldForReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.HeldForReview) {
		return nil, false
	}
	return o.HeldForReview, true
}

// HasHeldForReview returns a boolean if a field has been set.
func (o *NotificationComment) HasHeldForReview() bool {
	if o != nil && !IsNil(o.HeldForReview) {
		return true
	}

	return false
}

// SetHeldForReview gets a reference to the given bool and assigns it to the HeldForReview field.
func (o *NotificationComment) SetHeldForReview(v bool) {
	o.HeldForReview = &v
}

func (o NotificationComment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.HeldForReview) {
		toSerialize["heldForReview"] = o.HeldForReview
	}
	return toSerialize, nil
}

type NullableNotificationComment struct {
	value *NotificationComment
	isSet bool
}

func (v NullableNotificationComment) Get() *NotificationComment {
	return v.value
}

func (v *NullableNotificationComment) Set(val *NotificationComment) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationComment) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationComment(val *NotificationComment) *NullableNotificationComment {
	return &NullableNotificationComment{value: val, isSet: true}
}

func (v NullableNotificationComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
