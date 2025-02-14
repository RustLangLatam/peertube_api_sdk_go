/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// checks if the VideoUploadRequestCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoUploadRequestCommon{}

// VideoUploadRequestCommon struct for VideoUploadRequestCommon
type VideoUploadRequestCommon struct {
	// Video name
	Name string `json:"name"`
	// Channel id that will contain this video
	ChannelId int32            `json:"channelId"`
	Privacy   *VideoPrivacySet `json:"privacy,omitempty"`
	// category id of the video (see [/videos/categories](#operation/getCategories))
	Category *int32 `json:"category,omitempty"`
	// licence id of the video (see [/videos/licences](#operation/getLicences))
	Licence *int32 `json:"licence,omitempty"`
	// language id of the video (see [/videos/languages](#operation/getLanguages))
	Language *string `json:"language,omitempty"`
	// Video description
	Description *string `json:"description,omitempty"`
	// Whether or not we wait transcoding before publish the video
	WaitTranscoding *bool `json:"waitTranscoding,omitempty"`
	// **PeerTube >= 6.2** If enabled by the admin, automatically generate a subtitle of the video
	GenerateTranscription *bool `json:"generateTranscription,omitempty"`
	// A text tell the audience how to support the video creator
	Support *string `json:"support,omitempty"`
	// Whether or not this video contains sensitive content
	Nsfw *bool `json:"nsfw,omitempty"`
	// Video tags (maximum 5 tags each between 2 and 30 characters)
	Tags []string `json:"tags,omitempty"`
	// Deprecated in 6.2, use commentsPolicy instead
	// Deprecated
	CommentsEnabled *bool                   `json:"commentsEnabled,omitempty"`
	CommentsPolicy  *VideoCommentsPolicySet `json:"commentsPolicy,omitempty"`
	// Enable or disable downloading for this video
	DownloadEnabled *bool `json:"downloadEnabled,omitempty"`
	// Date when the content was originally published
	OriginallyPublishedAt *time.Time            `json:"originallyPublishedAt,omitempty"`
	ScheduleUpdate        *VideoScheduledUpdate `json:"scheduleUpdate,omitempty"`
	// Video thumbnail file
	Thumbnailfile **os.File `json:"thumbnailfile,omitempty"`
	// Video preview file
	Previewfile    **os.File `json:"previewfile,omitempty"`
	VideoPasswords []string  `json:"videoPasswords,omitempty"`
}

type _VideoUploadRequestCommon VideoUploadRequestCommon

// NewVideoUploadRequestCommon instantiates a new VideoUploadRequestCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoUploadRequestCommon(name string, channelId int32) *VideoUploadRequestCommon {
	this := VideoUploadRequestCommon{}
	this.Name = name
	this.ChannelId = channelId
	return &this
}

// NewVideoUploadRequestCommonWithDefaults instantiates a new VideoUploadRequestCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoUploadRequestCommonWithDefaults() *VideoUploadRequestCommon {
	this := VideoUploadRequestCommon{}
	return &this
}

// GetName returns the Name field value
func (o *VideoUploadRequestCommon) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VideoUploadRequestCommon) SetName(v string) {
	o.Name = v
}

