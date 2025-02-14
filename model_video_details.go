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

// checks if the VideoDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoDetails{}

// VideoDetails struct for VideoDetails
type VideoDetails struct {
	// object id for the video
	Id *int32 `json:"id,omitempty"`
	// universal identifier for the video, that can be used across instances
	Uuid *string `json:"uuid,omitempty" validate:"regexp=^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"`
	// translation of a uuid v4 with a bigger alphabet to have a shorter uuid
	ShortUUID *string `json:"shortUUID,omitempty"`
	IsLive    *bool   `json:"isLive,omitempty"`
	// time at which the video object was first drafted
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// time at which the video was marked as ready for playback (with restrictions depending on `privacy`). Usually set after a `state` evolution.
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	// last time the video's metadata was modified
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// used to represent a date of first publication, prior to the practical publication date of `publishedAt`
	OriginallyPublishedAt NullableTime `json:"originallyPublishedAt,omitempty"`
	// category in which the video is classified
	Category *VideoConstantNumberCategory `json:"category,omitempty"`
	// licence under which the video is distributed
	Licence *VideoConstantNumberLicence `json:"licence,omitempty"`
	// main language used in the video
	Language *VideoConstantStringLanguage `json:"language,omitempty"`
	// privacy policy used to distribute the video
	Privacy *VideoPrivacyConstant `json:"privacy,omitempty"`
	// truncated description of the video, written in Markdown.
	TruncatedDescription NullableString `json:"truncatedDescription,omitempty"`
	// duration of the video in seconds
	Duration *int32 `json:"duration,omitempty"`
	// **PeerTube >= 6.1** Aspect ratio of the video stream
	AspectRatio NullableFloat32 `json:"aspectRatio,omitempty"`
	IsLocal     *bool           `json:"isLocal,omitempty"`
	// title of the video
	Name            *string      `json:"name,omitempty"`
	ThumbnailPath   *string      `json:"thumbnailPath,omitempty"`
	PreviewPath     *string      `json:"previewPath,omitempty"`
	EmbedPath       *string      `json:"embedPath,omitempty"`
	Views           *int32       `json:"views,omitempty"`
	Likes           *int32       `json:"likes,omitempty"`
	Dislikes        *int32       `json:"dislikes,omitempty"`
	Nsfw            *bool        `json:"nsfw,omitempty"`
	WaitTranscoding NullableBool `json:"waitTranscoding,omitempty"`
	// represents the internal state of the video processing within the PeerTube instance
	State             *VideoStateConstant          `json:"state,omitempty"`
	ScheduledUpdate   NullableVideoScheduledUpdate `json:"scheduledUpdate,omitempty"`
	Blacklisted       NullableBool                 `json:"blacklisted,omitempty"`
	BlacklistedReason NullableString               `json:"blacklistedReason,omitempty"`
	Account           *Account                     `json:"account,omitempty"`
	Channel           *VideoChannel                `json:"channel,omitempty"`
	UserHistory       NullableVideoUserHistory     `json:"userHistory,omitempty"`
	// If the video is a live, you have the amount of current viewers
	Viewers *int32 `json:"viewers,omitempty"`
	// full description of the video, written in Markdown.
	Description NullableString `json:"description,omitempty"`
	// A text tell the audience how to support the video creator
	Support NullableString `json:"support,omitempty"`
	Tags    []string       `json:"tags,omitempty"`
	// Deprecated in 6.2, use commentsPolicy instead
	// Deprecated
	CommentsEnabled *bool                        `json:"commentsEnabled,omitempty"`
	CommentsPolicy  *VideoCommentsPolicyConstant `json:"commentsPolicy,omitempty"`
	DownloadEnabled *bool                        `json:"downloadEnabled,omitempty"`
	// Latest input file update. Null if the file has never been replaced since the original upload
	InputFileUpdatedAt NullableTime `json:"inputFileUpdatedAt,omitempty"`
	TrackerUrls        []string     `json:"trackerUrls,omitempty"`
	// Web compatible video files. If Web Video is disabled on the server:  - field will be empty - video files will be found in `streamingPlaylists[].files` field
	Files []VideoFile `json:"files,omitempty"`
	// HLS playlists/manifest files. If HLS is disabled on the server:  - field will be empty - video files will be found in `files` field
	StreamingPlaylists []VideoStreamingPlaylists `json:"streamingPlaylists,omitempty"`
}

// NewVideoDetails instantiates a new VideoDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoDetails() *VideoDetails {
	this := VideoDetails{}
	return &this
}

