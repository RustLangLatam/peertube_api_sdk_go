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

// checks if the VideosForXMLInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideosForXMLInner{}

// VideosForXMLInner struct for VideosForXMLInner
type VideosForXMLInner struct {
	// video watch page URL
	Link *string `json:"link,omitempty"`
	// video canonical URL
	Guid *string `json:"guid,omitempty"`
	// video publication date
	PubDate *time.Time `json:"pubDate,omitempty"`
	// video description
	Description *string `json:"description,omitempty"`
	// video description
	ContentEncoded *string `json:"content:encoded,omitempty"`
	// publisher user name
	DcCreator *string `json:"dc:creator,omitempty"`
	// video category (MRSS)
	MediaCategory  *int32                           `json:"media:category,omitempty"`
	MediaCommunity *VideosForXMLInnerMediaCommunity `json:"media:community,omitempty"`
	MediaEmbed     *VideosForXMLInnerMediaEmbed     `json:"media:embed,omitempty"`
	MediaPlayer    *VideosForXMLInnerMediaPlayer    `json:"media:player,omitempty"`
	MediaThumbnail *VideosForXMLInnerMediaThumbnail `json:"media:thumbnail,omitempty"`
	// see [media:title](https://www.rssboard.org/media-rss#media-title) (MRSS). We only use `plain` titles.
	MediaTitle       *string `json:"media:title,omitempty"`
	MediaDescription *string `json:"media:description,omitempty"`
	// see [media:rating](https://www.rssboard.org/media-rss#media-rating) (MRSS)
	MediaRating *string                     `json:"media:rating,omitempty"`
	Enclosure   *VideosForXMLInnerEnclosure `json:"enclosure,omitempty"`
	// list of streamable files for the video. see [media:peerLink](https://www.rssboard.org/media-rss#media-peerlink) and [media:content](https://www.rssboard.org/media-rss#media-content) or  (MRSS)
	MediaGroup []VideosForXMLInnerMediaGroupInner `json:"media:group,omitempty"`
}

// NewVideosForXMLInner instantiates a new VideosForXMLInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideosForXMLInner() *VideosForXMLInner {
	this := VideosForXMLInner{}
	return &this
}

// NewVideosForXMLInnerWithDefaults instantiates a new VideosForXMLInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideosForXMLInnerWithDefaults() *VideosForXMLInner {
	this := VideosForXMLInner{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *VideosForXMLInner) SetLink(v string) {
	o.Link = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *VideosForXMLInner) SetGuid(v string) {
	o.Guid = &v
}

// GetPubDate returns the PubDate field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetPubDate() time.Time {
	if o == nil || IsNil(o.PubDate) {
		var ret time.Time
		return ret
	}
	return *o.PubDate
}

// GetPubDateOk returns a tuple with the PubDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetPubDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PubDate) {
		return nil, false
	}
	return o.PubDate, true
}

// HasPubDate returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasPubDate() bool {
	if o != nil && !IsNil(o.PubDate) {
		return true
	}

	return false
}

// SetPubDate gets a reference to the given time.Time and assigns it to the PubDate field.
func (o *VideosForXMLInner) SetPubDate(v time.Time) {
	o.PubDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VideosForXMLInner) SetDescription(v string) {
	o.Description = &v
}

// GetContentEncoded returns the ContentEncoded field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetContentEncoded() string {
	if o == nil || IsNil(o.ContentEncoded) {
		var ret string
		return ret
	}
	return *o.ContentEncoded
}

// GetContentEncodedOk returns a tuple with the ContentEncoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetContentEncodedOk() (*string, bool) {
	if o == nil || IsNil(o.ContentEncoded) {
		return nil, false
	}
	return o.ContentEncoded, true
}

// HasContentEncoded returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasContentEncoded() bool {
	if o != nil && !IsNil(o.ContentEncoded) {
		return true
	}

	return false
}

