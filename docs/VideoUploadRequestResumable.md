# VideoUploadRequestResumable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Video name | 
**ChannelId** | **int32** | Channel id that will contain this video | 
**Filename** | **string** | Video filename including extension | 
**Privacy** | Pointer to [**VideoPrivacySet**](VideoPrivacySet.md) |  | [optional] 
**Category** | Pointer to **int32** | category id of the video (see [/videos/categories](#operation/getCategories)) | [optional] 
**Licence** | Pointer to **int32** | licence id of the video (see [/videos/licences](#operation/getLicences)) | [optional] 
**Language** | Pointer to **string** | language id of the video (see [/videos/languages](#operation/getLanguages)) | [optional] 
**Description** | Pointer to **string** | Video description | [optional] 
**WaitTranscoding** | Pointer to **bool** | Whether or not we wait transcoding before publish the video | [optional] 
**GenerateTranscription** | Pointer to **bool** | **PeerTube &gt;&#x3D; 6.2** If enabled by the admin, automatically generate a subtitle of the video | [optional] 
**Support** | Pointer to **string** | A text tell the audience how to support the video creator | [optional] 
**Nsfw** | Pointer to **bool** | Whether or not this video contains sensitive content | [optional] 
**Tags** | Pointer to **[]string** | Video tags (maximum 5 tags each between 2 and 30 characters) | [optional] 
**CommentsEnabled** | Pointer to **bool** | Deprecated in 6.2, use commentsPolicy instead | [optional] 
**CommentsPolicy** | Pointer to [**VideoCommentsPolicySet**](VideoCommentsPolicySet.md) |  | [optional] 
**DownloadEnabled** | Pointer to **bool** | Enable or disable downloading for this video | [optional] 
**OriginallyPublishedAt** | Pointer to **time.Time** | Date when the content was originally published | [optional] 
**ScheduleUpdate** | Pointer to [**VideoScheduledUpdate**](VideoScheduledUpdate.md) |  | [optional] 
**Thumbnailfile** | Pointer to ***os.File** | Video thumbnail file | [optional] 
**Previewfile** | Pointer to ***os.File** | Video preview file | [optional] 
**VideoPasswords** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVideoUploadRequestResumable

`func NewVideoUploadRequestResumable(name string, channelId int32, filename string, ) *VideoUploadRequestResumable`

NewVideoUploadRequestResumable instantiates a new VideoUploadRequestResumable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoUploadRequestResumableWithDefaults

`func NewVideoUploadRequestResumableWithDefaults() *VideoUploadRequestResumable`

NewVideoUploadRequestResumableWithDefaults instantiates a new VideoUploadRequestResumable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VideoUploadRequestResumable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoUploadRequestResumable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoUploadRequestResumable) SetName(v string)`

SetName sets Name field to given value.


### GetChannelId

`func (o *VideoUploadRequestResumable) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *VideoUploadRequestResumable) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *VideoUploadRequestResumable) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetFilename

`func (o *VideoUploadRequestResumable) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VideoUploadRequestResumable) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VideoUploadRequestResumable) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPrivacy

`func (o *VideoUploadRequestResumable) GetPrivacy() VideoPrivacySet`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *VideoUploadRequestResumable) GetPrivacyOk() (*VideoPrivacySet, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *VideoUploadRequestResumable) SetPrivacy(v VideoPrivacySet)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *VideoUploadRequestResumable) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetCategory

`func (o *VideoUploadRequestResumable) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VideoUploadRequestResumable) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VideoUploadRequestResumable) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *VideoUploadRequestResumable) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLicence

`func (o *VideoUploadRequestResumable) GetLicence() int32`

GetLicence returns the Licence field if non-nil, zero value otherwise.

### GetLicenceOk

`func (o *VideoUploadRequestResumable) GetLicenceOk() (*int32, bool)`

GetLicenceOk returns a tuple with the Licence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicence

`func (o *VideoUploadRequestResumable) SetLicence(v int32)`

SetLicence sets Licence field to given value.

### HasLicence

`func (o *VideoUploadRequestResumable) HasLicence() bool`

HasLicence returns a boolean if a field has been set.

### GetLanguage