// NewVideoDetailsWithDefaults instantiates a new VideoDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoDetailsWithDefaults() *VideoDetails {
	this := VideoDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoDetails) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VideoDetails) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VideoDetails) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VideoDetails) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VideoDetails) SetUuid(v string) {
	o.Uuid = &v
}

// GetShortUUID returns the ShortUUID field value if set, zero value otherwise.
func (o *VideoDetails) GetShortUUID() string {
	if o == nil || IsNil(o.ShortUUID) {
		var ret string
		return ret
	}
	return *o.ShortUUID
}

// GetShortUUIDOk returns a tuple with the ShortUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetShortUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.ShortUUID) {
		return nil, false
	}
	return o.ShortUUID, true
}

// HasShortUUID returns a boolean if a field has been set.
func (o *VideoDetails) HasShortUUID() bool {
	if o != nil && !IsNil(o.ShortUUID) {
		return true
	}

	return false
}

// SetShortUUID gets a reference to the given string and assigns it to the ShortUUID field.
func (o *VideoDetails) SetShortUUID(v string) {
	o.ShortUUID = &v
}

// GetIsLive returns the IsLive field value if set, zero value otherwise.
func (o *VideoDetails) GetIsLive() bool {
	if o == nil || IsNil(o.IsLive) {
		var ret bool
		return ret
	}
	return *o.IsLive
}

// GetIsLiveOk returns a tuple with the IsLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetIsLiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLive) {
		return nil, false
	}
	return o.IsLive, true
}

// HasIsLive returns a boolean if a field has been set.
func (o *VideoDetails) HasIsLive() bool {
	if o != nil && !IsNil(o.IsLive) {
		return true
	}

	return false
}

// SetIsLive gets a reference to the given bool and assigns it to the IsLive field.
func (o *VideoDetails) SetIsLive(v bool) {
	o.IsLive = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VideoDetails) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VideoDetails) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VideoDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *VideoDetails) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *VideoDetails) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *VideoDetails) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VideoDetails) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VideoDetails) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VideoDetails) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOriginallyPublishedAt returns the OriginallyPublishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetOriginallyPublishedAt() time.Time {
	if o == nil || IsNil(o.OriginallyPublishedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.OriginallyPublishedAt.Get()
}

// GetOriginallyPublishedAtOk returns a tuple with the OriginallyPublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetOriginallyPublishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginallyPublishedAt.Get(), o.OriginallyPublishedAt.IsSet()
}

// HasOriginallyPublishedAt returns a boolean if a field has been set.
func (o *VideoDetails) HasOriginallyPublishedAt() bool {
	if o != nil && o.OriginallyPublishedAt.IsSet() {
		return true
	}

	return false
}

// SetOriginallyPublishedAt gets a reference to the given NullableTime and assigns it to the OriginallyPublishedAt field.
func (o *VideoDetails) SetOriginallyPublishedAt(v time.Time) {
	o.OriginallyPublishedAt.Set(&v)
}

// SetOriginallyPublishedAtNil sets the value for OriginallyPublishedAt to be an explicit nil
func (o *VideoDetails) SetOriginallyPublishedAtNil() {
	o.OriginallyPublishedAt.Set(nil)
}

