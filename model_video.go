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

// checks if the Video type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Video{}

// Video struct for Video
type Video struct {
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
	Account           *AccountSummary              `json:"account,omitempty"`
	Channel           *VideoChannelSummary         `json:"channel,omitempty"`
	UserHistory       NullableVideoUserHistory     `json:"userHistory,omitempty"`
}

// NewVideo instantiates a new Video object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideo() *Video {
	this := Video{}
	return &this
}

// NewVideoWithDefaults instantiates a new Video object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWithDefaults() *Video {
	this := Video{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Video) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Video) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Video) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Video) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Video) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Video) SetUuid(v string) {
	o.Uuid = &v
}

// GetShortUUID returns the ShortUUID field value if set, zero value otherwise.
func (o *Video) GetShortUUID() string {
	if o == nil || IsNil(o.ShortUUID) {
		var ret string
		return ret
	}
	return *o.ShortUUID
}

// GetShortUUIDOk returns a tuple with the ShortUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetShortUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.ShortUUID) {
		return nil, false
	}
	return o.ShortUUID, true
}

// HasShortUUID returns a boolean if a field has been set.
func (o *Video) HasShortUUID() bool {
	if o != nil && !IsNil(o.ShortUUID) {
		return true
	}

	return false
}

// SetShortUUID gets a reference to the given string and assigns it to the ShortUUID field.
func (o *Video) SetShortUUID(v string) {
	o.ShortUUID = &v
}

// GetIsLive returns the IsLive field value if set, zero value otherwise.
func (o *Video) GetIsLive() bool {
	if o == nil || IsNil(o.IsLive) {
		var ret bool
		return ret
	}
	return *o.IsLive
}

// GetIsLiveOk returns a tuple with the IsLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetIsLiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLive) {
		return nil, false
	}
	return o.IsLive, true
}

// HasIsLive returns a boolean if a field has been set.
func (o *Video) HasIsLive() bool {
	if o != nil && !IsNil(o.IsLive) {
		return true
	}

	return false
}

// SetIsLive gets a reference to the given bool and assigns it to the IsLive field.
func (o *Video) SetIsLive(v bool) {
	o.IsLive = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Video) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Video) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Video) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *Video) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *Video) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *Video) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Video) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Video) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Video) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOriginallyPublishedAt returns the OriginallyPublishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetOriginallyPublishedAt() time.Time {
	if o == nil || IsNil(o.OriginallyPublishedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.OriginallyPublishedAt.Get()
}

// GetOriginallyPublishedAtOk returns a tuple with the OriginallyPublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetOriginallyPublishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginallyPublishedAt.Get(), o.OriginallyPublishedAt.IsSet()
}

// HasOriginallyPublishedAt returns a boolean if a field has been set.
func (o *Video) HasOriginallyPublishedAt() bool {
	if o != nil && o.OriginallyPublishedAt.IsSet() {
		return true
	}

	return false
}

// SetOriginallyPublishedAt gets a reference to the given NullableTime and assigns it to the OriginallyPublishedAt field.
func (o *Video) SetOriginallyPublishedAt(v time.Time) {
	o.OriginallyPublishedAt.Set(&v)
}

// SetOriginallyPublishedAtNil sets the value for OriginallyPublishedAt to be an explicit nil
func (o *Video) SetOriginallyPublishedAtNil() {
	o.OriginallyPublishedAt.Set(nil)
}