// SetContentEncoded gets a reference to the given string and assigns it to the ContentEncoded field.
func (o *VideosForXMLInner) SetContentEncoded(v string) {
	o.ContentEncoded = &v
}

// GetDcCreator returns the DcCreator field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetDcCreator() string {
	if o == nil || IsNil(o.DcCreator) {
		var ret string
		return ret
	}
	return *o.DcCreator
}

// GetDcCreatorOk returns a tuple with the DcCreator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetDcCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.DcCreator) {
		return nil, false
	}
	return o.DcCreator, true
}

// HasDcCreator returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasDcCreator() bool {
	if o != nil && !IsNil(o.DcCreator) {
		return true
	}

	return false
}

// SetDcCreator gets a reference to the given string and assigns it to the DcCreator field.
func (o *VideosForXMLInner) SetDcCreator(v string) {
	o.DcCreator = &v
}

// GetMediaCategory returns the MediaCategory field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaCategory() int32 {
	if o == nil || IsNil(o.MediaCategory) {
		var ret int32
		return ret
	}
	return *o.MediaCategory
}

// GetMediaCategoryOk returns a tuple with the MediaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaCategoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MediaCategory) {
		return nil, false
	}
	return o.MediaCategory, true
}

// HasMediaCategory returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaCategory() bool {
	if o != nil && !IsNil(o.MediaCategory) {
		return true
	}

	return false
}

// SetMediaCategory gets a reference to the given int32 and assigns it to the MediaCategory field.
func (o *VideosForXMLInner) SetMediaCategory(v int32) {
	o.MediaCategory = &v
}

// GetMediaCommunity returns the MediaCommunity field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaCommunity() VideosForXMLInnerMediaCommunity {
	if o == nil || IsNil(o.MediaCommunity) {
		var ret VideosForXMLInnerMediaCommunity
		return ret
	}
	return *o.MediaCommunity
}

// GetMediaCommunityOk returns a tuple with the MediaCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaCommunityOk() (*VideosForXMLInnerMediaCommunity, bool) {
	if o == nil || IsNil(o.MediaCommunity) {
		return nil, false
	}
	return o.MediaCommunity, true
}

// HasMediaCommunity returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaCommunity() bool {
	if o != nil && !IsNil(o.MediaCommunity) {
		return true
	}

	return false
}

// SetMediaCommunity gets a reference to the given VideosForXMLInnerMediaCommunity and assigns it to the MediaCommunity field.
func (o *VideosForXMLInner) SetMediaCommunity(v VideosForXMLInnerMediaCommunity) {
	o.MediaCommunity = &v
}

// GetMediaEmbed returns the MediaEmbed field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaEmbed() VideosForXMLInnerMediaEmbed {
	if o == nil || IsNil(o.MediaEmbed) {
		var ret VideosForXMLInnerMediaEmbed
		return ret
	}
	return *o.MediaEmbed
}

// GetMediaEmbedOk returns a tuple with the MediaEmbed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaEmbedOk() (*VideosForXMLInnerMediaEmbed, bool) {
	if o == nil || IsNil(o.MediaEmbed) {
		return nil, false
	}
	return o.MediaEmbed, true
}

// HasMediaEmbed returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaEmbed() bool {
	if o != nil && !IsNil(o.MediaEmbed) {
		return true
	}

	return false
}

// SetMediaEmbed gets a reference to the given VideosForXMLInnerMediaEmbed and assigns it to the MediaEmbed field.
func (o *VideosForXMLInner) SetMediaEmbed(v VideosForXMLInnerMediaEmbed) {
	o.MediaEmbed = &v
}

// GetMediaPlayer returns the MediaPlayer field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaPlayer() VideosForXMLInnerMediaPlayer {
	if o == nil || IsNil(o.MediaPlayer) {
		var ret VideosForXMLInnerMediaPlayer
		return ret
	}
	return *o.MediaPlayer
}

// GetMediaPlayerOk returns a tuple with the MediaPlayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaPlayerOk() (*VideosForXMLInnerMediaPlayer, bool) {
	if o == nil || IsNil(o.MediaPlayer) {
		return nil, false
	}
	return o.MediaPlayer, true
}

