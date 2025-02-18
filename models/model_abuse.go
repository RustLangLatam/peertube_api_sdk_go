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

// checks if the Abuse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Abuse{}

// Abuse struct for Abuse
type Abuse struct {
	Id                *int32              `json:"id,omitempty"`
	Reason            *string             `json:"reason,omitempty"`
	PredefinedReasons []string            `json:"predefinedReasons,omitempty"`
	ReporterAccount   *Account            `json:"reporterAccount,omitempty"`
	State             *AbuseStateConstant `json:"state,omitempty"`
	ModerationComment *string             `json:"moderationComment,omitempty"`
	Video             *VideoInfo          `json:"video,omitempty"`
	CreatedAt         *time.Time          `json:"createdAt,omitempty"`
}

// NewAbuse instantiates a new Abuse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbuse() *Abuse {
	this := Abuse{}
	return &this
}

// NewAbuseWithDefaults instantiates a new Abuse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbuseWithDefaults() *Abuse {
	this := Abuse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Abuse) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Abuse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Abuse) SetId(v int32) {
	o.Id = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Abuse) GetReason() string {
	if o == nil || utils.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Abuse) HasReason() bool {
	if o != nil && !utils.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Abuse) SetReason(v string) {
	o.Reason = &v
}

// GetPredefinedReasons returns the PredefinedReasons field value if set, zero value otherwise.
func (o *Abuse) GetPredefinedReasons() []string {
	if o == nil || utils.IsNil(o.PredefinedReasons) {
		var ret []string
		return ret
	}
	return o.PredefinedReasons
}

// GetPredefinedReasonsOk returns a tuple with the PredefinedReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetPredefinedReasonsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PredefinedReasons) {
		return nil, false
	}
	return o.PredefinedReasons, true
}

// HasPredefinedReasons returns a boolean if a field has been set.
func (o *Abuse) HasPredefinedReasons() bool {
	if o != nil && !utils.IsNil(o.PredefinedReasons) {
		return true
	}

	return false
}

// SetPredefinedReasons gets a reference to the given []string and assigns it to the PredefinedReasons field.
func (o *Abuse) SetPredefinedReasons(v []string) {
	o.PredefinedReasons = v
}

// GetReporterAccount returns the ReporterAccount field value if set, zero value otherwise.
func (o *Abuse) GetReporterAccount() Account {
	if o == nil || utils.IsNil(o.ReporterAccount) {
		var ret Account
		return ret
	}
	return *o.ReporterAccount
}

// GetReporterAccountOk returns a tuple with the ReporterAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetReporterAccountOk() (*Account, bool) {
	if o == nil || utils.IsNil(o.ReporterAccount) {
		return nil, false
	}
	return o.ReporterAccount, true
}

// HasReporterAccount returns a boolean if a field has been set.
func (o *Abuse) HasReporterAccount() bool {
	if o != nil && !utils.IsNil(o.ReporterAccount) {
		return true
	}

	return false
}

// SetReporterAccount gets a reference to the given Account and assigns it to the ReporterAccount field.
func (o *Abuse) SetReporterAccount(v Account) {
	o.ReporterAccount = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Abuse) GetState() AbuseStateConstant {
	if o == nil || utils.IsNil(o.State) {
		var ret AbuseStateConstant
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetStateOk() (*AbuseStateConstant, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Abuse) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given AbuseStateConstant and assigns it to the State field.
func (o *Abuse) SetState(v AbuseStateConstant) {
	o.State = &v
}

// GetModerationComment returns the ModerationComment field value if set, zero value otherwise.
func (o *Abuse) GetModerationComment() string {
	if o == nil || utils.IsNil(o.ModerationComment) {
		var ret string
		return ret
	}
	return *o.ModerationComment
}

// GetModerationCommentOk returns a tuple with the ModerationComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetModerationCommentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModerationComment) {
		return nil, false
	}
	return o.ModerationComment, true
}

// HasModerationComment returns a boolean if a field has been set.
func (o *Abuse) HasModerationComment() bool {
	if o != nil && !utils.IsNil(o.ModerationComment) {
		return true
	}

	return false
}

// SetModerationComment gets a reference to the given string and assigns it to the ModerationComment field.
func (o *Abuse) SetModerationComment(v string) {
	o.ModerationComment = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *Abuse) GetVideo() VideoInfo {
	if o == nil || utils.IsNil(o.Video) {
		var ret VideoInfo
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetVideoOk() (*VideoInfo, bool) {
	if o == nil || utils.IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *Abuse) HasVideo() bool {
	if o != nil && !utils.IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given VideoInfo and assigns it to the Video field.
func (o *Abuse) SetVideo(v VideoInfo) {
	o.Video = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Abuse) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Abuse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Abuse) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Abuse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o Abuse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Abuse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !utils.IsNil(o.PredefinedReasons) {
		toSerialize["predefinedReasons"] = o.PredefinedReasons
	}
	if !utils.IsNil(o.ReporterAccount) {
		toSerialize["reporterAccount"] = o.ReporterAccount
	}
	if !utils.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !utils.IsNil(o.ModerationComment) {
		toSerialize["moderationComment"] = o.ModerationComment
	}
	if !utils.IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAbuse struct {
	value *Abuse
	isSet bool
}

func (v NullableAbuse) Get() *Abuse {
	return v.value
}

func (v *NullableAbuse) Set(val *Abuse) {
	v.value = val
	v.isSet = true
}

func (v NullableAbuse) IsSet() bool {
	return v.isSet
}

func (v *NullableAbuse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbuse(val *Abuse) *NullableAbuse {
	return &NullableAbuse{value: val, isSet: true}
}

func (v NullableAbuse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbuse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