// UnsetOriginallyPublishedAt ensures that no value is present for OriginallyPublishedAt, not even an explicit nil
func (o *VideoDetails) UnsetOriginallyPublishedAt() {
	o.OriginallyPublishedAt.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *VideoDetails) GetCategory() VideoConstantNumberCategory {
	if o == nil || IsNil(o.Category) {
		var ret VideoConstantNumberCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetCategoryOk() (*VideoConstantNumberCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *VideoDetails) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given VideoConstantNumberCategory and assigns it to the Category field.
func (o *VideoDetails) SetCategory(v VideoConstantNumberCategory) {
	o.Category = &v
}

// GetLicence returns the Licence field value if set, zero value otherwise.
func (o *VideoDetails) GetLicence() VideoConstantNumberLicence {
	if o == nil || IsNil(o.Licence) {
		var ret VideoConstantNumberLicence
		return ret
	}
	return *o.Licence
}

// GetLicenceOk returns a tuple with the Licence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetLicenceOk() (*VideoConstantNumberLicence, bool) {
	if o == nil || IsNil(o.Licence) {
		return nil, false
	}
	return o.Licence, true
}

// HasLicence returns a boolean if a field has been set.
func (o *VideoDetails) HasLicence() bool {
	if o != nil && !IsNil(o.Licence) {
		return true
	}

	return false
}

// SetLicence gets a reference to the given VideoConstantNumberLicence and assigns it to the Licence field.
func (o *VideoDetails) SetLicence(v VideoConstantNumberLicence) {
	o.Licence = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *VideoDetails) GetLanguage() VideoConstantStringLanguage {
	if o == nil || IsNil(o.Language) {
		var ret VideoConstantStringLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetLanguageOk() (*VideoConstantStringLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *VideoDetails) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given VideoConstantStringLanguage and assigns it to the Language field.
func (o *VideoDetails) SetLanguage(v VideoConstantStringLanguage) {
	o.Language = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *VideoDetails) GetPrivacy() VideoPrivacyConstant {
	if o == nil || IsNil(o.Privacy) {
		var ret VideoPrivacyConstant
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetPrivacyOk() (*VideoPrivacyConstant, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *VideoDetails) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given VideoPrivacyConstant and assigns it to the Privacy field.
func (o *VideoDetails) SetPrivacy(v VideoPrivacyConstant) {
	o.Privacy = &v
}

// GetTruncatedDescription returns the TruncatedDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetTruncatedDescription() string {
	if o == nil || IsNil(o.TruncatedDescription.Get()) {
		var ret string
		return ret
	}
	return *o.TruncatedDescription.Get()
}

// GetTruncatedDescriptionOk returns a tuple with the TruncatedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetTruncatedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TruncatedDescription.Get(), o.TruncatedDescription.IsSet()
}

// HasTruncatedDescription returns a boolean if a field has been set.
func (o *VideoDetails) HasTruncatedDescription() bool {
	if o != nil && o.TruncatedDescription.IsSet() {
		return true
	}

	return false
}

// SetTruncatedDescription gets a reference to the given NullableString and assigns it to the TruncatedDescription field.
func (o *VideoDetails) SetTruncatedDescription(v string) {
	o.TruncatedDescription.Set(&v)
}

// SetTruncatedDescriptionNil sets the value for TruncatedDescription to be an explicit nil
func (o *VideoDetails) SetTruncatedDescriptionNil() {
	o.TruncatedDescription.Set(nil)
}

// UnsetTruncatedDescription ensures that no value is present for TruncatedDescription, not even an explicit nil
func (o *VideoDetails) UnsetTruncatedDescription() {
	o.TruncatedDescription.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VideoDetails) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VideoDetails) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *VideoDetails) SetDuration(v int32) {
	o.Duration = &v
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetAspectRatio() float32 {
	if o == nil || IsNil(o.AspectRatio.Get()) {
		var ret float32
		return ret
	}
	return *o.AspectRatio.Get()
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetAspectRatioOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AspectRatio.Get(), o.AspectRatio.IsSet()
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *VideoDetails) HasAspectRatio() bool {
	if o != nil && o.AspectRatio.IsSet() {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given NullableFloat32 and assigns it to the AspectRatio field.
func (o *VideoDetails) SetAspectRatio(v float32) {
	o.AspectRatio.Set(&v)
}

// SetAspectRatioNil sets the value for AspectRatio to be an explicit nil
func (o *VideoDetails) SetAspectRatioNil() {
	o.AspectRatio.Set(nil)
}

// UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
func (o *VideoDetails) UnsetAspectRatio() {
	o.AspectRatio.Unset()
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *VideoDetails) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *VideoDetails) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *VideoDetails) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VideoDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VideoDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VideoDetails) SetName(v string) {
	o.Name = &v
}

// GetThumbnailPath returns the ThumbnailPath field value if set, zero value otherwise.
func (o *VideoDetails) GetThumbnailPath() string {
	if o == nil || IsNil(o.ThumbnailPath) {
		var ret string
		return ret
	}
	return *o.ThumbnailPath
}

// GetThumbnailPathOk returns a tuple with the ThumbnailPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetThumbnailPathOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailPath) {
		return nil, false
	}
	return o.ThumbnailPath, true
}

// HasThumbnailPath returns a boolean if a field has been set.
func (o *VideoDetails) HasThumbnailPath() bool {
	if o != nil && !IsNil(o.ThumbnailPath) {
		return true
	}

	return false
}

// SetThumbnailPath gets a reference to the given string and assigns it to the ThumbnailPath field.
func (o *VideoDetails) SetThumbnailPath(v string) {
	o.ThumbnailPath = &v
}

