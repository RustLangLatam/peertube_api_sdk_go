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

// checks if the VideoStreamingPlaylists type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoStreamingPlaylists{}

// VideoStreamingPlaylists struct for VideoStreamingPlaylists
type VideoStreamingPlaylists struct {
	PlaylistUrl       *string `json:"playlistUrl,omitempty"`
	SegmentsSha256Url *string `json:"segmentsSha256Url,omitempty"`
	// Video files associated to this playlist.  The difference with the root `files` property is that these files are fragmented, so they can be used in this streaming playlist (HLS, etc.)
	Files        []VideoFile                                   `json:"files,omitempty"`
	Redundancies []VideoStreamingPlaylistsHLSRedundanciesInner `json:"redundancies,omitempty"`
	Id           *int32                                        `json:"id,omitempty"`
	// Playlist type: - `1`: HLS
	Type *int32 `json:"type,omitempty"`
}

// NewVideoStreamingPlaylists instantiates a new VideoStreamingPlaylists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoStreamingPlaylists() *VideoStreamingPlaylists {
	this := VideoStreamingPlaylists{}
	return &this
}

// NewVideoStreamingPlaylistsWithDefaults instantiates a new VideoStreamingPlaylists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoStreamingPlaylistsWithDefaults() *VideoStreamingPlaylists {
	this := VideoStreamingPlaylists{}
	return &this
}

// GetPlaylistUrl returns the PlaylistUrl field value if set, zero value otherwise.
func (o *VideoStreamingPlaylists) GetPlaylistUrl() string {
	if o == nil || IsNil(o.PlaylistUrl) {
		var ret string
		return ret
	}
	return *o.PlaylistUrl
}

// GetPlaylistUrlOk returns a tuple with the PlaylistUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStreamingPlaylists) GetPlaylistUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistUrl) {
		return nil, false
	}
	return o.PlaylistUrl, true
}

// HasPlaylistUrl returns a boolean if a field has been set.
func (o *VideoStreamingPlaylists) HasPlaylistUrl() bool {
	if o != nil && !IsNil(o.PlaylistUrl) {
		return true
	}

	return false
}

// SetPlaylistUrl gets a reference to the given string and assigns it to the PlaylistUrl field.
func (o *VideoStreamingPlaylists) SetPlaylistUrl(v string) {
	o.PlaylistUrl = &v
}

// GetSegmentsSha256Url returns the SegmentsSha256Url field value if set, zero value otherwise.
func (o *VideoStreamingPlaylists) GetSegmentsSha256Url() string {
	if o == nil || IsNil(o.SegmentsSha256Url) {
		var ret string
		return ret
	}
	return *o.SegmentsSha256Url
}

// GetSegmentsSha256UrlOk returns a tuple with the SegmentsSha256Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStreamingPlaylists) GetSegmentsSha256UrlOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentsSha256Url) {
		return nil, false
	}
	return o.SegmentsSha256Url, true
}

// HasSegmentsSha256Url returns a boolean if a field has been set.
func (o *VideoStreamingPlaylists) HasSegmentsSha256Url() bool {
	if o != nil && !IsNil(o.SegmentsSha256Url) {
		return true
	}

	return false
}

// SetSegmentsSha256Url gets a reference to the given string and assigns it to the SegmentsSha256Url field.
func (o *VideoStreamingPlaylists) SetSegmentsSha256Url(v string) {
	o.SegmentsSha256Url = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *VideoStreamingPlaylists) GetFiles() []VideoFile {
	if o == nil || IsNil(o.Files) {
		var ret []VideoFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStreamingPlaylists) GetFilesOk() ([]VideoFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *VideoStreamingPlaylists) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []VideoFile and assigns it to the Files field.
func (o *VideoStreamingPlaylists) SetFiles(v []VideoFile) {
	o.Files = v
}

// GetRedundancies returns the Redundancies field value if set, zero value otherwise.
func (o *VideoStreamingPlaylists) GetRedundancies() []VideoStreamingPlaylistsHLSRedundanciesInner {
	if o == nil || IsNil(o.Redundancies) {
		var ret []VideoStreamingPlaylistsHLSRedundanciesInner
		return ret
	}
	return o.Redundancies
}

// GetRedundanciesOk returns a tuple with the Redundancies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStreamingPlaylists) GetRedundanciesOk() ([]VideoStreamingPlaylistsHLSRedundanciesInner, bool) {
	if o == nil || IsNil(o.Redundancies) {
		return nil, false
	}
	return o.Redundancies, true
}

// HasRedundancies returns a boolean if a field has been set.
func (o *VideoStreamingPlaylists) HasRedundancies() bool {
	if o != nil && !IsNil(o.Redundancies) {
		return true
	}

	return false
}

// SetRedundancies gets a reference to the given []VideoStreamingPlaylistsHLSRedundanciesInner and assigns it to the Redundancies field.
func (o *VideoStreamingPlaylists) SetRedundancies(v []VideoStreamingPlaylistsHLSRedundanciesInner) {
	o.Redundancies = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoStreamingPlaylists) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStreamingPlaylists) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoStreamingPlaylists) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoStreamingPlaylists) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VideoStreamingPlaylists) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStreamingPlaylists) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VideoStreamingPlaylists) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *VideoStreamingPlaylists) SetType(v int32) {
	o.Type = &v
}

func (o VideoStreamingPlaylists) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoStreamingPlaylists) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaylistUrl) {
		toSerialize["playlistUrl"] = o.PlaylistUrl
	}
	if !IsNil(o.SegmentsSha256Url) {
		toSerialize["segmentsSha256Url"] = o.SegmentsSha256Url
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.Redundancies) {
		toSerialize["redundancies"] = o.Redundancies
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVideoStreamingPlaylists struct {
	value *VideoStreamingPlaylists
	isSet bool
}

func (v NullableVideoStreamingPlaylists) Get() *VideoStreamingPlaylists {
	return v.value
}

func (v *NullableVideoStreamingPlaylists) Set(val *VideoStreamingPlaylists) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoStreamingPlaylists) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoStreamingPlaylists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoStreamingPlaylists(val *VideoStreamingPlaylists) *NullableVideoStreamingPlaylists {
	return &NullableVideoStreamingPlaylists{value: val, isSet: true}
}

func (v NullableVideoStreamingPlaylists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoStreamingPlaylists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
