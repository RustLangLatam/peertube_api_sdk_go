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

// checks if the VideoFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoFile{}

// VideoFile struct for VideoFile
type VideoFile struct {
	Id *int32 `json:"id,omitempty"`
	// magnet URI allowing to resolve the video via BitTorrent without a metainfo file
	MagnetUri  *string                  `json:"magnetUri,omitempty" validate:"regexp=magnet:\\\\?xt=urn:[a-z0-9]+:[a-z0-9]{32}/i"`
	Resolution *VideoResolutionConstant `json:"resolution,omitempty"`
	// Video file size in bytes
	Size *int32 `json:"size,omitempty"`
	// Direct URL of the torrent file
	TorrentUrl *string `json:"torrentUrl,omitempty"`
	// URL endpoint that transfers the torrent file as an attachment (so that the browser opens a download dialog)
	TorrentDownloadUrl *string `json:"torrentDownloadUrl,omitempty"`
	// Direct URL of the video
	FileUrl *string `json:"fileUrl,omitempty"`
	// URL endpoint that transfers the video file as an attachment (so that the browser opens a download dialog)
	FileDownloadUrl *string `json:"fileDownloadUrl,omitempty"`
	// Frames per second of the video file
	Fps *float32 `json:"fps,omitempty"`
	// **PeerTube >= 6.1** Video stream width
	Width *float32 `json:"width,omitempty"`
	// **PeerTube >= 6.1** Video stream height
	Height *float32 `json:"height,omitempty"`
	// URL dereferencing the output of ffprobe on the file
	MetadataUrl *string `json:"metadataUrl,omitempty"`
	// **PeerTube >= 6.2** The file container has an audio stream
	HasAudio *bool `json:"hasAudio,omitempty"`
	// **PeerTube >= 6.2** The file container has a video stream
	HasVideo *bool        `json:"hasVideo,omitempty"`
	Storage  *FileStorage `json:"storage,omitempty"`
}

// NewVideoFile instantiates a new VideoFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoFile() *VideoFile {
	this := VideoFile{}
	return &this
}

// NewVideoFileWithDefaults instantiates a new VideoFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoFileWithDefaults() *VideoFile {
	this := VideoFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoFile) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoFile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoFile) SetId(v int32) {
	o.Id = &v
}

// GetMagnetUri returns the MagnetUri field value if set, zero value otherwise.
func (o *VideoFile) GetMagnetUri() string {
	if o == nil || IsNil(o.MagnetUri) {
		var ret string
		return ret
	}
	return *o.MagnetUri
}

// GetMagnetUriOk returns a tuple with the MagnetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetMagnetUriOk() (*string, bool) {
	if o == nil || IsNil(o.MagnetUri) {
		return nil, false
	}
	return o.MagnetUri, true
}

// HasMagnetUri returns a boolean if a field has been set.
func (o *VideoFile) HasMagnetUri() bool {
	if o != nil && !IsNil(o.MagnetUri) {
		return true
	}

	return false
}

// SetMagnetUri gets a reference to the given string and assigns it to the MagnetUri field.
func (o *VideoFile) SetMagnetUri(v string) {
	o.MagnetUri = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *VideoFile) GetResolution() VideoResolutionConstant {
	if o == nil || IsNil(o.Resolution) {
		var ret VideoResolutionConstant
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetResolutionOk() (*VideoResolutionConstant, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *VideoFile) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given VideoResolutionConstant and assigns it to the Resolution field.
func (o *VideoFile) SetResolution(v VideoResolutionConstant) {
	o.Resolution = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VideoFile) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VideoFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *VideoFile) SetSize(v int32) {
	o.Size = &v
}

// GetTorrentUrl returns the TorrentUrl field value if set, zero value otherwise.
func (o *VideoFile) GetTorrentUrl() string {
	if o == nil || IsNil(o.TorrentUrl) {
		var ret string
		return ret
	}
	return *o.TorrentUrl
}

// GetTorrentUrlOk returns a tuple with the TorrentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetTorrentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TorrentUrl) {
		return nil, false
	}
	return o.TorrentUrl, true
}

