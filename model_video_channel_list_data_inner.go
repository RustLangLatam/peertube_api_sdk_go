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

// checks if the VideoChannelListDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoChannelListDataInner{}

// VideoChannelListDataInner struct for VideoChannelListDataInner
type VideoChannelListDataInner struct {
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
	// editable name of the channel, displayed in its representations
	DisplayName *string        `json:"displayName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// text shown by default on all videos of this channel, to tell the audience how to support it
	Support      NullableString `json:"support,omitempty"`
	IsLocal      *bool          `json:"isLocal,omitempty"`
	Banners      []ActorImage   `json:"banners,omitempty"`
	OwnerAccount *Account       `json:"ownerAccount,omitempty"`
}

// NewVideoChannelListDataInner instantiates a new VideoChannelListDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoChannelListDataInner() *VideoChannelListDataInner {
	this := VideoChannelListDataInner{}
	return &this
}

// NewVideoChannelListDataInnerWithDefaults instantiates a new VideoChannelListDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoChannelListDataInnerWithDefaults() *VideoChannelListDataInner {
	this := VideoChannelListDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoChannelListDataInner) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VideoChannelListDataInner) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VideoChannelListDataInner) SetName(v string) {
	o.Name = &v
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetAvatars() []ActorImage {
	if o == nil || IsNil(o.Avatars) {
		var ret []ActorImage
		return ret
	}
	return o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetAvatarsOk() ([]ActorImage, bool) {
	if o == nil || IsNil(o.Avatars) {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasAvatars() bool {
	if o != nil && !IsNil(o.Avatars) {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given []ActorImage and assigns it to the Avatars field.
func (o *VideoChannelListDataInner) SetAvatars(v []ActorImage) {
	o.Avatars = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *VideoChannelListDataInner) SetHost(v string) {
	o.Host = &v
}

// GetHostRedundancyAllowed returns the HostRedundancyAllowed field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetHostRedundancyAllowed() bool {
	if o == nil || IsNil(o.HostRedundancyAllowed) {
		var ret bool
		return ret
	}
	return *o.HostRedundancyAllowed
}

// GetHostRedundancyAllowedOk returns a tuple with the HostRedundancyAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetHostRedundancyAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.HostRedundancyAllowed) {
		return nil, false
	}
	return o.HostRedundancyAllowed, true
}

// HasHostRedundancyAllowed returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasHostRedundancyAllowed() bool {
	if o != nil && !IsNil(o.HostRedundancyAllowed) {
		return true
	}

	return false
}

// SetHostRedundancyAllowed gets a reference to the given bool and assigns it to the HostRedundancyAllowed field.
func (o *VideoChannelListDataInner) SetHostRedundancyAllowed(v bool) {
	o.HostRedundancyAllowed = &v
}

// GetFollowingCount returns the FollowingCount field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetFollowingCount() int32 {
	if o == nil || IsNil(o.FollowingCount) {
		var ret int32
		return ret
	}
	return *o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetFollowingCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FollowingCount) {
		return nil, false
	}
	return o.FollowingCount, true
}

// HasFollowingCount returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasFollowingCount() bool {
	if o != nil && !IsNil(o.FollowingCount) {
		return true
	}

	return false
}

// SetFollowingCount gets a reference to the given int32 and assigns it to the FollowingCount field.
func (o *VideoChannelListDataInner) SetFollowingCount(v int32) {
	o.FollowingCount = &v
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount) {
		var ret int32
		return ret
	}
	return *o.FollowersCount
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetFollowersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FollowersCount) {
		return nil, false
	}
	return o.FollowersCount, true
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasFollowersCount() bool {
	if o != nil && !IsNil(o.FollowersCount) {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given int32 and assigns it to the FollowersCount field.
func (o *VideoChannelListDataInner) SetFollowersCount(v int32) {
	o.FollowersCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoChannelListDataInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VideoChannelListDataInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *VideoChannelListDataInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoChannelListDataInner) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoChannelListDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VideoChannelListDataInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VideoChannelListDataInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VideoChannelListDataInner) UnsetDescription() {
	o.Description.Unset()
}

// GetSupport returns the Support field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoChannelListDataInner) GetSupport() string {
	if o == nil || IsNil(o.Support.Get()) {
		var ret string
		return ret
	}
	return *o.Support.Get()
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoChannelListDataInner) GetSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Support.Get(), o.Support.IsSet()
}

// HasSupport returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasSupport() bool {
	if o != nil && o.Support.IsSet() {
		return true
	}

	return false
}

// SetSupport gets a reference to the given NullableString and assigns it to the Support field.
func (o *VideoChannelListDataInner) SetSupport(v string) {
	o.Support.Set(&v)
}

// SetSupportNil sets the value for Support to be an explicit nil
func (o *VideoChannelListDataInner) SetSupportNil() {
	o.Support.Set(nil)
}

// UnsetSupport ensures that no value is present for Support, not even an explicit nil
func (o *VideoChannelListDataInner) UnsetSupport() {
	o.Support.Unset()
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *VideoChannelListDataInner) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetBanners returns the Banners field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetBanners() []ActorImage {
	if o == nil || IsNil(o.Banners) {
		var ret []ActorImage
		return ret
	}
	return o.Banners
}

// GetBannersOk returns a tuple with the Banners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetBannersOk() ([]ActorImage, bool) {
	if o == nil || IsNil(o.Banners) {
		return nil, false
	}
	return o.Banners, true
}

// HasBanners returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasBanners() bool {
	if o != nil && !IsNil(o.Banners) {
		return true
	}

	return false
}

// SetBanners gets a reference to the given []ActorImage and assigns it to the Banners field.
func (o *VideoChannelListDataInner) SetBanners(v []ActorImage) {
	o.Banners = v
}

// GetOwnerAccount returns the OwnerAccount field value if set, zero value otherwise.
func (o *VideoChannelListDataInner) GetOwnerAccount() Account {
	if o == nil || IsNil(o.OwnerAccount) {
		var ret Account
		return ret
	}
	return *o.OwnerAccount
}

// GetOwnerAccountOk returns a tuple with the OwnerAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoChannelListDataInner) GetOwnerAccountOk() (*Account, bool) {
	if o == nil || IsNil(o.OwnerAccount) {
		return nil, false
	}
	return o.OwnerAccount, true
}

// HasOwnerAccount returns a boolean if a field has been set.
func (o *VideoChannelListDataInner) HasOwnerAccount() bool {
	if o != nil && !IsNil(o.OwnerAccount) {
		return true
	}

	return false
}

// SetOwnerAccount gets a reference to the given Account and assigns it to the OwnerAccount field.
func (o *VideoChannelListDataInner) SetOwnerAccount(v Account) {
	o.OwnerAccount = &v
}

func (o VideoChannelListDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoChannelListDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Avatars) {
		toSerialize["avatars"] = o.Avatars
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.HostRedundancyAllowed) {
		toSerialize["hostRedundancyAllowed"] = o.HostRedundancyAllowed
	}
	if !IsNil(o.FollowingCount) {
		toSerialize["followingCount"] = o.FollowingCount
	}
	if !IsNil(o.FollowersCount) {
		toSerialize["followersCount"] = o.FollowersCount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Support.IsSet() {
		toSerialize["support"] = o.Support.Get()
	}
	if !IsNil(o.IsLocal) {
		toSerialize["isLocal"] = o.IsLocal
	}
	if !IsNil(o.Banners) {
		toSerialize["banners"] = o.Banners
	}
	if !IsNil(o.OwnerAccount) {
		toSerialize["ownerAccount"] = o.OwnerAccount
	}
	return toSerialize, nil
}

type NullableVideoChannelListDataInner struct {
	value *VideoChannelListDataInner
	isSet bool
}

func (v NullableVideoChannelListDataInner) Get() *VideoChannelListDataInner {
	return v.value
}

func (v *NullableVideoChannelListDataInner) Set(val *VideoChannelListDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoChannelListDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoChannelListDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoChannelListDataInner(val *VideoChannelListDataInner) *NullableVideoChannelListDataInner {
	return &NullableVideoChannelListDataInner{value: val, isSet: true}
}

func (v NullableVideoChannelListDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoChannelListDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
