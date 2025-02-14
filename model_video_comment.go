/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"encoding/json"
	"time"
)

// checks if the VideoComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoComment{}

// VideoComment struct for VideoComment
type VideoComment struct {
	Id  *int32  `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// Text of the comment
	Text                        *string       `json:"text,omitempty"`
	ThreadId                    *int32        `json:"threadId,omitempty"`
	InReplyToCommentId          NullableInt32 `json:"inReplyToCommentId,omitempty"`
	VideoId                     *int32        `json:"videoId,omitempty"`
	CreatedAt                   *time.Time    `json:"createdAt,omitempty"`
	UpdatedAt                   *time.Time    `json:"updatedAt,omitempty"`
	DeletedAt                   NullableTime  `json:"deletedAt,omitempty"`
	IsDeleted                   *bool         `json:"isDeleted,omitempty"`
	HeldForReview               *bool         `json:"heldForReview,omitempty"`
	TotalRepliesFromVideoAuthor *int32        `json:"totalRepliesFromVideoAuthor,omitempty"`
	TotalReplies                *int32        `json:"totalReplies,omitempty"`
	Account                     *Account      `json:"account,omitempty"`
}

// NewVideoComment instantiates a new VideoComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoComment() *VideoComment {
	this := VideoComment{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// NewVideoCommentWithDefaults instantiates a new VideoComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoCommentWithDefaults() *VideoComment {
	this := VideoComment{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoComment) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoComment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoComment) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VideoComment) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VideoComment) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VideoComment) SetUrl(v string) {
	o.Url = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *VideoComment) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *VideoComment) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *VideoComment) SetText(v string) {
	o.Text = &v
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *VideoComment) GetThreadId() int32 {
	if o == nil || IsNil(o.ThreadId) {
		var ret int32
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *VideoComment) HasThreadId() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given int32 and assigns it to the ThreadId field.
func (o *VideoComment) SetThreadId(v int32) {
	o.ThreadId = &v
}

// GetInReplyToCommentId returns the InReplyToCommentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoComment) GetInReplyToCommentId() int32 {
	if o == nil || IsNil(o.InReplyToCommentId.Get()) {
		var ret int32
		return ret
	}
	return *o.InReplyToCommentId.Get()
}

// GetInReplyToCommentIdOk returns a tuple with the InReplyToCommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoComment) GetInReplyToCommentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InReplyToCommentId.Get(), o.InReplyToCommentId.IsSet()
}

// HasInReplyToCommentId returns a boolean if a field has been set.
func (o *VideoComment) HasInReplyToCommentId() bool {
	if o != nil && o.InReplyToCommentId.IsSet() {
		return true
	}

	return false
}

// SetInReplyToCommentId gets a reference to the given NullableInt32 and assigns it to the InReplyToCommentId field.
func (o *VideoComment) SetInReplyToCommentId(v int32) {
	o.InReplyToCommentId.Set(&v)
}

// SetInReplyToCommentIdNil sets the value for InReplyToCommentId to be an explicit nil
func (o *VideoComment) SetInReplyToCommentIdNil() {
	o.InReplyToCommentId.Set(nil)
}

// UnsetInReplyToCommentId ensures that no value is present for InReplyToCommentId, not even an explicit nil
func (o *VideoComment) UnsetInReplyToCommentId() {
	o.InReplyToCommentId.Unset()
}

// GetVideoId returns the VideoId field value if set, zero value otherwise.
func (o *VideoComment) GetVideoId() int32 {
	if o == nil || IsNil(o.VideoId) {
		var ret int32
		return ret
	}
	return *o.VideoId
}

// GetVideoIdOk returns a tuple with the VideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetVideoIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoId) {
		return nil, false
	}
	return o.VideoId, true
}

// HasVideoId returns a boolean if a field has been set.
func (o *VideoComment) HasVideoId() bool {
	if o != nil && !IsNil(o.VideoId) {
		return true
	}

	return false
}

// SetVideoId gets a reference to the given int32 and assigns it to the VideoId field.
func (o *VideoComment) SetVideoId(v int32) {
	o.VideoId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoComment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoComment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoComment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VideoComment) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VideoComment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VideoComment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoComment) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoComment) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *VideoComment) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *VideoComment) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *VideoComment) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *VideoComment) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *VideoComment) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *VideoComment) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *VideoComment) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetHeldForReview returns the HeldForReview field value if set, zero value otherwise.
func (o *VideoComment) GetHeldForReview() bool {
	if o == nil || IsNil(o.HeldForReview) {
		var ret bool
		return ret
	}
	return *o.HeldForReview
}

// GetHeldForReviewOk returns a tuple with the HeldForReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetHeldForReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.HeldForReview) {
		return nil, false
	}
	return o.HeldForReview, true
}

// HasHeldForReview returns a boolean if a field has been set.
func (o *VideoComment) HasHeldForReview() bool {
	if o != nil && !IsNil(o.HeldForReview) {
		return true
	}

	return false
}

// SetHeldForReview gets a reference to the given bool and assigns it to the HeldForReview field.
func (o *VideoComment) SetHeldForReview(v bool) {
	o.HeldForReview = &v
}

// GetTotalRepliesFromVideoAuthor returns the TotalRepliesFromVideoAuthor field value if set, zero value otherwise.
func (o *VideoComment) GetTotalRepliesFromVideoAuthor() int32 {
	if o == nil || IsNil(o.TotalRepliesFromVideoAuthor) {
		var ret int32
		return ret
	}
	return *o.TotalRepliesFromVideoAuthor
}

// GetTotalRepliesFromVideoAuthorOk returns a tuple with the TotalRepliesFromVideoAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetTotalRepliesFromVideoAuthorOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRepliesFromVideoAuthor) {
		return nil, false
	}
	return o.TotalRepliesFromVideoAuthor, true
}

// HasTotalRepliesFromVideoAuthor returns a boolean if a field has been set.
func (o *VideoComment) HasTotalRepliesFromVideoAuthor() bool {
	if o != nil && !IsNil(o.TotalRepliesFromVideoAuthor) {
		return true
	}

	return false
}

// SetTotalRepliesFromVideoAuthor gets a reference to the given int32 and assigns it to the TotalRepliesFromVideoAuthor field.
func (o *VideoComment) SetTotalRepliesFromVideoAuthor(v int32) {
	o.TotalRepliesFromVideoAuthor = &v
}

// GetTotalReplies returns the TotalReplies field value if set, zero value otherwise.
func (o *VideoComment) GetTotalReplies() int32 {
	if o == nil || IsNil(o.TotalReplies) {
		var ret int32
		return ret
	}
	return *o.TotalReplies
}

// GetTotalRepliesOk returns a tuple with the TotalReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetTotalRepliesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalReplies) {
		return nil, false
	}
	return o.TotalReplies, true
}

// HasTotalReplies returns a boolean if a field has been set.
func (o *VideoComment) HasTotalReplies() bool {
	if o != nil && !IsNil(o.TotalReplies) {
		return true
	}

	return false
}

// SetTotalReplies gets a reference to the given int32 and assigns it to the TotalReplies field.
func (o *VideoComment) SetTotalReplies(v int32) {
	o.TotalReplies = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *VideoComment) GetAccount() Account {
	if o == nil || IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComment) GetAccountOk() (*Account, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *VideoComment) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *VideoComment) SetAccount(v Account) {
	o.Account = &v
}

func (o VideoComment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	if o.InReplyToCommentId.IsSet() {
		toSerialize["inReplyToCommentId"] = o.InReplyToCommentId.Get()
	}
	if !IsNil(o.VideoId) {
		toSerialize["videoId"] = o.VideoId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deletedAt"] = o.DeletedAt.Get()
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.HeldForReview) {
		toSerialize["heldForReview"] = o.HeldForReview
	}
	if !IsNil(o.TotalRepliesFromVideoAuthor) {
		toSerialize["totalRepliesFromVideoAuthor"] = o.TotalRepliesFromVideoAuthor
	}
	if !IsNil(o.TotalReplies) {
		toSerialize["totalReplies"] = o.TotalReplies
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableVideoComment struct {
	value *VideoComment
	isSet bool
}

func (v NullableVideoComment) Get() *VideoComment {
	return v.value
}

func (v *NullableVideoComment) Set(val *VideoComment) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoComment) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoComment(val *VideoComment) *NullableVideoComment {
	return &NullableVideoComment{value: val, isSet: true}
}

func (v NullableVideoComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
