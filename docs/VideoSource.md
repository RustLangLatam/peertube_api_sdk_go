# VideoSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Deprecated in 6.1, use inputFilename instead | [optional] 
**InputFilename** | Pointer to **string** | Uploaded/imported filename | [optional] 
**FileDownloadUrl** | Pointer to **string** | **PeerTube &gt;&#x3D; 6.1** If enabled by the admin, the video source file is kept on the server and can be downloaded by the owner | [optional] 
**Resolution** | Pointer to [**VideoResolutionConstant**](VideoResolutionConstant.md) |  | [optional] 
**Size** | Pointer to **int32** | **PeerTube &gt;&#x3D; 6.1** Video file size in bytes | [optional] 
**Fps** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Frames per second of the video file | [optional] 
**Width** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Video stream width | [optional] 
**Height** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Video stream height | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVideoSource

`func NewVideoSource() *VideoSource`

NewVideoSource instantiates a new VideoSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSourceWithDefaults

`func NewVideoSourceWithDefaults() *VideoSource`

NewVideoSourceWithDefaults instantiates a new VideoSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *VideoSource) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VideoSource) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VideoSource) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *VideoSource) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetInputFilename

`func (o *VideoSource) GetInputFilename() string`

GetInputFilename returns the InputFilename field if non-nil, zero value otherwise.

### GetInputFilenameOk

`func (o *VideoSource) GetInputFilenameOk() (*string, bool)`

GetInputFilenameOk returns a tuple with the InputFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFilename

`func (o *VideoSource) SetInputFilename(v string)`

SetInputFilename sets InputFilename field to given value.

### HasInputFilename

`func (o *VideoSource) HasInputFilename() bool`

HasInputFilename returns a boolean if a field has been set.

### GetFileDownloadUrl

`func (o *VideoSource) GetFileDownloadUrl() string`

GetFileDownloadUrl returns the FileDownloadUrl field if non-nil, zero value otherwise.

### GetFileDownloadUrlOk

`func (o *VideoSource) GetFileDownloadUrlOk() (*string, bool)`

GetFileDownloadUrlOk returns a tuple with the FileDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadUrl

`func (o *VideoSource) SetFileDownloadUrl(v string)`

SetFileDownloadUrl sets FileDownloadUrl field to given value.

### HasFileDownloadUrl

`func (o *VideoSource) HasFileDownloadUrl() bool`

HasFileDownloadUrl returns a boolean if a field has been set.

### GetResolution

`func (o *VideoSource) GetResolution() VideoResolutionConstant`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VideoSource) GetResolutionOk() (*VideoResolutionConstant, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VideoSource) SetResolution(v VideoResolutionConstant)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VideoSource) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSize

`func (o *VideoSource) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoSource) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoSource) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoSource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFps

`func (o *VideoSource) GetFps() float32`

GetFps returns the Fps field if non-nil, zero value otherwise.

### GetFpsOk

`func (o *VideoSource) GetFpsOk() (*float32, bool)`

GetFpsOk returns a tuple with the Fps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFps

`func (o *VideoSource) SetFps(v float32)`

SetFps sets Fps field to given value.

### HasFps

`func (o *VideoSource) HasFps() bool`

HasFps returns a boolean if a field has been set.

### GetWidth

`func (o *VideoSource) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoSource) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoSource) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideoSource) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VideoSource) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoSource) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoSource) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideoSource) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoSource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoSource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoSource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoSource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