// GetPreviewPath returns the PreviewPath field value if set, zero value otherwise.
func (o *VideoDetails) GetPreviewPath() string {
	if o == nil || IsNil(o.PreviewPath) {
		var ret string
		return ret
	}
	return *o.PreviewPath
}

// GetPreviewPathOk returns a tuple with the PreviewPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetPreviewPathOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewPath) {
		return nil, false
	}
	return o.PreviewPath, true
}

// HasPreviewPath returns a boolean if a field has been set.
func (o *VideoDetails) HasPreviewPath() bool {
	if o != nil && !IsNil(o.PreviewPath) {
		return true
	}

	return false
}

// SetPreviewPath gets a reference to the given string and assigns it to the PreviewPath field.
func (o *VideoDetails) SetPreviewPath(v string) {
	o.PreviewPath = &v
}

// GetEmbedPath returns the EmbedPath field value if set, zero value otherwise.
func (o *VideoDetails) GetEmbedPath() string {
	if o == nil || IsNil(o.EmbedPath) {
		var ret string
		return ret
	}
	return *o.EmbedPath
}

// GetEmbedPathOk returns a tuple with the EmbedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetEmbedPathOk() (*string, bool) {
	if o == nil || IsNil(o.EmbedPath) {
		return nil, false
	}
	return o.EmbedPath, true
}

// HasEmbedPath returns a boolean if a field has been set.
func (o *VideoDetails) HasEmbedPath() bool {
	if o != nil && !IsNil(o.EmbedPath) {
		return true
	}

	return false
}

// SetEmbedPath gets a reference to the given string and assigns it to the EmbedPath field.
func (o *VideoDetails) SetEmbedPath(v string) {
	o.EmbedPath = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *VideoDetails) GetViews() int32 {
	if o == nil || IsNil(o.Views) {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetViewsOk() (*int32, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *VideoDetails) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *VideoDetails) SetViews(v int32) {
	o.Views = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise.
func (o *VideoDetails) GetLikes() int32 {
	if o == nil || IsNil(o.Likes) {
		var ret int32
		return ret
	}
	return *o.Likes
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Likes) {
		return nil, false
	}
	return o.Likes, true
}

// HasLikes returns a boolean if a field has been set.
func (o *VideoDetails) HasLikes() bool {
	if o != nil && !IsNil(o.Likes) {
		return true
	}

	return false
}

// SetLikes gets a reference to the given int32 and assigns it to the Likes field.
func (o *VideoDetails) SetLikes(v int32) {
	o.Likes = &v
}

// GetDislikes returns the Dislikes field value if set, zero value otherwise.
func (o *VideoDetails) GetDislikes() int32 {
	if o == nil || IsNil(o.Dislikes) {
		var ret int32
		return ret
	}
	return *o.Dislikes
}

// GetDislikesOk returns a tuple with the Dislikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetDislikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Dislikes) {
		return nil, false
	}
	return o.Dislikes, true
}

// HasDislikes returns a boolean if a field has been set.
func (o *VideoDetails) HasDislikes() bool {
	if o != nil && !IsNil(o.Dislikes) {
		return true
	}

	return false
}

// SetDislikes gets a reference to the given int32 and assigns it to the Dislikes field.
func (o *VideoDetails) SetDislikes(v int32) {
	o.Dislikes = &v
}

// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *VideoDetails) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *VideoDetails) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *VideoDetails) SetNsfw(v bool) {
	o.Nsfw = &v
}

// GetWaitTranscoding returns the WaitTranscoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetWaitTranscoding() bool {
	if o == nil || IsNil(o.WaitTranscoding.Get()) {
		var ret bool
		return ret
	}
	return *o.WaitTranscoding.Get()
}

// GetWaitTranscodingOk returns a tuple with the WaitTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetWaitTranscodingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WaitTranscoding.Get(), o.WaitTranscoding.IsSet()
}

// HasWaitTranscoding returns a boolean if a field has been set.
func (o *VideoDetails) HasWaitTranscoding() bool {
	if o != nil && o.WaitTranscoding.IsSet() {
		return true
	}

	return false
}

// SetWaitTranscoding gets a reference to the given NullableBool and assigns it to the WaitTranscoding field.
func (o *VideoDetails) SetWaitTranscoding(v bool) {
	o.WaitTranscoding.Set(&v)
}

// SetWaitTranscodingNil sets the value for WaitTranscoding to be an explicit nil
func (o *VideoDetails) SetWaitTranscodingNil() {
	o.WaitTranscoding.Set(nil)
}