// UnsetOriginallyPublishedAt ensures that no value is present for OriginallyPublishedAt, not even an explicit nil
func (o *Video) UnsetOriginallyPublishedAt() {
	o.OriginallyPublishedAt.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Video) GetCategory() VideoConstantNumberCategory {
	if o == nil || IsNil(o.Category) {
		var ret VideoConstantNumberCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetCategoryOk() (*VideoConstantNumberCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Video) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given VideoConstantNumberCategory and assigns it to the Category field.
func (o *Video) SetCategory(v VideoConstantNumberCategory) {
	o.Category = &v
}

// GetLicence returns the Licence field value if set, zero value otherwise.
func (o *Video) GetLicence() VideoConstantNumberLicence {
	if o == nil || IsNil(o.Licence) {
		var ret VideoConstantNumberLicence
		return ret
	}
	return *o.Licence
}

// GetLicenceOk returns a tuple with the Licence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetLicenceOk() (*VideoConstantNumberLicence, bool) {
	if o == nil || IsNil(o.Licence) {
		return nil, false
	}
	return o.Licence, true
}

// HasLicence returns a boolean if a field has been set.
func (o *Video) HasLicence() bool {
	if o != nil && !IsNil(o.Licence) {
		return true
	}

	return false
}

// SetLicence gets a reference to the given VideoConstantNumberLicence and assigns it to the Licence field.
func (o *Video) SetLicence(v VideoConstantNumberLicence) {
	o.Licence = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Video) GetLanguage() VideoConstantStringLanguage {
	if o == nil || IsNil(o.Language) {
		var ret VideoConstantStringLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetLanguageOk() (*VideoConstantStringLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Video) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given VideoConstantStringLanguage and assigns it to the Language field.
func (o *Video) SetLanguage(v VideoConstantStringLanguage) {
	o.Language = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *Video) GetPrivacy() VideoPrivacyConstant {
	if o == nil || IsNil(o.Privacy) {
		var ret VideoPrivacyConstant
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPrivacyOk() (*VideoPrivacyConstant, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *Video) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given VideoPrivacyConstant and assigns it to the Privacy field.
func (o *Video) SetPrivacy(v VideoPrivacyConstant) {
	o.Privacy = &v
}

// GetTruncatedDescription returns the TruncatedDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetTruncatedDescription() string {
	if o == nil || IsNil(o.TruncatedDescription.Get()) {
		var ret string
		return ret
	}
	return *o.TruncatedDescription.Get()
}

// GetTruncatedDescriptionOk returns a tuple with the TruncatedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetTruncatedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TruncatedDescription.Get(), o.TruncatedDescription.IsSet()
}

// HasTruncatedDescription returns a boolean if a field has been set.
func (o *Video) HasTruncatedDescription() bool {
	if o != nil && o.TruncatedDescription.IsSet() {
		return true
	}

	return false
}

// SetTruncatedDescription gets a reference to the given NullableString and assigns it to the TruncatedDescription field.
func (o *Video) SetTruncatedDescription(v string) {
	o.TruncatedDescription.Set(&v)
}

// SetTruncatedDescriptionNil sets the value for TruncatedDescription to be an explicit nil
func (o *Video) SetTruncatedDescriptionNil() {
	o.TruncatedDescription.Set(nil)
}

// UnsetTruncatedDescription ensures that no value is present for TruncatedDescription, not even an explicit nil
func (o *Video) UnsetTruncatedDescription() {
	o.TruncatedDescription.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Video) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Video) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *Video) SetDuration(v int32) {
	o.Duration = &v
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetAspectRatio() float32 {
	if o == nil || IsNil(o.AspectRatio.Get()) {
		var ret float32
		return ret
	}
	return *o.AspectRatio.Get()
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetAspectRatioOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AspectRatio.Get(), o.AspectRatio.IsSet()
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *Video) HasAspectRatio() bool {
	if o != nil && o.AspectRatio.IsSet() {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given NullableFloat32 and assigns it to the AspectRatio field.
func (o *Video) SetAspectRatio(v float32) {
	o.AspectRatio.Set(&v)
}

// SetAspectRatioNil sets the value for AspectRatio to be an explicit nil
func (o *Video) SetAspectRatioNil() {
	o.AspectRatio.Set(nil)
}

// UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
func (o *Video) UnsetAspectRatio() {
	o.AspectRatio.Unset()
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *Video) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *Video) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *Video) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Video) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Video) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Video) SetName(v string) {
	o.Name = &v
}