// GetChannelId returns the ChannelId field value
func (o *VideoUploadRequestCommon) GetChannelId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetChannelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *VideoUploadRequestCommon) SetChannelId(v int32) {
	o.ChannelId = v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetPrivacy() VideoPrivacySet {
	if o == nil || IsNil(o.Privacy) {
		var ret VideoPrivacySet
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetPrivacyOk() (*VideoPrivacySet, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given VideoPrivacySet and assigns it to the Privacy field.
func (o *VideoUploadRequestCommon) SetPrivacy(v VideoPrivacySet) {
	o.Privacy = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetCategory() int32 {
	if o == nil || IsNil(o.Category) {
		var ret int32
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetCategoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given int32 and assigns it to the Category field.
func (o *VideoUploadRequestCommon) SetCategory(v int32) {
	o.Category = &v
}

// GetLicence returns the Licence field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetLicence() int32 {
	if o == nil || IsNil(o.Licence) {
		var ret int32
		return ret
	}
	return *o.Licence
}

// GetLicenceOk returns a tuple with the Licence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetLicenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Licence) {
		return nil, false
	}
	return o.Licence, true
}

// HasLicence returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasLicence() bool {
	if o != nil && !IsNil(o.Licence) {
		return true
	}

	return false
}

// SetLicence gets a reference to the given int32 and assigns it to the Licence field.
func (o *VideoUploadRequestCommon) SetLicence(v int32) {
	o.Licence = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *VideoUploadRequestCommon) SetLanguage(v string) {
	o.Language = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VideoUploadRequestCommon) SetDescription(v string) {
	o.Description = &v
}

// GetWaitTranscoding returns the WaitTranscoding field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetWaitTranscoding() bool {
	if o == nil || IsNil(o.WaitTranscoding) {
		var ret bool
		return ret
	}
	return *o.WaitTranscoding
}

// GetWaitTranscodingOk returns a tuple with the WaitTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetWaitTranscodingOk() (*bool, bool) {
	if o == nil || IsNil(o.WaitTranscoding) {
		return nil, false
	}
	return o.WaitTranscoding, true
}

// HasWaitTranscoding returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasWaitTranscoding() bool {
	if o != nil && !IsNil(o.WaitTranscoding) {
		return true
	}

	return false
}

// SetWaitTranscoding gets a reference to the given bool and assigns it to the WaitTranscoding field.
func (o *VideoUploadRequestCommon) SetWaitTranscoding(v bool) {
	o.WaitTranscoding = &v
}

// GetGenerateTranscription returns the GenerateTranscription field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetGenerateTranscription() bool {
	if o == nil || IsNil(o.GenerateTranscription) {
		var ret bool
		return ret
	}
	return *o.GenerateTranscription
}

// GetGenerateTranscriptionOk returns a tuple with the GenerateTranscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetGenerateTranscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateTranscription) {
		return nil, false
	}
	return o.GenerateTranscription, true
}

// HasGenerateTranscription returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasGenerateTranscription() bool {
	if o != nil && !IsNil(o.GenerateTranscription) {
		return true
	}

	return false
}

// SetGenerateTranscription gets a reference to the given bool and assigns it to the GenerateTranscription field.
func (o *VideoUploadRequestCommon) SetGenerateTranscription(v bool) {
	o.GenerateTranscription = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetSupport() string {
	if o == nil || IsNil(o.Support) {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetSupportOk() (*string, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasSupport() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *VideoUploadRequestCommon) SetSupport(v string) {
	o.Support = &v
}

// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *VideoUploadRequestCommon) SetNsfw(v bool) {
	o.Nsfw = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VideoUploadRequestCommon) SetTags(v []string) {
	o.Tags = v
}

// GetCommentsEnabled returns the CommentsEnabled field value if set, zero value otherwise.
// Deprecated
func (o *VideoUploadRequestCommon) GetCommentsEnabled() bool {
	if o == nil || IsNil(o.CommentsEnabled) {
		var ret bool
		return ret
	}
	return *o.CommentsEnabled
}

// GetCommentsEnabledOk returns a tuple with the CommentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VideoUploadRequestCommon) GetCommentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CommentsEnabled) {
		return nil, false
	}
	return o.CommentsEnabled, true
}

// HasCommentsEnabled returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasCommentsEnabled() bool {
	if o != nil && !IsNil(o.CommentsEnabled) {
		return true
	}

	return false
}

// SetCommentsEnabled gets a reference to the given bool and assigns it to the CommentsEnabled field.
// Deprecated
func (o *VideoUploadRequestCommon) SetCommentsEnabled(v bool) {
	o.CommentsEnabled = &v
}

// GetCommentsPolicy returns the CommentsPolicy field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetCommentsPolicy() VideoCommentsPolicySet {
	if o == nil || IsNil(o.CommentsPolicy) {
		var ret VideoCommentsPolicySet
		return ret
	}
	return *o.CommentsPolicy
}