// UnsetWaitTranscoding ensures that no value is present for WaitTranscoding, not even an explicit nil
func (o *VideoDetails) UnsetWaitTranscoding() {
	o.WaitTranscoding.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VideoDetails) GetState() VideoStateConstant {
	if o == nil || IsNil(o.State) {
		var ret VideoStateConstant
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetStateOk() (*VideoStateConstant, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VideoDetails) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given VideoStateConstant and assigns it to the State field.
func (o *VideoDetails) SetState(v VideoStateConstant) {
	o.State = &v
}

// GetScheduledUpdate returns the ScheduledUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetScheduledUpdate() VideoScheduledUpdate {
	if o == nil || IsNil(o.ScheduledUpdate.Get()) {
		var ret VideoScheduledUpdate
		return ret
	}
	return *o.ScheduledUpdate.Get()
}

// GetScheduledUpdateOk returns a tuple with the ScheduledUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetScheduledUpdateOk() (*VideoScheduledUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledUpdate.Get(), o.ScheduledUpdate.IsSet()
}

// HasScheduledUpdate returns a boolean if a field has been set.
func (o *VideoDetails) HasScheduledUpdate() bool {
	if o != nil && o.ScheduledUpdate.IsSet() {
		return true
	}

	return false
}

// SetScheduledUpdate gets a reference to the given NullableVideoScheduledUpdate and assigns it to the ScheduledUpdate field.
func (o *VideoDetails) SetScheduledUpdate(v VideoScheduledUpdate) {
	o.ScheduledUpdate.Set(&v)
}

// SetScheduledUpdateNil sets the value for ScheduledUpdate to be an explicit nil
func (o *VideoDetails) SetScheduledUpdateNil() {
	o.ScheduledUpdate.Set(nil)
}

// UnsetScheduledUpdate ensures that no value is present for ScheduledUpdate, not even an explicit nil
func (o *VideoDetails) UnsetScheduledUpdate() {
	o.ScheduledUpdate.Unset()
}

// GetBlacklisted returns the Blacklisted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetBlacklisted() bool {
	if o == nil || IsNil(o.Blacklisted.Get()) {
		var ret bool
		return ret
	}
	return *o.Blacklisted.Get()
}

// GetBlacklistedOk returns a tuple with the Blacklisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetBlacklistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blacklisted.Get(), o.Blacklisted.IsSet()
}

// HasBlacklisted returns a boolean if a field has been set.
func (o *VideoDetails) HasBlacklisted() bool {
	if o != nil && o.Blacklisted.IsSet() {
		return true
	}

	return false
}

// SetBlacklisted gets a reference to the given NullableBool and assigns it to the Blacklisted field.
func (o *VideoDetails) SetBlacklisted(v bool) {
	o.Blacklisted.Set(&v)
}

// SetBlacklistedNil sets the value for Blacklisted to be an explicit nil
func (o *VideoDetails) SetBlacklistedNil() {
	o.Blacklisted.Set(nil)
}

// UnsetBlacklisted ensures that no value is present for Blacklisted, not even an explicit nil
func (o *VideoDetails) UnsetBlacklisted() {
	o.Blacklisted.Unset()
}

// GetBlacklistedReason returns the BlacklistedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetBlacklistedReason() string {
	if o == nil || IsNil(o.BlacklistedReason.Get()) {
		var ret string
		return ret
	}
	return *o.BlacklistedReason.Get()
}

// GetBlacklistedReasonOk returns a tuple with the BlacklistedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetBlacklistedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlacklistedReason.Get(), o.BlacklistedReason.IsSet()
}

// HasBlacklistedReason returns a boolean if a field has been set.
func (o *VideoDetails) HasBlacklistedReason() bool {
	if o != nil && o.BlacklistedReason.IsSet() {
		return true
	}

	return false
}

// SetBlacklistedReason gets a reference to the given NullableString and assigns it to the BlacklistedReason field.
func (o *VideoDetails) SetBlacklistedReason(v string) {
	o.BlacklistedReason.Set(&v)
}

// SetBlacklistedReasonNil sets the value for BlacklistedReason to be an explicit nil
func (o *VideoDetails) SetBlacklistedReasonNil() {
	o.BlacklistedReason.Set(nil)
}

