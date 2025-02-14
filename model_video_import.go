/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"encoding/json"
	"os"
	"time"
)

// checks if the VideoImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoImport{}

// VideoImport struct for VideoImport
type VideoImport struct {
	Id *int32 `json:"id,omitempty"`
	// remote URL where to find the import's source video
	TargetUrl *string `json:"targetUrl,omitempty"`
	// magnet URI allowing to resolve the import's source video
	MagnetUri *string `json:"magnetUri,omitempty" validate:"regexp=magnet:\\\\?xt=urn:[a-z0-9]+:[a-z0-9]{32}/i"`
	// Torrent file containing only the video file
	Torrentfile **os.File                 `json:"torrentfile,omitempty"`
	TorrentName *string                   `json:"torrentName,omitempty"`
	State       *VideoImportStateConstant `json:"state,omitempty"`
	Error       *string                   `json:"error,omitempty"`
	CreatedAt   *time.Time                `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time                `json:"updatedAt,omitempty"`
	Video       NullableVideo             `json:"video,omitempty"`
}

// NewVideoImport instantiates a new VideoImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoImport() *VideoImport {
	this := VideoImport{}
	return &this
}

// NewVideoImportWithDefaults instantiates a new VideoImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoImportWithDefaults() *VideoImport {
	this := VideoImport{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoImport) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoImport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoImport) SetId(v int32) {
	o.Id = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *VideoImport) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *VideoImport) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *VideoImport) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetMagnetUri returns the MagnetUri field value if set, zero value otherwise.
func (o *VideoImport) GetMagnetUri() string {
	if o == nil || IsNil(o.MagnetUri) {
		var ret string
		return ret
	}
	return *o.MagnetUri
}

// GetMagnetUriOk returns a tuple with the MagnetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetMagnetUriOk() (*string, bool) {
	if o == nil || IsNil(o.MagnetUri) {
		return nil, false
	}
	return o.MagnetUri, true
}

// HasMagnetUri returns a boolean if a field has been set.
func (o *VideoImport) HasMagnetUri() bool {
	if o != nil && !IsNil(o.MagnetUri) {
		return true
	}

	return false
}

// SetMagnetUri gets a reference to the given string and assigns it to the MagnetUri field.
func (o *VideoImport) SetMagnetUri(v string) {
	o.MagnetUri = &v
}

// GetTorrentfile returns the Torrentfile field value if set, zero value otherwise.
func (o *VideoImport) GetTorrentfile() *os.File {
	if o == nil || IsNil(o.Torrentfile) {
		var ret *os.File
		return ret
	}
	return *o.Torrentfile
}

// GetTorrentfileOk returns a tuple with the Torrentfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetTorrentfileOk() (**os.File, bool) {
	if o == nil || IsNil(o.Torrentfile) {
		return nil, false
	}
	return o.Torrentfile, true
}

// HasTorrentfile returns a boolean if a field has been set.
func (o *VideoImport) HasTorrentfile() bool {
	if o != nil && !IsNil(o.Torrentfile) {
		return true
	}

	return false
}

// SetTorrentfile gets a reference to the given *os.File and assigns it to the Torrentfile field.
func (o *VideoImport) SetTorrentfile(v *os.File) {
	o.Torrentfile = &v
}

// GetTorrentName returns the TorrentName field value if set, zero value otherwise.
func (o *VideoImport) GetTorrentName() string {
	if o == nil || IsNil(o.TorrentName) {
		var ret string
		return ret
	}
	return *o.TorrentName
}

// GetTorrentNameOk returns a tuple with the TorrentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetTorrentNameOk() (*string, bool) {
	if o == nil || IsNil(o.TorrentName) {
		return nil, false
	}
	return o.TorrentName, true
}

// HasTorrentName returns a boolean if a field has been set.
func (o *VideoImport) HasTorrentName() bool {
	if o != nil && !IsNil(o.TorrentName) {
		return true
	}

	return false
}

// SetTorrentName gets a reference to the given string and assigns it to the TorrentName field.
func (o *VideoImport) SetTorrentName(v string) {
	o.TorrentName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VideoImport) GetState() VideoImportStateConstant {
	if o == nil || IsNil(o.State) {
		var ret VideoImportStateConstant
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetStateOk() (*VideoImportStateConstant, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VideoImport) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given VideoImportStateConstant and assigns it to the State field.
func (o *VideoImport) SetState(v VideoImportStateConstant) {
	o.State = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *VideoImport) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *VideoImport) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *VideoImport) SetError(v string) {
	o.Error = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoImport) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoImport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoImport) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VideoImport) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoImport) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VideoImport) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VideoImport) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVideo returns the Video field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoImport) GetVideo() Video {
	if o == nil || IsNil(o.Video.Get()) {
		var ret Video
		return ret
	}
	return *o.Video.Get()
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoImport) GetVideoOk() (*Video, bool) {
	if o == nil {
		return nil, false
	}
	return o.Video.Get(), o.Video.IsSet()
}

// HasVideo returns a boolean if a field has been set.
func (o *VideoImport) HasVideo() bool {
	if o != nil && o.Video.IsSet() {
		return true
	}

	return false
}

// SetVideo gets a reference to the given NullableVideo and assigns it to the Video field.
func (o *VideoImport) SetVideo(v Video) {
	o.Video.Set(&v)
}

// SetVideoNil sets the value for Video to be an explicit nil
func (o *VideoImport) SetVideoNil() {
	o.Video.Set(nil)
}

// UnsetVideo ensures that no value is present for Video, not even an explicit nil
func (o *VideoImport) UnsetVideo() {
	o.Video.Unset()
}

func (o VideoImport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["targetUrl"] = o.TargetUrl
	}
	if !IsNil(o.MagnetUri) {
		toSerialize["magnetUri"] = o.MagnetUri
	}
	if !IsNil(o.Torrentfile) {
		toSerialize["torrentfile"] = o.Torrentfile
	}
	if !IsNil(o.TorrentName) {
		toSerialize["torrentName"] = o.TorrentName
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Video.IsSet() {
		toSerialize["video"] = o.Video.Get()
	}
	return toSerialize, nil
}

type NullableVideoImport struct {
	value *VideoImport
	isSet bool
}

func (v NullableVideoImport) Get() *VideoImport {
	return v.value
}

func (v *NullableVideoImport) Set(val *VideoImport) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoImport) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoImport(val *VideoImport) *NullableVideoImport {
	return &NullableVideoImport{value: val, isSet: true}
}

func (v NullableVideoImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
