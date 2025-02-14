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

// checks if the VideoBlacklist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoBlacklist{}

// VideoBlacklist struct for VideoBlacklist
type VideoBlacklist struct {
	Id          *int32     `json:"id,omitempty"`
	VideoId     *int32     `json:"videoId,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Uuid        *string    `json:"uuid,omitempty" validate:"regexp=^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"`
	Description *string    `json:"description,omitempty"`
	Duration    *int32     `json:"duration,omitempty"`
	Views       *int32     `json:"views,omitempty"`
	Likes       *int32     `json:"likes,omitempty"`
	Dislikes    *int32     `json:"dislikes,omitempty"`
	Nsfw        *bool      `json:"nsfw,omitempty"`
}

// NewVideoBlacklist instantiates a new VideoBlacklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoBlacklist() *VideoBlacklist {
	this := VideoBlacklist{}
	return &this
}

// NewVideoBlacklistWithDefaults instantiates a new VideoBlacklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoBlacklistWithDefaults() *VideoBlacklist {
	this := VideoBlacklist{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoBlacklist) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoBlacklist) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoBlacklist) SetId(v int32) {
	o.Id = &v
}

// GetVideoId returns the VideoId field value if set, zero value otherwise.
func (o *VideoBlacklist) GetVideoId() int32 {
	if o == nil || IsNil(o.VideoId) {
		var ret int32
		return ret
	}
	return *o.VideoId
}

// GetVideoIdOk returns a tuple with the VideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetVideoIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoId) {
		return nil, false
	}
	return o.VideoId, true
}

// HasVideoId returns a boolean if a field has been set.
func (o *VideoBlacklist) HasVideoId() bool {
	if o != nil && !IsNil(o.VideoId) {
		return true
	}

	return false
}

// SetVideoId gets a reference to the given int32 and assigns it to the VideoId field.
func (o *VideoBlacklist) SetVideoId(v int32) {
	o.VideoId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoBlacklist) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoBlacklist) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoBlacklist) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VideoBlacklist) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VideoBlacklist) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VideoBlacklist) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VideoBlacklist) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VideoBlacklist) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VideoBlacklist) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VideoBlacklist) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VideoBlacklist) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VideoBlacklist) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VideoBlacklist) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VideoBlacklist) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VideoBlacklist) SetDescription(v string) {
	o.Description = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VideoBlacklist) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VideoBlacklist) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *VideoBlacklist) SetDuration(v int32) {
	o.Duration = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *VideoBlacklist) GetViews() int32 {
	if o == nil || IsNil(o.Views) {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetViewsOk() (*int32, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *VideoBlacklist) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *VideoBlacklist) SetViews(v int32) {
	o.Views = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise.
func (o *VideoBlacklist) GetLikes() int32 {
	if o == nil || IsNil(o.Likes) {
		var ret int32
		return ret
	}
	return *o.Likes
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Likes) {
		return nil, false
	}
	return o.Likes, true
}

// HasLikes returns a boolean if a field has been set.
func (o *VideoBlacklist) HasLikes() bool {
	if o != nil && !IsNil(o.Likes) {
		return true
	}

	return false
}

// SetLikes gets a reference to the given int32 and assigns it to the Likes field.
func (o *VideoBlacklist) SetLikes(v int32) {
	o.Likes = &v
}

// GetDislikes returns the Dislikes field value if set, zero value otherwise.
func (o *VideoBlacklist) GetDislikes() int32 {
	if o == nil || IsNil(o.Dislikes) {
		var ret int32
		return ret
	}
	return *o.Dislikes
}

// GetDislikesOk returns a tuple with the Dislikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetDislikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Dislikes) {
		return nil, false
	}
	return o.Dislikes, true
}

// HasDislikes returns a boolean if a field has been set.
func (o *VideoBlacklist) HasDislikes() bool {
	if o != nil && !IsNil(o.Dislikes) {
		return true
	}

	return false
}

// SetDislikes gets a reference to the given int32 and assigns it to the Dislikes field.
func (o *VideoBlacklist) SetDislikes(v int32) {
	o.Dislikes = &v
}

// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *VideoBlacklist) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlacklist) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *VideoBlacklist) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *VideoBlacklist) SetNsfw(v bool) {
	o.Nsfw = &v
}

func (o VideoBlacklist) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoBlacklist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	if !IsNil(o.Likes) {
		toSerialize["likes"] = o.Likes
	}
	if !IsNil(o.Dislikes) {
		toSerialize["dislikes"] = o.Dislikes
	}
	if !IsNil(o.Nsfw) {
		toSerialize["nsfw"] = o.Nsfw
	}
	return toSerialize, nil
}

type NullableVideoBlacklist struct {
	value *VideoBlacklist
	isSet bool
}

func (v NullableVideoBlacklist) Get() *VideoBlacklist {
	return v.value
}

func (v *NullableVideoBlacklist) Set(val *VideoBlacklist) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoBlacklist) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoBlacklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoBlacklist(val *VideoBlacklist) *NullableVideoBlacklist {
	return &NullableVideoBlacklist{value: val, isSet: true}
}

func (v NullableVideoBlacklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoBlacklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