// UnsetBlacklistedReason ensures that no value is present for BlacklistedReason, not even an explicit nil
func (o *VideoDetails) UnsetBlacklistedReason() {
	o.BlacklistedReason.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *VideoDetails) GetAccount() Account {
	if o == nil || IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetAccountOk() (*Account, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *VideoDetails) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *VideoDetails) SetAccount(v Account) {
	o.Account = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *VideoDetails) GetChannel() VideoChannel {
	if o == nil || IsNil(o.Channel) {
		var ret VideoChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetChannelOk() (*VideoChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *VideoDetails) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given VideoChannel and assigns it to the Channel field.
func (o *VideoDetails) SetChannel(v VideoChannel) {
	o.Channel = &v
}

// GetUserHistory returns the UserHistory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetUserHistory() VideoUserHistory {
	if o == nil || IsNil(o.UserHistory.Get()) {
		var ret VideoUserHistory
		return ret
	}
	return *o.UserHistory.Get()
}

// GetUserHistoryOk returns a tuple with the UserHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetUserHistoryOk() (*VideoUserHistory, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserHistory.Get(), o.UserHistory.IsSet()
}

// HasUserHistory returns a boolean if a field has been set.
func (o *VideoDetails) HasUserHistory() bool {
	if o != nil && o.UserHistory.IsSet() {
		return true
	}

	return false
}

// SetUserHistory gets a reference to the given NullableVideoUserHistory and assigns it to the UserHistory field.
func (o *VideoDetails) SetUserHistory(v VideoUserHistory) {
	o.UserHistory.Set(&v)
}

// SetUserHistoryNil sets the value for UserHistory to be an explicit nil
func (o *VideoDetails) SetUserHistoryNil() {
	o.UserHistory.Set(nil)
}

// UnsetUserHistory ensures that no value is present for UserHistory, not even an explicit nil
func (o *VideoDetails) UnsetUserHistory() {
	o.UserHistory.Unset()
}

// GetViewers returns the Viewers field value if set, zero value otherwise.
func (o *VideoDetails) GetViewers() int32 {
	if o == nil || IsNil(o.Viewers) {
		var ret int32
		return ret
	}
	return *o.Viewers
}

// GetViewersOk returns a tuple with the Viewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetViewersOk() (*int32, bool) {
	if o == nil || IsNil(o.Viewers) {
		return nil, false
	}
	return o.Viewers, true
}

// HasViewers returns a boolean if a field has been set.
func (o *VideoDetails) HasViewers() bool {
	if o != nil && !IsNil(o.Viewers) {
		return true
	}

	return false
}

// SetViewers gets a reference to the given int32 and assigns it to the Viewers field.
func (o *VideoDetails) SetViewers(v int32) {
	o.Viewers = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VideoDetails) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VideoDetails) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VideoDetails) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VideoDetails) UnsetDescription() {
	o.Description.Unset()
}

// GetSupport returns the Support field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetSupport() string {
	if o == nil || IsNil(o.Support.Get()) {
		var ret string
		return ret
	}
	return *o.Support.Get()
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Support.Get(), o.Support.IsSet()
}

// HasSupport returns a boolean if a field has been set.
func (o *VideoDetails) HasSupport() bool {
	if o != nil && o.Support.IsSet() {
		return true
	}

	return false
}

// SetSupport gets a reference to the given NullableString and assigns it to the Support field.
func (o *VideoDetails) SetSupport(v string) {
	o.Support.Set(&v)
}

// SetSupportNil sets the value for Support to be an explicit nil
func (o *VideoDetails) SetSupportNil() {
	o.Support.Set(nil)
}

// UnsetSupport ensures that no value is present for Support, not even an explicit nil
func (o *VideoDetails) UnsetSupport() {
	o.Support.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VideoDetails) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VideoDetails) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VideoDetails) SetTags(v []string) {
	o.Tags = v
}

// GetCommentsEnabled returns the CommentsEnabled field value if set, zero value otherwise.
// Deprecated
func (o *VideoDetails) GetCommentsEnabled() bool {
	if o == nil || IsNil(o.CommentsEnabled) {
		var ret bool
		return ret
	}
	return *o.CommentsEnabled
}

// GetCommentsEnabledOk returns a tuple with the CommentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VideoDetails) GetCommentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CommentsEnabled) {
		return nil, false
	}
	return o.CommentsEnabled, true
}

// HasCommentsEnabled returns a boolean if a field has been set.
func (o *VideoDetails) HasCommentsEnabled() bool {
	if o != nil && !IsNil(o.CommentsEnabled) {
		return true
	}

	return false
}