// GetThumbnailPath returns the ThumbnailPath field value if set, zero value otherwise.
func (o *Video) GetThumbnailPath() string {
	if o == nil || IsNil(o.ThumbnailPath) {
		var ret string
		return ret
	}
	return *o.ThumbnailPath
}

// GetThumbnailPathOk returns a tuple with the ThumbnailPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetThumbnailPathOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailPath) {
		return nil, false
	}
	return o.ThumbnailPath, true
}

// HasThumbnailPath returns a boolean if a field has been set.
func (o *Video) HasThumbnailPath() bool {
	if o != nil && !IsNil(o.ThumbnailPath) {
		return true
	}

	return false
}

// SetThumbnailPath gets a reference to the given string and assigns it to the ThumbnailPath field.
func (o *Video) SetThumbnailPath(v string) {
	o.ThumbnailPath = &v
}

// GetPreviewPath returns the PreviewPath field value if set, zero value otherwise.
func (o *Video) GetPreviewPath() string {
	if o == nil || IsNil(o.PreviewPath) {
		var ret string
		return ret
	}
	return *o.PreviewPath
}

// GetPreviewPathOk returns a tuple with the PreviewPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPreviewPathOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewPath) {
		return nil, false
	}
	return o.PreviewPath, true
}

// HasPreviewPath returns a boolean if a field has been set.
func (o *Video) HasPreviewPath() bool {
	if o != nil && !IsNil(o.PreviewPath) {
		return true
	}

	return false
}

// SetPreviewPath gets a reference to the given string and assigns it to the PreviewPath field.
func (o *Video) SetPreviewPath(v string) {
	o.PreviewPath = &v
}

// GetEmbedPath returns the EmbedPath field value if set, zero value otherwise.
func (o *Video) GetEmbedPath() string {
	if o == nil || IsNil(o.EmbedPath) {
		var ret string
		return ret
	}
	return *o.EmbedPath
}

// GetEmbedPathOk returns a tuple with the EmbedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetEmbedPathOk() (*string, bool) {
	if o == nil || IsNil(o.EmbedPath) {
		return nil, false
	}
	return o.EmbedPath, true
}

// HasEmbedPath returns a boolean if a field has been set.
func (o *Video) HasEmbedPath() bool {
	if o != nil && !IsNil(o.EmbedPath) {
		return true
	}

	return false
}

// SetEmbedPath gets a reference to the given string and assigns it to the EmbedPath field.
func (o *Video) SetEmbedPath(v string) {
	o.EmbedPath = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *Video) GetViews() int32 {
	if o == nil || IsNil(o.Views) {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetViewsOk() (*int32, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *Video) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *Video) SetViews(v int32) {
	o.Views = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise.
func (o *Video) GetLikes() int32 {
	if o == nil || IsNil(o.Likes) {
		var ret int32
		return ret
	}
	return *o.Likes
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Likes) {
		return nil, false
	}
	return o.Likes, true
}

// HasLikes returns a boolean if a field has been set.
func (o *Video) HasLikes() bool {
	if o != nil && !IsNil(o.Likes) {
		return true
	}

	return false
}

// SetLikes gets a reference to the given int32 and assigns it to the Likes field.
func (o *Video) SetLikes(v int32) {
	o.Likes = &v
}

// GetDislikes returns the Dislikes field value if set, zero value otherwise.
func (o *Video) GetDislikes() int32 {
	if o == nil || IsNil(o.Dislikes) {
		var ret int32
		return ret
	}
	return *o.Dislikes
}

// GetDislikesOk returns a tuple with the Dislikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetDislikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Dislikes) {
		return nil, false
	}
	return o.Dislikes, true
}

// HasDislikes returns a boolean if a field has been set.
func (o *Video) HasDislikes() bool {
	if o != nil && !IsNil(o.Dislikes) {
		return true
	}

	return false
}

