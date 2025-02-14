# VideosForXMLInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | video watch page URL | [optional] 
**Guid** | Pointer to **string** | video canonical URL | [optional] 
**PubDate** | Pointer to **time.Time** | video publication date | [optional] 
**Description** | Pointer to **string** | video description | [optional] 
**ContentEncoded** | Pointer to **string** | video description | [optional] 
**DcCreator** | Pointer to **string** | publisher user name | [optional] 
**MediaCategory** | Pointer to **int32** | video category (MRSS) | [optional] 
**MediaCommunity** | Pointer to [**VideosForXMLInnerMediaCommunity**](VideosForXMLInnerMediaCommunity.md) |  | [optional] 
**MediaEmbed** | Pointer to [**VideosForXMLInnerMediaEmbed**](VideosForXMLInnerMediaEmbed.md) |  | [optional] 
**MediaPlayer** | Pointer to [**VideosForXMLInnerMediaPlayer**](VideosForXMLInnerMediaPlayer.md) |  | [optional] 
**MediaThumbnail** | Pointer to [**VideosForXMLInnerMediaThumbnail**](VideosForXMLInnerMediaThumbnail.md) |  | [optional] 
**MediaTitle** | Pointer to **string** | see [media:title](https://www.rssboard.org/media-rss#media-title) (MRSS). We only use &#x60;plain&#x60; titles. | [optional] 
**MediaDescription** | Pointer to **string** |  | [optional] 
**MediaRating** | Pointer to **string** | see [media:rating](https://www.rssboard.org/media-rss#media-rating) (MRSS) | [optional] 
**Enclosure** | Pointer to [**VideosForXMLInnerEnclosure**](VideosForXMLInnerEnclosure.md) |  | [optional] 
**MediaGroup** | Pointer to [**[]VideosForXMLInnerMediaGroupInner**](VideosForXMLInnerMediaGroupInner.md) | list of streamable files for the video. see [media:peerLink](https://www.rssboard.org/media-rss#media-peerlink) and [media:content](https://www.rssboard.org/media-rss#media-content) or  (MRSS) | [optional] 

## Methods

### NewVideosForXMLInner

`func NewVideosForXMLInner() *VideosForXMLInner`

NewVideosForXMLInner instantiates a new VideosForXMLInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideosForXMLInnerWithDefaults

`func NewVideosForXMLInnerWithDefaults() *VideosForXMLInner`

NewVideosForXMLInnerWithDefaults instantiates a new VideosForXMLInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *VideosForXMLInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *VideosForXMLInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *VideosForXMLInner) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *VideosForXMLInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetGuid

`func (o *VideosForXMLInner) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *VideosForXMLInner) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *VideosForXMLInner) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *VideosForXMLInner) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetPubDate

`func (o *VideosForXMLInner) GetPubDate() time.Time`

GetPubDate returns the PubDate field if non-nil, zero value otherwise.

### GetPubDateOk

`func (o *VideosForXMLInner) GetPubDateOk() (*time.Time, bool)`

GetPubDateOk returns a tuple with the PubDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubDate

`func (o *VideosForXMLInner) SetPubDate(v time.Time)`

SetPubDate sets PubDate field to given value.

### HasPubDate

`func (o *VideosForXMLInner) HasPubDate() bool`

HasPubDate returns a boolean if a field has been set.

### GetDescription

`func (o *VideosForXMLInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideosForXMLInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideosForXMLInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideosForXMLInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentEncoded

`func (o *VideosForXMLInner) GetContentEncoded() string`

GetContentEncoded returns the ContentEncoded field if non-nil, zero value otherwise.

### GetContentEncodedOk

`func (o *VideosForXMLInner) GetContentEncodedOk() (*string, bool)`

GetContentEncodedOk returns a tuple with the ContentEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEncoded

`func (o *VideosForXMLInner) SetContentEncoded(v string)`

SetContentEncoded sets ContentEncoded field to given value.

### HasContentEncoded

`func (o *VideosForXMLInner) HasContentEncoded() bool`

HasContentEncoded returns a boolean if a field has been set.

### GetDcCreator

`func (o *VideosForXMLInner) GetDcCreator() string`

GetDcCreator returns the DcCreator field if non-nil, zero value otherwise.

### GetDcCreatorOk

`func (o *VideosForXMLInner) GetDcCreatorOk() (*string, bool)`

GetDcCreatorOk returns a tuple with the DcCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcCreator

`func (o *VideosForXMLInner) SetDcCreator(v string)`

SetDcCreator sets DcCreator field to given value.

### HasDcCreator

`func (o *VideosForXMLInner) HasDcCreator() bool`

HasDcCreator returns a boolean if a field has been set.

### GetMediaCategory

`func (o *VideosForXMLInner) GetMediaCategory() int32`

GetMediaCategory returns the MediaCategory field if non-nil, zero value otherwise.

### GetMediaCategoryOk

`func (o *VideosForXMLInner) GetMediaCategoryOk() (*int32, bool)`

GetMediaCategoryOk returns a tuple with the MediaCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaCategory

`func (o *VideosForXMLInner) SetMediaCategory(v int32)`

SetMediaCategory sets MediaCategory field to given value.

### HasMediaCategory

`func (o *VideosForXMLInner) HasMediaCategory() bool`

HasMediaCategory returns a boolean if a field has been set.

### GetMediaCommunity

`func (o *VideosForXMLInner) GetMediaCommunity() VideosForXMLInnerMediaCommunity`

GetMediaCommunity returns the MediaCommunity field if non-nil, zero value otherwise.

### GetMediaCommunityOk

`func (o *VideosForXMLInner) GetMediaCommunityOk() (*VideosForXMLInnerMediaCommunity, bool)`

GetMediaCommunityOk returns a tuple with the MediaCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaCommunity

`func (o *VideosForXMLInner) SetMediaCommunity(v VideosForXMLInnerMediaCommunity)`

SetMediaCommunity sets MediaCommunity field to given value.

### HasMediaCommunity

`func (o *VideosForXMLInner) HasMediaCommunity() bool`

HasMediaCommunity returns a boolean if a field has been set.

### GetMediaEmbed

`func (o *VideosForXMLInner) GetMediaEmbed() VideosForXMLInnerMediaEmbed`

GetMediaEmbed returns the MediaEmbed field if non-nil, zero value otherwise.

### GetMediaEmbedOk

`func (o *VideosForXMLInner) GetMediaEmbedOk() (*VideosForXMLInnerMediaEmbed, bool)`

GetMediaEmbedOk returns a tuple with the MediaEmbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEmbed

`func (o *VideosForXMLInner) SetMediaEmbed(v VideosForXMLInnerMediaEmbed)`

SetMediaEmbed sets MediaEmbed field to given value.

### HasMediaEmbed

`func (o *VideosForXMLInner) HasMediaEmbed() bool`

HasMediaEmbed returns a boolean if a field has been set.

### GetMediaPlayer

`func (o *VideosForXMLInner) GetMediaPlayer() VideosForXMLInnerMediaPlayer`

GetMediaPlayer returns the MediaPlayer field if non-nil, zero value otherwise.

### GetMediaPlayerOk

`func (o *VideosForXMLInner) GetMediaPlayerOk() (*VideosForXMLInnerMediaPlayer, bool)`

GetMediaPlayerOk returns a tuple with the MediaPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaPlayer

`func (o *VideosForXMLInner) SetMediaPlayer(v VideosForXMLInnerMediaPlayer)`

SetMediaPlayer sets MediaPlayer field to given value.

### HasMediaPlayer

`func (o *VideosForXMLInner) HasMediaPlayer() bool`

HasMediaPlayer returns a boolean if a field has been set.

### GetMediaThumbnail

`func (o *VideosForXMLInner) GetMediaThumbnail() VideosForXMLInnerMediaThumbnail`

GetMediaThumbnail returns the MediaThumbnail field if non-nil, zero value otherwise.

### GetMediaThumbnailOk

`func (o *VideosForXMLInner) GetMediaThumbnailOk() (*VideosForXMLInnerMediaThumbnail, bool)`

GetMediaThumbnailOk returns a tuple with the MediaThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaThumbnail

`func (o *VideosForXMLInner) SetMediaThumbnail(v VideosForXMLInnerMediaThumbnail)`

SetMediaThumbnail sets MediaThumbnail field to given value.

### HasMediaThumbnail

`func (o *VideosForXMLInner) HasMediaThumbnail() bool`

HasMediaThumbnail returns a boolean if a field has been set.

### GetMediaTitle

`func (o *VideosForXMLInner) GetMediaTitle() string`

GetMediaTitle returns the MediaTitle field if non-nil, zero value otherwise.

### GetMediaTitleOk

`func (o *VideosForXMLInner) GetMediaTitleOk() (*string, bool)`

GetMediaTitleOk returns a tuple with the MediaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTitle

`func (o *VideosForXMLInner) SetMediaTitle(v string)`

SetMediaTitle sets MediaTitle field to given value.

### HasMediaTitle

`func (o *VideosForXMLInner) HasMediaTitle() bool`

HasMediaTitle returns a boolean if a field has been set.

### GetMediaDescription

`func (o *VideosForXMLInner) GetMediaDescription() string`

GetMediaDescription returns the MediaDescription field if non-nil, zero value otherwise.

### GetMediaDescriptionOk

`func (o *VideosForXMLInner) GetMediaDescriptionOk() (*string, bool)`

GetMediaDescriptionOk returns a tuple with the MediaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaDescription

`func (o *VideosForXMLInner) SetMediaDescription(v string)`

SetMediaDescription sets MediaDescription field to given value.

### HasMediaDescription

`func (o *VideosForXMLInner) HasMediaDescription() bool`

HasMediaDescription returns a boolean if a field has been set.

### GetMediaRating

`func (o *VideosForXMLInner) GetMediaRating() string`

GetMediaRating returns the MediaRating field if non-nil, zero value otherwise.

### GetMediaRatingOk

`func (o *VideosForXMLInner) GetMediaRatingOk() (*string, bool)`

GetMediaRatingOk returns a tuple with the MediaRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaRating

`func (o *VideosForXMLInner) SetMediaRating(v string)`

SetMediaRating sets MediaRating field to given value.

### HasMediaRating

`func (o *VideosForXMLInner) HasMediaRating() bool`

HasMediaRating returns a boolean if a field has been set.

### GetEnclosure

`func (o *VideosForXMLInner) GetEnclosure() VideosForXMLInnerEnclosure`

GetEnclosure returns the Enclosure field if non-nil, zero value otherwise.

### GetEnclosureOk

`func (o *VideosForXMLInner) GetEnclosureOk() (*VideosForXMLInnerEnclosure, bool)`

GetEnclosureOk returns a tuple with the Enclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosure

`func (o *VideosForXMLInner) SetEnclosure(v VideosForXMLInnerEnclosure)`

SetEnclosure sets Enclosure field to given value.

### HasEnclosure

`func (o *VideosForXMLInner) HasEnclosure() bool`

HasEnclosure returns a boolean if a field has been set.

### GetMediaGroup

`func (o *VideosForXMLInner) GetMediaGroup() []VideosForXMLInnerMediaGroupInner`

GetMediaGroup returns the MediaGroup field if non-nil, zero value otherwise.

### GetMediaGroupOk

`func (o *VideosForXMLInner) GetMediaGroupOk() (*[]VideosForXMLInnerMediaGroupInner, bool)`

GetMediaGroupOk returns a tuple with the MediaGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaGroup

`func (o *VideosForXMLInner) SetMediaGroup(v []VideosForXMLInnerMediaGroupInner)`

SetMediaGroup sets MediaGroup field to given value.

### HasMediaGroup

`func (o *VideosForXMLInner) HasMediaGroup() bool`

HasMediaGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