// SetCommentsEnabled gets a reference to the given bool and assigns it to the CommentsEnabled field.
// Deprecated
func (o *VideoDetails) SetCommentsEnabled(v bool) {
	o.CommentsEnabled = &v
}

// GetCommentsPolicy returns the CommentsPolicy field value if set, zero value otherwise.
func (o *VideoDetails) GetCommentsPolicy() VideoCommentsPolicyConstant {
	if o == nil || IsNil(o.CommentsPolicy) {
		var ret VideoCommentsPolicyConstant
		return ret
	}
	return *o.CommentsPolicy
}

// GetCommentsPolicyOk returns a tuple with the CommentsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetCommentsPolicyOk() (*VideoCommentsPolicyConstant, bool) {
	if o == nil || IsNil(o.CommentsPolicy) {
		return nil, false
	}
	return o.CommentsPolicy, true
}

// HasCommentsPolicy returns a boolean if a field has been set.
func (o *VideoDetails) HasCommentsPolicy() bool {
	if o != nil && !IsNil(o.CommentsPolicy) {
		return true
	}

	return false
}

// SetCommentsPolicy gets a reference to the given VideoCommentsPolicyConstant and assigns it to the CommentsPolicy field.
func (o *VideoDetails) SetCommentsPolicy(v VideoCommentsPolicyConstant) {
	o.CommentsPolicy = &v
}

// GetDownloadEnabled returns the DownloadEnabled field value if set, zero value otherwise.
func (o *VideoDetails) GetDownloadEnabled() bool {
	if o == nil || IsNil(o.DownloadEnabled) {
		var ret bool
		return ret
	}
	return *o.DownloadEnabled
}

// GetDownloadEnabledOk returns a tuple with the DownloadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetDownloadEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadEnabled) {
		return nil, false
	}
	return o.DownloadEnabled, true
}

// HasDownloadEnabled returns a boolean if a field has been set.
func (o *VideoDetails) HasDownloadEnabled() bool {
	if o != nil && !IsNil(o.DownloadEnabled) {
		return true
	}

	return false
}

// SetDownloadEnabled gets a reference to the given bool and assigns it to the DownloadEnabled field.
func (o *VideoDetails) SetDownloadEnabled(v bool) {
	o.DownloadEnabled = &v
}

// GetInputFileUpdatedAt returns the InputFileUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoDetails) GetInputFileUpdatedAt() time.Time {
	if o == nil || IsNil(o.InputFileUpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.InputFileUpdatedAt.Get()
}

// GetInputFileUpdatedAtOk returns a tuple with the InputFileUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoDetails) GetInputFileUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputFileUpdatedAt.Get(), o.InputFileUpdatedAt.IsSet()
}

// HasInputFileUpdatedAt returns a boolean if a field has been set.
func (o *VideoDetails) HasInputFileUpdatedAt() bool {
	if o != nil && o.InputFileUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetInputFileUpdatedAt gets a reference to the given NullableTime and assigns it to the InputFileUpdatedAt field.
func (o *VideoDetails) SetInputFileUpdatedAt(v time.Time) {
	o.InputFileUpdatedAt.Set(&v)
}

// SetInputFileUpdatedAtNil sets the value for InputFileUpdatedAt to be an explicit nil
func (o *VideoDetails) SetInputFileUpdatedAtNil() {
	o.InputFileUpdatedAt.Set(nil)
}

// UnsetInputFileUpdatedAt ensures that no value is present for InputFileUpdatedAt, not even an explicit nil
func (o *VideoDetails) UnsetInputFileUpdatedAt() {
	o.InputFileUpdatedAt.Unset()
}

// GetTrackerUrls returns the TrackerUrls field value if set, zero value otherwise.
func (o *VideoDetails) GetTrackerUrls() []string {
	if o == nil || IsNil(o.TrackerUrls) {
		var ret []string
		return ret
	}
	return o.TrackerUrls
}

// GetTrackerUrlsOk returns a tuple with the TrackerUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetTrackerUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrackerUrls) {
		return nil, false
	}
	return o.TrackerUrls, true
}

// HasTrackerUrls returns a boolean if a field has been set.
func (o *VideoDetails) HasTrackerUrls() bool {
	if o != nil && !IsNil(o.TrackerUrls) {
		return true
	}

	return false
}