// SetDislikes gets a reference to the given int32 and assigns it to the Dislikes field.
func (o *Video) SetDislikes(v int32) {
	o.Dislikes = &v
}

// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *Video) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *Video) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *Video) SetNsfw(v bool) {
	o.Nsfw = &v
}

// GetWaitTranscoding returns the WaitTranscoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetWaitTranscoding() bool {
	if o == nil || IsNil(o.WaitTranscoding.Get()) {
		var ret bool
		return ret
	}
	return *o.WaitTranscoding.Get()
}

// GetWaitTranscodingOk returns a tuple with the WaitTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetWaitTranscodingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WaitTranscoding.Get(), o.WaitTranscoding.IsSet()
}

// HasWaitTranscoding returns a boolean if a field has been set.
func (o *Video) HasWaitTranscoding() bool {
	if o != nil && o.WaitTranscoding.IsSet() {
		return true
	}

	return false
}

// SetWaitTranscoding gets a reference to the given NullableBool and assigns it to the WaitTranscoding field.
func (o *Video) SetWaitTranscoding(v bool) {
	o.WaitTranscoding.Set(&v)
}

// SetWaitTranscodingNil sets the value for WaitTranscoding to be an explicit nil
func (o *Video) SetWaitTranscodingNil() {
	o.WaitTranscoding.Set(nil)
}

// UnsetWaitTranscoding ensures that no value is present for WaitTranscoding, not even an explicit nil
func (o *Video) UnsetWaitTranscoding() {
	o.WaitTranscoding.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Video) GetState() VideoStateConstant {
	if o == nil || IsNil(o.State) {
		var ret VideoStateConstant
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetStateOk() (*VideoStateConstant, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Video) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given VideoStateConstant and assigns it to the State field.
func (o *Video) SetState(v VideoStateConstant) {
	o.State = &v
}

// GetScheduledUpdate returns the ScheduledUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetScheduledUpdate() VideoScheduledUpdate {
	if o == nil || IsNil(o.ScheduledUpdate.Get()) {
		var ret VideoScheduledUpdate
		return ret
	}
	return *o.ScheduledUpdate.Get()
}

// GetScheduledUpdateOk returns a tuple with the ScheduledUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetScheduledUpdateOk() (*VideoScheduledUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledUpdate.Get(), o.ScheduledUpdate.IsSet()
}

// HasScheduledUpdate returns a boolean if a field has been set.
func (o *Video) HasScheduledUpdate() bool {
	if o != nil && o.ScheduledUpdate.IsSet() {
		return true
	}

	return false
}

// SetScheduledUpdate gets a reference to the given NullableVideoScheduledUpdate and assigns it to the ScheduledUpdate field.
func (o *Video) SetScheduledUpdate(v VideoScheduledUpdate) {
	o.ScheduledUpdate.Set(&v)
}

// SetScheduledUpdateNil sets the value for ScheduledUpdate to be an explicit nil
func (o *Video) SetScheduledUpdateNil() {
	o.ScheduledUpdate.Set(nil)
}

// UnsetScheduledUpdate ensures that no value is present for ScheduledUpdate, not even an explicit nil
func (o *Video) UnsetScheduledUpdate() {
	o.ScheduledUpdate.Unset()
}

// GetBlacklisted returns the Blacklisted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetBlacklisted() bool {
	if o == nil || IsNil(o.Blacklisted.Get()) {
		var ret bool
		return ret
	}
	return *o.Blacklisted.Get()
}

// GetBlacklistedOk returns a tuple with the Blacklisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetBlacklistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blacklisted.Get(), o.Blacklisted.IsSet()
}

// HasBlacklisted returns a boolean if a field has been set.
func (o *Video) HasBlacklisted() bool {
	if o != nil && o.Blacklisted.IsSet() {
		return true
	}

	return false
}