// HasMediaPlayer returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaPlayer() bool {
	if o != nil && !IsNil(o.MediaPlayer) {
		return true
	}

	return false
}

// SetMediaPlayer gets a reference to the given VideosForXMLInnerMediaPlayer and assigns it to the MediaPlayer field.
func (o *VideosForXMLInner) SetMediaPlayer(v VideosForXMLInnerMediaPlayer) {
	o.MediaPlayer = &v
}

// GetMediaThumbnail returns the MediaThumbnail field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaThumbnail() VideosForXMLInnerMediaThumbnail {
	if o == nil || IsNil(o.MediaThumbnail) {
		var ret VideosForXMLInnerMediaThumbnail
		return ret
	}
	return *o.MediaThumbnail
}

// GetMediaThumbnailOk returns a tuple with the MediaThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaThumbnailOk() (*VideosForXMLInnerMediaThumbnail, bool) {
	if o == nil || IsNil(o.MediaThumbnail) {
		return nil, false
	}
	return o.MediaThumbnail, true
}

// HasMediaThumbnail returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaThumbnail() bool {
	if o != nil && !IsNil(o.MediaThumbnail) {
		return true
	}

	return false
}

// SetMediaThumbnail gets a reference to the given VideosForXMLInnerMediaThumbnail and assigns it to the MediaThumbnail field.
func (o *VideosForXMLInner) SetMediaThumbnail(v VideosForXMLInnerMediaThumbnail) {
	o.MediaThumbnail = &v
}

// GetMediaTitle returns the MediaTitle field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaTitle() string {
	if o == nil || IsNil(o.MediaTitle) {
		var ret string
		return ret
	}
	return *o.MediaTitle
}

// GetMediaTitleOk returns a tuple with the MediaTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaTitleOk() (*string, bool) {
	if o == nil || IsNil(o.MediaTitle) {
		return nil, false
	}
	return o.MediaTitle, true
}

// HasMediaTitle returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaTitle() bool {
	if o != nil && !IsNil(o.MediaTitle) {
		return true
	}

	return false
}

// SetMediaTitle gets a reference to the given string and assigns it to the MediaTitle field.
func (o *VideosForXMLInner) SetMediaTitle(v string) {
	o.MediaTitle = &v
}

// GetMediaDescription returns the MediaDescription field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaDescription() string {
	if o == nil || IsNil(o.MediaDescription) {
		var ret string
		return ret
	}
	return *o.MediaDescription
}

// GetMediaDescriptionOk returns a tuple with the MediaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MediaDescription) {
		return nil, false
	}
	return o.MediaDescription, true
}

// HasMediaDescription returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaDescription() bool {
	if o != nil && !IsNil(o.MediaDescription) {
		return true
	}

	return false
}

// SetMediaDescription gets a reference to the given string and assigns it to the MediaDescription field.
func (o *VideosForXMLInner) SetMediaDescription(v string) {
	o.MediaDescription = &v
}

// GetMediaRating returns the MediaRating field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaRating() string {
	if o == nil || IsNil(o.MediaRating) {
		var ret string
		return ret
	}
	return *o.MediaRating
}

// GetMediaRatingOk returns a tuple with the MediaRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaRatingOk() (*string, bool) {
	if o == nil || IsNil(o.MediaRating) {
		return nil, false
	}
	return o.MediaRating, true
}

// HasMediaRating returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaRating() bool {
	if o != nil && !IsNil(o.MediaRating) {
		return true
	}

	return false
}

// SetMediaRating gets a reference to the given string and assigns it to the MediaRating field.
func (o *VideosForXMLInner) SetMediaRating(v string) {
	o.MediaRating = &v
}

// GetEnclosure returns the Enclosure field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetEnclosure() VideosForXMLInnerEnclosure {
	if o == nil || IsNil(o.Enclosure) {
		var ret VideosForXMLInnerEnclosure
		return ret
	}
	return *o.Enclosure
}

