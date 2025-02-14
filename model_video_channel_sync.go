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

// checks if the VideoChannelSync type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoChannelSync{}

// VideoChannelSync struct for VideoChannelSync
type VideoChannelSync struct {
	Id                 *int32                 `json:"id,omitempty"`
	State              *VideoChannelSyncState `json:"state,omitempty"`
	ExternalChannelUrl *string                `json:"externalChannelUrl,omitempty"`
	CreatedAt          *time.Time             `json:"createdAt,omitempty"`
	LastSyncAt         NullableTime           `json:"lastSyncAt,omitempty"`
	Channel            *VideoChannel          `json:"channel,omitempty"`
}

// NewVideoChannelSync instantiates a new VideoChannelSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoChannelSync() *VideoChannelSync {
	this := VideoChannelSync{}
	return &this
}

// NewVideoChannelSyncWithDefaults instantiates a new VideoChannelSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoChannelSyncWithDefaults() *VideoChannelSync {
	this := VideoChannelSync{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoChannelSync) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelSync) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoChannelSync) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoChannelSync) SetId(v int32) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VideoChannelSync) GetState() VideoChannelSyncState {
	if o == nil || IsNil(o.State) {
		var ret VideoChannelSyncState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelSync) GetStateOk() (*VideoChannelSyncState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VideoChannelSync) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given VideoChannelSyncState and assigns it to the State field.
func (o *VideoChannelSync) SetState(v VideoChannelSyncState) {
	o.State = &v
}

// GetExternalChannelUrl returns the ExternalChannelUrl field value if set, zero value otherwise.
func (o *VideoChannelSync) GetExternalChannelUrl() string {
	if o == nil || IsNil(o.ExternalChannelUrl) {
		var ret string
		return ret
	}
	return *o.ExternalChannelUrl
}

// GetExternalChannelUrlOk returns a tuple with the ExternalChannelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelSync) GetExternalChannelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalChannelUrl) {
		return nil, false
	}
	return o.ExternalChannelUrl, true
}

// HasExternalChannelUrl returns a boolean if a field has been set.
func (o *VideoChannelSync) HasExternalChannelUrl() bool {
	if o != nil && !IsNil(o.ExternalChannelUrl) {
		return true
	}

	return false
}

// SetExternalChannelUrl gets a reference to the given string and assigns it to the ExternalChannelUrl field.
func (o *VideoChannelSync) SetExternalChannelUrl(v string) {
	o.ExternalChannelUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoChannelSync) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelSync) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoChannelSync) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoChannelSync) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastSyncAt returns the LastSyncAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoChannelSync) GetLastSyncAt() time.Time {
	if o == nil || IsNil(o.LastSyncAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSyncAt.Get()
}

// GetLastSyncAtOk returns a tuple with the LastSyncAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoChannelSync) GetLastSyncAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSyncAt.Get(), o.LastSyncAt.IsSet()
}

// HasLastSyncAt returns a boolean if a field has been set.
func (o *VideoChannelSync) HasLastSyncAt() bool {
	if o != nil && o.LastSyncAt.IsSet() {
		return true
	}

	return false
}

// SetLastSyncAt gets a reference to the given NullableTime and assigns it to the LastSyncAt field.
func (o *VideoChannelSync) SetLastSyncAt(v time.Time) {
	o.LastSyncAt.Set(&v)
}

// SetLastSyncAtNil sets the value for LastSyncAt to be an explicit nil
func (o *VideoChannelSync) SetLastSyncAtNil() {
	o.LastSyncAt.Set(nil)
}

// UnsetLastSyncAt ensures that no value is present for LastSyncAt, not even an explicit nil
func (o *VideoChannelSync) UnsetLastSyncAt() {
	o.LastSyncAt.Unset()
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *VideoChannelSync) GetChannel() VideoChannel {
	if o == nil || IsNil(o.Channel) {
		var ret VideoChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelSync) GetChannelOk() (*VideoChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *VideoChannelSync) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given VideoChannel and assigns it to the Channel field.
func (o *VideoChannelSync) SetChannel(v VideoChannel) {
	o.Channel = &v
}

func (o VideoChannelSync) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoChannelSync) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ExternalChannelUrl) {
		toSerialize["externalChannelUrl"] = o.ExternalChannelUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.LastSyncAt.IsSet() {
		toSerialize["lastSyncAt"] = o.LastSyncAt.Get()
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableVideoChannelSync struct {
	value *VideoChannelSync
	isSet bool
}

func (v NullableVideoChannelSync) Get() *VideoChannelSync {
	return v.value
}

func (v *NullableVideoChannelSync) Set(val *VideoChannelSync) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoChannelSync) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoChannelSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoChannelSync(val *VideoChannelSync) *NullableVideoChannelSync {
	return &NullableVideoChannelSync{value: val, isSet: true}
}

func (v NullableVideoChannelSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoChannelSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