// SetTrackerUrls gets a reference to the given []string and assigns it to the TrackerUrls field.
func (o *VideoDetails) SetTrackerUrls(v []string) {
	o.TrackerUrls = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *VideoDetails) GetFiles() []VideoFile {
	if o == nil || IsNil(o.Files) {
		var ret []VideoFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetFilesOk() ([]VideoFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *VideoDetails) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []VideoFile and assigns it to the Files field.
func (o *VideoDetails) SetFiles(v []VideoFile) {
	o.Files = v
}

// GetStreamingPlaylists returns the StreamingPlaylists field value if set, zero value otherwise.
func (o *VideoDetails) GetStreamingPlaylists() []VideoStreamingPlaylists {
	if o == nil || IsNil(o.StreamingPlaylists) {
		var ret []VideoStreamingPlaylists
		return ret
	}
	return o.StreamingPlaylists
}

// GetStreamingPlaylistsOk returns a tuple with the StreamingPlaylists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoDetails) GetStreamingPlaylistsOk() ([]VideoStreamingPlaylists, bool) {
	if o == nil || IsNil(o.StreamingPlaylists) {
		return nil, false
	}
	return o.StreamingPlaylists, true
}

// HasStreamingPlaylists returns a boolean if a field has been set.
func (o *VideoDetails) HasStreamingPlaylists() bool {
	if o != nil && !IsNil(o.StreamingPlaylists) {
		return true
	}

	return false
}

// SetStreamingPlaylists gets a reference to the given []VideoStreamingPlaylists and assigns it to the StreamingPlaylists field.
func (o *VideoDetails) SetStreamingPlaylists(v []VideoStreamingPlaylists) {
	o.StreamingPlaylists = v
}

func (o VideoDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsLive) {
		toSerialize["isLive"] = o.IsLive
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["publishedAt"] = o.PublishedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.OriginallyPublishedAt.IsSet() {
		toSerialize["originallyPublishedAt"] = o.OriginallyPublishedAt.Get()
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Licence) {
		toSerialize["licence"] = o.Licence
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if o.TruncatedDescription.IsSet() {
		toSerialize["truncatedDescription"] = o.TruncatedDescription.Get()
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if o.AspectRatio.IsSet() {
		toSerialize["aspectRatio"] = o.AspectRatio.Get()
	}
	if !IsNil(o.IsLocal) {
		toSerialize["isLocal"] = o.IsLocal
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ThumbnailPath) {
		toSerialize["thumbnailPath"] = o.ThumbnailPath
	}
	if !IsNil(o.PreviewPath) {
		toSerialize["previewPath"] = o.PreviewPath
	}
	if !IsNil(o.EmbedPath) {
		toSerialize["embedPath"] = o.EmbedPath
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
	if o.WaitTranscoding.IsSet() {
		toSerialize["waitTranscoding"] = o.WaitTranscoding.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.ScheduledUpdate.IsSet() {
		toSerialize["scheduledUpdate"] = o.ScheduledUpdate.Get()
	}
	if o.Blacklisted.IsSet() {
		toSerialize["blacklisted"] = o.Blacklisted.Get()
	}
	if o.BlacklistedReason.IsSet() {
		toSerialize["blacklistedReason"] = o.BlacklistedReason.Get()
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if o.UserHistory.IsSet() {
		toSerialize["userHistory"] = o.UserHistory.Get()
	}
	if !IsNil(o.Viewers) {
		toSerialize["viewers"] = o.Viewers
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Support.IsSet() {
		toSerialize["support"] = o.Support.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CommentsEnabled) {
		toSerialize["commentsEnabled"] = o.CommentsEnabled
	}
	if !IsNil(o.CommentsPolicy) {
		toSerialize["commentsPolicy"] = o.CommentsPolicy
	}
	if !IsNil(o.DownloadEnabled) {
		toSerialize["downloadEnabled"] = o.DownloadEnabled
	}
	if o.InputFileUpdatedAt.IsSet() {
		toSerialize["inputFileUpdatedAt"] = o.InputFileUpdatedAt.Get()
	}
	if !IsNil(o.TrackerUrls) {
		toSerialize["trackerUrls"] = o.TrackerUrls
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.StreamingPlaylists) {
		toSerialize["streamingPlaylists"] = o.StreamingPlaylists
	}
	return toSerialize, nil
}

type NullableVideoDetails struct {
	value *VideoDetails
	isSet bool
}

func (v NullableVideoDetails) Get() *VideoDetails {
	return v.value
}

func (v *NullableVideoDetails) Set(val *VideoDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoDetails(val *VideoDetails) *NullableVideoDetails {
	return &NullableVideoDetails{value: val, isSet: true}
}

func (v NullableVideoDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
