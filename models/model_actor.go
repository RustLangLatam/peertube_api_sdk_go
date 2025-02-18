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

// checks if the Actor type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Actor{}

// Actor struct for Actor
type Actor struct {
	Id  *int32  `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// immutable name of the actor, used to find or mention it
	Name    *string      `json:"name,omitempty" validate:"regexp=^[a-z0-9._]+$"`
	Avatars []ActorImage `json:"avatars,omitempty"`
	// server on which the actor is resident
	Host *string `json:"host,omitempty"`
	// whether this actor's host allows redundancy of its videos
	HostRedundancyAllowed *bool `json:"hostRedundancyAllowed,omitempty"`
	// number of actors subscribed to by this actor, as seen by this instance
	FollowingCount *int32 `json:"followingCount,omitempty"`
	// number of followers of this actor, as seen by this instance
	FollowersCount *int32     `json:"followersCount,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

// NewActor instantiates a new Actor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActor() *Actor {
	this := Actor{}
	return &this
}

// NewActorWithDefaults instantiates a new Actor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorWithDefaults() *Actor {
	this := Actor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Actor) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Actor) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Actor) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Actor) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Actor) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Actor) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Actor) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Actor) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Actor) SetName(v string) {
	o.Name = &v
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *Actor) GetAvatars() []ActorImage {
	if o == nil || utils.IsNil(o.Avatars) {
		var ret []ActorImage
		return ret
	}
	return o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetAvatarsOk() ([]ActorImage, bool) {
	if o == nil || utils.IsNil(o.Avatars) {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *Actor) HasAvatars() bool {
	if o != nil && !utils.IsNil(o.Avatars) {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given []ActorImage and assigns it to the Avatars field.
func (o *Actor) SetAvatars(v []ActorImage) {
	o.Avatars = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Actor) GetHost() string {
	if o == nil || utils.IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Actor) HasHost() bool {
	if o != nil && !utils.IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Actor) SetHost(v string) {
	o.Host = &v
}

// GetHostRedundancyAllowed returns the HostRedundancyAllowed field value if set, zero value otherwise.
func (o *Actor) GetHostRedundancyAllowed() bool {
	if o == nil || utils.IsNil(o.HostRedundancyAllowed) {
		var ret bool
		return ret
	}
	return *o.HostRedundancyAllowed
}

// GetHostRedundancyAllowedOk returns a tuple with the HostRedundancyAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetHostRedundancyAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HostRedundancyAllowed) {
		return nil, false
	}
	return o.HostRedundancyAllowed, true
}

// HasHostRedundancyAllowed returns a boolean if a field has been set.
func (o *Actor) HasHostRedundancyAllowed() bool {
	if o != nil && !utils.IsNil(o.HostRedundancyAllowed) {
		return true
	}

	return false
}

// SetHostRedundancyAllowed gets a reference to the given bool and assigns it to the HostRedundancyAllowed field.
func (o *Actor) SetHostRedundancyAllowed(v bool) {
	o.HostRedundancyAllowed = &v
}

// GetFollowingCount returns the FollowingCount field value if set, zero value otherwise.
func (o *Actor) GetFollowingCount() int32 {
	if o == nil || utils.IsNil(o.FollowingCount) {
		var ret int32
		return ret
	}
	return *o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetFollowingCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FollowingCount) {
		return nil, false
	}
	return o.FollowingCount, true
}

// HasFollowingCount returns a boolean if a field has been set.
func (o *Actor) HasFollowingCount() bool {
	if o != nil && !utils.IsNil(o.FollowingCount) {
		return true
	}

	return false
}

// SetFollowingCount gets a reference to the given int32 and assigns it to the FollowingCount field.
func (o *Actor) SetFollowingCount(v int32) {
	o.FollowingCount = &v
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise.
func (o *Actor) GetFollowersCount() int32 {
	if o == nil || utils.IsNil(o.FollowersCount) {
		var ret int32
		return ret
	}
	return *o.FollowersCount
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetFollowersCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FollowersCount) {
		return nil, false
	}
	return o.FollowersCount, true
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *Actor) HasFollowersCount() bool {
	if o != nil && !utils.IsNil(o.FollowersCount) {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given int32 and assigns it to the FollowersCount field.
func (o *Actor) SetFollowersCount(v int32) {
	o.FollowersCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Actor) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Actor) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Actor) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Actor) GetUpdatedAt() time.Time {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Actor) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Actor) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Actor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Actor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Avatars) {
		toSerialize["avatars"] = o.Avatars
	}
	if !utils.IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !utils.IsNil(o.HostRedundancyAllowed) {
		toSerialize["hostRedundancyAllowed"] = o.HostRedundancyAllowed
	}
	if !utils.IsNil(o.FollowingCount) {
		toSerialize["followingCount"] = o.FollowingCount
	}
	if !utils.IsNil(o.FollowersCount) {
		toSerialize["followersCount"] = o.FollowersCount
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableActor struct {
	value *Actor
	isSet bool
}

func (v NullableActor) Get() *Actor {
	return v.value
}

func (v *NullableActor) Set(val *Actor) {
	v.value = val
	v.isSet = true
}

func (v NullableActor) IsSet() bool {
	return v.isSet
}

func (v *NullableActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActor(val *Actor) *NullableActor {
	return &NullableActor{value: val, isSet: true}
}

func (v NullableActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