`func (o *VideoUploadRequestResumable) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VideoUploadRequestResumable) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VideoUploadRequestResumable) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VideoUploadRequestResumable) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetDescription

`func (o *VideoUploadRequestResumable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoUploadRequestResumable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoUploadRequestResumable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoUploadRequestResumable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWaitTranscoding

`func (o *VideoUploadRequestResumable) GetWaitTranscoding() bool`

GetWaitTranscoding returns the WaitTranscoding field if non-nil, zero value otherwise.

### GetWaitTranscodingOk

`func (o *VideoUploadRequestResumable) GetWaitTranscodingOk() (*bool, bool)`

GetWaitTranscodingOk returns a tuple with the WaitTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTranscoding

`func (o *VideoUploadRequestResumable) SetWaitTranscoding(v bool)`

SetWaitTranscoding sets WaitTranscoding field to given value.

### HasWaitTranscoding

`func (o *VideoUploadRequestResumable) HasWaitTranscoding() bool`

HasWaitTranscoding returns a boolean if a field has been set.

### GetGenerateTranscription

`func (o *VideoUploadRequestResumable) GetGenerateTranscription() bool`

GetGenerateTranscription returns the GenerateTranscription field if non-nil, zero value otherwise.

### GetGenerateTranscriptionOk

`func (o *VideoUploadRequestResumable) GetGenerateTranscriptionOk() (*bool, bool)`

GetGenerateTranscriptionOk returns a tuple with the GenerateTranscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateTranscription

`func (o *VideoUploadRequestResumable) SetGenerateTranscription(v bool)`

SetGenerateTranscription sets GenerateTranscription field to given value.

### HasGenerateTranscription

`func (o *VideoUploadRequestResumable) HasGenerateTranscription() bool`

HasGenerateTranscription returns a boolean if a field has been set.

### GetSupport

`func (o *VideoUploadRequestResumable) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *VideoUploadRequestResumable) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *VideoUploadRequestResumable) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *VideoUploadRequestResumable) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetNsfw

`func (o *VideoUploadRequestResumable) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *VideoUploadRequestResumable) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *VideoUploadRequestResumable) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *VideoUploadRequestResumable) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetTags

`func (o *VideoUploadRequestResumable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VideoUploadRequestResumable) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VideoUploadRequestResumable) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VideoUploadRequestResumable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCommentsEnabled

`func (o *VideoUploadRequestResumable) GetCommentsEnabled() bool`

GetCommentsEnabled returns the CommentsEnabled field if non-nil, zero value otherwise.

### GetCommentsEnabledOk

`func (o *VideoUploadRequestResumable) GetCommentsEnabledOk() (*bool, bool)`

GetCommentsEnabledOk returns a tuple with the CommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsEnabled

`func (o *VideoUploadRequestResumable) SetCommentsEnabled(v bool)`

SetCommentsEnabled sets CommentsEnabled field to given value.

### HasCommentsEnabled

`func (o *VideoUploadRequestResumable) HasCommentsEnabled() bool`

HasCommentsEnabled returns a boolean if a field has been set.

### GetCommentsPolicy

`func (o *VideoUploadRequestResumable) GetCommentsPolicy() VideoCommentsPolicySet`

GetCommentsPolicy returns the CommentsPolicy field if non-nil, zero value otherwise.

### GetCommentsPolicyOk

`func (o *VideoUploadRequestResumable) GetCommentsPolicyOk() (*VideoCommentsPolicySet, bool)`

GetCommentsPolicyOk returns a tuple with the CommentsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsPolicy

`func (o *VideoUploadRequestResumable) SetCommentsPolicy(v VideoCommentsPolicySet)`

SetCommentsPolicy sets CommentsPolicy field to given value.

### HasCommentsPolicy

`func (o *VideoUploadRequestResumable) HasCommentsPolicy() bool`

HasCommentsPolicy returns a boolean if a field has been set.

### GetDownloadEnabled

`func (o *VideoUploadRequestResumable) GetDownloadEnabled() bool`

GetDownloadEnabled returns the DownloadEnabled field if non-nil, zero value otherwise.

### GetDownloadEnabledOk

`func (o *VideoUploadRequestResumable) GetDownloadEnabledOk() (*bool, bool)`

GetDownloadEnabledOk returns a tuple with the DownloadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEnabled

`func (o *VideoUploadRequestResumable) SetDownloadEnabled(v bool)`

SetDownloadEnabled sets DownloadEnabled field to given value.

### HasDownloadEnabled

`func (o *VideoUploadRequestResumable) HasDownloadEnabled() bool`

HasDownloadEnabled returns a boolean if a field has been set.

### GetOriginallyPublishedAt

`func (o *VideoUploadRequestResumable) GetOriginallyPublishedAt() time.Time`

GetOriginallyPublishedAt returns the OriginallyPublishedAt field if non-nil, zero value otherwise.

### GetOriginallyPublishedAtOk

`func (o *VideoUploadRequestResumable) GetOriginallyPublishedAtOk() (*time.Time, bool)`

GetOriginallyPublishedAtOk returns a tuple with the OriginallyPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginallyPublishedAt

`func (o *VideoUploadRequestResumable) SetOriginallyPublishedAt(v time.Time)`

SetOriginallyPublishedAt sets OriginallyPublishedAt field to given value.

### HasOriginallyPublishedAt

`func (o *VideoUploadRequestResumable) HasOriginallyPublishedAt() bool`

HasOriginallyPublishedAt returns a boolean if a field has been set.

### GetScheduleUpdate

`func (o *VideoUploadRequestResumable) GetScheduleUpdate() VideoScheduledUpdate`

GetScheduleUpdate returns the ScheduleUpdate field if non-nil, zero value otherwise.

### GetScheduleUpdateOk

`func (o *VideoUploadRequestResumable) GetScheduleUpdateOk() (*VideoScheduledUpdate, bool)`

GetScheduleUpdateOk returns a tuple with the ScheduleUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUpdate

`func (o *VideoUploadRequestResumable) SetScheduleUpdate(v VideoScheduledUpdate)`

SetScheduleUpdate sets ScheduleUpdate field to given value.

### HasScheduleUpdate

`func (o *VideoUploadRequestResumable) HasScheduleUpdate() bool`

HasScheduleUpdate returns a boolean if a field has been set.

### GetThumbnailfile

`func (o *VideoUploadRequestResumable) GetThumbnailfile() *os.File`

GetThumbnailfile returns the Thumbnailfile field if non-nil, zero value otherwise.

### GetThumbnailfileOk

`func (o *VideoUploadRequestResumable) GetThumbnailfileOk() (**os.File, bool)`

GetThumbnailfileOk returns a tuple with the Thumbnailfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailfile

`func (o *VideoUploadRequestResumable) SetThumbnailfile(v *os.File)`

SetThumbnailfile sets Thumbnailfile field to given value.

### HasThumbnailfile

`func (o *VideoUploadRequestResumable) HasThumbnailfile() bool`

HasThumbnailfile returns a boolean if a field has been set.

### GetPreviewfile

`func (o *VideoUploadRequestResumable) GetPreviewfile() *os.File`

GetPreviewfile returns the Previewfile field if non-nil, zero value otherwise.

### GetPreviewfileOk

`func (o *VideoUploadRequestResumable) GetPreviewfileOk() (**os.File, bool)`

GetPreviewfileOk returns a tuple with the Previewfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewfile

`func (o *VideoUploadRequestResumable) SetPreviewfile(v *os.File)`

SetPreviewfile sets Previewfile field to given value.

### HasPreviewfile

`func (o *VideoUploadRequestResumable) HasPreviewfile() bool`

HasPreviewfile returns a boolean if a field has been set.

### GetVideoPasswords

`func (o *VideoUploadRequestResumable) GetVideoPasswords() []string`

GetVideoPasswords returns the VideoPasswords field if non-nil, zero value otherwise.

### GetVideoPasswordsOk

`func (o *VideoUploadRequestResumable) GetVideoPasswordsOk() (*[]string, bool)`

GetVideoPasswordsOk returns a tuple with the VideoPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoPasswords

`func (o *VideoUploadRequestResumable) SetVideoPasswords(v []string)`

SetVideoPasswords sets VideoPasswords field to given value.

### HasVideoPasswords

`func (o *VideoUploadRequestResumable) HasVideoPasswords() bool`

HasVideoPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