// GetEnclosureOk returns a tuple with the Enclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetEnclosureOk() (*VideosForXMLInnerEnclosure, bool) {
	if o == nil || IsNil(o.Enclosure) {
		return nil, false
	}
	return o.Enclosure, true
}

// HasEnclosure returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasEnclosure() bool {
	if o != nil && !IsNil(o.Enclosure) {
		return true
	}

	return false
}

// SetEnclosure gets a reference to the given VideosForXMLInnerEnclosure and assigns it to the Enclosure field.
func (o *VideosForXMLInner) SetEnclosure(v VideosForXMLInnerEnclosure) {
	o.Enclosure = &v
}

// GetMediaGroup returns the MediaGroup field value if set, zero value otherwise.
func (o *VideosForXMLInner) GetMediaGroup() []VideosForXMLInnerMediaGroupInner {
	if o == nil || IsNil(o.MediaGroup) {
		var ret []VideosForXMLInnerMediaGroupInner
		return ret
	}
	return o.MediaGroup
}

// GetMediaGroupOk returns a tuple with the MediaGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideosForXMLInner) GetMediaGroupOk() ([]VideosForXMLInnerMediaGroupInner, bool) {
	if o == nil || IsNil(o.MediaGroup) {
		return nil, false
	}
	return o.MediaGroup, true
}

// HasMediaGroup returns a boolean if a field has been set.
func (o *VideosForXMLInner) HasMediaGroup() bool {
	if o != nil && !IsNil(o.MediaGroup) {
		return true
	}

	return false
}

// SetMediaGroup gets a reference to the given []VideosForXMLInnerMediaGroupInner and assigns it to the MediaGroup field.
func (o *VideosForXMLInner) SetMediaGroup(v []VideosForXMLInnerMediaGroupInner) {
	o.MediaGroup = v
}

func (o VideosForXMLInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideosForXMLInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.PubDate) {
		toSerialize["pubDate"] = o.PubDate
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ContentEncoded) {
		toSerialize["content:encoded"] = o.ContentEncoded
	}
	if !IsNil(o.DcCreator) {
		toSerialize["dc:creator"] = o.DcCreator
	}
	if !IsNil(o.MediaCategory) {
		toSerialize["media:category"] = o.MediaCategory
	}
	if !IsNil(o.MediaCommunity) {
		toSerialize["media:community"] = o.MediaCommunity
	}
	if !IsNil(o.MediaEmbed) {
		toSerialize["media:embed"] = o.MediaEmbed
	}
	if !IsNil(o.MediaPlayer) {
		toSerialize["media:player"] = o.MediaPlayer
	}
	if !IsNil(o.MediaThumbnail) {
		toSerialize["media:thumbnail"] = o.MediaThumbnail
	}
	if !IsNil(o.MediaTitle) {
		toSerialize["media:title"] = o.MediaTitle
	}
	if !IsNil(o.MediaDescription) {
		toSerialize["media:description"] = o.MediaDescription
	}
	if !IsNil(o.MediaRating) {
		toSerialize["media:rating"] = o.MediaRating
	}
	if !IsNil(o.Enclosure) {
		toSerialize["enclosure"] = o.Enclosure
	}
	if !IsNil(o.MediaGroup) {
		toSerialize["media:group"] = o.MediaGroup
	}
	return toSerialize, nil
}

type NullableVideosForXMLInner struct {
	value *VideosForXMLInner
	isSet bool
}

func (v NullableVideosForXMLInner) Get() *VideosForXMLInner {
	return v.value
}

func (v *NullableVideosForXMLInner) Set(val *VideosForXMLInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVideosForXMLInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVideosForXMLInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideosForXMLInner(val *VideosForXMLInner) *NullableVideosForXMLInner {
	return &NullableVideosForXMLInner{value: val, isSet: true}
}

func (v NullableVideosForXMLInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideosForXMLInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