// SetBlacklisted gets a reference to the given NullableBool and assigns it to the Blacklisted field.
func (o *Video) SetBlacklisted(v bool) {
	o.Blacklisted.Set(&v)
}

// SetBlacklistedNil sets the value for Blacklisted to be an explicit nil
func (o *Video) SetBlacklistedNil() {
	o.Blacklisted.Set(nil)
}

// UnsetBlacklisted ensures that no value is present for Blacklisted, not even an explicit nil
func (o *Video) UnsetBlacklisted() {
	o.Blacklisted.Unset()
}

// GetBlacklistedReason returns the BlacklistedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetBlacklistedReason() string {
	if o == nil || IsNil(o.BlacklistedReason.Get()) {
		var ret string
		return ret
	}
	return *o.BlacklistedReason.Get()
}

// GetBlacklistedReasonOk returns a tuple with the BlacklistedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetBlacklistedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlacklistedReason.Get(), o.BlacklistedReason.IsSet()
}

// HasBlacklistedReason returns a boolean if a field has been set.
func (o *Video) HasBlacklistedReason() bool {
	if o != nil && o.BlacklistedReason.IsSet() {
		return true
	}

	return false
}

// SetBlacklistedReason gets a reference to the given NullableString and assigns it to the BlacklistedReason field.
func (o *Video) SetBlacklistedReason(v string) {
	o.BlacklistedReason.Set(&v)
}

// SetBlacklistedReasonNil sets the value for BlacklistedReason to be an explicit nil
func (o *Video) SetBlacklistedReasonNil() {
	o.BlacklistedReason.Set(nil)
}

// UnsetBlacklistedReason ensures that no value is present for BlacklistedReason, not even an explicit nil
func (o *Video) UnsetBlacklistedReason() {
	o.BlacklistedReason.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Video) GetAccount() AccountSummary {
	if o == nil || IsNil(o.Account) {
		var ret AccountSummary
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetAccountOk() (*AccountSummary, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Video) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountSummary and assigns it to the Account field.
func (o *Video) SetAccount(v AccountSummary) {
	o.Account = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *Video) GetChannel() VideoChannelSummary {
	if o == nil || IsNil(o.Channel) {
		var ret VideoChannelSummary
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetChannelOk() (*VideoChannelSummary, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *Video) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given VideoChannelSummary and assigns it to the Channel field.
func (o *Video) SetChannel(v VideoChannelSummary) {
	o.Channel = &v
}

// GetUserHistory returns the UserHistory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetUserHistory() VideoUserHistory {
	if o == nil || IsNil(o.UserHistory.Get()) {
		var ret VideoUserHistory
		return ret
	}
	return *o.UserHistory.Get()
}

// GetUserHistoryOk returns a tuple with the UserHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetUserHistoryOk() (*VideoUserHistory, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserHistory.Get(), o.UserHistory.IsSet()
}

// HasUserHistory returns a boolean if a field has been set.
func (o *Video) HasUserHistory() bool {
	if o != nil && o.UserHistory.IsSet() {
		return true
	}

	return false
}

// SetUserHistory gets a reference to the given NullableVideoUserHistory and assigns it to the UserHistory field.
func (o *Video) SetUserHistory(v VideoUserHistory) {
	o.UserHistory.Set(&v)
}

// SetUserHistoryNil sets the value for UserHistory to be an explicit nil
func (o *Video) SetUserHistoryNil() {
	o.UserHistory.Set(nil)
}

// UnsetUserHistory ensures that no value is present for UserHistory, not even an explicit nil
func (o *Video) UnsetUserHistory() {
	o.UserHistory.Unset()
}

func (o Video) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Video) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableVideo struct {
	value *Video
	isSet bool
}

func (v NullableVideo) Get() *Video {
	return v.value
}

func (v *NullableVideo) Set(val *Video) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo(val *Video) *NullableVideo {
	return &NullableVideo{value: val, isSet: true}
}

func (v NullableVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