// HasTorrentUrl returns a boolean if a field has been set.
func (o *VideoFile) HasTorrentUrl() bool {
	if o != nil && !IsNil(o.TorrentUrl) {
		return true
	}

	return false
}

// SetTorrentUrl gets a reference to the given string and assigns it to the TorrentUrl field.
func (o *VideoFile) SetTorrentUrl(v string) {
	o.TorrentUrl = &v
}

// GetTorrentDownloadUrl returns the TorrentDownloadUrl field value if set, zero value otherwise.
func (o *VideoFile) GetTorrentDownloadUrl() string {
	if o == nil || IsNil(o.TorrentDownloadUrl) {
		var ret string
		return ret
	}
	return *o.TorrentDownloadUrl
}

// GetTorrentDownloadUrlOk returns a tuple with the TorrentDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetTorrentDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TorrentDownloadUrl) {
		return nil, false
	}
	return o.TorrentDownloadUrl, true
}

// HasTorrentDownloadUrl returns a boolean if a field has been set.
func (o *VideoFile) HasTorrentDownloadUrl() bool {
	if o != nil && !IsNil(o.TorrentDownloadUrl) {
		return true
	}

	return false
}

// SetTorrentDownloadUrl gets a reference to the given string and assigns it to the TorrentDownloadUrl field.
func (o *VideoFile) SetTorrentDownloadUrl(v string) {
	o.TorrentDownloadUrl = &v
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise.
func (o *VideoFile) GetFileUrl() string {
	if o == nil || IsNil(o.FileUrl) {
		var ret string
		return ret
	}
	return *o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileUrl) {
		return nil, false
	}
	return o.FileUrl, true
}

// HasFileUrl returns a boolean if a field has been set.
func (o *VideoFile) HasFileUrl() bool {
	if o != nil && !IsNil(o.FileUrl) {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given string and assigns it to the FileUrl field.
func (o *VideoFile) SetFileUrl(v string) {
	o.FileUrl = &v
}

// GetFileDownloadUrl returns the FileDownloadUrl field value if set, zero value otherwise.
func (o *VideoFile) GetFileDownloadUrl() string {
	if o == nil || IsNil(o.FileDownloadUrl) {
		var ret string
		return ret
	}
	return *o.FileDownloadUrl
}

// GetFileDownloadUrlOk returns a tuple with the FileDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetFileDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileDownloadUrl) {
		return nil, false
	}
	return o.FileDownloadUrl, true
}

// HasFileDownloadUrl returns a boolean if a field has been set.
func (o *VideoFile) HasFileDownloadUrl() bool {
	if o != nil && !IsNil(o.FileDownloadUrl) {
		return true
	}

	return false
}

// SetFileDownloadUrl gets a reference to the given string and assigns it to the FileDownloadUrl field.
func (o *VideoFile) SetFileDownloadUrl(v string) {
	o.FileDownloadUrl = &v
}

// GetFps returns the Fps field value if set, zero value otherwise.
func (o *VideoFile) GetFps() float32 {
	if o == nil || IsNil(o.Fps) {
		var ret float32
		return ret
	}
	return *o.Fps
}

// GetFpsOk returns a tuple with the Fps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetFpsOk() (*float32, bool) {
	if o == nil || IsNil(o.Fps) {
		return nil, false
	}
	return o.Fps, true
}

// HasFps returns a boolean if a field has been set.
func (o *VideoFile) HasFps() bool {
	if o != nil && !IsNil(o.Fps) {
		return true
	}

	return false
}

// SetFps gets a reference to the given float32 and assigns it to the Fps field.
func (o *VideoFile) SetFps(v float32) {
	o.Fps = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VideoFile) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VideoFile) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *VideoFile) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VideoFile) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VideoFile) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *VideoFile) SetHeight(v float32) {
	o.Height = &v
}

// GetMetadataUrl returns the MetadataUrl field value if set, zero value otherwise.
func (o *VideoFile) GetMetadataUrl() string {
	if o == nil || IsNil(o.MetadataUrl) {
		var ret string
		return ret
	}
	return *o.MetadataUrl
}