// GetCommentsPolicyOk returns a tuple with the CommentsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetCommentsPolicyOk() (*VideoCommentsPolicySet, bool) {
	if o == nil || IsNil(o.CommentsPolicy) {
		return nil, false
	}
	return o.CommentsPolicy, true
}

// HasCommentsPolicy returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasCommentsPolicy() bool {
	if o != nil && !IsNil(o.CommentsPolicy) {
		return true
	}

	return false
}

// SetCommentsPolicy gets a reference to the given VideoCommentsPolicySet and assigns it to the CommentsPolicy field.
func (o *VideoUploadRequestCommon) SetCommentsPolicy(v VideoCommentsPolicySet) {
	o.CommentsPolicy = &v
}

// GetDownloadEnabled returns the DownloadEnabled field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetDownloadEnabled() bool {
	if o == nil || IsNil(o.DownloadEnabled) {
		var ret bool
		return ret
	}
	return *o.DownloadEnabled
}

// GetDownloadEnabledOk returns a tuple with the DownloadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetDownloadEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadEnabled) {
		return nil, false
	}
	return o.DownloadEnabled, true
}

// HasDownloadEnabled returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasDownloadEnabled() bool {
	if o != nil && !IsNil(o.DownloadEnabled) {
		return true
	}

	return false
}

// SetDownloadEnabled gets a reference to the given bool and assigns it to the DownloadEnabled field.
func (o *VideoUploadRequestCommon) SetDownloadEnabled(v bool) {
	o.DownloadEnabled = &v
}

// GetOriginallyPublishedAt returns the OriginallyPublishedAt field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetOriginallyPublishedAt() time.Time {
	if o == nil || IsNil(o.OriginallyPublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.OriginallyPublishedAt
}

// GetOriginallyPublishedAtOk returns a tuple with the OriginallyPublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetOriginallyPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OriginallyPublishedAt) {
		return nil, false
	}
	return o.OriginallyPublishedAt, true
}

// HasOriginallyPublishedAt returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasOriginallyPublishedAt() bool {
	if o != nil && !IsNil(o.OriginallyPublishedAt) {
		return true
	}

	return false
}

// SetOriginallyPublishedAt gets a reference to the given time.Time and assigns it to the OriginallyPublishedAt field.
func (o *VideoUploadRequestCommon) SetOriginallyPublishedAt(v time.Time) {
	o.OriginallyPublishedAt = &v
}

// GetScheduleUpdate returns the ScheduleUpdate field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetScheduleUpdate() VideoScheduledUpdate {
	if o == nil || IsNil(o.ScheduleUpdate) {
		var ret VideoScheduledUpdate
		return ret
	}
	return *o.ScheduleUpdate
}

// GetScheduleUpdateOk returns a tuple with the ScheduleUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetScheduleUpdateOk() (*VideoScheduledUpdate, bool) {
	if o == nil || IsNil(o.ScheduleUpdate) {
		return nil, false
	}
	return o.ScheduleUpdate, true
}

// HasScheduleUpdate returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasScheduleUpdate() bool {
	if o != nil && !IsNil(o.ScheduleUpdate) {
		return true
	}

	return false
}

// SetScheduleUpdate gets a reference to the given VideoScheduledUpdate and assigns it to the ScheduleUpdate field.
func (o *VideoUploadRequestCommon) SetScheduleUpdate(v VideoScheduledUpdate) {
	o.ScheduleUpdate = &v
}

// GetThumbnailfile returns the Thumbnailfile field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetThumbnailfile() *os.File {
	if o == nil || IsNil(o.Thumbnailfile) {
		var ret *os.File
		return ret
	}
	return *o.Thumbnailfile
}

// GetThumbnailfileOk returns a tuple with the Thumbnailfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetThumbnailfileOk() (**os.File, bool) {
	if o == nil || IsNil(o.Thumbnailfile) {
		return nil, false
	}
	return o.Thumbnailfile, true
}

// HasThumbnailfile returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasThumbnailfile() bool {
	if o != nil && !IsNil(o.Thumbnailfile) {
		return true
	}

	return false
}

