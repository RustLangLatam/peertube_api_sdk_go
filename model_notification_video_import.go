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

// checks if the NotificationVideoImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationVideoImport{}

// NotificationVideoImport struct for NotificationVideoImport
type NotificationVideoImport struct {
	Id          *int32         `json:"id,omitempty"`
	Video       *VideoInfo     `json:"video,omitempty"`
	TorrentName NullableString `json:"torrentName,omitempty"`
	// magnet URI allowing to resolve the import's source video
	MagnetUri *string        `json:"magnetUri,omitempty" validate:"regexp=magnet:\\\\?xt=urn:[a-z0-9]+:[a-z0-9]{32}/i"`
	TargetUri NullableString `json:"targetUri,omitempty"`
}

// NewNotificationVideoImport instantiates a new NotificationVideoImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationVideoImport() *NotificationVideoImport {
	this := NotificationVideoImport{}
	return &this
}

// NewNotificationVideoImportWithDefaults instantiates a new NotificationVideoImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationVideoImportWithDefaults() *NotificationVideoImport {
	this := NotificationVideoImport{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationVideoImport) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationVideoImport) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationVideoImport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NotificationVideoImport) SetId(v int32) {
	o.Id = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *NotificationVideoImport) GetVideo() VideoInfo {
	if o == nil || IsNil(o.Video) {
		var ret VideoInfo
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationVideoImport) GetVideoOk() (*VideoInfo, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *NotificationVideoImport) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given VideoInfo and assigns it to the Video field.
func (o *NotificationVideoImport) SetVideo(v VideoInfo) {
	o.Video = &v
}

// GetTorrentName returns the TorrentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationVideoImport) GetTorrentName() string {
	if o == nil || IsNil(o.TorrentName.Get()) {
		var ret string
		return ret
	}
	return *o.TorrentName.Get()
}

// GetTorrentNameOk returns a tuple with the TorrentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationVideoImport) GetTorrentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TorrentName.Get(), o.TorrentName.IsSet()
}

// HasTorrentName returns a boolean if a field has been set.
func (o *NotificationVideoImport) HasTorrentName() bool {
	if o != nil && o.TorrentName.IsSet() {
		return true
	}

	return false
}

// SetTorrentName gets a reference to the given NullableString and assigns it to the TorrentName field.
func (o *NotificationVideoImport) SetTorrentName(v string) {
	o.TorrentName.Set(&v)
}

// SetTorrentNameNil sets the value for TorrentName to be an explicit nil
func (o *NotificationVideoImport) SetTorrentNameNil() {
	o.TorrentName.Set(nil)
}

// UnsetTorrentName ensures that no value is present for TorrentName, not even an explicit nil
func (o *NotificationVideoImport) UnsetTorrentName() {
	o.TorrentName.Unset()
}

// GetMagnetUri returns the MagnetUri field value if set, zero value otherwise.
func (o *NotificationVideoImport) GetMagnetUri() string {
	if o == nil || IsNil(o.MagnetUri) {
		var ret string
		return ret
	}
	return *o.MagnetUri
}

// GetMagnetUriOk returns a tuple with the MagnetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationVideoImport) GetMagnetUriOk() (*string, bool) {
	if o == nil || IsNil(o.MagnetUri) {
		return nil, false
	}
	return o.MagnetUri, true
}

// HasMagnetUri returns a boolean if a field has been set.
func (o *NotificationVideoImport) HasMagnetUri() bool {
	if o != nil && !IsNil(o.MagnetUri) {
		return true
	}

	return false
}

// SetMagnetUri gets a reference to the given string and assigns it to the MagnetUri field.
func (o *NotificationVideoImport) SetMagnetUri(v string) {
	o.MagnetUri = &v
}

// GetTargetUri returns the TargetUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationVideoImport) GetTargetUri() string {
	if o == nil || IsNil(o.TargetUri.Get()) {
		var ret string
		return ret
	}
	return *o.TargetUri.Get()
}

// GetTargetUriOk returns a tuple with the TargetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationVideoImport) GetTargetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetUri.Get(), o.TargetUri.IsSet()
}

// HasTargetUri returns a boolean if a field has been set.
func (o *NotificationVideoImport) HasTargetUri() bool {
	if o != nil && o.TargetUri.IsSet() {
		return true
	}

	return false
}

// SetTargetUri gets a reference to the given NullableString and assigns it to the TargetUri field.
func (o *NotificationVideoImport) SetTargetUri(v string) {
	o.TargetUri.Set(&v)
}

// SetTargetUriNil sets the value for TargetUri to be an explicit nil
func (o *NotificationVideoImport) SetTargetUriNil() {
	o.TargetUri.Set(nil)
}

// UnsetTargetUri ensures that no value is present for TargetUri, not even an explicit nil
func (o *NotificationVideoImport) UnsetTargetUri() {
	o.TargetUri.Unset()
}

func (o NotificationVideoImport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationVideoImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if o.TorrentName.IsSet() {
		toSerialize["torrentName"] = o.TorrentName.Get()
	}
	if !IsNil(o.MagnetUri) {
		toSerialize["magnetUri"] = o.MagnetUri
	}
	if o.TargetUri.IsSet() {
		toSerialize["targetUri"] = o.TargetUri.Get()
	}
	return toSerialize, nil
}

type NullableNotificationVideoImport struct {
	value *NotificationVideoImport
	isSet bool
}

func (v NullableNotificationVideoImport) Get() *NotificationVideoImport {
	return v.value
}

func (v *NullableNotificationVideoImport) Set(val *NotificationVideoImport) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationVideoImport) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationVideoImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationVideoImport(val *NotificationVideoImport) *NullableNotificationVideoImport {
	return &NullableNotificationVideoImport{value: val, isSet: true}
}

func (v NullableNotificationVideoImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationVideoImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
