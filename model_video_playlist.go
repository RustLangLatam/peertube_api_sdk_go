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

// checks if the VideoPlaylist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoPlaylist{}

// VideoPlaylist struct for VideoPlaylist
type VideoPlaylist struct {
	Id   *int32  `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty" validate:"regexp=^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"`
	// translation of a uuid v4 with a bigger alphabet to have a shorter uuid
	ShortUUID     *string                       `json:"shortUUID,omitempty"`
	CreatedAt     *time.Time                    `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time                    `json:"updatedAt,omitempty"`
	Description   *string                       `json:"description,omitempty"`
	DisplayName   *string                       `json:"displayName,omitempty"`
	IsLocal       *bool                         `json:"isLocal,omitempty"`
	VideoLength   *int32                        `json:"videoLength,omitempty"`
	ThumbnailPath *string                       `json:"thumbnailPath,omitempty"`
	Privacy       *VideoPlaylistPrivacyConstant `json:"privacy,omitempty"`
	Type          *VideoPlaylistTypeConstant    `json:"type,omitempty"`
	OwnerAccount  *AccountSummary               `json:"ownerAccount,omitempty"`
	VideoChannel  *VideoChannelSummary          `json:"videoChannel,omitempty"`
}

// NewVideoPlaylist instantiates a new VideoPlaylist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoPlaylist() *VideoPlaylist {
	this := VideoPlaylist{}
	return &this
}

// NewVideoPlaylistWithDefaults instantiates a new VideoPlaylist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoPlaylistWithDefaults() *VideoPlaylist {
	this := VideoPlaylist{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoPlaylist) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoPlaylist) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoPlaylist) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VideoPlaylist) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VideoPlaylist) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VideoPlaylist) SetUuid(v string) {
	o.Uuid = &v
}

// GetShortUUID returns the ShortUUID field value if set, zero value otherwise.
func (o *VideoPlaylist) GetShortUUID() string {
	if o == nil || IsNil(o.ShortUUID) {
		var ret string
		return ret
	}
	return *o.ShortUUID
}

// GetShortUUIDOk returns a tuple with the ShortUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetShortUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.ShortUUID) {
		return nil, false
	}
	return o.ShortUUID, true
}

// HasShortUUID returns a boolean if a field has been set.
func (o *VideoPlaylist) HasShortUUID() bool {
	if o != nil && !IsNil(o.ShortUUID) {
		return true
	}

	return false
}

// SetShortUUID gets a reference to the given string and assigns it to the ShortUUID field.
func (o *VideoPlaylist) SetShortUUID(v string) {
	o.ShortUUID = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoPlaylist) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoPlaylist) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoPlaylist) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VideoPlaylist) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VideoPlaylist) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VideoPlaylist) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VideoPlaylist) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VideoPlaylist) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VideoPlaylist) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *VideoPlaylist) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *VideoPlaylist) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *VideoPlaylist) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *VideoPlaylist) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *VideoPlaylist) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *VideoPlaylist) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetVideoLength returns the VideoLength field value if set, zero value otherwise.
func (o *VideoPlaylist) GetVideoLength() int32 {
	if o == nil || IsNil(o.VideoLength) {
		var ret int32
		return ret
	}
	return *o.VideoLength
}

// GetVideoLengthOk returns a tuple with the VideoLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetVideoLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoLength) {
		return nil, false
	}
	return o.VideoLength, true
}

// HasVideoLength returns a boolean if a field has been set.
func (o *VideoPlaylist) HasVideoLength() bool {
	if o != nil && !IsNil(o.VideoLength) {
		return true
	}

	return false
}

// SetVideoLength gets a reference to the given int32 and assigns it to the VideoLength field.
func (o *VideoPlaylist) SetVideoLength(v int32) {
	o.VideoLength = &v
}

// GetThumbnailPath returns the ThumbnailPath field value if set, zero value otherwise.
func (o *VideoPlaylist) GetThumbnailPath() string {
	if o == nil || IsNil(o.ThumbnailPath) {
		var ret string
		return ret
	}
	return *o.ThumbnailPath
}

// GetThumbnailPathOk returns a tuple with the ThumbnailPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetThumbnailPathOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailPath) {
		return nil, false
	}
	return o.ThumbnailPath, true
}

// HasThumbnailPath returns a boolean if a field has been set.
func (o *VideoPlaylist) HasThumbnailPath() bool {
	if o != nil && !IsNil(o.ThumbnailPath) {
		return true
	}

	return false
}

// SetThumbnailPath gets a reference to the given string and assigns it to the ThumbnailPath field.
func (o *VideoPlaylist) SetThumbnailPath(v string) {
	o.ThumbnailPath = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *VideoPlaylist) GetPrivacy() VideoPlaylistPrivacyConstant {
	if o == nil || IsNil(o.Privacy) {
		var ret VideoPlaylistPrivacyConstant
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetPrivacyOk() (*VideoPlaylistPrivacyConstant, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *VideoPlaylist) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given VideoPlaylistPrivacyConstant and assigns it to the Privacy field.
func (o *VideoPlaylist) SetPrivacy(v VideoPlaylistPrivacyConstant) {
	o.Privacy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VideoPlaylist) GetType() VideoPlaylistTypeConstant {
	if o == nil || IsNil(o.Type) {
		var ret VideoPlaylistTypeConstant
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetTypeOk() (*VideoPlaylistTypeConstant, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VideoPlaylist) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given VideoPlaylistTypeConstant and assigns it to the Type field.
func (o *VideoPlaylist) SetType(v VideoPlaylistTypeConstant) {
	o.Type = &v
}

// GetOwnerAccount returns the OwnerAccount field value if set, zero value otherwise.
func (o *VideoPlaylist) GetOwnerAccount() AccountSummary {
	if o == nil || IsNil(o.OwnerAccount) {
		var ret AccountSummary
		return ret
	}
	return *o.OwnerAccount
}

// GetOwnerAccountOk returns a tuple with the OwnerAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetOwnerAccountOk() (*AccountSummary, bool) {
	if o == nil || IsNil(o.OwnerAccount) {
		return nil, false
	}
	return o.OwnerAccount, true
}

// HasOwnerAccount returns a boolean if a field has been set.
func (o *VideoPlaylist) HasOwnerAccount() bool {
	if o != nil && !IsNil(o.OwnerAccount) {
		return true
	}

	return false
}

// SetOwnerAccount gets a reference to the given AccountSummary and assigns it to the OwnerAccount field.
func (o *VideoPlaylist) SetOwnerAccount(v AccountSummary) {
	o.OwnerAccount = &v
}

// GetVideoChannel returns the VideoChannel field value if set, zero value otherwise.
func (o *VideoPlaylist) GetVideoChannel() VideoChannelSummary {
	if o == nil || IsNil(o.VideoChannel) {
		var ret VideoChannelSummary
		return ret
	}
	return *o.VideoChannel
}

// GetVideoChannelOk returns a tuple with the VideoChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoPlaylist) GetVideoChannelOk() (*VideoChannelSummary, bool) {
	if o == nil || IsNil(o.VideoChannel) {
		return nil, false
	}
	return o.VideoChannel, true
}

// HasVideoChannel returns a boolean if a field has been set.
func (o *VideoPlaylist) HasVideoChannel() bool {
	if o != nil && !IsNil(o.VideoChannel) {
		return true
	}

	return false
}

// SetVideoChannel gets a reference to the given VideoChannelSummary and assigns it to the VideoChannel field.
func (o *VideoPlaylist) SetVideoChannel(v VideoChannelSummary) {
	o.VideoChannel = &v
}

func (o VideoPlaylist) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoPlaylist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.ShortUUID) {
		toSerialize["shortUUID"] = o.ShortUUID
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IsLocal) {
		toSerialize["isLocal"] = o.IsLocal
	}
	if !IsNil(o.VideoLength) {
		toSerialize["videoLength"] = o.VideoLength
	}
	if !IsNil(o.ThumbnailPath) {
		toSerialize["thumbnailPath"] = o.ThumbnailPath
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.OwnerAccount) {
		toSerialize["ownerAccount"] = o.OwnerAccount
	}
	if !IsNil(o.VideoChannel) {
		toSerialize["videoChannel"] = o.VideoChannel
	}
	return toSerialize, nil
}

type NullableVideoPlaylist struct {
	value *VideoPlaylist
	isSet bool
}

func (v NullableVideoPlaylist) Get() *VideoPlaylist {
	return v.value
}

func (v *NullableVideoPlaylist) Set(val *VideoPlaylist) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoPlaylist) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoPlaylist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoPlaylist(val *VideoPlaylist) *NullableVideoPlaylist {
	return &NullableVideoPlaylist{value: val, isSet: true}
}

func (v NullableVideoPlaylist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoPlaylist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