// GetMetadataUrlOk returns a tuple with the MetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetMetadataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataUrl) {
		return nil, false
	}
	return o.MetadataUrl, true
}

// HasMetadataUrl returns a boolean if a field has been set.
func (o *VideoFile) HasMetadataUrl() bool {
	if o != nil && !IsNil(o.MetadataUrl) {
		return true
	}

	return false
}

// SetMetadataUrl gets a reference to the given string and assigns it to the MetadataUrl field.
func (o *VideoFile) SetMetadataUrl(v string) {
	o.MetadataUrl = &v
}

// GetHasAudio returns the HasAudio field value if set, zero value otherwise.
func (o *VideoFile) GetHasAudio() bool {
	if o == nil || IsNil(o.HasAudio) {
		var ret bool
		return ret
	}
	return *o.HasAudio
}

// GetHasAudioOk returns a tuple with the HasAudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetHasAudioOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAudio) {
		return nil, false
	}
	return o.HasAudio, true
}

// HasHasAudio returns a boolean if a field has been set.
func (o *VideoFile) HasHasAudio() bool {
	if o != nil && !IsNil(o.HasAudio) {
		return true
	}

	return false
}

// SetHasAudio gets a reference to the given bool and assigns it to the HasAudio field.
func (o *VideoFile) SetHasAudio(v bool) {
	o.HasAudio = &v
}

// GetHasVideo returns the HasVideo field value if set, zero value otherwise.
func (o *VideoFile) GetHasVideo() bool {
	if o == nil || IsNil(o.HasVideo) {
		var ret bool
		return ret
	}
	return *o.HasVideo
}

// GetHasVideoOk returns a tuple with the HasVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetHasVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.HasVideo) {
		return nil, false
	}
	return o.HasVideo, true
}

// HasHasVideo returns a boolean if a field has been set.
func (o *VideoFile) HasHasVideo() bool {
	if o != nil && !IsNil(o.HasVideo) {
		return true
	}

	return false
}

// SetHasVideo gets a reference to the given bool and assigns it to the HasVideo field.
func (o *VideoFile) SetHasVideo(v bool) {
	o.HasVideo = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *VideoFile) GetStorage() FileStorage {
	if o == nil || IsNil(o.Storage) {
		var ret FileStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetStorageOk() (*FileStorage, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *VideoFile) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given FileStorage and assigns it to the Storage field.
func (o *VideoFile) SetStorage(v FileStorage) {
	o.Storage = &v
}

func (o VideoFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MagnetUri) {
		toSerialize["magnetUri"] = o.MagnetUri
	}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.TorrentUrl) {
		toSerialize["torrentUrl"] = o.TorrentUrl
	}
	if !IsNil(o.TorrentDownloadUrl) {
		toSerialize["torrentDownloadUrl"] = o.TorrentDownloadUrl
	}
	if !IsNil(o.FileUrl) {
		toSerialize["fileUrl"] = o.FileUrl
	}
	if !IsNil(o.FileDownloadUrl) {
		toSerialize["fileDownloadUrl"] = o.FileDownloadUrl
	}
	if !IsNil(o.Fps) {
		toSerialize["fps"] = o.Fps
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.MetadataUrl) {
		toSerialize["metadataUrl"] = o.MetadataUrl
	}
	if !IsNil(o.HasAudio) {
		toSerialize["hasAudio"] = o.HasAudio
	}
	if !IsNil(o.HasVideo) {
		toSerialize["hasVideo"] = o.HasVideo
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	return toSerialize, nil
}

type NullableVideoFile struct {
	value *VideoFile
	isSet bool
}

func (v NullableVideoFile) Get() *VideoFile {
	return v.value
}

func (v *NullableVideoFile) Set(val *VideoFile) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoFile) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoFile(val *VideoFile) *NullableVideoFile {
	return &NullableVideoFile{value: val, isSet: true}
}

func (v NullableVideoFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