// SetThumbnailfile gets a reference to the given *os.File and assigns it to the Thumbnailfile field.
func (o *VideoUploadRequestCommon) SetThumbnailfile(v *os.File) {
	o.Thumbnailfile = &v
}

// GetPreviewfile returns the Previewfile field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetPreviewfile() *os.File {
	if o == nil || IsNil(o.Previewfile) {
		var ret *os.File
		return ret
	}
	return *o.Previewfile
}

// GetPreviewfileOk returns a tuple with the Previewfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetPreviewfileOk() (**os.File, bool) {
	if o == nil || IsNil(o.Previewfile) {
		return nil, false
	}
	return o.Previewfile, true
}

// HasPreviewfile returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasPreviewfile() bool {
	if o != nil && !IsNil(o.Previewfile) {
		return true
	}

	return false
}

// SetPreviewfile gets a reference to the given *os.File and assigns it to the Previewfile field.
func (o *VideoUploadRequestCommon) SetPreviewfile(v *os.File) {
	o.Previewfile = &v
}

// GetVideoPasswords returns the VideoPasswords field value if set, zero value otherwise.
func (o *VideoUploadRequestCommon) GetVideoPasswords() []string {
	if o == nil || IsNil(o.VideoPasswords) {
		var ret []string
		return ret
	}
	return o.VideoPasswords
}

// GetVideoPasswordsOk returns a tuple with the VideoPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUploadRequestCommon) GetVideoPasswordsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoPasswords) {
		return nil, false
	}
	return o.VideoPasswords, true
}

// HasVideoPasswords returns a boolean if a field has been set.
func (o *VideoUploadRequestCommon) HasVideoPasswords() bool {
	if o != nil && !IsNil(o.VideoPasswords) {
		return true
	}

	return false
}

// SetVideoPasswords gets a reference to the given []string and assigns it to the VideoPasswords field.
func (o *VideoUploadRequestCommon) SetVideoPasswords(v []string) {
	o.VideoPasswords = v
}

func (o VideoUploadRequestCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoUploadRequestCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["channelId"] = o.ChannelId
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.WaitTranscoding) {
		toSerialize["waitTranscoding"] = o.WaitTranscoding
	}
	if !IsNil(o.GenerateTranscription) {
		toSerialize["generateTranscription"] = o.GenerateTranscription
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !IsNil(o.Nsfw) {
		toSerialize["nsfw"] = o.Nsfw
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
	if !IsNil(o.OriginallyPublishedAt) {
		toSerialize["originallyPublishedAt"] = o.OriginallyPublishedAt
	}
	if !IsNil(o.ScheduleUpdate) {
		toSerialize["scheduleUpdate"] = o.ScheduleUpdate
	}
	if !IsNil(o.Thumbnailfile) {
		toSerialize["thumbnailfile"] = o.Thumbnailfile
	}
	if !IsNil(o.Previewfile) {
		toSerialize["previewfile"] = o.Previewfile
	}
	if !IsNil(o.VideoPasswords) {
		toSerialize["videoPasswords"] = o.VideoPasswords
	}
	return toSerialize, nil
}

func (o *VideoUploadRequestCommon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"channelId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVideoUploadRequestCommon := _VideoUploadRequestCommon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoUploadRequestCommon)

	if err != nil {
		return err
	}

	*o = VideoUploadRequestCommon(varVideoUploadRequestCommon)

	return err
}

type NullableVideoUploadRequestCommon struct {
	value *VideoUploadRequestCommon
	isSet bool
}

func (v NullableVideoUploadRequestCommon) Get() *VideoUploadRequestCommon {
	return v.value
}

func (v *NullableVideoUploadRequestCommon) Set(val *VideoUploadRequestCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoUploadRequestCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoUploadRequestCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoUploadRequestCommon(val *VideoUploadRequestCommon) *NullableVideoUploadRequestCommon {
	return &NullableVideoUploadRequestCommon{value: val, isSet: true}
}

func (v NullableVideoUploadRequestCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoUploadRequestCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
